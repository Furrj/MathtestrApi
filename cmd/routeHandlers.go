package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func LoginPost(db *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func RegisterPost(db *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerPayload RegisterPayload
		var registerResponse RegisterResponse
		registerResponse.Valid = false

		// Marshall JSON from request body
		if err := ctx.BindJSON(&registerPayload); err != nil {
			log.Printf("Error binding register payload:\n%+v\n", err)
			ctx.String(http.StatusNotFound, "Error")
			return
		}
		fmt.Printf("%+v\n", registerPayload)

		// Check if username exists
		user, err := FindByUsername(db, registerPayload.Username)
		if err != nil {
			log.Print("Error in FindByUsername")
			ctx.String(http.StatusBadRequest, "Error in FindByUsername")
			return
		}

		// If username doesn't exist
		if user.ID == -1 {
			createdUserClientData, err := CreateNewUser(db, registerPayload)
			if err != nil {
				log.Print("Error in CreateUser")
				ctx.String(http.StatusBadRequest, "Error in CreateUser")
				return
			}

			registerResponse.Valid = true
			registerResponse.User = createdUserClientData
			ctx.JSON(http.StatusOK, registerResponse)
			return
		}
		ctx.String(http.StatusBadRequest, "Error")
	}

}
