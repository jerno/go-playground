package modifyVariableThroughPointer

import "fmt"

func Run() {
	modifyVariableThroughPointer()
}

func modifyVariableThroughPointer() {
	fmt.Println("Setting V [V := 14]")
	V := 14
	fmt.Println("Creating pointer: PTR --> V [PTR := &V]")
	PTR := &V
	fmt.Printf("Reading PTR: %v, reading V: %v\n", *PTR, V)
	fmt.Println("Setting V2 [V2 := 55]")
	V2 := 55
	fmt.Println("Changing the value through the pointer to V2 [*PTR = V2]")
	*PTR = V2
	fmt.Printf("Reading PTR: %v, reading V: %v\n", *PTR, V)
}
