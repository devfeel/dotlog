package layout

import (
	"fmt"
	"testing"
)

func Test_CompileLayout(t *testing.T) {
	express := "{datetime} {message} {year}/{month}/{day}"
	val := CompileLayout(express)
	fmt.Println("express : ", express)
	fmt.Println("convert : ", val)
}
