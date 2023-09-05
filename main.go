package main

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"gitlhub.com/athoune/docker-dynasty-go/dynasty"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	d, err := dynasty.New(cli)
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		for _, l := range d.Tree() {
			fmt.Println(string(l.Code), l.Tags)
		}
	} else {
		ancestors, err := d.Ancestor(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println("Ancestors")
		for _, l := range ancestors {
			fmt.Println("\t", string(l.Code), l.Tags, l.Labels)
		}
		descendants, err := d.Descendant(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println("Descendants")
		for _, l := range descendants {
			fmt.Println("\t", string(l.Code), l.Tags, l.Labels)
		}
	}
}
