package animals_test

import (
	"sample-go-testing/animals"
	"sample-go-testing/foods/mock_foods"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestDuck_Eat_02(t *testing.T) {
	// create mock instance
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	food := mock_foods.NewMockFood(ctrl)

	// defined return value
	food.EXPECT().Name().Return("kyougyoku")

	duck := animals.NewDuck("taro")
	actual := duck.Eat(food)
	expected := "taro ate kyougyoku"
	if actual != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}
