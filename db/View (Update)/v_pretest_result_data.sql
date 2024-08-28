CREATE VIEW v_pretest_result_data (n_id, n_user_id, n_score, d_start, d_end, d_duration, j_answer, j_question_answered, c_created_by, c_updated_by, d_created_at, d_updated_at, c_global_id) AS  SELECT t_pretest_result_data.n_id,
    t_pretest_result_data.n_user_id,
    t_pretest_result_data.n_score,
    t_pretest_result_data.d_start,
    t_pretest_result_data.d_end,
    t_pretest_result_data.d_duration,
    t_pretest_result_data.j_answer,
    t_pretest_result_data.j_question_answered,
    t_pretest_result_data.c_created_by,
    t_pretest_result_data.c_updated_by,
    t_pretest_result_data.d_created_at,
    t_pretest_result_data.d_updated_at,
    t_pretest_result_data.c_global_id
   FROM t_pretest_result_data
  WHERE (((t_pretest_result_data.d_start IS NOT NULL) AND (t_pretest_result_data.d_end IS NULL)) OR ((t_pretest_result_data.d_start IS NOT NULL) AND (t_pretest_result_data.d_end IS NOT NULL)));
COMMENT ON VIEW v_pretest_result_data IS 'this is view for filtering invalid date in t_pretest_result';
