package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var password = flag.String("password", "", "password")
var cost = flag.Int("cost", 10, "cost")

func main() {
	flag.Parse()

	if *password == "" {
		fmt.Println("Usage: bcrypt -password=123456 -cost=10")
		flag.PrintDefaults()
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), *cost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(hashedPassword))

}
