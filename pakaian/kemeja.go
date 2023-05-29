package pakaian

import "fmt"

type Kemeja struct {
	Ukuran int
}

func (k Kemeja) Pakai() {
	fmt.Println("Pakai Kemeja")
}

func (k Kemeja) Cuci() {
	fmt.Println("Cuci Kemeja")
}

type Kaos struct {
	Ukuran int
}

func (k Kaos) Pakai() {
	fmt.Println("Pakai Kaos")
}

type Sweater struct {
	Ukuran int
}

func (k Sweater) Pakai() {
	fmt.Println("Pakai Swater")
}
