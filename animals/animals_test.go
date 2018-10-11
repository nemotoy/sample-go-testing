package animals

import "testing"

func TestDuck_name(t *testing.T) {
	duck := &Duck{"taro"}
	actual := duck.name
	expected := "taro"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
