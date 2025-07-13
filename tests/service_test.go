package tests

import (
    "testing"
    "github.com/Folombas/modern-go-app-structure/internal/service"
    "github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
    result := service.SayHello()
    assert.Equal(t, "Hello from Service!", result)
}