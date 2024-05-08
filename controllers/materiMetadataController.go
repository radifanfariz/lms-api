package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type MateriMetadataBody struct {
	ModuleID    int       `json:"module_id"`
	GlobalID    string    `json:"global_id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Src         string    `json:"src"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MateriMetadataCreate(ctx *gin.Context) {
	var body MateriMetadataBody

	ctx.Bind(&body)

	post := models.MateriMetadata{
		ModuleID:    body.ModuleID,
		GlobalID:    body.GlobalID,
		Name:        body.Name,
		Description: body.Description,
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata created successfully.", "data": post})
}

func MateriMetadataFindById(ctx *gin.Context) {
	var materiMetadata models.MateriMetadata

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&materiMetadata, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&materiMetadata, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Metadata not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": materiMetadata})
}
func MateriMetadataFindAll(ctx *gin.Context) {
	var materiMetadata []models.MateriMetadata
	result := initializers.DB.Find(&materiMetadata)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": materiMetadata})
}

func MateriMetadataUpdate(ctx *gin.Context) {
	var body MateriMetadataBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.MateriMetadata{
		ModuleID:    body.ModuleID,
		GlobalID:    body.GlobalID,
		Name:        body.Name,
		Description: body.Description,
		Src:         body.Src,
		UpdatedBy:   body.UpdatedBy,
		UpdatedAt:   body.UpdatedAt,
	}

	var current models.MateriMetadata
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Metadata not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Materi Metadata.",
		})
		return
	}

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResultAfterUpdate = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata updated successfully.", "data": &current})
}

func MateriMetadataUpsert(ctx *gin.Context) {
	var body MateriMetadataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.MateriMetadata
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.MateriMetadata{
				GlobalID:    body.GlobalID,
				ModuleID:    body.ModuleID,
				Name:        body.Name,
				Description: body.Description,
				Src:         body.Src,
				CreatedBy:   body.CreatedBy,
				CreatedAt:   body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Metadata Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.MateriMetadata{
				GlobalID:    id,
				ModuleID:    body.ModuleID,
				Name:        body.Name,
				Description: body.Description,
				Src:         body.Src,
				CreatedBy:   body.CreatedBy,
				CreatedAt:   body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Materi Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.MateriMetadata{
			ID:          current.ID,
			GlobalID:    current.GlobalID,
			ModuleID:    body.ModuleID,
			Name:        body.Name,
			Description: body.Description,
			Src:         body.Src,
			UpdatedBy:   body.UpdatedBy,
			UpdatedAt:   body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Materi Metadata.",
			})
			return
		}

		if govalidator.IsNumeric(ctx.Param("id")) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
		} else {
			id := ctx.Param("id")
			findByIdResultAfterUpdate = initializers.DB.First(&current, "c_global_id = ?", id)
		}

		if findByIdResultAfterUpdate.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Materi Metadata not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata updated successfully.", "data": &current})
	}
}

func MateriMetadataDelete(ctx *gin.Context) {
	var current models.MateriMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Metadata not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Materi Metadata.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Metadata deleted successfully.", "deletedData": &current})

}
