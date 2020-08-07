package facade

import "fmt"

func Api() Facade {
	return &facadeImpl{&aModuleImpl{}, &bModuleImpl{}, &cModuleImpl{}}
}

type Facade interface {
	work()
}

type facadeImpl struct {
	a AModule
	b BModule
	c CModule
}

func (f facadeImpl) work() {
	//假设这里的组合非常复杂,而如果这段复杂的逻辑放在客户端去完成,客户端有可能会出错
	//所以用外观模式,来简化一些复杂的工作

	f.a.work1()
	f.c.work5()
	f.b.work3()
	f.a.work2()
	f.c.work6()
	f.b.work4()
}

type AModule interface {
	work1()
	work2()
}

type aModuleImpl struct {
}

func (a aModuleImpl) work1() {
	fmt.Println("a module work1")
}

func (a aModuleImpl) work2() {
	fmt.Println("b module work2")
}

type BModule interface {
	work3()
	work4()
}

type bModuleImpl struct {
}

func (a bModuleImpl) work3() {
	fmt.Println("a module work3")
}

func (a bModuleImpl) work4() {
	fmt.Println("b module work4")
}

type CModule interface {
	work5()
	work6()
}

type cModuleImpl struct {
}

func (a cModuleImpl) work5() {
	fmt.Println("a module work5")
}

func (a cModuleImpl) work6() {
	fmt.Println("b module work6")
}
