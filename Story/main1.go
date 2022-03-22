package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var resto, kasir string

	// input nama resto
	fmt.Print("Masukan nama resto : ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	resto = scanner1.Text()

	// input nama kasir
	fmt.Print("Masukan nama kasir : ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	kasir = scanner2.Text()

	// tanggal
	timeNow := time.Now()
	tgl := timeNow.Format("02/01/2006 15:04:05")
	var menu map[string]int = make(map[string]int)
	x := true
	for x {
		var opsi string
		fmt.Print("Tambah menu [ya]/[exit] : ")
		fmt.Scanln(&opsi)
		switch opsi {
		case "ya":
			var nama, harga string
			// fmt.Print("Nama Menu : ")
			// fmt.Scanln(&nama)
			// input nama menu
			fmt.Print("Nama Menu : ")
			scanner3 := bufio.NewScanner(os.Stdin)
			scanner3.Scan()
			nama = scanner3.Text()

			// input harga menu
			fmt.Print("Harga : ")
			fmt.Scanln(&harga)

			var hrg int
			hrg, _ = strconv.Atoi(harga)
			menu[nama] = hrg
		case "exit":
			x = false
		default:
			fmt.Println("Keluar ketik [exit]")
			log.Fatal("Input Salah")
		}
	}

	fmt.Println()
	fmt.Println()

	// Cetak Resto
	// fmt.Println("Warung Makan", resto)
	if len(resto) > 30 {
		fmt.Println(insertLine(resto, 30))
	} else {
		fmt.Println(resto + strings.Repeat(" ", 30-len(resto)))
	}

	// Cetak Tanggal
	fmt.Println("Tanggal : ", tgl)

	// Cetak Kasir
	if len("Nama Kasir : "+kasir) > 30 {
		fmt.Println(insertLine("Nama Kasir : "+kasir, 30))
	} else {
		fmt.Println("Nama Kasir : " + strings.Repeat(" ", 30-len("Nama Kasir : ")-len(kasir)) + kasir)
	}

	fmt.Println(strings.Repeat("=", 30))

	// fmt.Println(menu)
	var total int
	for n, h := range menu {
		str := strconv.Itoa(h)
		// fmt.Print(n)
		// for i := len(n) + 1; i <= 27-len(str); i++ {
		// 	fmt.Print(".")
		// }
		// fmt.Printf("Rp.%d\n", h)
		leng := int(math.Abs(float64(30 - (len(n) % 30) - len(str) - 3)))
		buildString := n + strings.Repeat(".", leng) + "Rp." + str
		if len(buildString) == 30 {
			fmt.Println(buildString)
		} else {
			fmt.Println(insertLine(buildString, 30))
		}

		total = total + h
	}

	strTotal := strconv.Itoa(total)
	// fmt.Print("Total")
	// for i := 6; i <= 27-len(strTotal); i++ {
	// 	fmt.Print(".")
	// }
	// fmt.Printf("Rp.%d\n", total)
	fmt.Println("Total" + strings.Repeat(".", 22-len(strTotal)) + "Rp." + strTotal)

	fmt.Println()
}

func insertLine(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}
