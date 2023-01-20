package user

import (
	"github.com/gin-gonic/gin"
	"github/Dyngela/medium/src/utils"
	"net/http"
)

func login(c *gin.Context) {
	input := LoginUserRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ThrowExceptionBadArgument(c, err)
		return
	}

	user, err := LoginUser(&input, c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, nil)
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		err = errors.New("your email or password doesn't match")
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  err,
			"data":   nil,
		})
		return
	}

	jwt, err := utils.GenerateToken(user.Id)

	c.JSON(http.StatusOK, jwt)
}

func register(c *gin.Context) {
	input := RegisterUserRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ThrowExceptionBadArgument(c, err)
		return
	}
	err := utils.HashPassword(&input.Password)
	if err != nil {
		utils.ThrowExceptionBadArgument(c, err)
	}
	response, err := registerUser(&input, c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
