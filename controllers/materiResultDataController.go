package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type MateriResultDataBody struct {
	UserID    int              `json:"user_id"`
	MateriID  int              `json:"materi_id"`
	GlobalID  string           `json:"global_id"`
	Start     pgtype.Timestamp `json:"start"`
	End       pgtype.Timestamp `json:"end"`
	Duration  models.Duration  `json:"duration"`
	CreatedBy string           `json:"created_by"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedBy string           `json:"updated_by"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func MateriResultDataCreate(ctx *gin.Context) {
	var body MateriResultDataBody

	ctx.Bind(&body)

	post := models.MateriResultData{
		UserID:    body.UserID,
		MateriID:  body.MateriID,
		GlobalID:  body.GlobalID,
		Start:     body.Start,
		End:       body.End,
		Duration:  body.Duration,
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi ResultData created successfully.", "data": &post})
}

func MateriResultDataFindById(ctx *gin.Context) {
	var MateriResultData models.MateriResultData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&MateriResultData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&MateriResultData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi ResultData not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": MateriResultData})
}
func MateriResultDataFindAll(ctx *gin.Context) {
	var MateriResultData []models.MateriResultData
	result := initializers.DB.Find(&MateriResultData)

	// fmt.Println(MateriResultData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": MateriResultData})
}

func MateriResultDataUpdate(ctx *gin.Context) {
	var body MateriResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.MateriResultData{
		UserID:    body.UserID,
		MateriID:  body.MateriID,
		GlobalID:  body.GlobalID,
		Start:     body.Start,
		End:       body.End,
		Duration:  body.Duration,
		UpdatedBy: body.CreatedBy,
		UpdatedAt: body.CreatedAt,
	}

	var current models.MateriResultData
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
			"message": "Materi ResultData not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Materi ResultData.",
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
			"message": "Materi ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi ResultData updated successfully.", "data": &current})
}

func MateriResultDataUpsert(ctx *gin.Context) {
	var body MateriResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.MateriResultData
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
			upsert := models.MateriResultData{
				GlobalID:  body.GlobalID,
				UserID:    body.UserID,
				MateriID:  body.MateriID,
				Start:     body.Start,
				End:       body.End,
				Duration:  body.Duration,
				CreatedBy: body.CreatedBy,
				CreatedAt: body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Materi Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Materi Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.MateriResultData{
				GlobalID:  id,
				UserID:    body.UserID,
				MateriID:  body.MateriID,
				Start:     body.Start,
				End:       body.End,
				Duration:  body.Duration,
				CreatedBy: body.CreatedBy,
				CreatedAt: body.CreatedAt,
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
		upsert := models.MateriResultData{
			ID:        current.ID,
			GlobalID:  body.GlobalID,
			MateriID:  body.MateriID,
			UserID:    body.UserID,
			Start:     body.Start,
			End:       body.End,
			Duration:  body.Duration,
			UpdatedBy: body.UpdatedBy,
			UpdatedAt: body.UpdatedAt,
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

func MateriResultDataDelete(ctx *gin.Context) {
	var current models.MateriResultData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Materi ResultData not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Materi ResultData.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Materi ResultData deleted successfully.", "deletedData": &current})

}
