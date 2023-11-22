package handlers

type SignDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type MessageDTO struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
}
