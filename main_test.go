package main_test

import (
	"errors"
	"fmt"
	"testing"
)

func Pembagian(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Gak bisa dibagi 0")
	}
	return a / b, nil
}

func TestFunction(t *testing.T) {
	// bagi := Pembagian(8, 2)
	// assert.Equal(t, 4, bagi)

	// bagiKoma := Pembagian(9, 2)
	// assert.Equal(t, 4.5, bagiKoma)

	bagiNol, err := Pembagian(9, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bagiNol)
}
