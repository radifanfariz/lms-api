package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/utils"
)

type UploadFile struct {
	Title    string                `form:"title"`
	File     *multipart.FileHeader `form:"file"`
	Callback *string               `form:"callback"`
}

type BodyErrorResOSS struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
type DataSuccessResOSS struct {
	Filename   string `json:"fileName"`
	ClientID   int    `json:"clientId"`
	CreateDate string `json:"createdDate"`
	FileUrl    string `json:"fileUrl"`
	MediaId    int    `json:"mediaId"`
	Title      string `json:"title"`
}
type BodySuccessResOSS struct {
	Status string            `json:"status"`
	Data   DataSuccessResOSS `json:"data"`
}

func UploadFileCreate(ctx *gin.Context) {
	ossUrl := os.Getenv("OSS_URL")
	ossBasicAuthUsername := os.Getenv("OSS_BASIC_AUTH_USERNAME")
	ossBasicAuthPassword := os.Getenv("OSS_BASIC_AUTH_PASSWORD")

	var formData UploadFile

	// ctx.Bind(&formData)

	if err := ctx.ShouldBind(&formData); err != nil {
		log.Fatal(err)
		return
	}
	// fmt.Println(formData)
	file, header, errFile := ctx.Request.FormFile("file")
	if errFile != nil {
		http.Error(ctx.Writer, errFile.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// fmt.Println(file, header, errFile)

	form := map[string]string{"title": formData.Title, "file": header.Filename}
	ct, body, errForm := createForm(form, file)
	if errForm != nil {
		log.Fatal(errForm)
		return
	}
	req, err := http.NewRequest(http.MethodPost, ossUrl, body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", ct)
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(ossBasicAuthUsername, ossBasicAuthPassword))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		defer resp.Body.Close()
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	var jsonResultSuccess BodySuccessResOSS
	json.Unmarshal([]byte(responseBody), &jsonResultSuccess)

	var jsonResultFailed BodyErrorResOSS
	json.Unmarshal([]byte(responseBody), &jsonResultFailed)

	fmt.Println(resp.Status)
	fmt.Println(jsonResultSuccess)
	fmt.Println(string(responseBody))

	if strings.ToLower(jsonResultSuccess.Status) == "success" {
		// findByIdResult := initializers.DB.Where("c_nik = ?", credentials.NIK).First(&userData)

		// if findByIdResult.Error != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"message": "Invalid credentials !",
		// 	})
		// 	return
		// }
		ctx.JSON(http.StatusOK, jsonResultSuccess)
		return
	}
	ctx.JSON(http.StatusOK, jsonResultFailed)
}

func createForm(form map[string]string, file multipart.File) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if key == "file" {
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file)
		} else {
			mp.WriteField(key, val)
		}
	}
	return mp.FormDataContentType(), body, nil
}
