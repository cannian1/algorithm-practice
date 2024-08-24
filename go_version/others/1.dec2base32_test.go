package others

import (
	"fmt"
	"testing"
)

func TestDecToBase32(t *testing.T) {
	res := DecToBase32(32)
	fmt.Println(res)
}
