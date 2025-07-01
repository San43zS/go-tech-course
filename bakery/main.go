package main

import (
	"bakery/baker"
	"bakery/model"
	"bakery/packer"
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	bakingChan := make(chan model.Cake, model.K)
	packingChan := make(chan model.Cake, model.K)

	var wgBaker sync.WaitGroup
	var wgPacker sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cakeTasks := make(chan int, model.K)
	for i := 0; i < model.K; i++ {
		cakeTasks <- i
	}
	close(cakeTasks)

	for i := 0; i < model.N; i++ {
		wgBaker.Add(1)
		go baker.Baker(ctx, i, model.T1, cakeTasks, bakingChan, &wgBaker)
	}

	go func() {
		wgBaker.Wait()
		close(bakingChan)
	}()

	for i := 0; i < model.M; i++ {
		wgPacker.Add(1)
		go packer.Packer(ctx, i, model.T2, bakingChan, packingChan, &wgPacker)
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

	<-ctx.Done()
	fmt.Println("Завершение по сигналу...")
}
