package main

import "fmt"

func main(){
	
	tampilkan_pesan1()
	tampilkan_pesan2()
	// tampilkan_pesan3()
}

func tampilkan_pesan1(){
	fmt.Println("Pesan Diterima 1")
}

func tampilkan_pesan2(){
	return "Pesan Diterima 2"
}