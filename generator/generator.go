package generator

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// struct to unmarshal the template file
type Template struct {
	Event struct {
		Name     string `json:"name"`
		Date     string `json:"date"`
		SubName  string `json:"subname"`
		CertType string `json:"cert_type"`
		Volunteering bool `json:"volunteering"`
	} `json:"event"`

	Title struct {
		Font     string `json:"font"`
		Fontsize int    `json:"fontsize"`
		Align_x  string `json:"align_x"`
		Align_y  string `json:"align_y"`
		Offset_x int    `json:"offset_x"`
		Offset_y int    `json:"offset_y"`
	} `json:"title"`

	Serial struct {
		Font     string `json:"font"`
		Fontsize int    `json:"fontsize"`
		Align_x  string `json:"align_x"`
		Align_y  string `json:"align_y"`
		Offset_x int    `json:"offset_x"`
		Offset_y int    `json:"offset_y"`
	} `json:"serial"`
}

var (
	certificateDecoded       image.Image
	nameFontface             font.Face
	serialFontface           font.Face
	baseDir                  string
	err                      error
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

	nameFontface = truetype.NewFace(parsedFont, &truetype.Options{
		Size: float64(template.Title.Fontsize),
	})

	serialFontface = truetype.NewFace(parsedFont, &truetype.Options{
		Size: float64(template.Serial.Fontsize),
	})

	fmt.Println("fontfaces created")

	if _, err = os.Stat(baseDir + certificatesGeneratedDir); os.IsNotExist(err) {
		err = os.Mkdir(baseDir+certificatesGeneratedDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func GenerateImage(template Template, name string, index int) {
	fmt.Println("generating for", name)

	canvas := image.NewRGBA(certificateDecoded.Bounds())

	// draw blankCertificate onto the canvas
	draw.Draw(canvas, canvas.Rect, certificateDecoded, image.Point{0, 0}, draw.Src)
	fmt.Println("canvas drawn")

	// hinting := font.HintingNone
	nameDrawer := &font.Drawer{
		Dst:  canvas,
		Src:  image.Black,
		Face: nameFontface,
	}

	// textDrawer.Dot.X = fixed.Int26_6(canvas.Rect.Dx() * 26)
	// draw name
	x, y := getAlignment(template.Title.Align_x, template.Title.Align_y, canvas, name, nameFontface)
	nameDrawer.Dot.X = x + fixed.I(template.Title.Offset_x)
	nameDrawer.Dot.Y = y + fixed.I(template.Title.Offset_y)
	fmt.Println("Drawing name at:", x, y)
	nameDrawer.DrawString(name)

	serialDrawer := &font.Drawer{
		Dst:  canvas,
		Src:  image.Black,
		Face: serialFontface,
	}

	// draw serial
	// --------------------------------------------------------------------------
	// serial := getSerial(template)
	// fmt.Println("Serial name:", serial)
	serial := getSerial(template, index)
	x, y = getAlignment(template.Serial.Align_x, template.Serial.Align_y, canvas, serial, nameFontface)
	serialDrawer.Dot.X = x + fixed.I(template.Serial.Offset_x)
	serialDrawer.Dot.Y = y + fixed.I(template.Serial.Offset_y)
	fmt.Println("Drawing Serial at:", x, y)
	serialDrawer.DrawString(serial)

	// create file to save image to f"{baseDir}{name}.png"
	fmt.Println("saving to", baseDir+certificatesGeneratedDir+name+".png")
	outFile, err := os.Create(baseDir + certificatesGeneratedDir + name + ".png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outFileWriter := bufio.NewWriter(outFile)
	err = png.Encode(outFileWriter, canvas)
	if err != nil {
		fmt.Println("couldn't write to file")
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

func getSerial(template Template, index int) string {
	date := strings.SplitN(template.Event.Date, "/", 3)
	var serial string

	if !template.Event.Volunteering {
		serial = "IIC_" + date[0] + "/" + date[1] + "_" + template.Event.Name + "(" + date[2] + ")-" + template.Event.SubName + "-" + template.Event.CertType + fmt.Sprint(index)

	} else {
		serial = "IIC_" + date[1] + "_" + template.Event.Name + "(" + date[2] + ")" + "_VOL" + fmt.Sprint(index)
	}
	return serial
}
