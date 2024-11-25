package problem4

func iter(ch chan<- int, nums []int) {
	defer close(ch)
	for _, n := range nums {
		ch <- n
	}
}

func sum(nums []int) int {
	ch := make(chan int)

	go iter(ch, nums)

	var sum int
	sum = readChan(ch, sum)
	return sum
}

func readChan(ch chan int, sum int) int {
	for n := range ch {
		sum += n
	}
	return sum
}
