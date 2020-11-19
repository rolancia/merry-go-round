package main

import (
	"fmt"
	merry_go_round "merry-go-round"
	"sync"
)

func main() {
	i := 0
	pool := merry_go_round.NewPool(func() interface{} {
		rs := i
		i++
		return rs
	}, 64)

	result := map[int]int{}
	muResult := sync.Mutex{}
	wg := sync.WaitGroup{}
	shouldBe := 1024 * 64
	wg.Add(shouldBe)
	for i := 0; i < shouldBe; i++ {
		go func(result map[int]int) {
			rs := pool.Get().(int)
			pool.Put(rs)

			muResult.Lock()
			result[rs]++
			muResult.Unlock()

			wg.Done()
		}(result)
	}

	wg.Wait()
	resultSum := 0
	for _, num := range result {
		resultSum += num
	}

	fmt.Println("should:", shouldBe)
	fmt.Println("result:", resultSum)
}
