package jsfmt

import (
	"github.com/gopherjs/gopherjs/js"
	
	"strconv"
	"image/color"
)

var (
	outputID = ""
	backgroundColor = ""
	frontgroundColor = ""
	defaultFontColor = ""
	isSetID = false
)

// Print は fmt.Print() のJavaScript版です
func Print(a ...interface{}) {
	doPrint(a)
}

// Println は fmt.Println() のJavaScript版です
func Println(a ...interface{}) {
	doPrintln(a)
}

// Printf は fmt.Printf() のJavaScript版です
func Printf(format string, a ...interface{}) {
	// doPrintf(format, a)
}

// Color で Print系の関数の中で使えます 色を付けられます
func Color(fgColor color.RGBA, bgColor color.RGBA, a ...interface{}) []interface{} {
	frontgroundColor = "rgba(" + strconv.Itoa(int(fgColor.R)) + ", " + strconv.Itoa(int(fgColor.G)) + "," + strconv.Itoa(int(fgColor.B)) + "," + strconv.Itoa(int(fgColor.A)) + ")"
	backgroundColor = "rgba(" + strconv.Itoa(int(bgColor.R)) + ", " + strconv.Itoa(int(bgColor.G)) + "," + strconv.Itoa(int(bgColor.B)) + "," + strconv.Itoa(int(bgColor.A)) + ")"

	return a
}

// SetOutputID で 出力する場所を設定します
func SetOutputID(id string) {
	outputID = id
	isSetID = true
}

// SetFontFamily で css の font-family を設定します
func SetFontFamily(ff string) {
	if isSetID && outputID != "" {
		js.Global.Get("document").Call("getElementById", outputID).Get("style").Set("fontFamily", ff)
	}else {
		js.Global.Get("document").Get("body").Get("style").Set("fontFamily", ff)
	}
}

// SetScreenColor で 画面の背景色を設定します(bodyのbackgroundColor)
func SetScreenColor(c color.RGBA) {
	rgba := "rgba(" + strconv.Itoa(int(c.R)) + ", " + strconv.Itoa(int(c.G)) + "," + strconv.Itoa(int(c.B)) + "," + strconv.Itoa(int(c.A)) + ")"
	js.Global.Get("document").Get("body").Get("style").Set("backgroundColor", rgba)
}

// SetDefaultFontColor で デフォルトのフォントの色を設定します
func SetDefaultFontColor(c color.RGBA) {
	defaultFontColor := "rgba(" + strconv.Itoa(int(c.R)) + ", " + strconv.Itoa(int(c.G)) + "," + strconv.Itoa(int(c.B)) + "," + strconv.Itoa(int(c.A)) + ")"
	js.Global.Get("document").Get("body").Get("style").Set("color", defaultFontColor)
}

// SetColor フォントの色、背景色を設定します
func SetColor(fgColor color.RGBA, bgColor color.RGBA) {
	frontgroundColor = "rgba(" + strconv.Itoa(int(fgColor.R)) + ", " + strconv.Itoa(int(fgColor.G)) + "," + strconv.Itoa(int(fgColor.B)) + "," + strconv.Itoa(int(fgColor.A)) + ")"
	backgroundColor = "rgba(" + strconv.Itoa(int(bgColor.R)) + ", " + strconv.Itoa(int(bgColor.G)) + "," + strconv.Itoa(int(bgColor.B)) + "," + strconv.Itoa(int(bgColor.A)) + ")"
}

// ResetColor フォントの色をリセットできます
func ResetColor() {
	frontgroundColor = defaultFontColor
	backgroundColor = ""
}

func doPrint(a []interface{}) {
	doc := js.Global.Get("document")

	var out *js.Object
	if isSetID && outputID != "" {
		out = doc.Call("getElementById", outputID)
	}else {
		out = doc.Get("body")
	}
	
	span := doc.Call("createElement", "span")		
	
	for argNum, arg := range a {
		if argNum > 0 {
			span.Call("insertAdjacentHTML", "beforeend", " ")
		}

		span.Get("style").Set("color", frontgroundColor)
		span.Get("style").Set("backgroundColor", backgroundColor)

		span.Call("insertAdjacentHTML", "beforeend", arg)
		out.Call("appendChild", span)
	}

	ResetColor()
}


func doPrintln(a []interface{}) {
	doc := js.Global.Get("document")

	var out *js.Object
	if isSetID && outputID != "" {
		out = doc.Call("getElementById", outputID)
	}else {
		out = doc.Get("body")
	}

	span := doc.Call("createElement", "span")		
	
	for argNum, arg := range a {

		if argNum > 0 {
			span.Call("insertAdjacentHTML", "beforeend", " ")
		}

		span.Get("style").Set("color", frontgroundColor)
		span.Get("style").Set("background-color", backgroundColor)

		span.Call("insertAdjacentHTML", "beforeend", arg)
		
		out.Call("appendChild", span)
	}

	out.Call("insertAdjacentHTML", "beforeend", "<br>")

	ResetColor()
}