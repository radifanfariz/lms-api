package controllers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"github.com/radifanfariz/lms-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BodyDataSSO struct {
	Message string             `json:"message"`
	Status  string             `json:"status"`
	Data    models.UserDataSSO `json:"data"`
}

type BodyDataPortal struct {
	Message   string                `json:"message"`
	Status    string                `json:"status"`
	Data      models.UserDataPortal `json:"data"`
	Reff      string                `json:"reff"`
	UserToken *string               `json:"userToken"`
}

type UserDataBody struct {
	EmployeeID        int       `json:"employee_id"`
	Name              string    `json:"name"`
	NIK               string    `json:"nik"`
	MainCompany       string    `json:"main_company"`
	MainCompanyID     int       `json:"main_company_id"`
	Level             string    `json:"level"`
	LevelID           *int      `json:"level_id"`
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
	JoinDate          string    `json:"join_date"`
	AlternativeID     string    `json:"alternative_id"`
	Password          *string   `json:"password"`
	CreatedBy         string    `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedBy         string    `json:"updated_by"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func UserDataCreate(ctx *gin.Context) {
	var body UserDataBody

	ctx.Bind(&body)

	var joinDataValue *pgtype.Date = nil
	joinDateParsed, errJoinDateParsed := time.Parse("2006-01-02", body.JoinDate)
	if errJoinDateParsed != nil {
		log.Println("error parse joinDate time : ", errJoinDateParsed)
	}
	if !joinDateParsed.IsZero() {
		joinDataValue = &pgtype.Date{
			Time:  joinDateParsed,
			Valid: true,
		}
	}

	post := models.UserData{
		EmployeeID:        body.EmployeeID,
		Name:              body.Name,
		NIK:               body.NIK,
		MainCompany:       body.MainCompany,
		MainCompanyID:     body.MainCompanyID,
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
		JoinDate:          joinDataValue,
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

func UserDataLogin(ctx *gin.Context) {
	var userData models.UserData
	var credentials struct {
		NIK      string `json:"nik"`
		Password string `json:"password"`
	}

	ctx.Bind(&credentials)

	findByIdResult := initializers.DB.Where("c_nik = ? AND c_password = ?", credentials.NIK, credentials.Password).First(&userData)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid credentials !",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful !", "data": userData})
}

func UserDataLoginThroughSSO(ctx *gin.Context) {
	ssoUrl := os.Getenv("SSO_URL")
	ssoBasicAuthUsername := os.Getenv("SSO_BASIC_AUTH_USERNAME")
	ssoBasicAuthPassword := os.Getenv("SSO_BASIC_AUTH_PASSWORD")

	var userData models.UserData
	var credentials struct {
		NIK      string `json:"nik"`
		Password string `json:"password"`
	}

	ctx.Bind(&credentials)

	values := map[string]string{"nik": credentials.NIK, "password": credentials.Password}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	req, err := http.NewRequest(http.MethodPost, ssoUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(ssoBasicAuthUsername, ssoBasicAuthPassword))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	var jsonResult BodyDataSSO
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	// fmt.Println(string(responseBody))

	if strings.ToLower(jsonResult.Status) == "success" {
		findByIdResult := initializers.DB.Where("c_nik = ?", credentials.NIK).First(&userData)

		if findByIdResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid credentials !",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Login successful !", "data": userData})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": jsonResult.Message, "status": jsonResult.Status})

}
func UserDataLoginThroughPortal(ctx *gin.Context) {
	portalUrl := os.Getenv("PORTAL_URL")
	portalBasicAuthUsername := os.Getenv("PORTAL_BASIC_AUTH_USERNAME")
	portalBasicAuthPassword := os.Getenv("PORTAL_BASIC_AUTH_PASSWORD")

	var userData models.UserData
	var credentials struct {
		NIK      string `json:"nik"`
		Password string `json:"password"`
	}

	ctx.Bind(&credentials)

	values := map[string]string{"nik": credentials.NIK, "password": credentials.Password}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	req, err := http.NewRequest(http.MethodPost, portalUrl+"/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(portalBasicAuthUsername, portalBasicAuthPassword))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	var jsonResult BodyDataPortal
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	fmt.Println(string(responseBody))

	if strings.ToLower(jsonResult.Status) == "success" {
		findByIdResult := initializers.DB.Where("c_nik = ?", credentials.NIK).First(&userData)

		if findByIdResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid credentials !",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Login successful !", "data": userData, "dataFromPortal": jsonResult.Data})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": jsonResult.Message, "status": jsonResult.Status, "reff": jsonResult.Reff})

}
func UserDataLoginThroughPortalWithUserBu(ctx *gin.Context) {
	portalUrl := os.Getenv("PORTAL_URL")
	portalBasicAuthUsername := os.Getenv("PORTAL_BASIC_AUTH_USERNAME")
	portalBasicAuthPassword := os.Getenv("PORTAL_BASIC_AUTH_PASSWORD")

	var userData models.UserData
	var credentials struct {
		UserBU   *string `json:"userBu"`
		NIK      *string `json:"nik"`
		Password *string `json:"password"`
	}

	ctx.Bind(&credentials)

	values := map[string]*string{"userBu": credentials.UserBU, "nik": credentials.NIK, "password": credentials.Password}
	if credentials.UserBU == nil || *credentials.UserBU == "" {
		values = map[string]*string{"nik": credentials.NIK, "password": credentials.Password}
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	req, err := http.NewRequest(http.MethodPost, portalUrl+"/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(portalBasicAuthUsername, portalBasicAuthPassword))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong !"})
	}

	var jsonResult BodyDataPortal
	json.Unmarshal([]byte(responseBody), &jsonResult)

	// fmt.Println(resp.Status)
	// fmt.Println(jsonResult)
	fmt.Println(string(responseBody))

	if strings.ToLower(jsonResult.Status) == "success" {
		findByIdResult := initializers.DB.Where("n_employee_id = ?", jsonResult.Data.HRISID).First(&userData)

		if findByIdResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid credentials !",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Login successful !", "data": userData, "dataFromPortal": jsonResult.Data})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": jsonResult.Message, "status": jsonResult.Status, "reff": jsonResult.Reff, "userToken": &jsonResult.UserToken})

}

