package useragent

import (
    "github.com/mssola/useragent"
    "net"
    "runtime"
    "strings"
)

// ClientInfo содержит информацию о клиенте
type ClientInfo struct {
    DeviceType     string
    OS             string
    Browser        string
    Version        string
    Model          string
    Cores          int
    ConnectionType string
}

// GetClientInfo парсит User-Agent и возвращает данные
func GetClientInfo(userAgentString string, clientIP string) ClientInfo {
    ua := useragent.New(userAgentString)
    if ua == nil {
        return ClientInfo{} // Возвращает пустые данные при ошибке
    }

    deviceType := "Desktop"
    if ua.Mobile() {
        deviceType = "Mobile"
    }

    os := ua.OSInfo().Name
    model := ua.Model()
    browserName, browserVersion := ua.Browser()

    connectionType := getConnectionType(clientIP)

    return ClientInfo{
        DeviceType:     deviceType,
        OS:             os,
        Browser:        browserName,
        Version:        browserVersion,
        Model:          model,
        Cores:          runtime.NumCPU(),
        ConnectionType: connectionType,
    }
}

// getConnectionType определяет тип подключения
func getConnectionType(ip string) string {
    // Убираем порт из IP (например, "192.168.1.100:56789" → "192.168.1.100"
    ip = strings.Split(ip, ":")[0]

    ipObj := net.ParseIP(ip)
    if ipObj == nil {
        return "Не определено"
    }

    // Проверка частных IP-адресов (LAN/Wi-Fi)
    if ipObj.IsPrivate() || ipObj.IsLoopback() {
        return "LAN / Wi-Fi"
    }

    // Проверка IPv4-диапазона для 172.16.0.0–172.31.255.255
    if ipObj.To4() != nil {
        if strings.HasPrefix(ip, "192.168.") || 
			strings.HasPrefix(ip, "10.") ||
			(strings.HasPrefix(ip, "172.") && ipObj[0] == 172 && ipObj[1] >= 16 && ipObj[1] <= 31) {
            return "LAN / Wi-Fi"
        }
    }

    // Если IP публичный — это мобильный интернет
    return "Мобильный"
}