package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type CategoryDataBody struct {
	Domain    string    `json:"domain"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Seq       int       `json:"seq"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CategoryDataCreate(ctx *gin.Context) {
	var body CategoryDataBody

	ctx.Bind(&body)

	post := models.CategoryData{
		Domain:    body.Domain,
		Label:     body.Label,
		Value:     body.Value,
		Seq:       body.Seq,
		CreatedBy: body.CreatedBy,
		CreatedAt: body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category Data created successfully.", "data": &post})
}

func CategoryDataFindById(ctx *gin.Context) {
	var CategoryData models.CategoryData

	var findByIdResult *gorm.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult = initializers.DB.First(&CategoryData, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": CategoryData})
}
func CategoryDataFindByDomain(ctx *gin.Context) {
	var CategoryData []models.CategoryData

	var findByIdResult *gorm.DB

	label := ctx.Param("domain")
	findByIdResult = initializers.DB.First(&CategoryData, "c_domain ILIKE ?", label)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": CategoryData})
}
func CategoryDataFindByLabel(ctx *gin.Context) {
	var CategoryData []models.CategoryData

	var findByIdResult *gorm.DB

	label := ctx.Param("label")
	findByIdResult = initializers.DB.First(&CategoryData, "c_label ILIKE ?", label)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": CategoryData})
}
func CategoryDataFindAll(ctx *gin.Context) {
	var CategoryData []models.CategoryData
	result := initializers.DB.Find(&CategoryData)

	// fmt.Println(CategoryData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": CategoryData})
}

func CategoryDataUpdate(ctx *gin.Context) {
	var body CategoryDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.CategoryData{
		Domain:    body.Domain,
		Label:     body.Label,
		Value:     body.Value,
		Seq:       body.Seq,
		UpdatedBy: body.UpdatedBy,
		UpdatedAt: body.UpdatedAt,
	}

	var current models.CategoryData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult = initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Category Data.",
		})
		return
	}

	findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category Data updated successfully.", "data": &current})
}

func CategoryDataUpsert(ctx *gin.Context) {
	var body CategoryDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.CategoryData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult = initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		upsert := models.CategoryData{
			Domain:    body.Domain,
			Label:     body.Label,
			Value:     body.Value,
			Seq:       body.Seq,
			CreatedBy: body.CreatedBy,
			CreatedAt: body.CreatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Category Data.",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Category Data created successfully.", "data": &upsert})
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.CategoryData{
			ID:        current.ID,
			Domain:    body.Domain,
			Label:     body.Label,
			Value:     body.Value,
			Seq:       body.Seq,
			CreatedBy: current.CreatedBy,
			CreatedAt: current.CreatedAt,
			UpdatedBy: body.UpdatedBy,
			UpdatedAt: body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Category Data.",
			})
			return
		}

		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))

		if findByIdResultAfterUpdate.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Category Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Category Data updated successfully.", "data": &current})
	}
}

func CategoryDataDelete(ctx *gin.Context) {
	var current models.CategoryData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Category Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Category Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category Data deleted successfully.", "deletedData": &current})

}
