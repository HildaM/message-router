package model

import "errors"

const (
	TypeEmail             = "email"
	TypeWeChatTestAccount = "test"
	TypeWeChatCorpAccount = "corp_app"
	TypeCorp              = "corp"
	TypeLark              = "lark"
	TypeDing              = "ding"
	TypeTelegram          = "telegram"
	TypeDiscord           = "discord"
	TypeBark              = "bark"
	TypeClient            = "client"
	TypeNone              = "none"
	TypeOneBot            = "one_bot"
	TypeGroup             = "group"
	TypeLarkApp           = "lark_app"
	TypeCustom            = "custom"
	TypeTencentAlarm      = "tencent_alarm"
)

type Channel struct {
	Id          int    `json:"id"`
	Type        string `json:"type" gorm:"type:varchar(32)"`
	UserId      int    `json:"user_id" gorm:"uniqueIndex:name_user_id;index"`
	Name        string `json:"name" gorm:"type:varchar(32);uniqueIndex:name_user_id"`
	Description string `json:"description"`
	Status      int    `json:"status" gorm:"default:1"` // enabled, disabled
	Secret      string `json:"secret" gorm:"index"`
	AppId       string `json:"app_id"`
	AccountId   string `json:"account_id"`
	URL         string `json:"url" gorm:"column:url"`
	Other       string `json:"other"`
	CreatedTime int64  `json:"created_time" gorm:"bigint"`
}

func GetChannelByName(name string, userId int) (*Channel, error) {
	if name == "" || userId == 0 {
		return nil, errors.New("name 或 userId 为空！")
	}
	c := Channel{Name: name, UserId: userId}
	err := DB.Where(c).First(&c).Error
	return &c, err
}
