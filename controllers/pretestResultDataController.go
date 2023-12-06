package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

type PretestResultData struct {
	UserID           int              `json:"UserID"`
	Score            float64          `json:"Score"`
	Start            pgtype.Timestamp `json:"Start"`
	End              pgtype.Timestamp `json:"End"`
	Duration         models.Duration  `json:"Duration"`
	Answer           models.JSONB     `json:"Answer" gorm:"type:jsonb"`
	QuestionAnswered models.JSONB     `json:"QuestionAnswered" gorm:"type:jsonb"`
	CreatedBy        string           `json:"CreatedBy"`
	CreatedAt        time.Time        `json:"CreatedAt"`
	UpdatedBy        string           `json:"UpdatedBy"`
	UpdatedAt        time.Time        `json:"UpdatedAt"`
}

func PreTestResultDataCreate(ctx *gin.Context) {
	var body PretestResultData

	ctx.Bind(&body)

	post := models.PreTestResultData{
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
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "PreTest ResultData created successfully.", "Data": &post})
}

func PreTestResultDataFindById(ctx *gin.Context) {
	var PreTestResultData models.PreTestResultData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&PreTestResultData, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Data": PreTestResultData})
}
func PreTestResultDataFindAll(ctx *gin.Context) {
	var PreTestResultData []models.PreTestResultData
	result := initializers.DB.Find(&PreTestResultData)

	// fmt.Println(PreTestResultData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Data": PreTestResultData})
}

func PreTestResultDataUpdate(ctx *gin.Context) {
	var body PretestResultData

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.PreTestResultData{
		UserID:           body.UserID,
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

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

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

	findByIdResultAfterUpdate := initializers.DB.First(&current, uint(id))

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "PreTest ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset ResultData updated successfully.", "Data": &current})
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Preteset ResultData deleted successfully.", "deletedData": &current})

}
