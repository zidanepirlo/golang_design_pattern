package main

import (
	"golang_design_pattern/pattern"
	"golang_design_pattern/mypattern"
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func testBooks()  {

	var Book1 Books
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	fmt.Println(Book1)

	Book2 := new(Books)

	fmt.Println(Book2)
	fmt.Println(*Book2)
}

func testFactory()  {

	var pencilFactory pattern.PencilFactory
	pencilFactory.Produce().Write()
	fmt.Println(pencilFactory)
}

func modifyComputer(computer mypattern.Computer)  {
	computer.SetOs("win7")
}

func testMyFactory()  {

	computer := mypattern.Produce(mypattern.Mac{})

	//fmt.Println(computer)
	//modifyComputer(computer)
	//fmt.Println(computer)
	fmt.Println(computer.Configuration())

}

func main() {

	//var battery pattern.RechargeableBattery
	//
	//battery = pattern.AdapterNonToYes{pattern.NonRechargeableA{}}
	//battery.Use()
	//battery.Charge()
	//
	//battery = pattern.NonRechargeableB{}
	//battery.Use()
	//battery.Charge()

	//testFactory()
	testMyFactory()
	//testBooks()
}
