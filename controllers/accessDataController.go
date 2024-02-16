package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type AccessDataBody struct {
	ModuleMetaID int       `json:"module_meta_id"`
	GlobalID     string    `json:"global_id"`
	ArrayGradeID []int64   `json:"array_grade_id"`
	ArrayUserID  []int64   `json:"array_user_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func AccessDataCreate(ctx *gin.Context) {
	var body AccessDataBody

	ctx.Bind(&body)

	fmt.Println(&body)

	post := models.AccessData{
		GlobalID:     body.GlobalID,
		ModuleMetaID: body.ModuleMetaID,
		ArrayGradeID: body.ArrayGradeID,
		ArrayUserID:  body.ArrayUserID,
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Access Data created successfully.", "data": &post})
}

func AccessDataFindById(ctx *gin.Context) {
	var accessData models.AccessData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&accessData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&accessData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Access Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": accessData})
}
func AccessDataFindAll(ctx *gin.Context) {
	var accessData []models.AccessData
	result := initializers.DB.Find(&accessData)

	// fmt.Println(accessData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": accessData})
}

func AccessDataUpdate(ctx *gin.Context) {
	var body AccessDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.AccessData{
		GlobalID:     body.GlobalID,
		ModuleMetaID: body.ModuleMetaID,
		ArrayGradeID: body.ArrayGradeID,
		ArrayUserID:  body.ArrayUserID,
		UpdatedBy:    body.UpdatedBy,
		UpdatedAt:    body.UpdatedAt,
	}

	var current models.AccessData
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
			"message": "Access Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Access Data.",
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
			"message": "Access ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data updated successfully.", "data": &current})
}

func AccessDataUpsert(ctx *gin.Context) {
	var body AccessDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.AccessData
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

	fmt.Println(body.GlobalID)

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.AccessData{
				GlobalID:     body.GlobalID,
				ModuleMetaID: body.ModuleMetaID,
				ArrayGradeID: body.ArrayGradeID,
				ArrayUserID:  body.ArrayUserID,
				CreatedBy:    body.CreatedBy,
				CreatedAt:    body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Access Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Access Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.AccessData{
				GlobalID:     id,
				ModuleMetaID: body.ModuleMetaID,
				ArrayGradeID: body.ArrayGradeID,
				ArrayUserID:  body.ArrayUserID,
				CreatedBy:    body.CreatedBy,
				CreatedAt:    body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Access Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Access Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.AccessData{
			ID:           current.ID,
			GlobalID:     body.GlobalID,
			ModuleMetaID: body.ModuleMetaID,
			ArrayGradeID: body.ArrayGradeID,
			ArrayUserID:  body.ArrayUserID,
			UpdatedBy:    body.UpdatedBy,
			UpdatedAt:    body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Access Data.",
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
				"message": "Access Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Access Data updated successfully.", "data": &current})
	}
}

func AccessDataDelete(ctx *gin.Context) {
	var current models.AccessData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Access Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Access Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset Data deleted successfully.", "deletedData": &current})

}
