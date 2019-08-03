package handler

import (
	"github.com/ashishthakur913/project/article"
	"github.com/ashishthakur913/project/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
