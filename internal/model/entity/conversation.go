package entity

import . "github.com/google/uuid"

type Conversation struct {
	Id    UUID   `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Owner UUID   `json:"owner" binding:"required"`
	//Participants []dto.UserDTO `json:"participants" binding:"required"`
}
