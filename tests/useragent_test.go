package tests

import (
	"testing"

	"github.com/Folombas/modern-go-app-structure/internal/useragent"
	"github.com/stretchr/testify/assert"
)

func TestGetClientInfo(t *testing.T) {
	tests := []struct {
		name               string
		userAgentStr       string
		clientIP           string // Добавлено поле
		expectedDeviceType string
		expectedOS         string
		expectedBrowser    string
		expectedModel      string
		expectedConnection string
	}{
		{
			name:               "Mobile Device",
			userAgentStr:       "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
			clientIP:           "192.168.1.100", // Пример LAN/Wi-Fi
			expectedDeviceType: "Mobile",
			expectedOS:         "Android",
			expectedBrowser:    "Chrome",
			expectedModel:      "SM-A102U",
			expectedConnection: "LAN / Wi-Fi",
		},
		{
			name:               "Desktop Device",
			userAgentStr:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Safari/537.36",
			clientIP:           "8.8.8.8", // Публичный IP
			expectedDeviceType: "Desktop",
			expectedOS:         "Windows",
			expectedBrowser:    "Chrome",
			expectedModel:      "",
			expectedConnection: "Мобильный",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := useragent.GetClientInfo(tt.userAgentStr, tt.clientIP)

			assert.Equal(t, tt.expectedDeviceType, info.DeviceType)
			assert.Contains(t, info.OS, tt.expectedOS)
			assert.Contains(t, info.Browser, tt.expectedBrowser)
			assert.Equal(t, tt.expectedModel, info.Model)
			assert.Equal(t, tt.expectedConnection, info.ConnectionType)
		})
	}
}
