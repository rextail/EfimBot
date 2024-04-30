package telegram

type Update struct {
	ID      int `json:"update_id"`
	Message Message
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Message struct {
	From User
	Chat Chat
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
type Chat struct {
	ID int `json:"id"`
}
