package sqlenc

func ActorListByMovieCount(limit, offset int) string {
	sql := `
	SELECT tb_actor.id, tb_actor.stage_name_jp, tb_actor.stage_name_cn, tb_actor.portrait, COUNT(tb_actor.id) AS count 
	FROM tb_actor, tb_movie_actor 
	WHERE tb_actor.id = tb_movie_actor.actor_id 
	GROUP BY tb_actor.id 
	ORDER BY count DESC
	LIMIT ? 
	OFFSET ?;
	`
	return sql
}
func ActorListByMovieView(limit, offset int) string {
	sql := `
	SELECT tb_actor.id, tb_actor.stage_name_jp, tb_actor.stage_name_cn, tb_actor.portrait, COUNT(*) AS count 
	FROM tb_actor, tb_movie_actor, tb_user_view 
	WHERE tb_actor.id = tb_movie_actor.actor_id AND tb_movie_actor.movie_id = tb_user_view.movie_id 
	GROUP BY tb_actor.id 
	ORDER BY count DESC
	LIMIT ? 
	OFFSET ?;
	`
	return sql
}
func ActorListByMovieFavorite(limit, offset int) string { return "" }
