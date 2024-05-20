package channel

import (
	"errors"

	"github.com/HildaM/message-router/model"
)

// SendMessage TODO 发送消息的公共方法
func SendMessage(message *model.Message, user *model.User, channel_ *model.Channel) error {
	switch channel_.Type {
	case model.TypeEmail:
		return nil
	default:
		return errors.New("不支持的消息通道：" + channel_.Type)
	}
}
