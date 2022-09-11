package receiverPassedByValue

import "fmt"

func Run() {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println("Creating dog...")
	fmt.Printf("Dog: %+v\n", poodle)

	fmt.Printf("Make it speak: ")
	poodle.Speak()
	fmt.Println("")

	fmt.Println("Changing its sound...")
	poodle.Sound = "Arf!"

	fmt.Printf("Make it speak: ")
	poodle.Speak()
	fmt.Println("")

	fmt.Printf("Calling SpeakThreeTimes where we change its sound: ")
	poodle.SpeakThreeTimes()
	fmt.Println("")

	fmt.Printf("Calling SpeakThreeTimes where we change its sound: ")
	poodle.SpeakThreeTimes()
	fmt.Println("")
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Printf(d.Sound)
}

// SpeakThreeTimes is how the dog speaks loudly
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Printf(d.Sound)
}
