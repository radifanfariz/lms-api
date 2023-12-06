package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

type ModuleDataBody struct {
	ID             int    `json:"ID"`
	Slug           string `json:"Slug"`
	ModuleMetaID   int    `json:"ModuleMetaID"`
	PretestMetaID  int    `json:"PretestMetaID"`
	PretestID      int    `json:"PretestID"`
	MateriMetaID   int    `json:"MateriMetaID"`
	MateriID       int    `json:"MateriID"`
	PosttestMetaID int    `json:"PosttestMetaID"`
	PosttestID     int    `json:"PosttestID"`
	UserID         int    `json:"UserID"`
	GradeID        int    `json:"GradeID"`
}

func ModuleDataCreate(ctx *gin.Context) {
	var body ModuleDataBody

	ctx.Bind(&body)

	post := models.ModuleData{
		ID:             body.ID,
		Slug:           body.Slug,
		ModuleMetaID:   body.ModuleMetaID,
		PretestMetaID:  body.PretestMetaID,
		PretestID:      body.PretestID,
		MateriMetaID:   body.MateriMetaID,
		PosttestMetaID: body.PosttestMetaID,
		PosttestID:     body.PosttestID,
		UserID:         body.UserID,
		GradeID:        body.GradeID,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Data created successfully.", "data": &post})
}

func ModuleDataFindById(ctx *gin.Context) {
	var ModuleData models.ModuleData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&ModuleData, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ModuleData})
}
func ModuleDataFindAll(ctx *gin.Context) {
	var ModuleData []models.ModuleData
	result := initializers.DB.Find(&ModuleData)

	// fmt.Println(ModuleData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ModuleData})
}

func ModuleDataUpdate(ctx *gin.Context) {
	var body ModuleDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.ModuleData{
		ID:             body.ID,
		Slug:           body.Slug,
		ModuleMetaID:   body.ModuleMetaID,
		PretestMetaID:  body.PretestMetaID,
		PretestID:      body.PretestID,
		MateriMetaID:   body.MateriMetaID,
		PosttestMetaID: body.PosttestMetaID,
		PosttestID:     body.PosttestID,
		UserID:         body.UserID,
		GradeID:        body.GradeID,
	}

	var current models.ModuleData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Module Data.",
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

func ModuleDataDelete(ctx *gin.Context) {
	var current models.ModuleData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Module Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data deleted successfully.", "deletedData": &current})

}
