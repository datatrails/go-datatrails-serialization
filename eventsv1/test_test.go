package eventsv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {

	foo := Foo()
	assert.Equal(t, true, foo)

}
