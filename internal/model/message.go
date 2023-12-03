package model

import "fmt"

type Message struct {
	ID               int                    `json:"id"`
	Level            string                 `json:"level"`
	Text             string                 `json:"text"`
	MicroServiceName string                 `json:"ms_name"`
	Time             string                 `json:"time"`
	Other            map[string]interface{} `json:"other,omitempty"`
}

func (msg *Message) String() string {
	res := fmt.Sprintf("Message:{\n")
	res = res + fmt.Sprintf("\tID: %d\n", msg.ID)
	res = res + fmt.Sprintf("\tLevel: %s\n", msg.Level)
	res = res + fmt.Sprintf("\tText: %s\n", msg.Text)
	res = res + fmt.Sprintf("\tMicroService Name: %s\n", msg.MicroServiceName)
	res = res + fmt.Sprintf("\tTime: %s\n", msg.Time)
	res = res + fmt.Sprintf("\tOther: %s\n", msg.Other)
	res = res + fmt.Sprintf("}")
	return res
}
