package useragent

import (
	"github.com/mssola/useragent"
	"runtime"
)

// ClientInfo содержит информацию о клиенте
type ClientInfo struct {
	DeviceType string
	OS         string
	Browser    string
	Version    string
	Model      string
	Cores      int
}

// GetClientInfo парсит User-Agent и возвращает данные
func GetClientInfo(userAgentString string) ClientInfo {
	// Используем New для создания объекта UserAgent
	ua := useragent.New(userAgentString)
	if ua == nil {
		return ClientInfo{} // Возвращает пустые данные при ошибке
	}

	// Определение типа устройства
	deviceType := "Desktop"
	if ua.Mobile() {
		deviceType = "Mobile"
	}

	// Получение информации об операционной системе
	os := ua.OSInfo().Name
	model := ua.Model()

	// Получение информации о браузере
	browserName, browserVersion := ua.Browser()

	return ClientInfo{
		DeviceType: deviceType,
		OS:         os,
		Browser:    browserName,
		Version:    browserVersion,
		Model:      model,
		Cores:      runtime.NumCPU(),
	}
}
