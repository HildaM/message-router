package controller

import (
	"io"
	"sync"

	"github.com/HildaM/message-router/model"
	"github.com/gin-gonic/gin"
)

var messageChanBufferSize = 10

var messageChanStore struct {
	Map   map[int]*chan *model.Message // 一个存储着 message channel 的 Map 结构，主键是userId
	Mutex sync.RWMutex
}

func init() {
	messageChanStore.Map = make(map[int]*chan *model.Message)
}

func messageChanStoreAdd(messageChan *chan *model.Message, userId int) {
	messageChanStore.Mutex.Lock()
	defer messageChanStore.Mutex.Unlock()

	messageChanStore.Map[userId] = messageChan
}

func messageChanStoreRemove(userId int) {
	messageChanStore.Mutex.Lock()
	defer messageChanStore.Mutex.Unlock()

	delete(messageChanStore.Map, userId)
}

func syncMessageToUser(message *model.Message, userId int) {
	messageChanStore.Mutex.RLock()
	defer messageChanStore.Mutex.RUnlock()

	messageChan, ok := messageChanStore.Map[userId]
	if !ok {
		return
	}
	*messageChan <- message
}

func GetNewMessages(c *gin.Context) {
	userId := c.GetInt("id")
	messageChan := make(chan *model.Message, messageChanBufferSize)
	messageChanStoreAdd(&messageChan, userId)

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-messageChan; ok {
			c.SSEvent("message", *msg)
			return true
		}
		return false
	})
	messageChanStoreRemove(userId)
	close(messageChan)
}
