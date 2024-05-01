package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	queue := make(chan string, len(emails))

	for _, i := range emails {
		fmt.Println("this is i %s", i)
		queue <- i
	}

	return queue
}
