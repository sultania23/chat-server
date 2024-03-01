package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/dto"
	"net/http"
)

func (h *Handler) initConversationRoutes(api *gin.RouterGroup) {
	chat := api.Group("/conversation", h.userIdentity)
	{
		chat.POST("/", h.createConversation)
		chat.GET("/", h.getUserConversations)
		chat.GET("/:id", h.getConversationById)
		chat.DELETE("/:id", h.deleteConversationById)
	}
}

// @Summary 	Create Conversation
// @Security 	Bearer
// @Tags		conversation
// @Description	creating new conversation
// @ID			createConversation
// @Accept		json
// @Produce		json
// @Param       input    body      	dto.ConversationDTO  true  "conversation information"
// @Success     201
// @Failure		400		{object}  	errorResponse
// @Failure     500		{object}  	errorResponse
// @Failure     default	{object}  	errorResponse
// @Router      /conversation 		[post]
func (h *Handler) createConversation(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var conversationDTO dto.ConversationDTO
	if err := c.BindJSON(&conversationDTO); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.conversationService.CreateConversation(c, id, conversationDTO)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary 	Get Conversations
// @Security 	Bearer
// @Tags 		conversation
// @Description gets all conversations
// @ID 			allConversations
// @Accept  	json
// @Produce  	json
// @Success 	200 	{array}		dto.ConversationDTO
// @Failure 	500 	{object} 	errorResponse
// @Failure 	default {object} 	errorResponse
// @Router 		/conversation 		[get]
func (h *Handler) getUserConversations(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	conversations, err := h.conversationService.GetByOwnerId(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if conversations != nil {
		c.JSON(http.StatusOK, conversations)
	}
}

// @Summary 	GET Conversation By ID
// @Security 	Bearer
// @Tags 		conversation
// @Description gets conversation by ID
// @ID 			getConversationById
// @Accept  	json
// @Produce  	json
// @Param id path string true "Conversation ID"
// @Success 	200 	{object} 	dto.ConversationDTO
// @Failure 	500 	{object} 	errorResponse
// @Failure 	default {object} 	errorResponse
// @Router 		/conversation/{id} 	[get]
func (h *Handler) getConversationById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorMessage := fmt.Sprintf("Illegal format of ID [%s]", id)
		newErrorResponse(c, http.StatusBadRequest, errorMessage)
		return
	}

	conversation, err := h.conversationService.GetById(c, id)
	if err != nil && err.Error() == "sql: no rows in result set" {
		errorMessage := fmt.Sprintf("Conversation not found by ID [%s]", id)
		newErrorResponse(c, http.StatusNotFound, errorMessage)
		return
	}

	c.JSON(http.StatusOK, conversation)
}

// @Summary 	Delete Conversation By ID
// @Security 	Bearer
// @Tags 		conversation
// @Description deletes conversation by ID
// @ID 			deleteConversationById
// @Accept  	json
// @Produce  	json
// @Param id path string true "Conversation ID"
// @Success 	204
// @Failure 	500 {object} 		errorResponse
// @Failure 	default {object} 	errorResponse
// @Router 		/conversation/{id}	[delete]
func (h *Handler) deleteConversationById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorMessage := fmt.Sprintf("Illegal format of ID [%s]", id)
		newErrorResponse(c, http.StatusBadRequest, errorMessage)
		return
	}

	if err := h.conversationService.RemoveById(c, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
