package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(gctx *gin.Context) {
	var payload CreateUserReq

	if err := gctx.ShouldBindJSON(&payload); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(gctx.Request.Context(), &payload)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, res)
}

func (h *Handler) Login(gctx *gin.Context) {
	var payload LoginUserReq

	if err := gctx.ShouldBindJSON(&payload); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Service.Login(gctx.Request.Context(), &payload)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gctx.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{
		Username: u.Username,
		ID: u.ID,
	}

	gctx.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(gctx *gin.Context) {
	gctx.SetCookie("jwt", "", -1, "", "", false, true)
	gctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}