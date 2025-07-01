package packer

import (
	"bakery/model"
	"math/rand"
	"sync"
	"time"
)

func Packer(id, t2 int, in <-chan model.Cake, out chan<- model.Cake, wg *sync.WaitGroup) {
	defer wg.Done()

	for cake := range in {
		timeSpent := rand.Intn(t2*2) + 1
		time.Sleep(time.Duration(timeSpent) * time.Millisecond)

		cake.PackedBy = id
		cake.PackTime = timeSpent

		out <- cake
	}
}
