package main

import (
	"os"
	"log"
)

type SVG struct {
	FileName string
	ptr *os.File
}

func (svg SVG) Create() {
	ptr, err := os.Create(svg.FileName)
	if err != nil {
		log.Fatal("Panic!")
	}
	svg.ptr = ptr
}

func (svg SVG) Write(s string) {
	if svg.ptr == nil {
		log.Fatal("Cannot write string: file is not created")
	}
	svg.ptr.WriteString(s)
}

func (svg SVG) Close() {
	svg.ptr.Sync()
	svg.ptr.Close()
}

func main()  {
	svg := SVG{FileName: "content.html"}
	svg.Create()
	svg.Write("<svg width='600' height='400'>")
	svg.Write("</svg")
	svg.Close()

	//file, err := os.Create("file.txt")
	//if err != nil {
	//	os.Exit(1)
	//}
	//file.WriteString("Some string\n")
	//file.Sync()
	//file.Close()
}
