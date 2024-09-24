package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func cpuHeavyWork() {
	// Loop untuk membuat beban CPU
	for i := 0; i < 1000000000000; i++ {
		math.Sqrt(float64(i)) // Komputasi berat untuk mengisi CPU
	}
}

func main() {
	// Menggunakan semua core yang tersedia
	numCPU := runtime.NumCPU()
	fmt.Printf("Menggunakan %d CPU\n", numCPU)
	runtime.GOMAXPROCS(numCPU)

	// Menjalankan beban kerja di banyak goroutine
	for i := 0; i < numCPU; i++ {
		go cpuHeavyWork()
	}

	// Memberi waktu untuk goroutine berjalan
	time.Sleep(10 * time.Second)

	fmt.Println("Selesai!")
}
