package main

import (
	"context"
	"fmt"
)

func main() {
	conohaSettings, err := NewConohaSettings()
	if err != nil {
		fmt.Println(err)
		return
	}
	conohaClient, err := NewConohaClient(conohaSettings)
	if err != nil {
		fmt.Println(err)
		return
	}

	got, _, err := conohaClient.ComputeFlavors(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range got {
		fmt.Println(*image)
	}
	return
}
