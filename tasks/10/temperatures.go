package main

import (
	"fmt"
	"sync"
)

func Linear(temperatures []float32) map[int][]float32 {
	groups := make(map[int][]float32)

	for _, value := range temperatures {
		key := (int(value) / 10) * 10
		groups[key] = append(groups[key], value)
	}

	return groups
}

func Parallel(temperatures []float32) map[int][]float32 {
	var wg sync.WaitGroup
	var mu sync.Mutex

	groups := make(map[int][]float32)

	for _, value := range temperatures {
		wg.Add(1)

		go func(value float32) {
			defer wg.Done()

			if value > 0 {
				for aboveZero := 0; aboveZero < 100; aboveZero += 10 {
					if aboveZero <= int(value) && int(value) < aboveZero+10 {
						mu.Lock()
						groups[int(aboveZero)] = append(groups[int(aboveZero)], value)
						mu.Unlock()

						break
					}
				}
			} else {
				for belowZero := 0; belowZero > -100; belowZero -= 10 {
					if belowZero >= int(value) && belowZero-10 < int(value) {
						mu.Lock()
						groups[int(belowZero)] = append(groups[int(belowZero)], value)
						mu.Unlock()

						break
					}
				}
			}
		}(value)
	}

	wg.Wait()

	return groups
}

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupsLinear, groupsParallel := Linear(temperatures), Parallel(temperatures)

	fmt.Printf("Linear  : %v\n", groupsLinear)
	fmt.Printf("Parallel: %v", groupsParallel)
}
