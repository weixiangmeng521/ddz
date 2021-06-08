package api

import (
	"ddz/app/cards"
	"encoding/json"
	"fmt"
)

type Message struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewMessage() *Message {
	return &Message{
		Code:    1,
		Message: "",
		Data:    "",
	}
}

func (t *Message) Success() *Message {
	t.Code = 1
	t.Message = "success"
	return t
}

func (t *Message) Error() *Message {
	t.Code = -1
	t.Message = "error"
	return t
}

func (t *Message) Warn() *Message {
	t.Code = 0
	t.Message = "warn"
	return t
}

func (t *Message) SetCode(code int8) *Message {
	t.Code = code
	return t
}

func (t *Message) GetCode() int8 {
	return t.Code
}

func (t *Message) SetMessage(msg string) *Message {
	t.Message = msg
	return t
}

func (t *Message) GetMessage() string {
	return t.Message
}

func (t *Message) SetData(data interface{}) *Message {
	t.Data = data
	return t
}

func (t *Message) GetData() interface{} {
	return t.Data
}

// message 转化
func ParseMessage(i interface{}) (msg *Message) {
	var b []byte
	switch i.(type) {
	case string:
		b = []byte(i.(string))
	case []byte:
		b = i.([]byte)
	case map[string]interface{}:
		b, _ = json.Marshal(i)
	default:
		fmt.Println(i, " it's not string type.")
		return
	}
	if err := json.Unmarshal(b, &msg); err != nil {
		fmt.Println("Unmarshall err: ", err)
		return
	}
	return
}

// cards info
type PlayCardsData struct {
	Type   string        `json:"type"`
	Option string        `json:"option"`
	Cards  []*cards.Card `json:"cards"`
}

// 解析出牌的信息
func ParsePlayCardsMessage(i interface{}) {
	var b []byte
	switch i.(type) {
	case string:
		b = []byte(i.(string))
	case []byte:
		b = i.([]byte)
	case map[string]interface{}:
		b, _ = json.Marshal(i)
	default:
		fmt.Println(i, " it's not string type.")
		return
	}

	msg := &struct {
		Data *PlayCardsData `json:"data"`
		*Message
	}{}

	if err := json.Unmarshal(b, &msg); err != nil {
		return
	}
	return
}
