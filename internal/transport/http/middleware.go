package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/pkg/auth"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	strId, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	id, err := Parse(strId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Set(userCtx, id)
}

func getUserId(c *gin.Context) (id UUID, err error) {
	ctxId, ok := c.Get(userCtx)
	if !ok {
		err = errors.New("user doesn't exist in context")
	}

	id = ctxId.(UUID)
	return
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.ParseToken(auth.Token(headerParts[1]))
}
