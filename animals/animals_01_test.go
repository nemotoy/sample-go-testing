package animals_test

import (
	"sample-go-testing/animals"
	"testing"
)

func TestDuck_Say(t *testing.T) {
	duck := animals.NewDuck("taro")
	actual := duck.Say()
	expected := "taro says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
