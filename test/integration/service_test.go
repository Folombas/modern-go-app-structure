package unit

import (
	"testing"
	"github.com/Folombas/modern-go-app-structure/internal/service"
)

func TestPaymentService_ProcessPayment(t *testing.T) {
	paymentService := service.NewPaymentService()

	tests := []struct {
		name    string
		amount  int
		wantErr bool
	}{
		{"Valid payment", 1000, false},
		{"Zero amount", 0, true},
		{"Negative amount", -100, true},
		{"Large amount", 1000000000, false},
		{"Minimum amount", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paymentService.ProcessPayment(tt.amount)
			
			hasError := err != nil
			if hasError != tt.wantErr {
				t.Errorf(
					"ProcessPayment(%d) error = %v, wantErr %v", 
					tt.amount, 
					err, 
					tt.wantErr,
				)
			}
			
			// Проверка текста ошибки (если нужна)
			if tt.wantErr && err != nil {
				expectedErr := "invalid payment amount"
				if err.Error() != expectedErr {
					t.Errorf(
						"ProcessPayment(%d) error text = '%s', want '%s'", 
						tt.amount, 
						err.Error(), 
						expectedErr,
					)
				}
			}
		})
	}
}