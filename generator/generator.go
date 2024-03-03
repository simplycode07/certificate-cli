package generator

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// struct to unmarshal the template file
type Template struct {
	Event struct {
		Name	 string	`json:"name"`
		Date	 string	`json:"date"`
	} `json:"event"`

	Title struct {
		Font	 string `json:"font"`
		Fontsize int	`json:"fontsize"`
		Align_x  string `json:"align_x"`
		Align_y	 string `json:"align_y"`
		Offset_x int	`json:"offset_x"`
		Offset_y int	`json:"offset_y"`
	} `json:"title"`

	Serial struct {
		Font	 string `json:"font"`
		Fontsize int	`json:"fontsize"`
		Align_x  string `json:"align_x"`
		Align_y	 string `json:"align_y"`
		Offset_x int	`json:"offset_x"`
		Offset_y int	`json:"offset_y"`
	} `json:"serial"`
}

func GenerateImage(template Template, blankCertificate io.Reader, parsedFont *truetype.Font, name string, baseDir string) {
	fmt.Println("generating for", template.Event.Name)

	certificateDecoded, err := png.Decode(blankCertificate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	canvas := image.NewRGBA(certificateDecoded.Bounds())

	// draw blankCertificate onto the canvas
	draw.Draw(canvas, canvas.Rect, certificateDecoded, image.Point{0, 0}, draw.Src)

	fontface := truetype.NewFace(parsedFont, &truetype.Options{
			Size: float64(template.Title.Fontsize),
		})

	// hinting := font.HintingNone
	textDrawer := &font.Drawer{
		Dst: canvas,
		Src: image.Black,
		Face: fontface,
	}
	
	// textDrawer.Dot.X = fixed.Int26_6(canvas.Rect.Dx() * 26)
	textDrawer.Dot.X = fixed.Int26_6((fixed.I(canvas.Rect.Dx()) - font.MeasureString(fontface, "testing...")) / 2 + fixed.I(template.Title.Offset_x))
	textDrawer.Dot.Y = fixed.Int26_6((fixed.I(canvas.Rect.Dx()) - font.MeasureString(fontface, "testing...")) / 2 + fixed.I(template.Title.Offset_y))
	textDrawer.DrawString("testing...")
	
	// create file to save image to f"{baseDir}{name}.png" 
	fmt.Println("saving to", baseDir + name + ".png")
	outFile, err := os.Create(baseDir + name + ".png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outFileWriter := bufio.NewWriter(outFile)
	err = png.Encode(outFileWriter, canvas)
	if err != nil {
		fmt.Println("couldnt write to file")
	}

	outFileWriter.Flush()
}
