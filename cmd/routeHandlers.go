package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {
	var user User
	user.ID = 1
	ctx.BindJSON(&user)
	fmt.Printf("%+v\n", user)
}
