package simple_factory

import (
	"fmt"
	"testing"
)

func TestNewShape(t *testing.T) {
	s := NewShape(1)

	fmt.Println(s.Draw())
}
