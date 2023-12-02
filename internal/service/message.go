package service

import (
	"encoding/json"
	"io"
	"logger-bot/internal/model"
	"net/http"
)

type MessageService struct {
	urlapi string
}

func (m *MessageService) GetLogs() ([]*model.Message, error) {
	resp, err := http.Get(m.urlapi)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var msgs []*model.Message

	err = json.Unmarshal(body, &msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil

}

func (m *MessageService) GetLog(id string) (*model.Message, error) {
	resp, err := http.Get(m.urlapi + "/" + id)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var msg *model.Message

	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func NewMessageService(urlapi string) *MessageService {
	return &MessageService{urlapi: urlapi}
}

var _ MethodLogger = (*MessageService)(nil)
