package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var resto, kasir string
	fmt.Print("Masukan nama resto : ")
	fmt.Scanln(&resto)
	fmt.Print("Masukan nama kasir : ")
	fmt.Scanln(&kasir)
	timeNow := time.Now()
	tgl := timeNow.Format("02/01/2006 15:04:05")
	var menu map[string]int = make(map[string]int)
	x := true
	for x {
		var opsi string
		fmt.Print("Tambah menu [1.Ya] [2.Tidak] : ")
		fmt.Scanln(&opsi)
		switch opsi {
		case "1":
			var nama, harga string
			fmt.Print("Nama Menu : ")
			fmt.Scanln(&nama)
			fmt.Print("Harga : ")
			fmt.Scanln(&harga)
			var hrg int
			hrg, _ = strconv.Atoi(harga)
			menu[nama] = hrg
		default:
			x = false
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Warung Makan", resto)
	fmt.Println("Tanggal : ", tgl)
	fmt.Print("Nama Kasir :")
	for i := 13; i <= 30-len(kasir); i++ {
		fmt.Print(" ")
	}
	fmt.Println(kasir)

	for i := 0; i <= 30; i++ {
		if i == 30 {
			fmt.Println("")
		} else {
			fmt.Print("=")
		}
	}

	// fmt.Println(menu)
	var total int
	for n, h := range menu {
		fmt.Print(n)
		str := strconv.Itoa(h)
		for i := len(n) + 1; i <= 27-len(str); i++ {
			fmt.Print(".")
		}
		fmt.Printf("Rp.%d\n", h)
		total = total + h
	}

	fmt.Print("Total")
	strTotal := strconv.Itoa(total)
	for i := 6; i <= 27-len(strTotal); i++ {
		fmt.Print(".")
	}
	fmt.Printf("Rp.%d\n", total)

	fmt.Println()
}
