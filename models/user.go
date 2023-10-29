// Package models for users
package models

import (
	"regexp"
	"strings"
)

type Users []UserResponse
type User struct {
	BaseModel
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u User) Validate() string {
	var errMessage []string
	if re := regexp.MustCompile(`^[A-Za-z]{3,}$`); !re.MatchString(u.Name) {
		errMessage = append(errMessage, "Name should consist of only English letters with a minimum length of 3 characters.")
	}
	if re := regexp.MustCompile(`^[A-Za-z]{3,}$`); !re.MatchString(u.Surname) {
		errMessage = append(errMessage, "Surname should consist of only English letters with a minimum length of 3 characters.")
	}
	if re := regexp.MustCompile(`^[A-Za-z0-9]{3,}$`); !re.MatchString(u.Username) {
		errMessage = append(errMessage, "Username should consist of only English letters and digits with a minimum length of 3 characters.")
	}
	if re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`); !re.MatchString(u.Email) {
		errMessage = append(errMessage, "Email should be in a valid format (e.g., john@example.com)")
	}
	if re := regexp.MustCompile(`^\+7[89]\d{9}$`); !re.MatchString(u.Phone) {
		errMessage = append(errMessage, "Phone number should be validated according to specific phone number format requirements (example: +79251115599).")
	}

	return strings.Join(errMessage, "\n")
}
