package transforms

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"teste-apache-beam/entities"
)

func init() {
	beam.RegisterFunction(printPixel)
}

func PrintPixels(s beam.Scope, col beam.PCollection) {
	fmt.Println("pixels: ------------------")
	beam.ParDo0(s, printPixel, col)
}

func printPixel(key string, value entities.Pixel) {
	fmt.Println(key, value)
}
