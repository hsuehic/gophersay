package main

import (
	"fmt"
	gopher "github/hsuehic/gophersay/gophers"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"embed"
)

// Hey, I want to embed "gophers" folder in the executable binary
// Use embed go 1.16 new feature (for embed gophers static files)
//
//go:embed gophers
var embedGopherFiles embed.FS

func main() {
	application := app.New()
	win := application.NewWindow("Gophersay")

	str := "Gopher"
	text := canvas.NewText(str, color.Black)
	text.Alignment = fyne.TextAlignCenter

	buffer, _ := embedGopherFiles.ReadFile("gophers/1.png")
	res := fyne.NewStaticResource("gopher", buffer)
	img := canvas.NewImageFromResource(res)
	img.SetMinSize(fyne.NewSize(500, 500))

	box := container.NewVBox(text, img)
	win.SetContent(box)
	win.ShowAndRun()
}

func say() {

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
