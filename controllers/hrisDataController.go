package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/models"
	"github.com/radifanfariz/lms-api/utils"
)

func EmployeeDataFindByParams(ctx *gin.Context) {
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	var employeeId string
	if ctx.Query("employeeId") != "" {
		employeeId = url.QueryEscape(ctx.Query("employeeId"))
	}
	var mainCompanyId string
	if ctx.Query("mainCompanyId") != "" {
		mainCompanyId = url.QueryEscape(ctx.Query("mainCompanyId"))
	}
	var limit string = "999999"
	if ctx.Query("limit") != "" {
		mainCompanyId = url.QueryEscape(ctx.Query("limit"))
	}
	var employeeStatus string = "active"
	if ctx.Query("employeeStatus") != "" {
		employeeStatus = url.QueryEscape(ctx.Query("employeeStatus"))
	}
	var employeeName string
	if ctx.Query("employeeName") != "" {
		employeeName = url.QueryEscape(ctx.Query("employeeName"))
	}
	var employeeNik string
	if ctx.Query("employeeNik") != "" {
		employeeNik = url.QueryEscape(ctx.Query("employeeNik"))
	}
	var employeePosition string
	if ctx.Query("employeePosition") != "" {
		employeePosition = url.QueryEscape(ctx.Query("employeePosiition"))
	}
	var mainCompanyName string
	if ctx.Query("mainCompanyName") != "" {
		mainCompanyName = url.QueryEscape(ctx.Query("mainCompanyName"))
	}
	var joinDate string
	if ctx.Query("joinDate") != "" {
		joinDate = url.QueryEscape(ctx.Query("joinDate"))
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+
		"/smartmulia/employee?mainCompanyId="+mainCompanyId+"&mainCompanyName="+mainCompanyName+"&employeeId="+
		employeeId+"&employeeName="+employeeName+"&employeeStatus="+employeeStatus+
		"&employeeNik="+employeeNik+"&employeePosition="+employeePosition+"&joinDate="+joinDate+"&limit="+limit,
		nil)
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
	req.Header.Add("Access-Control-Allow-Origin", "*")
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
		mainCompanyId = url.QueryEscape(ctx.Query("mainCompanyId"))
	}
	limit := ""
	if ctx.Query("limit") != "" {
		limit = url.QueryEscape(ctx.Query("limit"))
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/employee/grade?mainCompanyId="+mainCompanyId+"&limit="+limit, nil)
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
	req.Header.Add("Access-Control-Allow-Origin", "*")
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

func PositionDataFindByParams(ctx *gin.Context) {
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	mainCompanyId := ""
	if ctx.Query("mainCompanyId") != "" {
		mainCompanyId = url.QueryEscape(ctx.Query("mainCompanyId"))
	}
	limit := ""
	if ctx.Query("limit") != "" {
		limit = url.QueryEscape(ctx.Query("limit"))
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/position?mainCompanyId="+mainCompanyId+"&limit="+limit, nil)
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
	req.Header.Add("Access-Control-Allow-Origin", "*")
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

	var jsonResult models.PositionDataFromHris
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	// fmt.Println(string(responseBody))

	ctx.JSON(http.StatusOK, jsonResult)

}
