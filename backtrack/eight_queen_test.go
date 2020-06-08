package backtrack

import (
	"fmt"
	"testing"
)

func TestPrintQueen(t *testing.T) {
	Set8Queen(0)
	fmt.Println("---------")
	printQueen()
}
