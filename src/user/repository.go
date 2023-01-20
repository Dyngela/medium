package user

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github/Dyngela/medium/src/utils"
)

func LoginUser(request *LoginUserRequest, c *gin.Context) (LoginResponse, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return LoginResponse{}, err
	}
	var user LoginResponse
	query := `select user_id, email, password from users where email = $1`
	err = pgxscan.Get(c, conn, &user, query, request.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	return user, nil
}

func registerUser(registerRequest *RegisterUserRequest, c *gin.Context) (string, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return "", err
	}

	query := `insert into users (username, email, password, created_at) values ($1, $2, $3, now());`
	_, err = conn.Query(c, query, registerRequest.Username, registerRequest.Email, pass)
	if err != nil {
		return "", err
	}
	return "Account successfully created", nil
}