package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func UserLogin(db *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var submittedUserInfo, validatedUserInfo User
		invalidUser := User{
			ID:       -1,
			Username: "",
			Password: "",
		}

		// Bind data to JSON
		if err := ctx.BindJSON(&submittedUserInfo); err != nil {
			log.Printf("Error unmarshalling user login info: %+v\n", err)
			ctx.JSON(http.StatusBadRequest, invalidUser)
			return
		}

		// Check database for user
		validatedUserInfo, err := FindByUsername(db, submittedUserInfo)
		if err != nil {
			log.Printf("Error searching database for user info:\n%+v\n", err)
			ctx.JSON(http.StatusNotFound, invalidUser)
			return
		}

		fmt.Printf("%+v\n", validatedUserInfo)
		ctx.JSON(http.StatusOK, validatedUserInfo)
	}

}

func UserRegister(db *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bytes, _ := ctx.GetRawData()
		data := string(bytes[:])
		fmt.Printf("%+v\n", data)
	}

}
