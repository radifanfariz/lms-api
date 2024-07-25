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

type MateriDataBody struct {
	ModuleID     int       `json:"module_id"`
	MateriMetaID int       `json:"materi_meta_id"`
	GlobalID     string    `json:"global_id"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Type         string    `json:"type"`
	Src          string    `json:"src"`
	IsPublished  *bool     `json:"is_publishing"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func MateriDataCreate(ctx *gin.Context) {
	var body MateriDataBody

	ctx.Bind(&body)

	post := models.MateriData{
		ModuleID:     body.ModuleID,
		MateriMetaID: body.MateriMetaID,
		GlobalID:     body.GlobalID,
		Name:         body.Name,
		Description:  body.Description,
		Type:         body.Type,
		Src:          body.Src,
		IsPublished:  body.IsPublished,
		CreatedBy:    body.CreatedBy,
		CreatedAt:    body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data created successfully.", "data": &post})
}

func MateriDataFindById(ctx *gin.Context) {
	var materiData []models.MateriData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Preload("Metadata").Find(&materiData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.Preload("Metadata").Find(&materiData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": materiData})
}
func MateriDataFindAll(ctx *gin.Context) {
	var materiData []models.MateriData
	result := initializers.DB.Preload("Metadata").Find(&materiData)

	// fmt.Println(materiData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": materiData})
}

func MateriDataUpdate(ctx *gin.Context) {
	var body MateriDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.MateriData{
		ModuleID:     body.ModuleID,
		MateriMetaID: body.MateriMetaID,
		GlobalID:     body.GlobalID,
		Name:         body.Name,
		Description:  body.Description,
		Type:         body.Type,
		Src:          body.Src,
		IsPublished:  body.IsPublished,
		UpdatedBy:    body.UpdatedBy,
		UpdatedAt:    body.UpdatedAt,
	}

	var current models.MateriData
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
			"message": "Materi Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Materi Data.",
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data updated successfully.", "data": &current})
}

func MateriDataUpsert(ctx *gin.Context) {
	var body MateriDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.MateriData
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
			upsert := models.MateriData{
				GlobalID:     body.GlobalID,
				ModuleID:     body.ModuleID,
				MateriMetaID: body.MateriMetaID,
				Name:         body.Name,
				Description:  body.Description,
				Type:         body.Type,
				Src:          body.Src,
				IsPublished:  body.IsPublished,
				CreatedBy:    body.UpdatedBy,
				CreatedAt:    body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PreTest Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.MateriData{
				GlobalID:     id,
				ModuleID:     body.ModuleID,
				MateriMetaID: body.MateriMetaID,
				Name:         body.Name,
				Description:  body.Description,
				Type:         body.Type,
				Src:          body.Src,
				IsPublished:  body.IsPublished,
				CreatedBy:    body.UpdatedBy,
				CreatedAt:    body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Materi Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.MateriData{
			ID:           current.ID,
			GlobalID:     current.GlobalID,
			ModuleID:     body.ModuleID,
			MateriMetaID: body.MateriMetaID,
			Name:         body.Name,
			Description:  body.Description,
			Type:         body.Type,
			Src:          body.Src,
			IsPublished:  body.IsPublished,
			CreatedBy:    current.CreatedBy,
			CreatedAt:    current.CreatedAt,
			UpdatedBy:    body.UpdatedBy,
			UpdatedAt:    body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Materi Data.",
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

		ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data updated successfully.", "data": &current})
	}
}

func MateriDataDelete(ctx *gin.Context) {
	var current models.MateriData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Materi Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data deleted successfully.", "deletedData": &current})

}
