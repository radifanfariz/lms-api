package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

func TotalInsightsDataFindAll(ctx *gin.Context) {

	/* this query string just affected to n_total_enrolled_learning_modules and n_total_graduated_learning_modules */
	minModule := ctx.DefaultQuery("min_module", "1")

	var totalInsightsData models.TotalInsightsData

	/* if min. learning module is n */
	findAllResult := initializers.DB.Raw(`
			WITH user_action_counts AS (
			SELECT 
				n_user_id,
				COUNT(n_user_id) AS action_count
			FROM 
				t_user_action_data
			WHERE 
				n_module_accessed = 1
			GROUP BY 
				n_user_id
		),
		graduated_counts AS (
			SELECT 
				n_user_id,
				COUNT(n_user_id) AS action_count
			FROM 
				t_user_action_data
			WHERE 
				n_module_accessed = 1 
				AND n_pretest_accessed = 1 
				AND n_materi_accessed = 1 
				AND n_posttest_accessed = 1
			GROUP BY 
				n_user_id
		)

		SELECT 
			COUNT(DISTINCT t_user_data.n_id) AS "n_total_users",
			(SELECT COUNT(DISTINCT t_module_data.n_id) FROM public.t_module_data) AS "n_total_learning_modules", 
			(SELECT COUNT(DISTINCT t_pretest_result_data.n_user_id) FROM public.t_pretest_result_data) AS "n_total_pretest_participants", 
			(SELECT COUNT(DISTINCT t_materi_result_data.n_user_id) FROM public.t_materi_result_data) AS "n_total_materi_participants", 
			(SELECT COUNT(DISTINCT t_posttest_result_data.n_user_id) FROM public.t_posttest_result_data) AS "n_total_posttest_participants",
			(SELECT COUNT(DISTINCT n_user_id) FROM user_action_counts WHERE action_count > ` + minModule + `) AS "n_total_enrolled_learning_modules",
			(SELECT COUNT(DISTINCT n_user_id) FROM graduated_counts WHERE action_count > ` + minModule + `) AS "n_total_graduated_learning_modules"
		FROM 
			public.t_user_data;
	`).Scan(&totalInsightsData)
	/*--------------------------------------*/

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Total insight data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": totalInsightsData})
}

func TotalUserPerMonthInsightsDataFindAll(ctx *gin.Context) {
	var totalUserPerMonthInsightsData []models.TotalUserPerMonthInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT TO_CHAR(t_user_action_data.d_updated_at, 'Mon-YY') AS "c_date",
       COUNT(DISTINCT t_user_action_data.n_user_id) AS "n_total_users"
	FROM public.t_user_action_data 
	GROUP BY TO_CHAR(t_user_action_data.d_updated_at, 'Mon-YY'),
         EXTRACT(YEAR FROM t_user_action_data.d_updated_at),
         EXTRACT(MONTH FROM t_user_action_data.d_updated_at)
	ORDER BY EXTRACT(YEAR FROM t_user_action_data.d_updated_at) ASC,
         EXTRACT(MONTH FROM t_user_action_data.d_updated_at) ASC;
	`).Scan(&totalUserPerMonthInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Total user per month data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": totalUserPerMonthInsightsData})
}

func ModuleInsightsDataFindAll(ctx *gin.Context) {
	var moduleInsightsData []models.ModuleInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT t_user_action_data.c_global_id, t_module_metadata.c_name AS "c_module_name", 
	COUNT(t_user_action_data.c_global_id) AS "n_total_users" FROM public.t_user_action_data 
	JOIN public.t_module_metadata ON t_user_action_data.c_global_id = t_module_metadata.c_global_id 
	GROUP BY t_user_action_data.c_global_id, t_module_metadata.c_name 
	ORDER BY COUNT(t_user_action_data.c_global_id) DESC
	`).Scan(&moduleInsightsData)

	if findAllResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module insight data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": moduleInsightsData})
}

func UserInsightsDataFindAll(ctx *gin.Context) {
	var userInsightsData []models.UserInsightsData

	findAllResult := initializers.DB.Raw(`
	SELECT t_user_data.n_id, t_user_data.c_nik,t_user_data.c_name, t_user_data.c_learning_journey, t_module_metadata.c_name as c_module_name, pretest_result.n_score as n_pretest_score, t_posttest_result_data.n_score as n_posttest_score , t_posttest_metadata.n_min_score as n_posttest_min_score from public.t_user_action_data 
	JOIN public.t_user_data ON t_user_data.n_id = t_user_action_data.n_user_id
	JOIN public.t_module_metadata ON t_module_metadata.c_global_id = t_user_action_data.c_global_id 
	JOIN (SELECT t_pretest_result_data.c_global_id,t_pretest_result_data.n_user_id, FIRST_VALUE(t_pretest_result_data.n_score) OVER (PARTITION BY t_pretest_result_data.c_global_id,t_pretest_result_data.n_user_id ORDER BY t_pretest_result_data.n_id DESC) AS n_score FROM public.t_pretest_result_data) pretest_result ON pretest_result.c_global_id = t_user_action_data.c_global_id AND pretest_result.n_user_id = t_user_action_data.n_user_id
	JOIN public.t_posttest_result_data ON t_posttest_result_data.c_global_id = t_user_action_data.c_global_id AND t_posttest_result_data.n_user_id = t_user_action_data.n_user_id
	JOIN public.t_posttest_metadata ON t_posttest_metadata.c_global_id = t_user_action_data.c_global_id
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
