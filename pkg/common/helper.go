package common

import "fmt"

type Helper struct{}

func (h *Helper) testFunc() {
	fmt.Println("helper test func")
}
