package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

type PretestDataBody struct {
	ModuleID      int          `json:"ModuleID"`
	PreTestMetaID int          `json:"PreTestMetaID"`
	Slug          string       `json:"Slug"`
	Question      models.JSONB `json:"Question" gorm:"type:jsonb"`
	IsPublished   bool         `json:"IsPublished"`
	CreatedBy     string       `json:"CreatedBy"`
	CreatedAt     time.Time    `json:"CreatedAt"`
	UpdatedBy     string       `json:"UpdatedBy"`
	UpdatedAt     time.Time    `json:"UpdatedAt"`
}

func PreTestDataCreate(ctx *gin.Context) {
	var body PretestDataBody

	ctx.Bind(&body)

	post := models.PreTestData{
		ModuleID:      body.ModuleID,
		PreTestMetaID: body.PreTestMetaID,
		Slug:          body.Slug,
		Question:      body.Question,
		IsPublished:   body.IsPublished,
		CreatedBy:     body.CreatedBy,
		CreatedAt:     body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "PreTest Data created successfully.", "data": &post})
}

func PreTestDataFindById(ctx *gin.Context) {
	var PreTestData models.PreTestData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&PreTestData, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": PreTestData})
}
func PreTestDataFindAll(ctx *gin.Context) {
	var PreTestData []models.PreTestData
	result := initializers.DB.Find(&PreTestData)

	// fmt.Println(PreTestData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": PreTestData})
}

func PreTestDataUpdate(ctx *gin.Context) {
	var body PretestDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PreTestData{
		ModuleID:      body.ModuleID,
		PreTestMetaID: body.PreTestMetaID,
		Slug:          body.Slug,
		Question:      body.Question,
		IsPublished:   body.IsPublished,
		UpdatedBy:     body.UpdatedBy,
		UpdatedAt:     body.UpdatedAt,
	}

	var current models.PreTestData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PreTest Data.",
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

func PreTestDataDelete(ctx *gin.Context) {
	var current models.PreTestData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PreTest Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data deleted successfully.", "deletedData": &current})

}
