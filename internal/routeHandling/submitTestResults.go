package routeHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
	"net/http"
)

func (r *RouteHandler) SubmitTestResults(ctx *gin.Context) {
	var testResultsPayload schemas.TestResults

	// Bind testResult payload
	if err := ctx.BindJSON(&testResultsPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", testResultsPayload)

	if err := r.dbHandler.InsertTestResults(testResultsPayload); err != nil {
		fmt.Printf("Error inserting test results: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}

	ctx.String(http.StatusOK, "Submitted test results succesfully")
}
