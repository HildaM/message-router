package channel

import (
	"github.com/HildaM/message-router/common"
	"github.com/HildaM/message-router/model"
)

var AsyncMessageQueue chan int // 异步消息发送队列
var AsyncMessageQueueSize = 20 // 队列长度
var AsyncMessageSenderNum = 2  // 队列数量

func init() {
	AsyncMessageQueue = make(chan int, AsyncMessageQueueSize)
	for i := 0; i < AsyncMessageSenderNum; i++ {
		go asyncMessageSender()
	}
}

// LoadAsyncMessages 初始化异步信息发送的依赖（必须得等数据库准备就绪才能使用异步信息）
func LoadAsyncMessages() {
	ids, err := model.GetAsyncPendingMessageIds()
	if err != nil {
		// TODO
	}

	for _, id := range ids {
		AsyncMessageQueue <- id
	}
}

func asyncMessageSender() {
	for {
		id := <-AsyncMessageQueue
		message, err := model.GetMessageByID(id)
		if err != nil {
			// TODO log
			continue
		}

		status := common.MessageSendStatusFailed
		if err = asyncMessageSenderHelper(message); err != nil {
			// TODO log
		} else {
			status = common.MessageSendStatusSent
		}

		if err = message.UpdateStatus(status); err != nil {
			// TODO log
		}
	}
}

func asyncMessageSenderHelper(message *model.Message) error {
	user, err := model.GetUserById(message.Id, false)
	if err != nil {
		return err
	}

	channel_, err := model.GetChannelByName(message.Channel, user.Id)
	if err != nil {
		return err
	}
	return SendMessage(message, user, channel_)
}
