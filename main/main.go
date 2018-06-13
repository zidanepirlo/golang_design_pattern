package main

import (
	"golang_design_pattern/pattern"
	"fmt"
)

func testFactory()  {

	var pencilFactory pattern.PencilFactory
	pencilFactory.Produce().Write()
	fmt.Println(pencilFactory)
}

func testMyFactory()  {

	var computer *pattern.Computer



	fmt.Println(computer)


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
}
