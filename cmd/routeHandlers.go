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

		if err := ctx.BindJSON(&registerPayload); err != nil {
			log.Printf("Error binding register payload:\n%+v\n", err)
			ctx.String(http.StatusNotFound, "Error")
			return
		}
		fmt.Printf("%+v\n", registerPayload)

		user, err := FindByUsername(db, registerPayload.Username)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error")
			return
		}
		if user.ID == -1 {
			if err := InsertUser(db, registerPayload); err != nil {
				ctx.String(http.StatusConflict, "Error")
			}
			registerResponse.Valid = true
			registerResponse.User.ID = 1
			registerResponse.User.Username = registerPayload.Username
			ctx.JSON(http.StatusOK, registerResponse)
			return
		}
		ctx.JSON(http.StatusOK, registerResponse)
	}

}
