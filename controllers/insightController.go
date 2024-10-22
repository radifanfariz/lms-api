package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

func TotalInsightsDataFindAll(ctx *gin.Context) {
	var totalInsightsData models.TotalInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT COUNT(DISTINCT t_user_data.n_id) AS "n_total_users",
	(SELECT COUNT(DISTINCT t_module_data.n_id) FROM public.t_module_data) AS "n_total_learning_modules", 
	(SELECT COUNT(DISTINCT t_pretest_result_data.n_user_id) FROM public.t_pretest_result_data) AS "n_total_pretest_participants", 
	(SELECT COUNT(DISTINCT t_materi_result_data.n_user_id) FROM public.t_materi_result_data) AS "n_total_materi_participants", 
	(SELECT COUNT(DISTINCT t_posttest_result_data.n_user_id) FROM public.t_posttest_result_data) AS "n_total_posttest_participants",
	(SELECT COUNT(DISTINCT t_user_action_data.n_id) AS "n_total_enrolled_learning_modules" FROM t_user_action_data where t_user_action_data.n_module_accessed = 1),
	(SELECT COUNT(DISTINCT t_user_action_data.n_id) AS "n_total_graduated_learning_modules" FROM t_user_action_data where t_user_action_data.n_module_accessed = 1 AND t_user_action_data.n_pretest_accessed = 1 AND t_user_action_data.n_materi_accessed = 1 AND t_user_action_data.n_posttest_accessed = 1)
	FROM public.t_user_data
	`).Scan(&totalInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Total insight data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": totalInsightsData})
}

func TotalEnrolledPerMonthInsightsDataFindAll(ctx *gin.Context) {
	var totalEnrolledPerMonthInsightsData []models.TotalEnrolledPerMonthInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT TO_CHAR(t_user_action_data.d_updated_at, 'Mon-YY') AS "c_date",
	COUNT(t_user_action_data.n_id) AS "n_total_enrolled" FROM public.t_user_action_data 
	GROUP BY TO_CHAR(t_user_action_data.d_updated_at, 'Mon-YY')
	`).Scan(&totalEnrolledPerMonthInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Total enrolled per month data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": totalEnrolledPerMonthInsightsData})
}

func EnrolledInsightsDataFindAll(ctx *gin.Context) {
	var enrolledInsightsData []models.EnrolledInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT t_user_action_data.c_global_id, t_module_metadata.c_name AS "c_module_name", 
	COUNT(t_user_action_data.c_global_id) AS "n_total_enrolled" FROM public.t_user_action_data 
	JOIN public.t_module_metadata ON t_user_action_data.c_global_id = t_module_metadata.c_global_id 
	GROUP BY t_user_action_data.c_global_id, t_module_metadata.c_name 
	ORDER BY COUNT(t_user_action_data.c_global_id) DESC
	`).Scan(&enrolledInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Enrolled insight data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": enrolledInsightsData})
}

func UserInsightsDataFindAll(ctx *gin.Context) {
	var userInsightsData []models.UserInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT t_user_data.n_id, t_user_data.c_nik,t_user_data.c_name, t_user_data.c_learning_journey, t_module_metadata.c_name AS "c_module_name",
	t_pretest_result_data.n_score AS n_pretest_score, t_posttest_result_data.n_score AS "n_posttest_score", t_posttest_metadata.n_min_score AS "n_posttest_min_score" FROM public.t_user_action_data
	JOIN public.t_user_data ON t_user_data.n_id = t_user_action_data.n_user_id
	JOIN public.t_module_metadata ON t_module_metadata.c_global_id = t_user_action_data.c_global_id 
	JOIN public.t_pretest_result_data ON t_pretest_result_data.c_global_id = t_user_action_data.c_global_id
	JOIN public.t_materi_result_data ON t_materi_result_data.c_global_id = t_user_action_data.c_global_id
	JOIN public.t_posttest_result_data ON t_posttest_result_data.c_global_id = t_user_action_data.c_global_id 
	JOIN public.t_posttest_metadata ON t_posttest_metadata.c_global_id = t_user_action_data.c_global_id
	GROUP BY t_user_data.n_id, t_user_data.c_nik,t_user_data.c_name, t_user_data.c_learning_journey, t_module_metadata.c_name, t_pretest_result_data.n_score, t_posttest_result_data.n_score, t_posttest_metadata.n_min_score
	ORDER BY t_user_data.c_name
	`).Scan(&userInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User insight data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userInsightsData})
}
