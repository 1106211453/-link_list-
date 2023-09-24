package main

import "fmt"

func Ended() {
	var sfd [5]int
	sfd[0] = 1
	Sad(sfd[:])
	fmt.Print(sfd[0])
}
func Sad(sfd []int) {
	sfd[0] = 2
}
