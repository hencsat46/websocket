package handlers

type SignDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type MessageDTO struct {
	UserId  int    `json:"userId"`
	Message string `json:"message"`
}
