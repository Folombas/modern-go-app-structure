package greeter_test

import (
	"testing"
	"github.com/Folombas/modern-go-app-structure/internal/greeter"
)

func TestSayHello(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid name",
			input:    "Тест",
			expected: "Привет, Тест! Добро пожаловать в мир Go-программирования!",
			wantErr:  false,
		},
		{
			name:     "Empty name",
			input:    "",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Name with spaces",
			input:    "   ",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := greeter.SayHello(tc.input)
			
			if tc.wantErr {
				if err == nil {
					t.Errorf("Ожидалась ошибка, но её нет")
				}
			} else {
				if err != nil {
					t.Errorf("Неожиданная ошибка: %v", err)
				}
				if result != tc.expected {
					t.Errorf("Ожидалось: %s, Получено: %s", tc.expected, result)
				}
			}
		})
	}
}