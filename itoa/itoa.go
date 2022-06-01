package main

import "strconv"

func main() {
	d, e := strconv.ParseInt("1,001,001", 10, 32)
	println("d", d, "error", e)
}
