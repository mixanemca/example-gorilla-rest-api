// Package v1 for endpoints
package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixanemca/example-gorilla-rest-api/internal/app/api/utils"
	"github.com/mixanemca/example-gorilla-rest-api/models"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go

type UserRepository interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
type UserRepo struct {
	db         *pgxpool.Pool
	validate   *validator.Validate
	translator ut.Translator
}

func NewUserRepository(db *pgxpool.Pool) *UserRepo {
	validator, translator := utils.NewValidator()
	return &UserRepo{
		db:         db,
		validate:   validator,
		translator: translator,
	}
}

// CreateUser method for create new user
// @Summary Create user
// @Tags users
// @Description Create user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.IDInfo
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user [post]
func (u UserRepo) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest
	// Parse the JSON data from the request body and store it in the user variable
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := u.validate.Struct(user); err != nil {
		jsonErrResponse(w, u.translator, err)
		return
	}

	var userID models.IDInfo
	dbUser := user.ConvertToEntity()
	query := "INSERT INTO users (created_at, updated_at, deleted_at, name, surname, username, email, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := u.db.QueryRow(r.Context(), query, time.Now(), time.Now(), nil, dbUser.Name, dbUser.Surname, dbUser.Username, dbUser.Email, dbUser.Phone).Scan(&userID.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusOK, userID)
}

// GetUsers method for get all users
// @Summary Get users
// @Tags users
// @Description  Get users
// @Success 200 {object} models.Users
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user/list [get]
func (u UserRepo) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users models.Users
	query := "SELECT id, name, surname, username, email, phone FROM users WHERE deleted_at IS NULL;"
	rows, err := u.db.Query(r.Context(), query)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserResponse{}
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Email, &user.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	jsonResponse(w, http.StatusOK, users)
}

// GetUserByID method for get user by id
// @Summary Get user by ID
// @Tags users
// @Description Get user by ID
// @Param id path string true "user ID"
// @Success 200 {object} models.User
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user/{id} [get]
func (u UserRepo) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user models.UserResponse
	id := mux.Vars(r)["id"]

	rows, err := u.getUserByID(w, r, id)
	if err != nil {
		return
	}

	if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Email, &user.Phone); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rows.Close()
	jsonResponse(w, http.StatusOK, user)
}

// getUserByID method is used in other methods where it is necessary to check if a user exists before the main query
func (u UserRepo) getUserByID(w http.ResponseWriter, r *http.Request, id string) (pgx.Rows, error) {
	if !utils.ValidUUID(id) {
		http.Error(w, "Invalid UUID format for user ID", http.StatusBadRequest)
		return nil, errors.New("invalid UUID format for user ID")
	}
	query := "SELECT id, name, surname, username, email, phone FROM users WHERE id=$1 LIMIT 1;"
	rows, err := u.db.Query(r.Context(), query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	if !rows.Next() {
		rows.Close()
		http.Error(w, "User not found", http.StatusNotFound)
		return nil, pgx.ErrNoRows
	}
	return rows, nil
}

// UpdateUser method for update user
// @Summary Update user
// @Tags users
// @Description Update user
// @Param id path string true "user ID"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 204
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user/{id} [put]
func (u UserRepo) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest
	id := mux.Vars(r)["id"]

	// сheck if the user exists
	rows, err := u.getUserByID(w, r, id)
	if err != nil {
		return
	}
	rows.Close()

	// Parse the JSON data from the request body and store it in the user variable
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := u.validate.Struct(user); err != nil {
		jsonErrResponse(w, u.translator, err)
		return
	}

	dbUser := user.ConvertToEntity()
	query := "UPDATE  users SET updated_at=$1, name=$2, surname=$3, username=$4, email=$5, phone=$6  WHERE id=$7;"
	_, err = u.db.Exec(r.Context(), query, time.Now(), dbUser.Name, dbUser.Surname, dbUser.Username, dbUser.Email, dbUser.Phone, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusNoContent, nil)
}

// DeleteUser method for delete user
// @Summary Delete user
// @Tags users
// @Description Delete user
// @Param id path string true "user ID"
// @Success 204
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user/{id} [delete]
func (u UserRepo) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// сheck if the user exists
	rows, err := u.getUserByID(w, r, id)
	if err != nil {
		return
	}
	rows.Close()

	query := "UPDATE  users SET deleted_at=$1 WHERE id=$2;"
	_, err = u.db.Exec(r.Context(), query, time.Now(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusNoContent, nil)
}
