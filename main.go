package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
)

func main() {
	img1, _ := createColorGridImage("alice@server.tld")
	fmt.Println(img1)

	img2, _ := createColorGridImage("bob@server.tld")
	fmt.Println(img2)

	img3, _ := createColorGridImage("chris@server.tld")
	fmt.Println(img3)

	img4, _ := createColorGridImage("alice@bob.tld")
	fmt.Println(img4)

	img5, _ := createColorGridImage("bob@chris.tld")
	fmt.Println(img5)
}

func createColorGridImage(s string) (string, error) {

	size := 64
	// Create an image with the specified size
	img := image.NewNRGBA(image.Rect(0, 0, size, size))

	// Seed the random number generator with the string
	rnd := rand.New(rand.NewSource(int64(hash(s))))

	// Generate the color grid
	for i := 0; i < size; i += 10 {
		for j := 0; j < size; j += 10 {
			c := color.NRGBA{
				R: uint8(rnd.Intn(256)),
				G: uint8(rnd.Intn(256)),
				B: uint8(rnd.Intn(256)),
				A: 255,
			}
			for x := i; x < i+10; x++ {
				for y := j; y < j+10; y++ {
					img.Set(x, y, c)
				}
			}
		}
	}

	// Encode the image as a PNG
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return "", err
	}

	// Return the data encoded image string
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// hash is a helper function that returns a deterministic hash of a string
func hash(s string) int {
	h := 0
	for _, c := range s {
		h = 31*h + int(c)
	}
	return h
}
