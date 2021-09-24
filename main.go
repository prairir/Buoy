package main

import (
	"fmt"

	"github.com/songgao/water"
)

const BuffSize = 21

func main() {
	fmt.Println("hi")

	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: "by1",
		},
	}

	f, err := water.New(config)
	if err != nil {
		fmt.Println("FUCK!: %s", err)
		panic(1)
	}

	defer f.Close()

	buff := make([]byte, BuffSize)

	for {
		readSize, err := f.Read(buff)
		if err != nil {
			fmt.Printf("FUCK: %s\n", err)
			panic(1)
		}
		fmt.Println([]byte(buff[:readSize]))
	}
}
