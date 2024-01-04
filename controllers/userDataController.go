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

type UserDataBody struct {
	EmployeeID        int       `json:"employee_id"`
	Name              string    `json:"name"`
	NIK               string    `json:"nik"`
	Level             string    `json:"level"`
	LevelID           int       `json:"level_id"`
	Grade             string    `json:"grade"`
	GradeID           int       `json:"grade_id"`
	Department        string    `json:"department"`
	DepartmentID      int       `json:"department_id"`
	LearningJourney   string    `json:"learning_journey"`
	LearningJourneyID int       `json:"learning_journey_id"`
	Role              string    `json:"role"`
	RoleID            int       `json:"role_id"`
	Status            string    `json:"status"`
	StatusID          int       `json:"status_id"`
	IsActive          *bool     `json:"is_active"`
	Position          string    `json:"position"`
	PositionID        int       `json:"position_id"`
	AlternativeID     string    `json:"alternative_id"`
	Password          string    `json:"password"`
	CreatedBy         string    `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedBy         string    `json:"updated_by"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func UserDataCreate(ctx *gin.Context) {
	var body UserDataBody

	ctx.Bind(&body)

	post := models.UserData{
		EmployeeID:        body.EmployeeID,
		Name:              body.Name,
		NIK:               body.NIK,
		Level:             body.Level,
		LevelID:           body.LevelID,
		Grade:             body.Grade,
		GradeID:           body.GradeID,
		Department:        body.Department,
		DepartmentID:      body.EmployeeID,
		LearningJourney:   body.LearningJourney,
		LearningJourneyID: body.LearningJourneyID,
		Role:              body.Role,
		RoleID:            body.RoleID,
		Status:            body.Status,
		StatusID:          body.StatusID,
		IsActive:          body.IsActive,
		Position:          body.Position,
		PositionID:        body.PositionID,
		AlternativeID:     body.AlternativeID,
		Password:          body.Password,
		CreatedBy:         body.CreatedBy,
		CreatedAt:         body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &post})
}

func UserDataFindById(ctx *gin.Context) {
	var userData models.UserData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&userData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&userData, "c_alternative_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}
func UserDataFindAll(ctx *gin.Context) {
	var userData []models.UserData
	result := initializers.DB.Find(&userData)

	// fmt.Println(userData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}

func UserDataUpdate(ctx *gin.Context) {
	var body UserDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.UserData{
		EmployeeID:        body.EmployeeID,
		Name:              body.Name,
		NIK:               body.NIK,
		Level:             body.Level,
		LevelID:           body.LevelID,
		Grade:             body.Grade,
		GradeID:           body.GradeID,
		Department:        body.Department,
		DepartmentID:      body.EmployeeID,
		LearningJourney:   body.LearningJourney,
		LearningJourneyID: body.LearningJourneyID,
		Role:              body.Role,
		RoleID:            body.RoleID,
		Status:            body.Status,
		StatusID:          body.StatusID,
		IsActive:          body.IsActive,
		Position:          body.Position,
		PositionID:        body.PositionID,
		AlternativeID:     body.AlternativeID,
		Password:          body.Password,
		UpdatedBy:         body.UpdatedBy,
		UpdatedAt:         body.UpdatedAt,
	}

	var current models.UserData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_alternative_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating User Data.",
		})
		return
	}

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResultAfterUpdate = initializers.DB.First(&current, "c_alternative_id = ?", id)
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User ResultData not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Data updated successfully.", "data": &current})
}

func UserDataUpsert(ctx *gin.Context) {
	var body UserDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.UserData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_alternative_id = ?", id)
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.UserData{
				// GlobalID:   body.GlobalID,
				EmployeeID:        body.EmployeeID,
				Name:              body.Name,
				NIK:               body.NIK,
				Level:             body.Level,
				LevelID:           body.LevelID,
				Grade:             body.Grade,
				GradeID:           body.GradeID,
				Department:        body.Department,
				DepartmentID:      body.EmployeeID,
				LearningJourney:   body.LearningJourney,
				LearningJourneyID: body.LearningJourneyID,
				Role:              body.Role,
				RoleID:            body.RoleID,
				Status:            body.Status,
				StatusID:          body.StatusID,
				IsActive:          body.IsActive,
				Position:          body.Position,
				PositionID:        body.PositionID,
				AlternativeID:     body.AlternativeID,
				Password:          body.Password,
				CreatedBy:         body.CreatedBy,
				CreatedAt:         body.CreatedAt,
				UpdatedBy:         body.UpdatedBy,
				UpdatedAt:         body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating User Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.UserData{
				EmployeeID:        body.EmployeeID,
				Name:              body.Name,
				NIK:               body.NIK,
				Level:             body.Level,
				LevelID:           body.LevelID,
				Grade:             body.Grade,
				GradeID:           body.GradeID,
				Department:        body.Department,
				DepartmentID:      body.EmployeeID,
				LearningJourney:   body.LearningJourney,
				LearningJourneyID: body.LearningJourneyID,
				Role:              body.Role,
				RoleID:            body.RoleID,
				Status:            body.Status,
				StatusID:          body.StatusID,
				IsActive:          body.IsActive,
				Position:          body.Position,
				PositionID:        body.PositionID,
				AlternativeID:     id,
				Password:          body.Password,
				CreatedBy:         body.CreatedBy,
				CreatedAt:         body.CreatedAt,
				UpdatedBy:         body.UpdatedBy,
				UpdatedAt:         body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating User Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.UserData{
			ID:                current.ID,
			EmployeeID:        body.EmployeeID,
			Name:              body.Name,
			NIK:               body.NIK,
			Level:             body.Level,
			LevelID:           body.LevelID,
			Grade:             body.Grade,
			GradeID:           body.GradeID,
			Department:        body.Department,
			DepartmentID:      body.EmployeeID,
			LearningJourney:   body.LearningJourney,
			LearningJourneyID: body.LearningJourneyID,
			Role:              body.Role,
			RoleID:            body.RoleID,
			Status:            body.Status,
			StatusID:          body.StatusID,
			IsActive:          body.IsActive,
			Position:          body.Position,
			PositionID:        body.PositionID,
			AlternativeID:     body.AlternativeID,
			Password:          body.Password,
			CreatedBy:         body.CreatedBy,
			CreatedAt:         body.CreatedAt,
			UpdatedBy:         body.UpdatedBy,
			UpdatedAt:         body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating User Data.",
			})
			return
		}

		if govalidator.IsNumeric(ctx.Param("id")) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
		} else {
			id := ctx.Param("id")
			findByIdResultAfterUpdate = initializers.DB.First(&current, "c_alternative_id = ?", id)
		}

		if findByIdResultAfterUpdate.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "User Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User Data updated successfully.", "data": &current})
	}
}

func UserDataDelete(ctx *gin.Context) {
	var current models.UserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting User Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Data deleted successfully.", "deletedData": &current})

}
