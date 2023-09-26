package model

type Message struct {
	Type   string
	Data   interface{}
	SentAt int
	SentBy int
	SentTo int
}

type SessionInfo struct {
	UserId   int
	ServerIp string
}
