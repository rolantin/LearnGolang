package main

import "fmt"

type foodData struct {
	fish int
	fishPrice int
	apple int
	applePrice int
}

func (f *foodData) setData() {
	f.fish=10
	f.fishPrice=4
	f.applePrice=3
	f.apple =3
	//return f
}

func (f foodData) tollaPrice() int  {
	f.setData()
	a:= f.apple*f.applePrice
	b:= f.fishPrice*f.fish
	return a+b
}

func main()  {
    var food *foodData = new(foodData)
	//food := new(foodData)
	fmt.Println(food.tollaPrice())

	//fmt.Println(foodData.tollaPrice(foodData{2,2,2,2}))
}
