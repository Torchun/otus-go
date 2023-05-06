package main

import "fmt"
import "golang.org/x/example/stringutil"

func main() {
	reversed := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reversed)
}
