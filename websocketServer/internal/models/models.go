package models

type Response struct {
	Status  int         `json:"Status"`
	Payload interface{} `json:"Payload"`
}

type Users struct {
	User_id  uint `gorm:"primaryKey"`
	Username string
	Passwd   string
}

type Chatrooms struct {
	Chatroom_id   uint `gorm:"primaryKey"`
	Chatroom_name string
}

type Chatroom_users struct {
	Chat_users_id   uint `gorm:"primarKey"`
	Chat_users_user int
	Chat_users_chat int
}

type Messages struct {
	Message_id    uint `gorm:"primaryKey"`
	Message_text  string
	Message_date  string
	Message_owner int
}

type MessageSender struct {
	Connection interface{}
	Message    string
}
