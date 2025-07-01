package packer

import (
	"bakery/model"
	"context"
	"math/rand"
	"sync"
	"time"
)

func Packer(ctx context.Context, id, t2 int, in <-chan model.Cake, out chan<- model.Cake, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():

			return
		case cake, ok := <-in:

			if !ok {
				return
			}

			timeSpent := rand.Intn(t2*2) + 1
			time.Sleep(time.Duration(timeSpent) * time.Millisecond)

			cake.PackedBy = id
			cake.PackTime = timeSpent

			out <- cake
		}
	}
}
