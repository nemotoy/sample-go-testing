package animals_test

import (
	"sample-go-testing/animals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuck_Say(t *testing.T) {
	duck := animals.NewDuck("taro")
	actual := duck.Say()
	expected := "taro says quack"
	assert.Equal(t, expected, actual)
}
