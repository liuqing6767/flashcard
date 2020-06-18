package router

import (
	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/controller/auth"
	"github.com/liuximu/flashcard/controller/card"
	"github.com/liuximu/flashcard/controller/knows"
	"github.com/liuximu/flashcard/controller/template"
	"github.com/liuximu/flashcard/shared"
)

func Route(e *echo.Echo) {
	root := e.Group("/api")

	authGroup := root.Group("/auth")
	authGroup.POST("/login", shared.Mine2Handler(auth.Login))
	authGroup.POST("/register", shared.Mine2Handler(auth.Register))
	authGroup.GET("/register/active", shared.Mine2Handler(auth.RegisterActive))
	authGroup.POST("/reset_email", shared.Mine2Handler(auth.Register))
	authGroup.GET("/reset_email/active", shared.Mine2Handler(auth.Register))

	knowGroup := root.Group("/know", auth.MustLoginHandlerFunc)
	knowGroup.GET("", shared.Mine2Handler(knows.Tree))
	knowGroup.POST("", shared.Mine2Handler(knows.Create))
	knowGroup.POST("/modify", shared.Mine2Handler(knows.Modify))

	templateGroup := root.Group("/template", auth.MustLoginHandlerFunc)
	templateGroup.GET("/:tid", shared.Mine2Handler(template.One))
	templateGroup.POST("/:tid", shared.Mine2Handler(template.Update))
	templateGroup.POST("/:tid/del", shared.Mine2Handler(template.Delete))
	templateGroup.GET("", shared.Mine2Handler(template.List))
	templateGroup.POST("", shared.Mine2Handler(template.Create))

	cardGroup := root.Group("/card", auth.MustLoginHandlerFunc)
	cardGroup.GET("/know/:kid", shared.Mine2Handler(card.List))
	cardGroup.GET("/:id", shared.Mine2Handler(card.One))
	cardGroup.POST("/:id/del", shared.Mine2Handler(card.Delete))
	cardGroup.POST("", shared.Mine2Handler(card.Upsert))

	cardGroup.GET("/learning/:kid", shared.Mine2Handler(card.LearningList))
	cardGroup.POST("/progress", shared.Mine2Handler(card.UpdateLearnProgress))
	cardGroup.GET("/progress", shared.Mine2Handler(card.LearnProgressStatis))
}
