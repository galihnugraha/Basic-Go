package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Mahasiswa struct {
	Nama string
	Nim  string
	Ipk  int
}

func main() {
	var Data [99]Mahasiswa
	var menu int
	var banyak int
	var jumlah = 0
	var ulang1 = true
	var ulang2 string
	var ulang3 = true

	for ulang1 {

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Print("===MENU AWAL===\n")
		fmt.Print("1. Input Data\n")
		fmt.Print("2. Tampilkan Data\n")
		fmt.Print("3. Keluar\n")

		fmt.Print("\nPilih : ")

		fmt.Scan(&menu)

		switch menu {
		case 1:
			{
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()

				fmt.Print("===INPUT DATA===\n\n")

				fmt.Print("Banyak Data : ")
				fmt.Scan(&banyak)

				for i := 0; i < banyak; i++ {
					fmt.Print("\n--Data ", i+1, "--\n")
					fmt.Print("Nama : ")
					fmt.Scan(&Data[i+jumlah].Nama)
					fmt.Print("NIM  : ")
					fmt.Scan(&Data[i+jumlah].Nim)
					fmt.Print("IPK  : ")
					fmt.Scan(&Data[i+jumlah].Ipk)
				}

				jumlah = banyak + jumlah

				fmt.Print("\nInput Data Selesai\n")
			}
		case 2:
			{
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()

				fmt.Print("===DAFTAR DATA===\n")

				for i := 0; i < jumlah; i++ {
					fmt.Print("\n--Data ", i+1, "--\n")
					fmt.Print("Nama : " + Data[i].Nama + "\n")
					fmt.Print("NIM  : " + Data[i].Nim + "\n")
					fmt.Print("IPK  : ", Data[i].Ipk, "\n")
				}

				fmt.Print("\nSelesai\n")
			}
		case 3:
			{
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()

				os.Exit(0)
			}
		default:
			{
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()

				fmt.Print("input yang anda masukan salah.\n")
				fmt.Print("Silakan Ulangi.\n")
			}
		}

		ulang3 = true
		for ulang3 {

			fmt.Print("\nKembali ke MENU AWAL? (y/n) ")
			fmt.Scan(&ulang2)

			if ulang2 == "y" {
				ulang1 = true
				ulang3 = false
			} else if ulang2 == "n" {
				ulang1 = false
				ulang3 = false

				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
			} else {
				ulang3 = true
			}
		}
	}
}
