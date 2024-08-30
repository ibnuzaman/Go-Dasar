package main

import (
	"fmt"
	"time"
)

func main() {

	// Run()
	Proses()

}

func Login() {
	fmt.Println("Berhasil Login")
}

func Loading() {
	defer Login()
	fmt.Println("Sedang proses")
}

func Masuk(masuk time.Time) {
	fmt.Print("Anda Login pada: ", time.DateOnly)
}

func Proses() {
	login := time.Now()
	defer Masuk(login)
	fmt.Println("Proses")
}

func Waktu(start time.Time) {
	fmt.Print("Total Waktu dalam detik: ", time.Since(start).Seconds())
}

func Run() {
	start := time.Now()
	defer Waktu(start)
	time.Sleep(2 * time.Second)
	fmt.Println("sleep")
}
