// ./mandelbrot | ./converter jpeg  > manderbrot.jpeg

package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	outFormat := os.Args[1]
	if err := convert(outFormat, os.Stdout, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// convertToPNG converts from any recognized format to PNG.
func convert(outFormat string, w io.Writer, r io.Reader) error {
	img, kind, err := image.Decode(r)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input Format = ", kind)
	switch outFormat {
	case "png":
		return png.Encode(w, img)
	case "jpeg":
		return jpeg.Encode(w, img, &jpeg.Options{})
	case "gif":
		return gif.Encode(w, img, &gif.Options{})
	default:
		fmt.Fprintln(os.Stderr, "unsupported")
		return err
	}
}
