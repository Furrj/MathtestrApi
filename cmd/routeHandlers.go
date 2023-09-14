package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func UserLogin(db *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var submittedUserInfo, validatedUserInfo User
		invalidUser := User{
			ID:       0,
			Username: "",
			Password: "",
		}

		if err := ctx.BindJSON(&submittedUserInfo); err != nil {
			log.Printf("Error unmarshalling user login info: %+v\n", err)
			ctx.JSON(http.StatusBadRequest, invalidUser)
			return
		}

		validatedUserInfo, err := FindByUsername(db, submittedUserInfo)
		if err != nil {
			log.Printf("Error searching database for user info:\n%+v\n", err)
			ctx.JSON(http.StatusNotFound, invalidUser)
			return
		}

		ctx.JSON(http.StatusOK, validatedUserInfo)
	}

}
