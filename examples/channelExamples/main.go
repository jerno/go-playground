package channelExamples

import (
	"fmt"
	"time"
)

func ChannelExample() {
	ch := make(chan int)
	chs := make(chan string)

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	val := <-ch // receive
	fmt.Printf("got %d\n", val)

	fmt.Println("-----")

	// Send multiple
	const count = 3
	go func() {
		for i := 1; i <= count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

	values := []string{"one ", " two", "three"}
	// close to signal we're done
	go func() {
		for i := 0; i < len(values); i++ {
			fmt.Printf("sending %s\n", values[i])
			chs <- values[i]
			time.Sleep(time.Second)
		}
		close(chs)
	}()

	for i := range chs {
		fmt.Printf("received %s\n", i)
	}

	fmt.Println("-----")

	chb := make(chan int, 1) // buffered channel
	chb <- 19
	val2 := <-chb
	fmt.Println(val2)
}
