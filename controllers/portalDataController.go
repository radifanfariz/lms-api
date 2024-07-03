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

func CompanyDataFindAll(ctx *gin.Context) {
	portalUrl := os.Getenv("PORTAL_URL")
	portalBasicAuthUsername := os.Getenv("PORTAL_BASIC_AUTH_USERNAME")
	portalBasicAuthPassword := os.Getenv("PORTAL_BASIC_AUTH_PASSWORD")
	portalToken := ctx.Request.Header["X-Portal-Token"]

	req, err := http.NewRequest(http.MethodGet, portalUrl+
		"/bu/all",
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
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(portalBasicAuthUsername, portalBasicAuthPassword))
	req.Header.Add("X-Portal-Token", portalToken[0])
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

	var jsonResult models.CompanyDataFromPortal
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	// fmt.Println(string(responseBody))

	ctx.JSON(http.StatusOK, jsonResult)

}
