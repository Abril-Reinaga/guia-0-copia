package estructurasrepetitivas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	resultado := Factorial(5)
	assert.Equal(t, 120, resultado)
}

func TestFactorialFallido(t *testing.T) {
	assert.Panics(t, func() { Factorial(-5) }, "Debe ser un número positivo o 0")
}
