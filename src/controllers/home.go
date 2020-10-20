package controllers

import (
	"github.com/kataras/iris/v12"
)

// Home ... Represents a home
type Home struct {
	Message string `json:"message"`
}

// GreetUser godoc
// @Summary Greets a user
// @Produces json
// @Success 200 {object} Home
// @Router /home [get]
func GreetUser(ctx iris.Context) {
	home := Home{"Hello, user!"}
	ctx.JSON(home)
}
