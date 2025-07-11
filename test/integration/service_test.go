package integration

import (
	"context"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paymentService.ProcessPayment(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessPayment(%d) error = %v, wantErr %v", tt.amount, err, tt.wantErr)
			}
		})
	}
}