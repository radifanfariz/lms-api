CREATE VIEW v_materi_result_data (n_id, n_user_id, c_global_id, d_start, d_end, d_duration, c_created_by, c_updated_by, d_created_at, d_updated_at, n_materi_id) AS  SELECT t_materi_result_data.n_id,
    t_materi_result_data.n_user_id,
    t_materi_result_data.c_global_id,
    t_materi_result_data.d_start,
    t_materi_result_data.d_end,
    t_materi_result_data.d_duration,
    t_materi_result_data.c_created_by,
    t_materi_result_data.c_updated_by,
    t_materi_result_data.d_created_at,
    t_materi_result_data.d_updated_at,
    t_materi_result_data.n_materi_id
   FROM t_materi_result_data
  WHERE ((t_materi_result_data.d_start IS NOT NULL) AND (t_materi_result_data.d_end IS NOT NULL));
