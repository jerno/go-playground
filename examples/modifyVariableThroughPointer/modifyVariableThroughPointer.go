package modifyVariableThroughPointer

import "fmt"

func Run() {
	fmt.Println("Setting V [V := 14]")
	V := 14
	fmt.Println("Creating pointer: PTR --> V [PTR := &V]")

	var PTR *int

	PTR = &V
	fmt.Printf("Reading PTR: %v, reading V: %v\n", *PTR, V)

	modifyVariableThroughPointer(PTR)

	fmt.Printf("Reading PTR: %v, reading V: %v\n", *PTR, V)
}

func modifyVariableThroughPointer(PTR *int) {
	fmt.Println("Setting V2 [V2 := 55]")
	V2 := 55
	fmt.Println("Changing the value through the pointer to V2 [*PTR = V2]")
	*PTR = V2
}
