package interfaces

type EmailRequestBody struct {
	Template string    `json:"template"`
	Vars     variables `json:"vars"`
}

type variables struct {
	Name 	string `json:"name"`
	Content string `json:"content"`
}
