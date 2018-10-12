package animals_test

import (
	"sample-go-testing/animals"
	"sample-go-testing/foods"
	"testing"
	"time"
)

func TestDuck_05(t *testing.T) {
	t.Run("it says quack", func(t *testing.T) {
		duck := createInstance()

		actual := duck.Say()
		expected := "taro says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		duck := createInstance()

		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple)
		expected := "taro ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}

func createInstance() *animals.Duck {
	duck := animals.NewDuck("taro")
	time.Sleep(3 * time.Second)

	return duck
}
