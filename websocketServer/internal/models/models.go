package models

type Response struct {
	Status  int         `json:"Status"`
	Payload interface{} `json:"Payload"`
}

type Users struct {
	User_id  uint `gorm:"primary key"`
	Username string
	Passwd   string
}

type Chatrooms struct {
	Chatroom_id   uint `gorm:"primary key"`
	Chatroom_name string
}

type Chatroom_users struct {
	Chat_users_id   uint `gorm:"primary key"`
	Chat_users_user int
	Chat_users_chat int
}

type Messages struct {
	Message_id    uint `gorm:"primary key"`
	Message_text  string
	Message_date  string
	Message_owner int
}
