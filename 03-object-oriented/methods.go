package _3_object_oriented

import (
	"fmt"
	"golang-essential-training-course/03-object-oriented/structs"
	"os"
)

func TradeMethodsUsage() {
	t, err := structs.NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("ERROR: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}

func PointMethodsUsage() {
	p := &structs.Point{}
	p.Move(15, 15)
	fmt.Printf("%+v\n", p)
}
