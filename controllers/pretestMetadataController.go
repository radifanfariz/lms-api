package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

func PreTestMetadataCreate(ctx *gin.Context) {
	var body struct {
		ModuleID    int       `json:"ModuleID"`
		Slug        string    `json:"Slug"`
		Name        string    `json:"Name" binding:"required"`
		Description string    `json:"Description" binding:"required"`
		MaxAccess   int       `json:"MaxAccess" binding:"required"`
		MinScore    float64   `json:"MinScore" binding:"required"`
		CreatedBy   string    `json:"CreatedBy"`
		CreatedAt   time.Time `json:"CreatedAt"`
	}

	ctx.Bind(&body)

	post := models.PreTestMetadata{
		ModuleID:    body.ModuleID,
		Slug:        body.Slug,
		Name:        body.Name,
		Description: body.Description,
		MaxAccess:   body.MaxAccess,
		MinScore:    body.MinScore,
		CreatedBy:   body.CreatedBy,
		CreatedAt:   body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "PreTest Metadata created successfully.", "data": post})
}

func PreTestMetadataFindById(ctx *gin.Context) {
	var PreTestMetadata models.PreTestMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&PreTestMetadata, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Metadata not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": PreTestMetadata})
}
func PreTestMetadataFindAll(ctx *gin.Context) {
	var PreTestMetadata []models.PreTestMetadata
	result := initializers.DB.Find(&PreTestMetadata)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": PreTestMetadata})
}

func PreTestMetadataUpdate(ctx *gin.Context) {
	var body struct {
		ModuleID    int       `json:"ModuleID"`
		Slug        string    `json:"Slug"`
		Name        string    `json:"Name" binding:"required"`
		Description string    `json:"Description" binding:"required"`
		MaxAccess   int       `json:"MaxAccess" binding:"required"`
		MinScore    float64   `json:"MinScore" binding:"required"`
		UpdatedBy   string    `json:"UpdatedBy"`
		UpdatedAt   time.Time `json:"UpdatedAt"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PreTestMetadata{
		ModuleID:    body.ModuleID,
		Slug:        body.Slug,
		Name:        body.Name,
		Description: body.Description,
		MaxAccess:   body.MaxAccess,
		MinScore:    body.MinScore,
		UpdatedBy:   body.UpdatedBy,
		UpdatedAt:   body.UpdatedAt,
	}

	var current models.PreTestMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Metadata not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PreTest Metadata.",
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Metadata updated successfully.", "data": &current})
}

func PreTestMetadataDelete(ctx *gin.Context) {
	var current models.PreTestMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Metadata not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PreTest Metadata.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Metadata deleted successfully.", "deletedData": &current})

}
