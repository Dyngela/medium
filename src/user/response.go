package user

type LoginResponse struct {
	Id       uint   `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
