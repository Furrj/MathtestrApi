package routeHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
	"net/http"
)

func (r *RouteHandler) ProfilePageInfo(ctx *gin.Context) {
	var requestPayload schemas.ProfilePagePayload
	var responsePayload schemas.ProfilePageResponse
	var tests []schemas.TestResultsResponse

	// Bind request body
	if err := ctx.BindJSON(&requestPayload); err != nil {
		fmt.Printf("Error binding requestPayload json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", requestPayload)

	// Get teacher name
	teacherInfo, err := r.dbHandler.GetBasicUserInfoByID(requestPayload.TeacherID)
	if err != nil {
		fmt.Printf("Error getting teacher name: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	responsePayload.TeacherName = fmt.Sprintf("%s %s", teacherInfo.FirstName, teacherInfo.LastName)

	// Get tests
	testResults, err := r.dbHandler.GetTestResultsByUserID(int(requestPayload.ID))
	if err != nil {
		fmt.Printf("Error querying test results: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}

	for _, v := range testResults {
		tests = append(tests, schemas.TestResultsResponse{
			Score:         v.Score,
			Min:           v.Min,
			Max:           v.Max,
			QuestionCount: v.QuestionCount,
			Operations:    v.Operations,
			Timestamp:     v.Timestamp,
		})
	}
	responsePayload.Tests = tests

	ctx.JSON(http.StatusOK, responsePayload)
}
