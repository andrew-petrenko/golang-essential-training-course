package _3_object_oriented

import "fmt"
import "golang-essential-training-course/03-object-oriented/structs"

func StructUsage() {
	t1 := structs.Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := structs.Trade{
		Symbol: "AMZN",
		Volume: 15,
		Price:  74.59,
		Buy:    false,
	}
	fmt.Printf("%+v\n", t2)

	t3 := structs.Trade{}
	fmt.Printf("%+v\n", t3)
}
