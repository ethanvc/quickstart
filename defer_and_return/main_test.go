package defer_and_return

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeferAndReturn(t *testing.T) {
	assert.Equal(t, "RealResult", f1().Error())
}

func f1() (err error) {
	defer func(err *error) {
		*err = errors.New("RealResult")
	}(&err)
	return errors.New("Result")
}
