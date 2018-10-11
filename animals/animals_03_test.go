package animals_test

import (
	"os"
	"sample-go-testing/animals"
	"sample-go-testing/foods"
	"testing"
)

func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)

}

func TestDuck_Say_03(t *testing.T) {
	duck := animals.NewDuck("taro")
	actual := duck.Say()
	expected := "taro says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestDuck_Eat(t *testing.T) {
	duck := animals.NewDuck("taro")
	apple := foods.NewApple("sunfuji")

	actual := duck.Eat(apple)
	expected := "taro ate sunfuji"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
