package channels

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Printf("Creating channels\n")

	ch := make(chan int)
	chs := make(chan string)

	fmt.Println("")
	fmt.Println("----- Example with a single value -----")
	fmt.Println("")

	fmt.Printf("[Provider] Define go routine\n")
	go func() {
		val := 353
		fmt.Printf("[Provider] Constructing value and sending to channel: %d\n", val)
		// Send number of the channel
		ch <- val
	}()

	fmt.Printf("[Consumer] Requesting value from channel\n")
	val := <-ch // receive
	fmt.Printf("[Consumer] Got value from channel: %d\n", val)

	fmt.Println("")
	fmt.Println("----- Example with a multiple integer values: 1, 2, 3 -----")
	fmt.Println("")

	// Send multiple
	const count = 3
	go func() {
		for i := 1; i <= count; i++ {
			fmt.Printf("[Provider] Sending value to channel %d\n", i)
			ch <- i
			fmt.Printf("[Provider] Value %d picked up from channel, sleeping...\n", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		fmt.Printf("[Consumer] Awaiting value from channel\n")
		val := <-ch
		fmt.Printf("[Consumer] Got value from channel: %d\n", val)
	}

	fmt.Println("")
	fmt.Println("----- Example with a multiple string values: 1️⃣ , 2️⃣ , 3️⃣ -----")
	fmt.Println("")

	values := []string{"1️⃣ ", "2️⃣ ", "3️⃣ "}
	// close to signal we're done
	go func() {
		for i := 0; i < len(values); i++ {
			fmt.Printf("[Provider] Sending value to channel %s\n", values[i])
			chs <- values[i]
			fmt.Printf("[Provider] Value %s picked up from channel, sleeping...\n", values[i])
			time.Sleep(time.Second)
		}

		fmt.Printf("[Provider] Closing channel\n")
		close(chs)
	}()

	fmt.Printf("[Consumer] Awaiting value from channel\n")
	for i := range chs {
		fmt.Printf("[Consumer] Got value from channel %s\n", i)
	}

	fmt.Println("")
	fmt.Println("-----")
	fmt.Println("")

	fmt.Printf("[Provider] Creating buffered channel with size 1\n")

	chb := make(chan int, 1) // buffered channel

	fmt.Printf("[Provider] Sending value 19\n")
	chb <- 19

	fmt.Printf("[Consumer] Awaiting value\n")
	val2 := <-chb

	fmt.Printf("[Consumer] Got value from channel %d\n", val2)
	fmt.Println(val2)
}
