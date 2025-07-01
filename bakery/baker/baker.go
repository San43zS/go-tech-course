package baker

import (
	"bakery/model"
	"math/rand"
	"sync"
	"time"
)

func Baker(id, t1 int, in <-chan int, out chan<- model.Cake, wg *sync.WaitGroup) {
	defer wg.Done()

	for cakeID := range in {
		timeSpent := rand.Intn(t1*2) + 1
		time.Sleep(time.Duration(timeSpent) * time.Millisecond)

		out <- model.Cake{ID: cakeID, BakedBy: id, BakeTime: timeSpent}
	}
}
