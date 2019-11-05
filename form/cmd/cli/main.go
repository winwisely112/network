package main

import (
	"flag"
	"fmt"
	"os"

	// "github.com/winwisely99/network/form"
	"github.com/winwisely99/network/form"
)

var welcomeSignature = `
Usage of Form made By Rohit
_______________________________
	< Form >
-------------------------------

-name string:
	Set the name

`

func main() {
	var (
		name = flag.String("name", "", "Set the name where you want to generate a identicon for")
	)
	flag.Parse()

	if *name == "" {
		flag.Usage = func() {
			fmt.Println(welcomeSignature)
		}
		flag.Usage()
		os.Exit(0)
	}

	generatedIdenticon := identicon.Generate([]byte(*name))

	f, err := os.Create(generatedIdenticon.Name + ".png")
	if err != nil {
		fmt.Printf("error: failed creating file for output png")
		return
	}
	defer f.Close()

	if err := generatedIdenticon.WriteImage(f); err != nil {
		fmt.Printf("error failed writing image to file")
		return
	}
}
