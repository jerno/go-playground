# go-playground

An interactive playground with various examples in GO

## Usage

```
$ go run ./playground.go
=======================================================================
Available use cases:
   1) 🌐 HTTP GET String data
   2) 🌐 HTTP GET JSON data
   3) 🌐 HTTP POST JSON data
   4) 🏢 Start HTTP server
   5) 🕦 Get current time formatted
   6) ⌨️ Calculator
   7) 📝 Receiver passed by value
   8) 📄 Receiver passed by reference
   9) 🔗 Modifying variable through a pointer
  10) 🔎 Map checking missing item (A.K.A. CommaOK)
  11) 😱 Panic
  12) 🤕 Panic and recover
  13) 📅 Get next business day
  14) ⏳ Time and duration
  15) ⏳ Sync WaitGroup
  16) 📻 Channels
  17) 🪟 Window powered by astilectron
Please select a use case (type its number or its name):
```

## Examples

There are a couple of examples covering the language basics, as well as some advanced concepts I was experimenting while working with GO

### HTTP GET String data

Using `justHttp ("github.com/jerno/just-http/basic")` – a simple HTTP client I made – this example fetches some data as a string from a public HTTP API.

### HTTP GET JSON data

Using `justHttp` this example fetches some data by sending a HTTP GET request, with additional query params, and returns it as a parsed JSON data.

### HTTP POST JSON data

Using `justHttp` this example that sends a HTTP POST request, with additional query params, a `map[string]string` body, and returns a parsed JSON data.

### Start HTTP server

This example starts up a simple HTTP server. The available endpoints are below:

```log
Available endpoints:
  /health
    curl http://localhost:8081/health
  /math
    curl -d '{"Left": 91, "Right": 8, "Op": "+"}' -H "Content-Type: application/json" -X POST http://localhost:8081/math
  /vue/
    curl http://localhost:8081/vue/
  /animals
    curl http://localhost:8081/animals
  /animals/refresh
    curl http://localhost:8081/animals/refresh
```

- The `/health` endpoint simply returns an OK message.

- The `/math` endpoint features a simple calculator where you can provide an operator (`+-*/`) plus two operands, and you can get the result in the HTTP payload.

- Under `/animals` there was a proxy to the publicly available zoo API (https://zoo-animal-api.herokuapp.com/animals)

  > 🚧 Unfortunatelly the Zoo animal API is down at the moment

- Under the `/vue` endpoint, you can find a simple Vue.js application that displays a zoo, built on top of the `/animals` API.
  > 🚧 Unfortunatelly the Zoo animal API is down at the moment

### Get current time formatted

Shows how to use the built-in formatter to format a date time object

### Calculator

A simple calculator where you can provide an operator (`+-*/`) plus two operands, and you can get back the result

### Receiver passed by value

> 💡 In GO, since there are no traditional OOP classes, you can create "instance methods" like mechanisms by declaring receivers on a custom type (struct).

In this example there is a Dog struct, that features a `Sound` property.

```golang
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}
```

There is a **value receiver** named `Speak` that interacts with the dog's data, and logs the `Sound` property's value. Here the receiver operates on a copy of the original dog object.

```golang
func (d Dog) Speak() {
	fmt.Printf(d.Sound)
}
```

### Receiver passed by reference

Same example as the one above, with one main difference: instead of a **value receiver** I am using a **pointer receiver** that operates on the given instance of a dog, instead of the copy.

### Modifying variable through a pointer

Basic example that shows how the pointers are woring in _golang_. By passing a pointer (referencing an other variable), to an other function, we can actually modify the value of the original variable.

### Map checking missing item (A.K.A. CommaOK)

When accessing items of a map (e.g `map[string]int`), we can use the index operator `prices["Apple"]`. By default, if the element does not exists, GO returns the default value of a given type, in our case `0` for an uninitialized `int`.

If we want to know whether the given key exists in a map, we have a mechanism called `comma ok`. When accessing elements of a map, GO can return a second variable called `ok` along with the value. If we do so, we can have the information about the presence of the element.

```golang
price, ok := prices["Apple"]
```

### Panic

A simple example demonstrating what happens when we have an "unhandled exception" - a Panic when running the application. In this example I am simply trying to access a non-existing element in an array, resulting in a Panic.

```golang
func Panic() {
	nums := []int{1}
	fmt.Printf("Trying to access second to last item of %#v\n", nums)
	fmt.Println(secondToLast(nums)) // will panic
}

func secondToLast(nums []int) int {
	return nums[len(nums)-2]
}
```

### Panic and recover

As a continuation of the previous example, I added a new function `safeSecondToLast`, where I demonstrate the recover mechanism of GO.

```golang
func PanicWithRecover() {
	nums := []int{1}
	fmt.Printf("Trying to access second to last item of %#v\n", nums)
	result, err := safeSecondToLast(nums)
	fmt.Printf("Result: %v, error: %v\n", result, err)
}

func safeSecondToLast(nums []int) (i int, err error) {
	defer func() {
		if e := recover(); e != nil { // e is interface{}
			err = fmt.Errorf("%v", e)
		}
	}()

	return secondToLast(nums), nil
}

func secondToLast(nums []int) int {
	return nums[len(nums)-2]
}
```

### Get next business day

Simply shows the next business day

### Time and duration

A converter that can covert between timezones, featuring GO's built-in time handling utilities.

### Sync WaitGroup

By creating a sample HTTP server with a 200ms latency, this example compares a syncronous, and an asyncronous execution of 5 HTTP requests, demonstrating `Channels`, and `WaitGroups`.

### Channels

This example is using two channels, an `int` one, and a `string` one, and shows how a _Provider_ and a _Consumer_ can interact with each other.

- The _Provider_ is creating resources in a Goroutine
- The _Consumer_ is reading the channel using the loop `for i := range c` syntax

> 💡 A goroutine is a lightweight thread managed by the Go runtime.

### Window powered by astilectron

Using `go-astikit` and `go-astilectron`, this example shows how to craete a simple electron application.

- It opens a new window, displaying a local html file as a content.
- It shows how to communicate from the direction of GO towards the Javascript application
- Shows how to communicate from the direction of the Javascript application towards the GO app
