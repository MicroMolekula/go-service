package controller

import (
	"github.com/MicroMolekula/gpt-service/internal/models"
	"github.com/MicroMolekula/gpt-service/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserMiddleware struct {
	userRepository *repository.UserRepository
}

func NewUserMiddleware(userRepository *repository.UserRepository) *UserMiddleware {
	return &UserMiddleware{userRepository}
}

func (m *UserMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetHeader("x-user-id")
		if userId == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		id, err := strconv.Atoi(userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		user, err := m.userRepository.FindOneById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func (m *UserMiddleware) Profile(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	ctx.JSON(http.StatusOK, user)
}
