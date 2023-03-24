package main

import (
	"fmt"
	"strings"
)
func main() {
	var product = []string{
		"Pensil",
		"Bulpoin",
		"Buku",
	}
	var price = []int{
		2500,
		3000,
		5000,
	}

	for {
		var total = 0
		for {
			fmt.Println("Selamat datang di toko saya ")
			for i := 0; i < len(product); i++ {
				fmt.Println(i+1, product[i], "Rp.", price[i])
			}
			var choiceProduct int
			fmt.Print("Masukkan no barang yang ingin anda beli ? : ")
			fmt.Scan(&choiceProduct)
			fmt.Println("Anda memilih ", product[choiceProduct-1], "dengan harga ", price[choiceProduct-1])
			total = price[choiceProduct-1] + total
			fmt.Print("Apakah anda ingin menambah barang lagi ? (y/t) : ")
			var choice string
			fmt.Scan(&choice)
			if strings.ToLower(choice) == "t" {
				break
			}
		}
		fmt.Println("Total pembelian anda Rp.", total)
		var money int
		for {
			fmt.Print("Masukkan uang anda : ")
			fmt.Scan(&money)
			var change = money - total
			if total > money {
				fmt.Println("Maaf uang anda kurang Rp.", change)
			}
			if total <= money {
				fmt.Println("Kembaliannya adalah Rp.", change)
				break
			}
		}
		var choiceTransaction string
		fmt.Print("mau melakukan transaksi lagi y/t ? \t: ")
		fmt.Scan(&choiceTransaction)
		if strings.ToLower(choiceTransaction) == "t" {
			break
		}
	}
	fmt.Println("Terima Kasih dan sampai jumpa")
}
