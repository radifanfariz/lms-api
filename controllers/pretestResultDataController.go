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

type PreTestResultDataBody struct {
	UserID           int              `json:"user_id"`
	GlobalID         string           `json:"global_id"`
	Score            float64          `json:"score"`
	Start            pgtype.Timestamp `json:"start"`
	End              pgtype.Timestamp `json:"end"`
	Duration         models.Duration  `json:"duration"`
	Answer           models.JSONB     `json:"answer" gorm:"type:jsonb"`
	QuestionAnswered models.JSONB     `json:"question_answered" gorm:"type:jsonb"`
	CreatedBy        string           `json:"created_by"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedBy        string           `json:"updated_by"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

func PreTestResultDataCreate(ctx *gin.Context) {
	var body PreTestResultDataBody

	ctx.Bind(&body)

	post := models.PreTestResultData{
		UserID:           body.UserID,
		GlobalID:         body.GlobalID,
		Score:            body.Score,
		Start:            body.Start,
		End:              body.End,
		Duration:         body.Duration,
		Answer:           body.Answer,
		QuestionAnswered: body.QuestionAnswered,
		CreatedBy:        body.CreatedBy,
		CreatedAt:        body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "PreTest ResultData created successfully.", "data": &post})
}

func PreTestResultDataFindById(ctx *gin.Context) {
	var preTestResultData []models.PreTestResultData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Find(&preTestResultData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.Find(&preTestResultData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": preTestResultData})
}
func PreTestResultDataFindByIdAndUserId(ctx *gin.Context) {
	var preTestResultData []models.PreTestResultData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).Find(&preTestResultData, uint(id))
	} else {
		id := ctx.Param("id")
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).Find(&preTestResultData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": preTestResultData})
}
func PreTestResultDataFindAll(ctx *gin.Context) {
	var preTestResultData []models.PreTestResultData
	result := initializers.DB.Find(&preTestResultData)

	// fmt.Println(preTestResultData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": preTestResultData})
}

func PreTestResultDataUpdate(ctx *gin.Context) {
	var body PreTestResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PreTestResultData{
		UserID:           body.UserID,
		GlobalID:         body.GlobalID,
		Score:            body.Score,
		Start:            body.Start,
		End:              body.End,
		Duration:         body.Duration,
		Answer:           body.Answer,
		QuestionAnswered: body.QuestionAnswered,
		UpdatedBy:        body.CreatedBy,
		UpdatedAt:        body.CreatedAt,
	}

	var current models.PreTestResultData
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
			"message": "PreTest ResultData not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PreTest ResultData.",
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
			"message": "PreTest ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pretest ResultData updated successfully.", "data": &current})
}

func PreTestResultDataUpsert(ctx *gin.Context) {
	var body PreTestResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.PreTestResultData
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
			upsert := models.PreTestResultData{
				GlobalID:         body.GlobalID,
				UserID:           body.UserID,
				Score:            body.Score,
				Start:            body.Start,
				End:              body.End,
				Duration:         body.Duration,
				Answer:           body.Answer,
				QuestionAnswered: body.QuestionAnswered,
				CreatedBy:        body.CreatedBy,
				CreatedAt:        body.CreatedAt,
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
			upsert := models.PreTestResultData{
				GlobalID:         id,
				UserID:           body.UserID,
				Score:            body.Score,
				Start:            body.Start,
				End:              body.End,
				Duration:         body.Duration,
				Answer:           body.Answer,
				QuestionAnswered: body.QuestionAnswered,
				CreatedBy:        body.CreatedBy,
				CreatedAt:        body.CreatedAt,
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
		upsert := models.PreTestResultData{
			ID:               current.ID,
			GlobalID:         body.GlobalID,
			UserID:           body.UserID,
			Score:            body.Score,
			Start:            body.Start,
			End:              body.End,
			Duration:         body.Duration,
			Answer:           body.Answer,
			QuestionAnswered: body.QuestionAnswered,
			UpdatedBy:        body.UpdatedBy,
			UpdatedAt:        body.UpdatedAt,
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

func PreTestResultDataDelete(ctx *gin.Context) {
	var current models.PreTestResultData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PreTest ResultData.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pretest ResultData deleted successfully.", "deletedData": &current})

}
