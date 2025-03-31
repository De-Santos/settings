package domain

import "encoding/json"

type DataProvider interface {
	Get(key string) (json.RawMessage, error)
	Update(key string, val json.RawMessage) error
	Insert(key string, val json.RawMessage) error
	Delete(key string) error
}

type Logger interface {
	Debugf(format string, args ...any)
	Errorf(format string, args ...any)
}

type SettingService interface {
	Get(key string) ([]byte, error)
	GetAs(key string, out any) error
	Update(key string, val any) error
	Insert(key string, val any) error
	Delete(key string) error
}
