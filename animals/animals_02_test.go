package animals_test

import (
	"sample-go-testing/animals"
	"sample-go-testing/foods"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuck(t *testing.T) {
	duck := animals.NewDuck("taro")

	t.Run("it say quack", func(t *testing.T) {
		actual := duck.Say()
		expected := "taro says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		apple := foods.NewApple("tsugaru")
		actual := duck.Eat(apple)
		expected := "taro ate tsugaru"
		assert.Equal(t, actual, expected)
	})
}
