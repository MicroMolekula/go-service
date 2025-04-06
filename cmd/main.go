package main

import (
	"github.com/MicroMolekula/gpt-service/internal/config"
	"github.com/MicroMolekula/gpt-service/internal/controller"
	"github.com/MicroMolekula/gpt-service/internal/database"
	"github.com/MicroMolekula/gpt-service/internal/mongo"
	"github.com/MicroMolekula/gpt-service/internal/repository"
	"github.com/MicroMolekula/gpt-service/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	mongoClient, err := mongo.NewMongoClient(cfg.Mongo.URI, cfg.Mongo.DBName)
	if err != nil {
		log.Fatal(err)
	}
	userPlanCollection := mongoClient.GetCollection("user_plans")
	userPlanRepository := repository.NewUserPlanRepository(userPlanCollection)
	userRepository := repository.NewUserRepository(db)
	messageRepository := repository.NewMessageRepository(db)
	gptService := service.NewGptService(cfg)
	gptController := controller.NewGPTController(gptService)
	fitnessService := service.NewFitnessService(gptService, cfg, userPlanRepository)
	chatService := service.NewChatService(gptService, messageRepository)
	fitnessController := controller.NewFitnessController(fitnessService)
	chatController := controller.NewChatController(chatService)
	userMiddleware := controller.NewUserMiddleware(userRepository)
	engine := gin.Default()
	engine.Use(userMiddleware.Middleware())
	engine.POST("/query", gptController.GetAnswer)
	engine.GET("/plan/generate", fitnessController.GetPlan)
	engine.GET("/profile", userMiddleware.Profile)
	engine.GET("/plan", fitnessController.FindPlanByUserId)
	engine.POST("/chat/send", chatController.SendMessage)
	engine.GET("/chat", chatController.GetMessages)

	if err = engine.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
