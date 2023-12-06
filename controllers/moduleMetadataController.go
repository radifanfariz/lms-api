package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

type ModuleMetadataBody struct {
	ID              int       `json:"ID"`
	Slug            string    `json:"Slug"`
	Name            string    `json:"Name" binding:"required"`
	Description     string    `json:"Description" binding:"required"`
	Src             string    `json:"Src" binding:"required"`
	LearningJourney string    `json:"LearningJourney" binding:"required"`
	Category        string    `json:"Category" binding:"required"`
	MaxMonth        int       `json:"MaxMonth" binding:"required"`
	CreatedBy       string    `json:"CreatedBy"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedBy       string    `json:"UpdatedBy"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
}

func ModuleMetadataCreate(ctx *gin.Context) {
	var body ModuleMetadataBody

	ctx.Bind(&body)

	post := models.ModuleMetadata{
		Slug:            body.Slug,
		Name:            body.Name,
		Description:     body.Description,
		Src:             body.Src,
		LearningJourney: body.LearningJourney,
		Category:        body.Category,
		MaxMonth:        body.MaxMonth,
		CreatedBy:       body.CreatedBy,
		CreatedAt:       body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata created successfully.", "data": &post})
}

func ModuleMetadataFindById(ctx *gin.Context) {
	var ModuleMetadata models.ModuleMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&ModuleMetadata, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ModuleMetadata})
}
func ModuleMetadataFindAll(ctx *gin.Context) {
	var ModuleMetadata []models.ModuleMetadata
	result := initializers.DB.Find(&ModuleMetadata)

	// fmt.Println(ModuleMetadata)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ModuleMetadata})
}

func ModuleMetadataUpdate(ctx *gin.Context) {
	var body ModuleMetadataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.ModuleMetadata{
		Slug:            body.Slug,
		Name:            body.Name,
		Description:     body.Description,
		Src:             body.Src,
		LearningJourney: body.LearningJourney,
		Category:        body.Category,
		MaxMonth:        body.MaxMonth,
		UpdatedBy:       body.UpdatedBy,
		UpdatedAt:       body.UpdatedAt,
	}

	var current models.ModuleMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Module Metadata.",
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

func ModuleMetadataDelete(ctx *gin.Context) {
	var current models.ModuleMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Module Metadata.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data deleted successfully.", "deletedData": &current})

}
