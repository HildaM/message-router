package model

import (
	"errors"

	"github.com/HildaM/message-router/common"
)

type Message struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url" gorm:"column:url"`
	Channel     string `json:"channel"`
	Token       string `json:"token" gorm:"-:all"`
	HTMLContent string `json:"html_content"  gorm:"-:all"`
	Timestamp   int64  `json:"timestamp" gorm:"type:bigint"`
	Link        string `json:"link" gorm:"unique;index"`
	To          string `json:"to" gorm:"column:to"`           // if specified, will send to this user(s)
	Status      int    `json:"status" gorm:"default:0;index"` // pending, sent, failed
	OpenId      string `json:"openid" gorm:"-:all"`           // alias for to
	Desp        string `json:"desp" gorm:"-:all"`             // alias for content
	Short       string `json:"short" gorm:"-:all"`            // alias for description
	Async       bool   `json:"async" gorm:"-"`                // if true, will send message asynchronously
}

func GetAsyncPendingMessageIds() (ids []int, err error) {
	err = DB.Model(&Message{}).Where("status = ?", common.MessageSendStatusAsyncPending).Pluck("id", &ids).Error
	return ids, err
}

func GetMessageByID(id int) (*Message, error) {
	if id == 0 {
		return nil, errors.New("ID 为空！")
	}

	message := Message{Id: id}
	err := DB.Where(message).First(&message).Error
	return &message, err
}

func (message *Message) UpdateStatus(status int) error {
	err := DB.Model(message).Update("status", status).Error
	return err
}
