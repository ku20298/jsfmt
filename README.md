# jsfmt
ブラウザの画面にfmt.Print()をしたい！(GopherJS)

'''
package main

import (
	"golang.org/x/image/colornames"
	"github.com/ku20298/jsfmt"
)

func main() {
	jsfmt.Println("Hello")
	jsfmt.Println(jsfmt.Color(colornames.Red, colornames.Skyblue, "World"))
}

'''
