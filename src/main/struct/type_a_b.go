package main

import "fmt"

func main() {
	fakeFish := FakeFish{}
	// 不用了Fish的方法
	// fakeFish.swim()
	fakeFish.FakeSwim()

	// 类型转换
	fakeFishToSwim := Fish(fakeFish)
	fakeFishToSwim.Swim()

	strongFish := StrongFish{}
	strongFish.Swim()

	strongFishCover := Fish(strongFish)
	strongFishCover.Swim()
}

type FakeFish Fish

func (F FakeFish) FakeSwim() {
	fmt.Printf("我是甲鱼，哈哈哈\n")
}

type StrongFish Fish

func (f StrongFish) Swim() {
	fmt.Printf("我是肌肉鱼，哦吼吼吼\n")
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼本鱼，会游哦\n")
}
