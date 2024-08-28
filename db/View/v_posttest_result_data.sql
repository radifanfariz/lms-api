CREATE VIEW v_posttest_result_data (n_id, n_user_id, n_score, d_start, d_end, d_duration, j_answer, j_question_answered, c_created_by, c_updated_by, d_created_at, d_updated_at, c_global_id) AS  SELECT t_posttest_result_data.n_id,
    t_posttest_result_data.n_user_id,
    t_posttest_result_data.n_score,
    t_posttest_result_data.d_start,
    t_posttest_result_data.d_end,
    t_posttest_result_data.d_duration,
    t_posttest_result_data.j_answer,
    t_posttest_result_data.j_question_answered,
    t_posttest_result_data.c_created_by,
    t_posttest_result_data.c_updated_by,
    t_posttest_result_data.d_created_at,
    t_posttest_result_data.d_updated_at,
    t_posttest_result_data.c_global_id
   FROM t_posttest_result_data
  WHERE ((t_posttest_result_data.d_start IS NOT NULL) AND (t_posttest_result_data.d_end IS NOT NULL));
