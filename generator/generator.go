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

var (
	certificateDecoded image.Image
	fontface font.Face
	baseDir string
	err error
	certificatesGeneratedDir string = "certificatesGenerated/"
)

func Initialize(template Template, blankCertificate io.Reader, parsedFont *truetype.Font, baseDirectory string) {
	baseDir = baseDirectory
	certificateDecoded, err = png.Decode(blankCertificate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("certificate decoded")

	fontface = truetype.NewFace(parsedFont, &truetype.Options{
			Size: float64(template.Title.Fontsize),
		})
	fmt.Println("new fontface created")

	if _, err = os.Stat(baseDir + certificatesGeneratedDir); os.IsNotExist(err) {
		err = os.Mkdir(baseDir + certificatesGeneratedDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func GenerateImage(template Template, name string) {
	fmt.Println("generating for", name)

	canvas := image.NewRGBA(certificateDecoded.Bounds())

	// draw blankCertificate onto the canvas
	draw.Draw(canvas, canvas.Rect, certificateDecoded, image.Point{0, 0}, draw.Src)
	fmt.Println("canvas drawn")


	// hinting := font.HintingNone
	textDrawer := &font.Drawer{
		Dst: canvas,
		Src: image.Black,
		Face: fontface,
	}

	// textDrawer.Dot.X = fixed.Int26_6(canvas.Rect.Dx() * 26)
	x, y := getAlignment(template.Title.Align_x, template.Title.Align_y, canvas, name, fontface)
	textDrawer.Dot.X = x + fixed.I(template.Title.Offset_x)
	textDrawer.Dot.Y = y + fixed.I(template.Title.Offset_y)
	fmt.Println("Drawing at:", x, y)
	textDrawer.DrawString(name)
	
	// create file to save image to f"{baseDir}{name}.png" 
	fmt.Println("saving to", baseDir + certificatesGeneratedDir + name + ".png")
	outFile, err := os.Create(baseDir + certificatesGeneratedDir + name + ".png")
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

func getAlignment(Align_x string, Align_y string, canvas *image.RGBA, text string, fontface font.Face) (fixed.Int26_6, fixed.Int26_6) {
	var (
		x fixed.Int26_6
		y fixed.Int26_6
	)

	switch Align_x {
	case "left":
		x = 0

	case "right":
		x = fixed.I(canvas.Rect.Dx()) - font.MeasureString(fontface, text)

	case "center":
		x = (fixed.I(canvas.Rect.Dx()) - font.MeasureString(fontface, text)) / 2
	
	}

	switch Align_y {
	case "top":
		y = 0 + fontface.Metrics().Ascent

	case "bottom":
		y = fixed.I(canvas.Rect.Dy()) - fontface.Metrics().Descent

	case "center":
		// this is not exactly centered because of how fonts are drawn
		y = (fixed.I(canvas.Rect.Dy())) / 2
	}

	return x, y
}
