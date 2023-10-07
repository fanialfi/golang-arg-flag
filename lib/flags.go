package lib

import (
	"flag"
	"fmt"
)

func Flags() {
	// cara pertama
	name := flag.String("nama", "anonymous", "type your name")
	age := flag.Int("umur", 18, "type your age")

	// cara kedua
	var nama string
	flag.StringVar(&nama, "gender", "male", "masukkan nama")

	// memparsing semua command line flag from os.Args[1:]
	// flag.Parse()

	// gunakan
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
	fmt.Println("Halo, jenis kelamin saya saya", nama)
}
