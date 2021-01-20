package models

type Respond struct {
	Message Message `json:"message"`
	Result  bool    `json:"result"`
	Data    Authen  `json:"data"`
}

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
