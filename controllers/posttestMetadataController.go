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

type PostTestMetadataBody struct {
	ModuleID    int       `json:"module_id"`
	GlobalID    string    `json:"global_id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	MaxAccess   int       `json:"max_access" binding:"required"`
	MinScore    *float64  `json:"min_score" binding:"required"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PostTestMetadataCreate(ctx *gin.Context) {
	var body PostTestMetadataBody

	ctx.Bind(&body)

	post := models.PostTestMetadata{
		ModuleID:    body.ModuleID,
		GlobalID:    body.GlobalID,
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

	ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Metadata created successfully.", "data": post})
}

func PostTestMetadataFindById(ctx *gin.Context) {
	var postTestMetadata models.PostTestMetadata

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&postTestMetadata, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&postTestMetadata, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest Metadata not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestMetadata})
}
func PostTestMetadataFindAll(ctx *gin.Context) {
	var postTestMetadata []models.PostTestMetadata
	result := initializers.DB.Find(&postTestMetadata)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestMetadata})
}

func PostTestMetadataUpdate(ctx *gin.Context) {
	var body PostTestMetadataBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PostTestMetadata{
		ModuleID:    body.ModuleID,
		GlobalID:    body.GlobalID,
		Name:        body.Name,
		Description: body.Description,
		MaxAccess:   body.MaxAccess,
		MinScore:    body.MinScore,
		UpdatedBy:   body.UpdatedBy,
		UpdatedAt:   body.UpdatedAt,
	}

	var current models.PostTestMetadata
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
			"message": "PostTest Metadata not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PostTest Metadata.",
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
			"message": "PostTest Metadata not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Metadata updated successfully.", "data": &current})
}

func PostTestMetadataUpsert(ctx *gin.Context) {
	var body PostTestMetadataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.PostTestMetadata
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
			upsert := models.PostTestMetadata{
				GlobalID:    body.GlobalID,
				ModuleID:    body.ModuleID,
				Name:        body.Name,
				Description: body.Description,
				MaxAccess:   body.MaxAccess,
				MinScore:    body.MinScore,
				CreatedBy:   body.CreatedBy,
				CreatedAt:   body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PostTest Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Metadata created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.PostTestMetadata{
				GlobalID:    id,
				ModuleID:    body.ModuleID,
				Name:        body.Name,
				Description: body.Description,
				MaxAccess:   body.MaxAccess,
				MinScore:    body.MinScore,
				UpdatedBy:   body.UpdatedBy,
				UpdatedAt:   body.UpdatedAt,
				CreatedBy:   body.CreatedBy,
				CreatedAt:   body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PostTest Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Metadata created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.PostTestMetadata{
			ID:          current.ID,
			GlobalID:    body.GlobalID,
			ModuleID:    body.ModuleID,
			Name:        body.Name,
			Description: body.Description,
			MaxAccess:   body.MaxAccess,
			MinScore:    body.MinScore,
			UpdatedBy:   body.UpdatedBy,
			UpdatedAt:   body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating PostTest Metadata.",
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
				"message": "PostTest Metadata not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Metadata updated successfully.", "data": &current})
	}
}

func PostTestMetadataDelete(ctx *gin.Context) {
	var current models.PostTestMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest Metadata not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PostTest Metadata.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Metadata deleted successfully.", "deletedData": &current})

}
