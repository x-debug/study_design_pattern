package factory_method

import "testing"

func TestRun(t *testing.T) {
	Run(&DuckFactory{}, &DogFactory{})
}
