package main

import (
	"fmt"
	"sync"
	"time"
)

func masak(menu string, durasi time.Duration, ch chan string, wg *sync.WaitGroup) {

	defer wg.Done() // Menandai bahwa goroutine ini telah selesai
	time.Sleep(durasi)

	ch <- fmt.Sprintf("Menu %s sudah siap dalam %v", menu, durasi) // Mengirim hasil ke channel
}

func mainLatihan() {
	var wg sync.WaitGroup
	ch := make(chan string)

	menu := []struct {
		nama   string
		durasi time.Duration
	}{
		{"Nasi Goreng", 2 * time.Second},
		{"Mie Goreng", 3 * time.Second},
		{"Ayam Bakar", 4 * time.Second},
	}

	wg.Add(len(menu))

	for _, m := range menu {
		go masak(m.nama, m.durasi, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch) // Menutup channel setelah semua goroutine selesai
	}()

	for hasil := range ch {
		fmt.Println(hasil)
	}

	fmt.Print("semua masakan selesai")
}
