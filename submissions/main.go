package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kzh/noob/lib/message"
	"github.com/kzh/noob/lib/model"
	noobsess "github.com/kzh/noob/lib/sessions"
)

func handleSubmit(ctx *gin.Context) {
	var submission model.Submission
	if err := ctx.ShouldBind(&submission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	submission.ID = uuid.New().String()

	if err := message.Schedule(submission); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func main() {
	log.Println("Noob: Submissions MS is starting...")

	r := gin.Default()
	r.Use(noobsess.Sessions())
	r.Use(noobsess.LoggedIn(false))

	r.POST("/submit", handleSubmit)

	if err := r.Run(); err != nil {
		log.Println(err)
	}
}
