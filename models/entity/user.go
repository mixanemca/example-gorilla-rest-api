package entity

type User struct {
	BaseModel
	Firstname string `db:"firstname"`
	Surname   string `db:"surname"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}
