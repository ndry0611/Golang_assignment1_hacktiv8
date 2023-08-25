package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	ID 		  int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Absen(nama string, people []Person) {
	var orang Person
	for _, person := range people {
		if strings.ToLower(person.Nama) == strings.ToLower(nama) {
			orang = person
			break
		}
	}
	if orang.Nama == "" {
		fmt.Printf("Nama %s tidak ditemukan!\n", nama)
	} else {
		fmt.Println("ID :", orang.ID)
		fmt.Println("Nama :", orang.Nama)
		fmt.Println("Alamat :", orang.Alamat)
		fmt.Println("Pekerjaan :", orang.Pekerjaan)
		fmt.Println("Alasan :", orang.Alasan)
	}
}

func main() {
	var people = []Person {
		{
			ID: 0,
			Nama: "Thomas",
			Alamat: "Jalan A",
			Pekerjaan: "Programmer",
			Alasan: "Pemuja Winter Tech",
		},
		{
			ID: 1,
			Nama: "Andry",
			Alamat: "Jalan B",
			Pekerjaan: "Tutor",
			Alasan: "Saya suka mengajari, benci diajari",
		},
		{
			ID: 2,
			Nama: "Jason",
			Alamat: "Jalan C",
			Pekerjaan: "Siswa",
			Alasan: "Karna sebagian besar mahasiswa adalah entertainer Huwu",
		},
		{
			ID: 3,
			Nama: "Boulets",
			Alamat: "Jalan D",
			Pekerjaan: "Tentara",
			Alasan: "Berperang akan membutuhkan MORE BOULETS",
		},
	}

	var inputan = os.Args
	if len(inputan) >1 {
		Absen(inputan[1], people)
	} else {
		fmt.Println("Isi nama yang ingin diabsensi!")
	}
}
