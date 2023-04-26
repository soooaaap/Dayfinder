package dayfinder

import (
	"fmt"
	"time"
)

func day60daysago() {
	before60 := time.Now().AddDate(0, 0, -60)
	y, m, d := before60.Date()
	time60 := fmt.Sprint(y, "-", int(m), "-", d)
	fmt.Println(time60)
}
