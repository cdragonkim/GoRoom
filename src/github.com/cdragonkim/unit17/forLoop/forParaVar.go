package forLoop

import (
	"fmt"
)

func ForParaVar1() {
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}