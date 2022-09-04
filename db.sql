CREATE DATABASE IF NOT EXISTS db_porn DEFAULT CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS tb_user (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    ip VARCHAR(12) NOT NULL DEFAULT '0.0.0.0',
    nickname VARCHAR(24) NOT NULL DEFAULT 'Guest',
    devie_name VARCHAR(128),
    devie_code VARCHAR(128),
    email VARCHAR(32),
    email_verify TINYINT(1) NOT NULL DEFAULT 0,
    email_verify_at DATETIME,
    enable TINYINT(1) NOT NULL DEFAULT 1,
    create_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS tb_movie (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(12) NOT NULL,
    title VARCHAR(32) NOT NULL,
    cover_cross VARCHAR(128) NOT NULL,
    cover_shaft VARCHAR(128) NOT NULL,
    release_date DATETIME NOT NULL,
    enable TINYINT(1) NOT NULL DEFAULT 0,
    create_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS tb_actor (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    stage_name_jp VARCHAR(12) NOT NULL,
    stage_name_cn VARCHAR(12),
    portrait VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS tb_movie_actor (
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES tb_movie(id),
    FOREIGN KEY(actor_id) REFERENCES tb_actor(id)
);

CREATE TABLE IF NOT EXISTS tb_tag (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    named VARCHAR(12) NOT NULL,
    cover VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS tb_movie_tag (
    movie_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY(movie_id) REFERENCES tb_movie(id),
    FOREIGN KEY(tag_id) REFERENCES tb_tag(id)
);

CREATE TABLE IF NOT EXISTS tb_user_view (
    user_id INT NOT NULL,
    movie_id INT NOT NULL,
    create_at DATETIME NOT NULL,
    FOREIGN KEY(user_id) REFERENCES tb_user(id),
    FOREIGN KEY(movie_id) REFERENCES tb_movie(id)
);

CREATE TABLE IF NOT EXISTS tb_user_favorite (
    user_id INT NOT NULL,
    movie_id INT NOT NULL,
    create_at DATETIME NOT NULL,
    FOREIGN KEY(user_id) REFERENCES tb_user(id),
    FOREIGN KEY(movie_id) REFERENCES tb_movie(id)
);