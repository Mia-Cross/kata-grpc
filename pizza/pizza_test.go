package pizza

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func QuatreFromages() Pizza {
	return Pizza{"4fromages", 1}
}

func TestShouldReturnTheBestPizzaKind(t *testing.T) {
	pizzas := NewPizza("4fromages", 1)
	expected := QuatreFromages()
	assert.Equal(t, expected, pizzas)
}
