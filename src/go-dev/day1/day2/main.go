package main

import (
	"fmt"
	"go-dev/day1/day2/ch4"
	"log"
	"os"
)

func main() {
	ch4.MarshalDemo()
	fmt.Println("调用get方法")
	result, err := ch4.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
