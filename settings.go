package settings

import (
	"encoding/json"
	"settings/config"
)

func New(c config.Config) *SettingService {
	return &SettingService{
		c,
	}
}

type SettingService struct {
	config.Config
}

func (s *SettingService) Get(key string) ([]byte, error) {
	s.Log.Debugf("Getting setting by key: %s", key)
	v, err := s.DP.Get(key)
	if err != nil {
		s.Log.Errorf("Failed to update setting by key: %s, error: %s", key, err)
	}
	return v, err
}

func (s *SettingService) GetAs(key string, out any) error {
	b, err := s.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (s *SettingService) Update(key string, val any) error {
	s.Log.Debugf("Updating setting by key: %s, val: %s", key, val)
	b, err := s.marshal(val)
	if err != nil {
		return err
	}
	err = s.DP.Update(key, b)
	if err != nil {
		s.Log.Errorf("Failed to get setting by key: %s, error: %s", key, err)
	}
	return err
}

func (s *SettingService) Insert(key string, val any) error {
	s.Log.Debugf("Insert setting by key: %s, val: %T", key, val)
	b, err := s.marshal(val)
	if err != nil {
		return err
	}
	err = s.DP.Insert(key, b)
	if err != nil {
		s.Log.Errorf("Failed to insert setting key: %s, val: %T, err: %s", key, val, err)
		return err
	}
	return nil
}

func (s *SettingService) Delete(key string) error {
	s.Log.Debugf("Delete setting by key: %s", key)
	err := s.DP.Delete(key)
	if err != nil {
		s.Log.Errorf("Failed to delete setting key: %s, err: %s", key, err)
		return err
	}
	return nil
}

func (s *SettingService) marshal(val any) ([]byte, error) {
	b, err := json.Marshal(val)
	if err != nil {
		s.Log.Errorf("Failed to marshal obj: %T, err: %s", val, err)
		return nil, err
	}
	return b, nil
}
