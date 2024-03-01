package dto

import . "github.com/google/uuid"

type ConversationDTO struct {
	Name  string `json:"name" binding:"required"`
	Owner UUID   `json:"owner"`
	//Participant UserDTO `json:"participants" binding:"required"`
}
