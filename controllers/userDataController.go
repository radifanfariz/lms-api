package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

func UserDataCreate(ctx *gin.Context) {
	var body struct {
		EmployeeID int       `json:"EmployeeID"`
		Name       string    `json:"Name"`
		NIK        string    `json:"NIK"`
		Level      string    `json:"Level"`
		LevelID    int       `json:"LevelID"`
		Grade      string    `json:"Grade"`
		GradeID    int       `json:"GradeID"`
		CreatedBy  string    `json:"CreatedBy"`
		CreatedAt  time.Time `json:"CreatedAt"`
	}

	ctx.Bind(&body)

	post := models.UserData{
		EmployeeID: body.EmployeeID,
		Name:       body.Name,
		NIK:        body.NIK,
		Level:      body.Level,
		LevelID:    body.LevelID,
		Grade:      body.Grade,
		GradeID:    body.GradeID,
		// CreatedBy:  body.CreatedBy,
		// CreatedAt:  body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &post})
}

func UserDataFindById(ctx *gin.Context) {
	var UserData models.UserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&UserData, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UserData})
}
func UserDataFindAll(ctx *gin.Context) {
	var UserData []models.UserData
	result := initializers.DB.Find(&UserData)

	// fmt.Println(UserData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UserData})
}

func UserDataUpdate(ctx *gin.Context) {
	var body struct {
		EmployeeID int       `json:"EmployeeID"`
		Name       string    `json:"Name"`
		NIK        string    `json:"NIK"`
		Level      string    `json:"Level"`
		LevelID    int       `json:"LevelID"`
		Grade      string    `json:"Grade"`
		GradeID    int       `json:"GradeID"`
		UpdatedBy  string    `json:"CreatedBy"`
		UpdatedAt  time.Time `json:"CreatedAt"`
	}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.UserData{
		EmployeeID: body.EmployeeID,
		Name:       body.Name,
		NIK:        body.NIK,
		Level:      body.Level,
		LevelID:    body.LevelID,
		Grade:      body.Grade,
		GradeID:    body.GradeID,
		// UpdatedBy:  body.UpdatedBy,
		// UpdatedAt:  body.UpdatedAt,
	}

	var current models.UserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating User Data.",
		})
		return
	}

	findByIdResultAfterUpdate := initializers.DB.First(&current, uint(id))

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data updated successfully.", "data": &current})
}

func UserDataDelete(ctx *gin.Context) {
	var current models.UserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting User Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data deleted successfully.", "deletedData": &current})

}
