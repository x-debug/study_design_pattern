package factory_method

import "fmt"

func Run(factory1 BaseFactory, factory2 BaseFactory) {
	item1 := factory1.Create()
	item2 := factory2.Create()

	item1.Jump()
	item2.Jump()
}

//抽象产品接口
type Item interface {
	Jump()
}

//抽象工厂接口
type BaseFactory interface {
	Create() Item
}

type Duck struct {
}

func (d Duck) Jump() {
	fmt.Println("duck jump")
}

type DuckFactory struct {
}

func (df DuckFactory) Create() Item {
	return &Duck{}
}

type Dog struct {
}

type DogFactory struct {
}

func (d Dog) Jump() {
	fmt.Println("dog jump")
}

func (df DogFactory) Create() Item {
	return &Dog{}
}
