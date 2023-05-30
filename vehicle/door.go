package vehicle

import "fmt"

type Door struct {
	Side string
}

func (d Door) Open(side string) {
	if side == "left" {
		fmt.Println("")
	} else {
		fmt.Println("")
	}
}

func (d Door) Knock(side string) {
	if side == "left" {
		fmt.Println("")
	} else {
		fmt.Println("")
	}
}
