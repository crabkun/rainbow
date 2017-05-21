## Rainbow
let your console print colorful character （让golang支持在控制台输出彩色字符）

## Attention！
 only support Windows!!

## Installation
```
go get github.com/crabkun/rainbow
```
### Usage
![test](http://crab.pub/rainbow.png)
```go
package main

import (
    "fmt"
	"rainbow"
)
func main(){
	fmt.Println("default color")

	rainbow.SetFontDarkColor(rainbow.Red)
	fmt.Println("This is dark red font!")

	rainbow.SetFontColor(rainbow.Red)
	fmt.Println("This is red font!")

	rainbow.SetBackgroundColor(rainbow.Cyan)
	fmt.Println("This is red font and cyan background!")

	rainbow.SetBackgroundDarkColor(rainbow.Cyan)
	fmt.Println("This is dark red font and cyan background!")

	rainbow.SetFontDarkColor(rainbow.Red)
	fmt.Println("This is dark red font and dark cyan background!")

	rainbow.CleanBackgroundColor()
	fmt.Println("Deleted background color")

	rainbow.CleanFontColor()
	fmt.Println("Deleted font color")

	rainbow.SetFontColor(rainbow.Pink)
	rainbow.SetBackgroundColor(rainbow.Blue)
	fmt.Println("test")

	rainbow.CleanColor()
	fmt.Println("clean all color")
}
```