package routeHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
	"net/http"
)

func (r *RouteHandler) GetTestResults(ctx *gin.Context) {
	var requestPayload schemas.ValidationPayload
	var responsePayload []schemas.TestResultsResponse

	// Bind request body
	if err := ctx.BindJSON(&requestPayload); err != nil {
		fmt.Printf("Error binding requestPayload json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", requestPayload)

	testResults, err := r.dbHandler.GetTestResultsByUserID(int(requestPayload.ID))
	if err != nil {
		fmt.Printf("Error querying test results: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}

	for _, v := range testResults {
		responsePayload = append(responsePayload, schemas.TestResultsResponse{
			Score:         v.Score,
			Min:           v.Min,
			Max:           v.Max,
			QuestionCount: v.QuestionCount,
			Operations:    v.Operations,
			Timestamp:     v.Timestamp,
		})
	}

	ctx.JSON(http.StatusOK, responsePayload)
}
