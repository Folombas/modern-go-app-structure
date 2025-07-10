package utils

import "math/rand"

// GenerateID - генерирует случайный ID
func GenerateID(prefix string) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return prefix + "-" + string(b)
}

// CalculateTotal - рассчитывает итоговую сумму с комиссией
func CalculateTotal(amount, feePercent int) int {
	fee := amount * feePercent / 100
	return amount + fee
}
