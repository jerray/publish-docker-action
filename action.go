package main

import "fmt"

func setOutput(name string, value string) {
	fmt.Printf("::set-output name=%s::%s", name, value)
}
