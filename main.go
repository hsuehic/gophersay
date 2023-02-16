package main

import (
	"fmt"
	gopher "github/hsuehic/gophersay/gophers"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"embed"
)

// Hey, I want to embed "gophers" folder in the executable binary
// Use embed go 1.16 new feature (for embed gophers static files)
//
//go:embed gophers
var embedGopherFiles embed.FS

func main() {

	// Display usage/help message
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-h") || (len(os.Args) == 2 && os.Args[1] == "--help") {
		usage := "GopherSay is inspired by Cowsay program.\nGopherSay allow you to display a message said by a cute random Gopher.\n\nUsage:\n   gophersay MESSAGE\n\nExample:\n   gophersay hello Gopher lovers"

		fmt.Println(usage)
		return
	} else if len(os.Args) > 1 {

		message := strings.Join(os.Args[1:], " ")
		nbChar := len(message)

		line := " "
		for i := 0; i <= nbChar; i++ {
			line += "-"
		}

		fmt.Println(line)
		fmt.Println("< " + message + " >")
		fmt.Println(line)
		fmt.Println("        \\")
		fmt.Println("         \\")

		// Generate a random integer depending on get the number of ascii files
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(getNbOfGopherFiles() - 1)

		ascii := gopher.Convert("gophers/" + strconv.Itoa(randInt) + ".png")
		fmt.Print(ascii)
	}
}

func getNbOfGopherFiles() int {

	files, err := embedGopherFiles.ReadDir("gophers")
	if err != nil {
		log.Fatal("Error during reading gophers folder", err)
	}

	nbOfFiles := 0
	for _, file := range files {
		name := file.Name()
		if strings.Contains(name, ".png") {
			nbOfFiles++
		}

	}
	return nbOfFiles
}
