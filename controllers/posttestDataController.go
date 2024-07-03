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

type PostTestDataBody struct {
	ModuleID       int          `json:"module_id"`
	PostTestMetaID int          `json:"posttest_meta_id"`
	GlobalID       string       `json:"global_id"`
	Question       models.JSONB `json:"question" gorm:"type:jsonb"`
	IsPublished    *bool        `json:"is_published"`
	CreatedBy      string       `json:"created_by"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedBy      string       `json:"updated_by"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

func PostTestDataCreate(ctx *gin.Context) {
	var body PostTestDataBody

	ctx.Bind(&body)

	post := models.PostTestData{
		ModuleID:       body.ModuleID,
		PostTestMetaID: body.PostTestMetaID,
		GlobalID:       body.GlobalID,
		Question:       body.Question,
		IsPublished:    body.IsPublished,
		CreatedBy:      body.CreatedBy,
		CreatedAt:      body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Data created successfully.", "data": &post})
}

func PostTestDataFindById(ctx *gin.Context) {
	var postTestData models.PostTestData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Preload("Metadata").First(&postTestData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.Preload("Metadata").First(&postTestData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestData})
}
func PostTestDataFindAll(ctx *gin.Context) {
	var postTestData []models.PostTestData
	result := initializers.DB.Preload("Metadata").Find(&postTestData)

	// fmt.Println(postTestData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestData})
}

func PostTestDataUpdate(ctx *gin.Context) {
	var body PostTestDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PostTestData{
		ModuleID:       body.ModuleID,
		PostTestMetaID: body.PostTestMetaID,
		GlobalID:       body.GlobalID,
		Question:       body.Question,
		IsPublished:    body.IsPublished,
		UpdatedBy:      body.UpdatedBy,
		UpdatedAt:      body.UpdatedAt,
	}

	var current models.PostTestData
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
			"message": "PostTest Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PostTest Data.",
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
			"message": "PostTest Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Posttest Data updated successfully.", "data": &current})
}

func PostTestDataUpsert(ctx *gin.Context) {
	var body PostTestDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.PostTestData
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
			upsert := models.PostTestData{
				GlobalID:       body.GlobalID,
				ModuleID:       body.ModuleID,
				PostTestMetaID: body.PostTestMetaID,
				Question:       body.Question,
				IsPublished:    body.IsPublished,
				CreatedBy:      body.CreatedBy,
				CreatedAt:      body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PostTest Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.PostTestData{
				GlobalID:       id,
				ModuleID:       body.ModuleID,
				PostTestMetaID: body.PostTestMetaID,
				Question:       body.Question,
				IsPublished:    body.IsPublished,
				CreatedBy:      body.CreatedBy,
				CreatedAt:      body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating PostTest Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.PostTestData{
			ID:             current.ID,
			GlobalID:       body.GlobalID,
			ModuleID:       body.ModuleID,
			PostTestMetaID: body.PostTestMetaID,
			Question:       body.Question,
			IsPublished:    body.IsPublished,
			UpdatedBy:      body.UpdatedBy,
			UpdatedAt:      body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating PostTest Data.",
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
				"message": "PostTest Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Data updated successfully.", "data": &current})
	}
}

func PostTestDataDelete(ctx *gin.Context) {
	var current models.PostTestData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PostTest Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Posttest Data deleted successfully.", "deletedData": &current})

}
