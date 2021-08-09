package main

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2021-12-13"
	var url = " Code = %s & endDate = %d "
	var target_url = fmt.Sprintf(url, enddate, stockcode)
	fmt.Println(target_url)
	//如果我想表示1.020怎么办
	var a float32 = 1.00120
	fmt.Println(a)

}
