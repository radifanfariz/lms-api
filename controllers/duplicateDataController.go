package controllers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

func DuplicateLearningModuleAllData(ctx *gin.Context) {
	var moduleData models.ModuleData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.
			Preload("Metadata").
			Preload("UserData").
			Preload("PreTestMetadata").
			Preload("MateriMetadata").
			Preload("PostTestMetadata").
			Preload("PreTestData").
			Preload("MateriData").
			Preload("PostTestData").
			Where("b_ispublished = ?", true).
			Find(&moduleData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.
			Preload("Metadata").
			Preload("UserData").
			Preload("PreTestMetadata").
			Preload("MateriMetadata").
			Preload("PostTestMetadata").
			Preload("PreTestData").
			Preload("MateriData").
			Preload("PostTestData").
			Where("b_ispublished = ?", true).
			Find(&moduleData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
		return
	}

	// Note the use of tx as the database handle once you are within a transaction
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	globalId := uuid.New()

	moduleData.Metadata.GlobalID = globalId.String()
	if err := tx.Model(models.ModuleMetadata{}).
		Omit("ID").
		Create(&moduleData.Metadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PreTestMetadata.GlobalID = globalId.String()
	if err := tx.Model(models.PreTestMetadata{}).
		Omit("ID").
		Create(&moduleData.PreTestMetadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PreTestData.GlobalID = globalId.String()
	if err := tx.Model(models.PreTestData{}).
		Omit("ID").
		Create(&moduleData.PreTestData).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.MateriMetadata.GlobalID = globalId.String()
	if err := tx.Omit("ID").
		Create(&moduleData.MateriMetadata).
		Model(models.MateriData{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	for _, element := range moduleData.MateriData {
		element.GlobalID = globalId.String()
		if err := tx.Model(models.MateriData{}).
			Omit("ID").
			Create(&element).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
			})
			return
		}
	}

	moduleData.PostTestMetadata.GlobalID = globalId.String()
	if err := tx.Model(models.PostTestMetadata{}).
		Omit("ID").
		Create(&moduleData.PostTestMetadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PostTestData.GlobalID = globalId.String()
	if err := tx.Model(models.PostTestData{}).
		Omit("ID").
		Create(&moduleData.PostTestData).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Learning Module duplicated successfully."})
}
