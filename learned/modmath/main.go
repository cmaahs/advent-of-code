package main

import "fmt"

func main() {
	i := -20 % 40
	fmt.Printf("Mod -20 %% 40 : %d\n", i)
	i = -19 % 40
	fmt.Printf("Mod -19 %% 40 : %d\n", i)
	i = 0 % 40
	fmt.Printf("Mod 0 %% 40 : %d\n", i)
	i = 40 % 40
	fmt.Printf("Mod 40 %% 40 : %d\n", i)
	i = 80 % 40
	fmt.Printf("Mod 80 %% 40 : %d\n", i)

}
