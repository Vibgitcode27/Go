package main

import "fmt"

func main() {

	emails := []string{
		"Email1",
		"Email2",
		"Email3",
	}

	ok := addEmailsToQueue(emails)

	for i := 0; i < len(ok); i++ {
		email := <-ok
		fmt.Println(email)
	}
}
