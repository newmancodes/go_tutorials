package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type engine interface {
	milesLeft() uint8
}

func main() {
	var myGasEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myGasEngine.mpg, myGasEngine.gallons)
	fmt.Printf("Total miles left in the tank: %v.\n", myGasEngine.milesLeft())
	canMakeIt(myGasEngine, 50)

	var myElectricEngine = electricEngine{mpkwh: 15, kwh: 10}
	fmt.Println(myElectricEngine.mpkwh, myElectricEngine.kwh)
	fmt.Printf("Total miles left in the battery: %v.\n", myElectricEngine.milesLeft())
	canMakeIt(myElectricEngine, 50)
}