func UserDataFindById(ctx *gin.Context) {
	var userData models.UserData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&userData, "n_employee_id = ?", uint(id))
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
func UserDataFindByEmployeeIds(ctx *gin.Context) {
	var userData []models.UserData

	var findByIdResult *gorm.DB

	var body struct {
		EmployeeIDs []int `json:"employee_id"`
	}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findByIdResult = initializers.DB.Where("n_employee_id IN ?", body.EmployeeIDs).Find(&userData)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}

func UserDataFindByParams(ctx *gin.Context) {
	var userData []models.UserData

	var findByIdResult *gorm.DB

	companyId := ctx.Query("company_id")

	if companyId != "" {
		findByIdResult = initializers.DB.Where("n_main_company_id = ?", companyId).Find(&userData)
	} else {
		findByIdResult = initializers.DB.Find(&userData)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}

func UserDataFindByPaging(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort")
	filter := ctx.Query("filter")
	filterColumn := ctx.Query("filter_column")
	params := utils.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	var userData []models.UserData
	/* this to get total all data (total all rows) and total pages in pagination (without limit) */
	if filter != "" && filterColumn != "" {

		if filterColumn == "d_join_date" {
			filterColumn = "to_char(" + filterColumn + ", 'YYYY-MM-DD')"
		}
		findUserByParams := initializers.DB.Model(userData).Scopes(utils.Paginate(userData, &params, initializers.DB)).Where(filterColumn+" ILIKE ?", "%"+filter+"%").Find(&userData)

		if findUserByParams.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "User Data not found.",
			})
			return
		}
		var userDataAll []models.UserData
		totalRows := initializers.DB.Where(filterColumn+" ILIKE ?", "%"+filter+"%").Find(&userDataAll).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	} else {
		findUserByParams := initializers.DB.Model(userData).Scopes(utils.Paginate(userData, &params, initializers.DB)).Find(&userData)

		if findUserByParams.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "User Data not found.",
			})
			return
		}
		var userDataAll []models.UserData
		totalRows := initializers.DB.Find(&userDataAll).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	}
	/*------------------------------------------------------------------------------*/

	params.Data = userData

	ctx.JSON(http.StatusOK, params)
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

	var joinDataValue *pgtype.Date = nil
	joinDateParsed, errJoinDateParsed := time.Parse("2006-01-02", body.JoinDate)
	if errJoinDateParsed != nil {
		log.Println("error parse joinDate time : ", errJoinDateParsed)
	}
	if !joinDateParsed.IsZero() {
		joinDataValue = &pgtype.Date{
			Time:  joinDateParsed,
			Valid: true,
		}
	}
	if !joinDateParsed.IsZero() {
		joinDataValue = &pgtype.Date{
			Time:  joinDateParsed,
			Valid: true,
		}
	}

	updates := models.UserData{
		EmployeeID:        body.EmployeeID,
		Name:              body.Name,
		NIK:               body.NIK,
		MainCompany:       body.MainCompany,
		MainCompanyID:     body.MainCompanyID,
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
		JoinDate:          joinDataValue,
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
		findByIdResult = initializers.DB.First(&current, "n_employee_id = ?", uint(id))
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
		findByIdResultAfterUpdate = initializers.DB.First(&current, "n_employee_id = ?", uint(id))
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
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uniqueConstraintsName string = "tuserdata_ix1"

	var joinDataValue *pgtype.Date = nil
	joinDateParsed, errJoinDateParsed := time.Parse("2006-01-02", body.JoinDate)
	if errJoinDateParsed != nil {
		log.Println("error parse joinDate time : ", errJoinDateParsed)
	}
	if !joinDateParsed.IsZero() {
		joinDataValue = &pgtype.Date{
			Time:  joinDateParsed,
			Valid: true,
		}
	}

	var current models.UserData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, "n_employee_id = ?", uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_alternative_id = ?", id)
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then alternative_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.UserData{
				// GlobalID:   body.GlobalID,
				EmployeeID:        body.EmployeeID,
				Name:              body.Name,
				NIK:               body.NIK,
				MainCompany:       body.MainCompany,
				MainCompanyID:     body.MainCompanyID,
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
				JoinDate:          joinDataValue,
				AlternativeID:     body.AlternativeID,
				Password:          body.Password,
				CreatedBy:         body.CreatedBy,
				CreatedAt:         body.CreatedAt,
				UpdatedBy:         body.UpdatedBy,
				UpdatedAt:         body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Clauses(clause.OnConflict{
				OnConstraint: uniqueConstraintsName,
				DoUpdates: clause.AssignmentColumns([]string{
					"c_name",
					"n_level_id",
					"c_level",
					"n_grade_id",
					"c_grade",
					"n_department_id",
					"c_department",
					"n_position_id",
					"c_position",
					"n_main_company_id",
					"c_main_company",
					"d_join_date",
					"c_created_by",
					"c_updated_by",
					"d_created_at",
					"d_updated_at",
				}),
			}).Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating User Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is alternative_id then alternative_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.UserData{
				EmployeeID:        body.EmployeeID,
				Name:              body.Name,
				NIK:               body.NIK,
				MainCompany:       body.MainCompany,
				MainCompanyID:     body.MainCompanyID,
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
				JoinDate:          joinDataValue,
				AlternativeID:     id,
				Password:          body.Password,
				CreatedBy:         body.CreatedBy,
				CreatedAt:         body.CreatedAt,
				UpdatedBy:         body.UpdatedBy,
				UpdatedAt:         body.UpdatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Clauses(clause.OnConflict{
				OnConstraint: uniqueConstraintsName,
				DoUpdates: clause.AssignmentColumns([]string{
					"c_name",
					"n_level_id",
					"c_level",
					"n_grade_id",
					"c_grade",
					"n_department_id",
					"c_department",
					"n_position_id",
					"c_position",
					"n_main_company_id",
					"c_main_company",
					"d_join_date",
					"c_created_by",
					"c_updated_by",
					"d_created_at",
					"d_updated_at",
				}),
			}).Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating User Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update alternative_id, so dont need to provide alternative_id in JSON Body req */
		upsert := models.UserData{
			ID:                current.ID,
			EmployeeID:        body.EmployeeID,
			Name:              body.Name,
			NIK:               body.NIK,
			MainCompany:       body.MainCompany,
			MainCompanyID:     body.MainCompanyID,
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
			JoinDate: &pgtype.Date{
				Time:  joinDateParsed,
				Valid: !joinDateParsed.IsZero(),
			},
			AlternativeID: body.AlternativeID,
			Password:      body.Password,
			CreatedBy:     current.CreatedBy,
			CreatedAt:     current.CreatedAt,
			UpdatedBy:     body.UpdatedBy,
			UpdatedAt:     body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Clauses(clause.OnConflict{
			OnConstraint: uniqueConstraintsName,
			DoUpdates: clause.AssignmentColumns([]string{
				"c_name",
				"n_level_id",
				"c_level",
				"n_grade_id",
				"c_grade",
				"n_department_id",
				"c_department",
				"n_position_id",
				"c_position",
				"n_main_company_id",
				"c_main_company",
				"d_join_date",
				"c_created_by",
				"c_updated_by",
				"d_created_at",
				"d_updated_at",
			}),
		}).Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating User Data.",
			})
			return
		}

		if govalidator.IsNumeric(ctx.Param("id")) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			findByIdResultAfterUpdate = initializers.DB.First(&current, "n_employee_id = ?", uint(id))
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

func UserDataBulkUpsert(ctx *gin.Context) {
	/* before use this function, make sure to add unique constraint in the database (for upsert purpose) */
	/*-------------------make request to another api-----------------------------*/
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	var limit string = "99999999"
	if ctx.Query("limit") != "" {
		limit = url.QueryEscape(ctx.Query("limit"))
	}
	var employeeStatus string = "active"
	if ctx.Query("employeeStatus") != "" {
		employeeStatus = url.QueryEscape(ctx.Query("employeeStatus"))
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/employee?limit="+limit+"&employeeStatus="+employeeStatus, nil)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(hrisBasicAuthUsername, hrisBasicAuthPassword))
	req.Header.Add("key-access", hrisKeyAccess)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	var employeeJsonResult models.EmployeeDataFromHris
	json.Unmarshal([]byte(responseBody), &employeeJsonResult)
	// fmt.Println("employee data: ", employeeJsonResult)
	/*-----------------------------------------------------------------------------*/
	/*
		using struct []UserDataBody has problem
		ERROR: relation "user_data_bodies" does not exist (SQLSTATE 42P01)
		because tablename have not been configured (reference in models/configTable.go)
		so instead using []models.UserData
	*/
	/*----------------------------------------do upsert------------------*/

	/* whether need manual input (for now the input is from another api) */
	// var body []models.UserData

	// if err := ctx.ShouldBind(&body); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	/*----------------------*/

	var uniqueConstraintsName string = "tuserdata_ix1"
	var createdBy string
	if ctx.Query("createdBy") != "" {
		createdBy = ctx.Query("createdBy")
	}
	var createdAt string
	if ctx.Query("createdAt") != "" {
		createdAt = ctx.Query("createdAt")
	}
	var updatedBy string
	if ctx.Query("updatedBy") != "" {
		updatedBy = ctx.Query("updatedBy")
	}
	var updatedAt string
	if ctx.Query("updatedAt") != "" {
		updatedAt = ctx.Query("updatedAt")
	}

	/*------------------formatting data from HRIS api----------------*/
	// Remove duplicates based on the ID
	uniqueEmployeeData := utils.RemoveDuplicates(employeeJsonResult.Data, "EmployeeID").([]models.EmployeeData)
	var formattedEmployeeData []models.UserData
	for _, v := range uniqueEmployeeData {
		employeeIdInt, _ := strconv.Atoi(v.EmployeeID)
		mainCompanyIdInt, _ := strconv.Atoi(v.MainCOmpanyID)
		positionIdInt, _ := strconv.Atoi(v.PositionID)
		gradeIdInt, _ := strconv.Atoi(v.GradeID)
		pangkatIdInt, _ := strconv.Atoi(v.PangkatID)
		departmentIdInt, _ := strconv.Atoi(v.DepartmentID)
		isActiveBool := models.UserData{IsActive: func() *bool { b := true; return &b }()}
		joinDateParsed, errJoinDateParsed := time.Parse("2006-01-02", v.JoinDate)
		if errJoinDateParsed != nil {
			log.Println("error parse joinDate time : ", errJoinDateParsed)
		}
		createdAtTime, errCreatedAtTime := time.Parse("2006-01-02T15:04:05.999 07:00", createdAt)
		if errCreatedAtTime != nil && createdAt != "" {
			log.Println("error parse createdAt time : ", errCreatedAtTime)
		}
		updatedAtTime, errUpdatedAtTime := time.Parse("2006-01-02T15:04:05.999 07:00", updatedAt)
		if errUpdatedAtTime != nil && updatedAt != "" {
			log.Println("error parse updatedAt time : ", errUpdatedAtTime)
		}
		uuidString := func() (uuid string) {

			b := make([]byte, 16)
			_, err := rand.Read(b)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

			return
		}

		formattedEmployeeData = append(formattedEmployeeData, models.UserData{
			ID:                nil,
			EmployeeID:        employeeIdInt,
			Name:              v.EmployeeName,
			NIK:               v.EmployeeNik,
			MainCompany:       v.MainCompanyName,
			MainCompanyID:     mainCompanyIdInt,
			Level:             v.PangkatName,
			LevelID:           &pangkatIdInt,
			Position:          v.EmployeePosition,
			PositionID:        positionIdInt,
			Grade:             v.GradeName,
			GradeID:           gradeIdInt,
			Department:        v.DepartmentName,
			DepartmentID:      departmentIdInt,
			LearningJourney:   "foundational",
			LearningJourneyID: 1,
			Role:              "user",
			RoleID:            3,
			Status:            "active",
			StatusID:          1,
			IsActive:          isActiveBool.IsActive,
			JoinDate: &pgtype.Date{
				Time:  joinDateParsed,
				Valid: !joinDateParsed.IsZero(),
			},
			CreatedBy:     createdBy,
			CreatedAt:     createdAtTime,
			UpdatedBy:     updatedBy,
			UpdatedAt:     updatedAtTime,
			AlternativeID: uuidString(),
		})
	}

	/*----------------------------------------------------------------*/

	bulkUpsertResult := initializers.DB.Clauses(clause.OnConflict{
		OnConstraint: uniqueConstraintsName,
		DoUpdates: clause.AssignmentColumns([]string{
			"n_employee_id",
			"c_name",
			"n_level_id",
			"c_level",
			"n_grade_id",
			"c_grade",
			"n_department_id",
			"c_department",
			"n_position_id",
			"c_position",
			"n_main_company_id",
			"c_main_company",
			"d_join_date",
			"c_created_by",
			"c_updated_by",
			"d_created_at",
			"d_updated_at",
		}),
	}).CreateInBatches(&formattedEmployeeData, 100)

	if bulkUpsertResult.Error != nil {
		log.Println(bulkUpsertResult.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    true,
			"errorMsg": bulkUpsertResult.Error.Error(),
			"message":  "Error upserting bulk User Data.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success upserting bulk User Data.",
	})
	/*-------------------------------------------------------------------------------------------------*/
}
func UserDataBulkUpsertResetAll(ctx *gin.Context) {
	/* before use this function, make sure to add unique constraint in the database (for upsert purpose) */
	/*-------------------make request to another api-----------------------------*/
	hrisUrl := os.Getenv("HRIS_URL")
	hrisBasicAuthUsername := os.Getenv("HRIS_BASIC_AUTH_USERNAME")
	hrisBasicAuthPassword := os.Getenv("HRIS_BASIC_AUTH_PASSWORD")
	hrisKeyAccess := os.Getenv("HRIS_KEY_ACCESS")

	var limit string = "99999999"
	if ctx.Query("limit") != "" {
		limit = url.QueryEscape(ctx.Query("limit"))
	}
	var employeeStatus string = "active"
	if ctx.Query("employeeStatus") != "" {
		employeeStatus = url.QueryEscape(ctx.Query("employeeStatus"))
	}

	req, err := http.NewRequest(http.MethodGet, hrisUrl+"/smartmulia/employee?limit="+limit+"&employeeStatus="+employeeStatus, nil)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	// // appending to existing query args
	// q := req.URL.Query()
	// q.Add("foo", "bar")

	// // assign encoded query string to http request
	// req.URL.RawQuery = q.Encode()

	client := &http.Client{
		CheckRedirect: utils.RedirectPolicyFunc,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	req.Header.Add("Authorization", "Basic "+utils.BasicAuth(hrisBasicAuthUsername, hrisBasicAuthPassword))
	req.Header.Add("key-access", hrisKeyAccess)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong !",
		})
		return
	}

	var employeeJsonResult models.EmployeeDataFromHris
	json.Unmarshal([]byte(responseBody), &employeeJsonResult)
	// fmt.Println("employee data: ", employeeJsonResult)
	/*-----------------------------------------------------------------------------*/
	/*
		using struct []UserDataBody has problem
		ERROR: relation "user_data_bodies" does not exist (SQLSTATE 42P01)
		because tablename have not been configured (reference in models/configTable.go)
		so instead using []models.UserData
	*/
	/*----------------------------------------do upsert------------------*/

	/* whether need manual input (for now the input is from another api) */
	// var body []models.UserData

	// if err := ctx.ShouldBind(&body); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	/*----------------------*/

	var uniqueConstraintsName string = "tuserdata_ix1"
	var createdBy string
	if ctx.Query("createdBy") != "" {
		createdBy = ctx.Query("createdBy")
	}
	var createdAt string
	if ctx.Query("createdAt") != "" {
		createdAt = ctx.Query("createdAt")
	}
	var updatedBy string
	if ctx.Query("updatedBy") != "" {
		updatedBy = ctx.Query("updatedBy")
	}
	var updatedAt string
	if ctx.Query("updatedAt") != "" {
		updatedAt = ctx.Query("updatedAt")
	}

	/*------------------formatting data from HRIS api----------------*/
	// Remove duplicates based on the ID
	uniqueEmployeeData := utils.RemoveDuplicates(employeeJsonResult.Data, "EmployeeID").([]models.EmployeeData)
	var formattedEmployeeData []models.UserData
	for _, v := range uniqueEmployeeData {
		employeeIdInt, _ := strconv.Atoi(v.EmployeeID)
		mainCompanyIdInt, _ := strconv.Atoi(v.MainCOmpanyID)
		positionIdInt, _ := strconv.Atoi(v.PositionID)
		gradeIdInt, _ := strconv.Atoi(v.GradeID)
		pangkatIdInt, _ := strconv.Atoi(v.PangkatID)
		departmentIdInt, _ := strconv.Atoi(v.DepartmentID)
		isActiveBool := models.UserData{IsActive: func() *bool { b := true; return &b }()}
		joinDateParsed, errJoinDateParsed := time.Parse("2006-01-02", v.JoinDate)
		if errJoinDateParsed != nil {
			log.Println("error parse joinDate time : ", errJoinDateParsed)
		}
		createdAtTime, errCreatedAtTime := time.Parse("2006-01-02T15:04:05.999 07:00", createdAt)
		if errCreatedAtTime != nil && createdAt != "" {
			log.Println("error parse createdAt time : ", errCreatedAtTime)
		}
		updatedAtTime, errUpdatedAtTime := time.Parse("2006-01-02T15:04:05.999 07:00", updatedAt)
		if errUpdatedAtTime != nil && updatedAt != "" {
			log.Println("error parse updatedAt time : ", errUpdatedAtTime)
		}
		uuidString := func() (uuid string) {

			b := make([]byte, 16)
			_, err := rand.Read(b)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

			return
		}

		formattedEmployeeData = append(formattedEmployeeData, models.UserData{
			ID:                nil,
			EmployeeID:        employeeIdInt,
			Name:              v.EmployeeName,
			NIK:               v.EmployeeNik,
			MainCompany:       v.MainCompanyName,
			MainCompanyID:     mainCompanyIdInt,
			Level:             v.PangkatName,
			LevelID:           &pangkatIdInt,
			Position:          v.EmployeePosition,
			PositionID:        positionIdInt,
			Grade:             v.GradeName,
			GradeID:           gradeIdInt,
			Department:        v.DepartmentName,
			DepartmentID:      departmentIdInt,
			LearningJourney:   "foundational",
			LearningJourneyID: 1,
			Role:              "user",
			RoleID:            3,
			Status:            "active",
			StatusID:          1,
			IsActive:          isActiveBool.IsActive,
			JoinDate: &pgtype.Date{
				Time:  joinDateParsed,
				Valid: !joinDateParsed.IsZero(),
			},
			CreatedBy:     createdBy,
			CreatedAt:     createdAtTime,
			UpdatedBy:     updatedBy,
			UpdatedAt:     updatedAtTime,
			AlternativeID: uuidString(),
		})
	}

	/*----------------------------------------------------------------*/

	bulkUpsertResult := initializers.DB.Clauses(clause.OnConflict{
		OnConstraint: uniqueConstraintsName,
	}).CreateInBatches(&formattedEmployeeData, 100)

	if bulkUpsertResult.Error != nil {
		log.Println(bulkUpsertResult.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    true,
			"errorMsg": bulkUpsertResult.Error.Error(),
			"message":  "Error upserting bulk User Data.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success upserting bulk User Data.",
	})
	/*-------------------------------------------------------------------------------------------------*/
}

func UserDataDelete(ctx *gin.Context) {
	var current models.UserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, "n_employee_id = ?", uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, "n_employee_id = ?", uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting User Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Data deleted successfully.", "deletedData": &current})

}
