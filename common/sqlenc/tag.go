package sqlenc

func TagListSql() string {
	return `
	SELECT tb_tag.id, tb_tag.named, tb_tag.cover, COUNT(*) AS count 
	FROM tb_tag, tb_movie_tag 
	WHERE tb_tag.id = tb_movie_tag.tag_id 
	GROUP BY tb_tag.id 
	ORDER BY count DESC;
	`
}
