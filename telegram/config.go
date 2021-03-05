package telegram

import (
	"encoding/json"
	"io"
)

type TapirConfig struct {
	NthMediaRule struct {
		Count int
		ReadonlyDuration int
	}
}

type TapirConfigManager struct {
	DefaultConfig *TapirConfig
	ChatConfigs map[string]*TapirConfig
}

func NewTapirConfigManager(reader io.Reader) (*TapirConfigManager, error) {
	var configManager TapirConfigManager
	err := json.NewDecoder(reader).Decode(&configManager)
	if err != nil {
		return nil, err
	}
	return &configManager, nil
}

func (m *TapirConfigManager) GetConfig(chatName string) *TapirConfig {
	config, ok := m.ChatConfigs[chatName]
	if ok {
		return config
	} else {
		return m.DefaultConfig
	}
}