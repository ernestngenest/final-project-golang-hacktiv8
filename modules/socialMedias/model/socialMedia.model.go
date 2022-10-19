package model

import (
	"final_project_hacktiv8/modules/users/model"
	"time"
)

type SocialMedia struct {
	ID             uint       `json:"id gorm:primaryKey"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Name           string     `json:"name" gorm:"unique;not null"`
	SocialMediaUrl string     `json:"social_media_url" gorm:"unique;not null"`
	UserID         uint       `json:"user_id"`
	User           model.User `json:"user"`
}

/*
-- POSTGRESQL SYNTAX --

	CREATE SEQUENCE social_media_id_seq;

	CREATE TABLE IF NOT EXISTS social_media (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		name VARCHAR(255) NOT NULL,
		social_media_url VARCHAR(255) NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

*/
