package model

import "errors"

// User if you add sensitive fields, don't forget to clean them in setupLogin function.
// Otherwise, the sensitive information will be saved on local storage in plain text!
type User struct {
	Id                    int    `json:"id"`
	Username              string `json:"username" gorm:"unique;index" validate:"max=12"`
	Password              string `json:"password" gorm:"not null;" validate:"min=8,max=20"`
	DisplayName           string `json:"display_name" gorm:"index" validate:"max=20"`
	Role                  int    `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status                int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	Token                 string `json:"token"`
	Email                 string `json:"email" gorm:"index" validate:"max=50"`
	GitHubId              string `json:"github_id" gorm:"column:github_id;index"`
	WeChatId              string `json:"wechat_id" gorm:"column:wechat_id;index"`
	VerificationCode      string `json:"verification_code" gorm:"-:all"` // this field is only for Email verification, don't save it to database!
	Channel               string `json:"channel"`
	SendEmailToOthers     int    `json:"send_email_to_others" gorm:"type:int;default:0"`
	SaveMessageToDatabase int    `json:"save_message_to_database" gorm:"type:int;default:0"`
}

func GetUserById(id int, selectAll bool) (*User, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	user := User{Id: id}
	var err error = nil
	if selectAll {
		err = DB.First(&user, "id = ?", id).Error
	} else {
		err = DB.Select([]string{"id", "username", "display_name", "role", "status", "email", "wechat_id", "github_id",
			"channel", "token", "save_message_to_database",
		}).First(&user, "id = ?", id).Error
	}
	return &user, err
}
