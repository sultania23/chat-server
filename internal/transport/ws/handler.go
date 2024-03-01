package ws

import (
	"github.com/gin-gonic/gin"
	. "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/tuxoo/idler/internal/config"
	"github.com/tuxoo/idler/internal/service"
	"github.com/tuxoo/idler/pkg/cache"
	"net/http"
)

type Handler struct {
	HubCache            cache.Cache[UUID, Pool]
	Upgrader            *websocket.Upgrader
	MessageService      service.Messages
	ConversationService service.Conversations
}

func NewHandler(cfg config.WSConfig, hubCache cache.Cache[UUID, Pool], messageService service.Messages, conversationService service.Conversations) *Handler {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.ReadBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Handler{
		Upgrader:            upgrader,
		HubCache:            hubCache,
		MessageService:      messageService,
		ConversationService: conversationService,
	}
}

func (h *Handler) InitWSConversations() http.Handler {
	handler := gin.New()
	handler.GET("/conversation", h.CreateWSConversation)

	return handler
}

func (h *Handler) CreateWSConversation(c *gin.Context) {
	params := c.Request.URL.Query()

	id, err := Parse(params.Get("id"))
	if err != nil {
		return
	}

	if _, err := h.ConversationService.GetById(c, id); err != nil {
		return
	}

	conn, err := h.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	pool, err := h.HubCache.Get(c.Request.Context(), id)
	if err != nil && err.Error() == "cache: value not found" {
		pool = NewPool(id)
		err := h.HubCache.Set(c.Request.Context(), id, pool)
		if err != nil {
			return
		}
		go pool.Run()
	}

	client := NewClient(params.Get("user"), conn, pool, h.MessageService)
	go client.HandleMessage(c.Request.Context())
}
