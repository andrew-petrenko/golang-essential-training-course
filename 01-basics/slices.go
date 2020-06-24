package _1_basics

import "fmt"

func Slices() {
	brands := []string{"Google", "Apple", "Microsoft"}

	fmt.Printf("brands = %v (type %T)\n", brands, brands)

	fmt.Println(len(brands))

	fmt.Println("----")

	fmt.Println(brands[1])

	fmt.Println(brands[1:])

	fmt.Println("----")

	for i := 0; i < len(brands); i++ {
		fmt.Println(brands[i])
	}

	fmt.Println("----")

	for i := range brands {
		fmt.Println(i)
	}

	fmt.Println("----")

	for i, name := range brands {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")

	for _, name := range brands {
		fmt.Println(name)
	}

	fmt.Println("----")

	brands = append(brands, "Facebook")
	fmt.Printf("brands = %v", brands)
}
