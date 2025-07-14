package useragent

import (
    "github.com/mssola/useragent"
    "runtime"
    "strings"
    "net"
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
        return ClientInfo{}
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
    ip = strings.Split(ip, ":")[0]

    ipObj := net.ParseIP(ip)
    if ipObj == nil {
        return "Не определено"
    }

    if ipObj.IsPrivate() || ipObj.IsLoopback() {
        return "LAN / Wi-Fi"
    }

    if ipObj.To4() != nil {
        if strings.HasPrefix(ip, "192.168.") || 
            strings.HasPrefix(ip, "10.") || 
            (strings.HasPrefix(ip, "172.") && ip[4] >= '1' && ip[4] <= '3') {
            return "LAN / Wi-Fi"
        }
    }

    return "Мобильный"
}