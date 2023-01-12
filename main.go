package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
)

func main() {

	// one route for requests, send everything to handlerfunc
	http.HandleFunc("/", handlerFunc)

	// API on port 8080
	http.ListenAndServe(":8080", nil)
}

// handlerfunc takes a string in a get request from the URL and returns an image as text
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	// get the string from the URL path
	s := r.URL.Path[1:]

	// generate an image based on the string
	result, err := createAbstractImage(s)

	// check for errors
	if err != nil {
		// on error set the status code to internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if all is well set the status code to OK
	w.WriteHeader(http.StatusOK)

	// return the image text as the response
	w.Write([]byte(result))
}

// hash returns a deterministic hash of a string, by combining all characters
// of the input string using Horner's method with a prime number to help avoid
// collisions and increase the distribution of the results
func hash(s string) int {

	// initialize h as 0
	h := 0

	// 31 is used to make computation faster, but at the risk of collission
	// given that this is just an image, it's fine. let's hope.
	prime := 31

	// using horner's method combine each character of the input string
	// with the current hash value and multiply by a prime number
	for _, c := range s {
		h = h*prime + int(c)
	}
	return h
}

// createAbstractImage takes a string, generates the image, then returns
// that image as a string that can be rendered in an <img> as a URL.
func createAbstractImage(s string) (string, error) {

	// image size in pixels
	size := 88

	// create a new image with size
	img := image.NewNRGBA(image.Rect(0, 0, size, size))

	// use a hash to seed the random number generator
	rnd := rand.New(rand.NewSource(int64(hash(s))))

	// loop through the image and draw the shapes
	for i := 0; i < size; i += 10 {
		for j := 0; j < size; j += 10 {

			// for each shape, pick a random color at full opacity
			// intially we also did alpha transparency but it just
			// made everything look like easter shapes.
			c := color.NRGBA{
				R: uint8(rnd.Intn(256)),
				G: uint8(rnd.Intn(256)),
				B: uint8(rnd.Intn(256)),
				A: uint8(255),
			}

			// pick a shape
			shapeType := rnd.Intn(2)
			if shapeType == 0 {

				// draw a circle
				drawCircle(img, c, i, j)
			} else {

				// draw a square
				drawSquare(img, c, i, j)
			}
		}
	}

	// buffer to hold the image data
	buf := new(bytes.Buffer)

	// encode as png
	if err := png.Encode(buf, img); err != nil {
		return "", err
	}

	// Return the image data as a base64-encoded string
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()), nil

}

// drawCircle uses image, color, and coordinates to draw curves.
// initially it was an attempt to draw a circle but that didn't work out
// and what it resulted in looked cool, so it stayed.
func drawCircle(img *image.NRGBA, c color.NRGBA, x int, y int) {

	// first we got our i passed in as arguments.
	// next, we loop through the pixels in an 8x8 area where the circle will be drawn.
	for i := x; i < x+8; i++ {
		for j := y; j < y+8; j++ {

			// then check if the current pixel is within the circle by using the
			// distance from center of circle to current pixel
			if ((i-x)*(i-x) + (j-y)*(j-y)) <= 64 {

				// if it's within the circle, we color that pixel with the passed in color.
				img.Set(i, j, c)
			}
		}
	}
}

// drawSquare takes in three arguments: an image, a color and two coords.
// it starts iterating over the pixels in the 8x8 square starting from the
// given x and y coordinates and sets the color of every pixel in this square
// to the given color.
func drawSquare(img *image.NRGBA, c color.NRGBA, x int, y int) {

	// For every pixel in the x range from x to x+8
	for i := x; i < x+8; i++ {

		// For every pixel in the y range from y to y+8
		for j := y; j < y+8; j++ {

			// Set the pixel's color to the desired color
			img.Set(i, j, c)
		}
	}
}
