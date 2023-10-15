package v1

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mixanemca/example-gorilla-rest-api/models"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewUserRepository(db)

	tests := []struct {
		name          string
		input         models.User
		mockBehavior  func()
		expected      string
		expectedError bool
	}{
		{
			name: "ok",
			input: models.User{
				Name:     "Peter",
				Surname:  "Parker",
				Username: "spider",
				Email:    "p.parker@gmail.com",
				Phone:    "+79267775511",
			},
			mockBehavior: func() {
				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "users"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Peter", "Parker", "spider", "p.parker@gmail.com", "+79267775511").
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
				mock.ExpectCommit()
			},
		},
		{
			name: "database error",
			input: models.User{
				Name:     "Peter",
				Surname:  "Parker",
				Username: "spider",
				Email:    "p.parker@gmail.com",
				Phone:    "+79267775511",
			},
			mockBehavior: func() {
				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Peter", "Parker", "spider", "p.parker@gmail.com", "+79267775511").
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior()

			got, err := r.CreateUser(test.input)
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}
