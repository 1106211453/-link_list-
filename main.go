package main

import "fmt"

func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31 %d"
	var url = "Code= %s &endDate= %d "
	var target_url = fmt.Sprintf(url, enddate, stockcode)
	fmt.Println(target_url)
}
