# go-aalib

[AAlib](http://aa-project.sourceforge.net/aalib/) binding for Go.

## Sample Code

```go
package main

import "github.com/syohex/go-aalib"
import "image"
import _ "image/png"
import "fmt"
import "os"

func main() {
	file, err := os.Open("sample.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the image.
	goPng, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	handle, _ := aalib.Init(80, 60, aalib.AA_REVERSE)
	handle.PutImage(goPng)
	handle.Render(nil, 0, 0, 96, 96)
	aaStr := handle.Text()
	fmt.Println(aaStr)
}
```
