package models

type Respond struct {
	Message Message
	Result  bool "true"
	Data    Authen
}

type Message struct {
	Th string	"ล็อกอินสำเร็จ"
	En string	"login success"
	Bu string
}
