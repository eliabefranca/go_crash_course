package main

import (
	"fmt"
	"math"

	"github.com/eliabefranca/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(29.127))
	fmt.Println(math.Ceil(29.127))
	fmt.Println(math.Sqrt(29.127))
	fmt.Println(strutil.Reverse("Eliabo franca"))
}
