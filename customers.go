package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cmmentMap := make(map[int][]int)
	conterNum := 3
	customerConter := 5
	for ID := 1; ID <= conterNum; ID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for customerID := (id-1)*customerConter + 1; customerID <= conterNum; customerID++ {

				mu.Lock()
				cmmentMap[id] = append(cmmentMap[id], ID)
				mu.Unlock()
				fmt.Printf("柜台%d正在处理顾客%d\n", id, customerID)
			}
		}(ID)
	}
	wg.Wait()
	fmt.Println("最终结果：%v\n", cmmentMap)
}
