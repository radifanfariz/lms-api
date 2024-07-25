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

type PreTestDataBody struct {
	ModuleID      int          `json:"module_id"`
	PreTestMetaID int          `json:"pretest_meta_id"`
	GlobalID      string       `json:"global_id"`
	Question      models.JSONB `json:"question" gorm:"type:jsonb"`
	IsPublished   *bool        `json:"is_published"`
	CreatedBy     string       `json:"created_by"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedBy     string       `json:"updated_by"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

func PreTestDataCreate(ctx *gin.Context) {
	var body PreTestDataBody

	ctx.Bind(&body)

	post := models.PreTestData{
		ModuleID:      body.ModuleID,
		PreTestMetaID: body.PreTestMetaID,
		GlobalID:      body.GlobalID,
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
	var preTestData models.PreTestData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Preload("Metadata").First(&preTestData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.Preload("Metadata").First(&preTestData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Data not found.",
		})
		return
	}

	// fmt.Println(preTestData)

	ctx.JSON(http.StatusOK, gin.H{"data": preTestData})
}
func PreTestDataFindAll(ctx *gin.Context) {
	var preTestData []models.PreTestData
	result := initializers.DB.Preload("Metadata").Find(&preTestData)

	// fmt.Println(preTestData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": preTestData})
}

func PreTestDataUpdate(ctx *gin.Context) {
	var body PreTestDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PreTestData{
		ModuleID:      body.ModuleID,
		PreTestMetaID: body.PreTestMetaID,
		GlobalID:      body.GlobalID,
		Question:      body.Question,
		IsPublished:   body.IsPublished,
		UpdatedBy:     body.UpdatedBy,
		UpdatedAt:     body.UpdatedAt,
	}

	var current models.PreTestData
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

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResultAfterUpdate = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pretest Data updated successfully.", "data": &current})
}
func PreTestDataUpsert(ctx *gin.Context) {
	var body PreTestDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.PreTestData
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
			upsert := models.PreTestData{
				GlobalID:      body.GlobalID,
				ModuleID:      body.ModuleID,
				PreTestMetaID: body.PreTestMetaID,
				Question:      body.Question,
				IsPublished:   body.IsPublished,
				CreatedBy:     body.CreatedBy,
				CreatedAt:     body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PreTest Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PreTest Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.PreTestData{
				GlobalID:      id,
				ModuleID:      body.ModuleID,
				PreTestMetaID: body.PreTestMetaID,
				Question:      body.Question,
				IsPublished:   body.IsPublished,
				CreatedBy:     body.CreatedBy,
				CreatedAt:     body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PreTest Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PreTest Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.PreTestData{
			ID:            current.ID,
			GlobalID:      body.GlobalID,
			ModuleID:      body.ModuleID,
			PreTestMetaID: body.PreTestMetaID,
			Question:      body.Question,
			IsPublished:   body.IsPublished,
			CreatedBy:     current.CreatedBy,
			CreatedAt:     current.CreatedAt,
			UpdatedBy:     body.UpdatedBy,
			UpdatedAt:     body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating PreTest Data.",
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
				"message": "PreTest Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "PreTest Data updated successfully.", "data": &current})
	}
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Pretest Data deleted successfully.", "deletedData": &current})

}
