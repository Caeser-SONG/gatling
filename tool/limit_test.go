package tool

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimiter(1)
	var count int
	for i := 0; i < 120; i++ {
		if l.Allow() {
			count++
			fmt.Printf("success  count = %d", count)

		} else {
			fmt.Println("failed")
		}

	}

}
