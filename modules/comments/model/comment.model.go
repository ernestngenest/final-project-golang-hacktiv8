package model

import (
	Photo "final_project_hacktiv8/modules/photos/model"
	User "final_project_hacktiv8/modules/users/model"
	"time"
)

type Comment struct {
	ID        uint         `json:"id" gorm:"primaryKey"`
	UserID    uint         `json:"user_id"`
	PhotoID   uint         `json:"photo_id"`
	Message   string       `json:"message" gorm:"not null"`
	User      *User.User   `gorm:"foreignKey:UserID"`
	Photo     *Photo.Photo `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

/*
	-- Postgres syntax --

-- create sequence comment_id_seq;
	CREATE SEQUENCE comments_id_seq;

-- create table - comments
	CREATE TABLE COMMENTS (
		ID 			INT 		NOT NULL 	DEFAULT nextval('comments_id_seq'::regclass),
		User_id 	INT 		REFERENCES USERS(ID)	NOT NULL,
		Photo_id 	INT 		REFERENCES PHOTOS(ID)	NOT NULL,
		Message 	text 								NOT NULL,
		Created_at 	timestamp 							NOT NULL,
		Updated_at 	timestamp 							NOT NULL
	);
*/
