package user

import (
	"github.com/stretchr/testify/assert"
	"mock/mocks"
	"testing"
)

func TestUser_Walk(t *testing.T) {
	u := &mocks.Man{}
	u = NewUser(u)
	u.(*mocks.Man).On("Walk", "Hello").Return("Hello")
	u.(*mocks.Man).On("Walk", "HeyYou").Return("Hia")
	walkOk := u.Walk("Hello")
	assert.Equal(t, "Hello", walkOk)
}

func TestUser_Talk(t *testing.T) {

}
