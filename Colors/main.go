package Colors

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func GetColorsFromImage(fileinname string, fileoutname string) {
	reader, err := os.Open(fileinname)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// arra := [4000][3000]color.Color{}
	f, err := os.OpenFile(fileoutname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			// arra[x][y] = color

			temp := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
			if err != nil {
				panic(err)
			}
			if _, err = f.WriteString(temp); err != nil {
				panic(err)
			}
		}
	}
	defer f.Close()

}
