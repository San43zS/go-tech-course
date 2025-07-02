package baker

import (
	"bakery/model"
	"context"
	"math/rand"
	"sync"
	"time"
)

func Baker(ctx context.Context, id, t1 int, in <-chan int, out chan<- model.Cake, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case cakeID, ok := <-in:
			if !ok {
				return
			}

			timeSpent := rand.Intn(t1*2) + 1
			time.Sleep(time.Duration(timeSpent) * time.Millisecond)

			out <- model.Cake{ID: cakeID, BakedBy: id, BakeTime: timeSpent}
		}
	}
}
