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

type PostTestResultDataBody struct {
	UserID           int              `json:"user_id"`
	GlobalID         string           `json:"global_id"`
	Score            float64          `json:"score"`
	Start            pgtype.Timestamp `json:"start"`
	End              pgtype.Timestamp `json:"end"`
	Duration         models.Duration  `json:"duration"`
	Answer           models.JSONB     `json:"answer" gorm:"type:jsonb"`
	QuestionAnswered models.JSONB     `json:"question_answer" gorm:"type:jsonb"`
	CreatedBy        string           `json:"created_by"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedBy        string           `json:"updated_by"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

func PostTestResultDataCreate(ctx *gin.Context) {
	var body PostTestResultDataBody

	ctx.Bind(&body)

	post := models.PostTestResultData{
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

	ctx.JSON(http.StatusOK, gin.H{"message": "PostTest ResultData created successfully.", "data": &post})
}

func PostTestResultDataFindById(ctx *gin.Context) {
	var postTestResultData models.PostTestResultData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&postTestResultData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&postTestResultData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest ResultData not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestResultData})
}
func PostTestResultDataFindAll(ctx *gin.Context) {
	var postTestResultData []models.PostTestResultData
	result := initializers.DB.Find(&postTestResultData)

	// fmt.Println(postTestResultData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postTestResultData})
}

func PostTestResultDataUpdate(ctx *gin.Context) {
	var body PostTestResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PostTestResultData{
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

	var current models.PostTestResultData
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
			"message": "PostTest ResultData not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating PostTest ResultData.",
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
			"message": "PostTest ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset ResultData updated successfully.", "data": &current})
}

func PostTestResultDataUpsert(ctx *gin.Context) {
	var body PostTestResultDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.PostTestResultData
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
			upsert := models.PostTestResultData{
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
					"message": "Error updating PostTest Result Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Result Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.PostTestResultData{
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
					"message": "Error updating PostTest Result Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Result Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.PostTestResultData{
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
				"message": "Error updating PostTest Result Data.",
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
				"message": "PostTest Result Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "PostTest Result Data updated successfully.", "data": &current})
	}
}

func PostTestResultDataDelete(ctx *gin.Context) {
	var current models.PostTestResultData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PostTest ResultData not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting PostTest ResultData.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset ResultData deleted successfully.", "deletedData": &current})

}
