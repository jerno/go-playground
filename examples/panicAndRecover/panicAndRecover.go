package panicAndRecover

import "fmt"

func Panic() {
	nums := []int{1}
	fmt.Printf("Trying to access second to last item of %#v\n", nums)
	fmt.Println(secondToLast(nums)) // will panic
}

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
