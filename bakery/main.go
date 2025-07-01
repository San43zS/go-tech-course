package main

import (
	"bakery/baker"
	"bakery/model"
	"bakery/packer"
	"fmt"
	"sync"
)

func main() {
	K := 100
	N := 8
	M := 5
	t1 := 100
	t2 := 150

	bakingChan := make(chan model.Cake, K)
	packingChan := make(chan model.Cake, K)

	var wgBaker sync.WaitGroup
	var wgPacker sync.WaitGroup

	cakeTasks := make(chan int, K)
	for i := 0; i < K; i++ {
		cakeTasks <- i
	}
	close(cakeTasks)

	for i := 0; i < N; i++ {
		wgBaker.Add(1)
		go baker.Baker(i, t1, cakeTasks, bakingChan, &wgBaker)
	}

	go func() {
		wgBaker.Wait()
		close(bakingChan)
	}()

	for i := 0; i < M; i++ {
		wgPacker.Add(1)
		go packer.Packer(i, t2, bakingChan, packingChan, &wgPacker)
	}

	go func() {
		wgPacker.Wait()
		close(packingChan)
	}()

	var cakes []model.Cake
	for cake := range packingChan {
		cakes = append(cakes, cake)
	}

	for i, cake := range cakes {
		fmt.Printf("Торт %d: Выпек %d за %dмс, Упаковал %d за %dмс\n",
			i+1, cake.BakedBy, cake.BakeTime, cake.PackedBy, cake.PackTime)
	}
}
