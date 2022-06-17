package main

import "fmt"

// 接口是一种类型,一般以er结尾
// 在编程中有一种场景，我不关心一个变量是什么类型，我只关心能调用它的什么方法
// 网上商城可能使用支付宝、微信、银联等方式去在线支付，他们是不同的变量，但都具有支付的方法
// 一个变量如果实现了接口中规定的所有方法，那么这个变量就成为了这个接口类型
type speaker interface { //只要是实现了speak方法的变量都是speaker类型
	speak()
}
type dog struct {
}
type cat struct {
}

func (d *dog) speak() {
	fmt.Println("汪汪汪")
}
func (c *cat) speak() {
	fmt.Println("喵喵喵")
}
func play(x speaker) {
	x.speak()
}

// 示例：不管是什么牌子车都能跑
type car interface {
	run()
}
type falali struct {
	brand string
}
type benz struct {
	brand string
}

func (f *falali) run() {
	fmt.Printf("%v会跑\n", f.brand)
}
func (b *benz) run() {
	fmt.Printf("%v会跑\n", b.brand)
}
func drive(c car) {
	c.run()
}

func main() {
	var (
		c1 *cat
		d1 *dog
	)
	play(c1)
	play(d1)
	f1 := &falali{
		brand: "法拉利",
	}
	b1 := &benz{
		brand: "奔驰",
	}
	drive(f1)
	drive(b1)

}
