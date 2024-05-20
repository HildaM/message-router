package controller

import (
	"github.com/HildaM/message-router/model"
	"github.com/gin-gonic/gin"
)

func GetPushMessage(c *gin.Context) {
	message := model.Message{
		Title:       c.Query("title"),
		Description: c.Query("description"),
		Content:     c.Query("content"),
		URL:         c.Query("url"),
		Channel:     c.Query("channel"),
		Token:       c.Query("token"),
		To:          c.Query("to"),
		Desp:        c.Query("desp"),
		Short:       c.Query("short"),
		OpenId:      c.Query("openid"),
		Async:       c.Query("async") == "true",
	}
	keepCompatible(&message)
	pushMessageHelper(c, &message)
}

func keepCompatible(message *model.Message) {
	// Keep compatible with ServerChan: https://sct.ftqq.com/sendkey
	if message.Description == "" {
		message.Description = message.Short
	}
	if message.Content == "" {
		message.Content = message.Desp
	}
	if message.To == "" {
		message.To = message.OpenId
	}
}

func pushMessageHelper(c *gin.Context, message *model.Message) {
	// TODO
}
