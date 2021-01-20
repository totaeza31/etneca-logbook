package models

type RespondAuthen struct {
	Result  bool    `json:"result"`
	Message Message `json:"message"`
	Data    Authen  `json:"data"`
}

type RespondUser struct {
	Message Message `json:"message"`
	Result  bool    `json:"result"`
	Data    User    `json:"data"`
}

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
