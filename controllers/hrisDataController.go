package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/models"
	"github.com/radifanfariz/lms-api/utils"
)

type BodyDataHris struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    UserDataPortal `json:"data"`
}

func EmployeeDataFindByParams(ctx *gin.Context) {
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	mainCompanyId := ""
	if ctx.Query("mainCompanyId") != "" {
		mainCompanyId = ctx.Query("mainCompanyId")
	}
	employeeId := ""
	if ctx.Query("employeeId") != "" {
		employeeId = ctx.Query("employeeId")
	}
	employeeName := ""
	if ctx.Query("employeeName") != "" {
		employeeName = ctx.Query("employeeName")
	}
	employeeStatus := ""
	if ctx.Query("employeeStatus") != "" {
		employeeStatus = ctx.Query("employeeStatus")
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/employee?mainCompanyId="+mainCompanyId+"&employeeId="+employeeId+"&employeeName="+employeeName+"&employeeStatus="+employeeStatus, nil)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(hrisBasicAuthUsername, hrisBasicAuthPassword))
	req.Header.Add("key-access", hrisKeyAccess)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	var jsonResult models.EmployeeDataFromHris
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	// fmt.Println(string(responseBody))

	ctx.JSON(http.StatusOK, jsonResult)

}

func GradeDataFindByParams(ctx *gin.Context) {
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	mainCompanyId := ""
	if ctx.Query("mainCompanyId") != "" {
		mainCompanyId = ctx.Query("mainCompanyId")
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/employee/grade?mainCompanyId="+mainCompanyId, nil)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(hrisBasicAuthUsername, hrisBasicAuthPassword))
	req.Header.Add("key-access", hrisKeyAccess)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	var jsonResult models.GradeDataFromHris
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	// fmt.Println(string(responseBody))

	ctx.JSON(http.StatusOK, jsonResult)

}
