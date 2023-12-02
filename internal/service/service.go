package service

import "logger-bot/internal/model"

type MethodLogger interface {
	GetLogs() ([]*model.Message, error)
	GetLog(id string) (*model.Message, error)
}
