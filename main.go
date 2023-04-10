package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	name    string
	address string
	job     string
	reason  string
}

var students = []student{
	{name: "Radiffa", address: "Bekasi", job: "Penulis", reason: "Mencari Raden Saleh"},
	{name: "Fikri", address: "Depok", job: "Pebisnis", reason: "Mejadi Kaya"},
	{name: "Reza", address: "Bogor", job: "IT Support", reason: "Mencari Kesalahan"},
	{name: "Soni", address: "Tangerang", job: "Programmer", reason: "Menjadi Bug Hunter"},
}

func main() {
	var args = os.Args[1]
	intVar, _ := strconv.Atoi(args)

	if intVar > len(students) {
		fmt.Println("Tidak ada data")
	} else {
		findPerson(intVar)
	}
}

func findPerson(args int) {
	var i int = args - 1

	fmt.Println("Nama:", students[i].name)
	fmt.Println("Alamat:", students[i].address)
	fmt.Println("Pekerjaan:", students[i].job)
	fmt.Println("Alasan:", students[i].reason)
}
