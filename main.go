package main

import "fmt"

func main() {
	var username string
	var password string

	fmt.Println("Enter your username")
	fmt.Scan(&username)
	fmt.Println("username:", username)

	fmt.Println("Enter your password")
	fmt.Scan(&password)
	fmt.Println("password:", password)

	var all = username + " " + password

	fmt.Println(all)
}
