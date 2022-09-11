package sqlenc

func ActorListByMovieCount() string {
	sql := `
	SELECT tb_actor.id, tb_actor.stage_name_jp, tb_actor.stage_name_cn, tb_actor.portrait, COUNT(*) AS movie_count 
	FROM tb_actor, tb_movie_actor 
	WHERE tb_actor.id = tb_movie_actor.actor_id 
	GROUP BY tb_actor.id 
	ORDER BY movie_count DESC
	LIMIT ? 
	OFFSET ?;
	`
	return sql
}
func ActorListByMovieView() string {
	sql := `
	SELECT tb_actor.id, tb_actor.stage_name_jp, tb_actor.stage_name_cn, tb_actor.portrait, COUNT(*) AS view_count
	FROM tb_actor, tb_movie_actor, tb_user_view 
	WHERE tb_actor.id = tb_movie_actor.actor_id AND tb_movie_actor.movie_id = tb_user_view.movie_id
	GROUP BY tb_actor.id 
	ORDER BY view_count DESC 
	LIMIT ? 
	OFFSET ?;
	`
	return sql
}
func ActorListByMovieFavorite() string {
	sql := `
	SELECT tb_actor.id, tb_actor.stage_name_jp, tb_actor.stage_name_cn, tb_actor.portrait, COUNT(*) AS favorite_count
	FROM tb_actor, tb_movie_actor, tb_user_favorite 
	WHERE tb_actor.id = tb_movie_actor.actor_id AND tb_movie_actor.movie_id = tb_user_favorite.movie_id
	GROUP BY tb_actor.id 
	ORDER BY favorite_count DESC 
	LIMIT ? 
	OFFSET ?;
	`
	return sql
}
