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

type UserActionDataBody struct {
	UserID           int       `json:"user_id"`
	GlobalID         string    `json:"global_id"`
	IsStartCourse    *bool     `json:"is_startcourse"`
	ModuleAccessed   int       `json:"module_accessed"`
	PretestAccessed  int       `json:"pretest_accessed"`
	MateriAccessed   int       `json:"materi_accessed"`
	PosttestAccessed int       `json:"posttest_accessed"`
	CreatedBy        string    `json:"created_by"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedBy        string    `json:"updated_by"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func UserActionDataCreate(ctx *gin.Context) {
	var body UserActionDataBody

	ctx.Bind(&body)

	post := models.UserActionData{
		UserID:           body.UserID,
		GlobalID:         body.GlobalID,
		IsStartCourse:    body.IsStartCourse,
		ModuleAccessed:   body.ModuleAccessed,
		PretestAccessed:  body.PretestAccessed,
		MateriAccessed:   body.MateriAccessed,
		PosttestAccessed: body.PosttestAccessed,
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

	ctx.JSON(http.StatusOK, gin.H{"message": "User Action Data created successfully.", "data": &post})
}

func UserActionDataFindById(ctx *gin.Context) {
	var UserActionData models.UserActionData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&UserActionData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&UserActionData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Action Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UserActionData})
}
func UserActionDataFindByIdAndUserId(ctx *gin.Context) {
	var UserActionData models.UserActionData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).First(&UserActionData, uint(id))
	} else {
		id := ctx.Param("id")
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).First(&UserActionData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Action Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UserActionData})
}
func UserActionDataFindAll(ctx *gin.Context) {
	var UserActionData []models.UserActionData
	result := initializers.DB.Find(&UserActionData)

	// fmt.Println(UserActionData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UserActionData})
}

func UserActionDataUpdate(ctx *gin.Context) {
	var body UserActionDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.UserActionData{
		UserID:           body.UserID,
		GlobalID:         body.GlobalID,
		IsStartCourse:    body.IsStartCourse,
		ModuleAccessed:   body.ModuleAccessed,
		PretestAccessed:  body.PretestAccessed,
		MateriAccessed:   body.MateriAccessed,
		PosttestAccessed: body.PosttestAccessed,
		UpdatedBy:        body.CreatedBy,
		UpdatedAt:        body.CreatedAt,
	}

	var current models.UserActionData
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
			"message": "User Action Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating User Action Data.",
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
			"message": "User Action Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Action Data updated successfully.", "data": &current})
}

func UserActionDataUpsert(ctx *gin.Context) {
	var body UserActionDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.UserActionData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.UserActionData{
				GlobalID:         body.GlobalID,
				UserID:           body.UserID,
				IsStartCourse:    body.IsStartCourse,
				ModuleAccessed:   body.ModuleAccessed,
				PretestAccessed:  body.PretestAccessed,
				MateriAccessed:   body.MateriAccessed,
				PosttestAccessed: body.PosttestAccessed,
				CreatedBy:        body.CreatedBy,
				CreatedAt:        body.CreatedAt,
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
			upsert := models.UserActionData{
				GlobalID:         id,
				UserID:           body.UserID,
				IsStartCourse:    body.IsStartCourse,
				ModuleAccessed:   body.ModuleAccessed,
				PretestAccessed:  body.PretestAccessed,
				MateriAccessed:   body.MateriAccessed,
				PosttestAccessed: body.PosttestAccessed,
				CreatedBy:        body.CreatedBy,
				CreatedAt:        body.CreatedAt,
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
		upsert := models.UserActionData{
			ID:               current.ID,
			GlobalID:         body.GlobalID,
			UserID:           body.UserID,
			IsStartCourse:    body.IsStartCourse,
			ModuleAccessed:   body.ModuleAccessed,
			PretestAccessed:  body.PretestAccessed,
			MateriAccessed:   body.MateriAccessed,
			PosttestAccessed: body.PosttestAccessed,
			UpdatedBy:        body.UpdatedBy,
			UpdatedAt:        body.UpdatedAt,
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

func UserActionDataDelete(ctx *gin.Context) {
	var current models.UserActionData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Action Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting User Action Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Action Data deleted successfully.", "deletedData": &current})

}
