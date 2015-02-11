package vision

import (
	"os"
	"testing"

	//	"fmt"
)

func TestLoad(t *testing.T) {
	datasets := [...]struct {
		width, height int
		fn            string
	}{
		{177, 177, "gray-177x177.jpg"},
		{500, 374, "ycbcr-500x374.jpg"},
		{500, 374, "pal-500x374.gif"},
		{16, 16, "nrgba-16x16.png"},
	}

	os.MkdirAll("testout", 0755)

	for _, d := range datasets {
		img, err := ImageFromFile("testdata/" + d.fn)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("Image loaded: %v", d.fn)

		var m GrayImage
		m.SetImage(img)

		if m.Width != d.width || m.Height != d.height {
			t.Errorf("(GrayImage) Wrong size: %d, %d, expected %d, %d",
				m.Width, m.Height, d.width, d.height)
		}
		outImg := m.AsImage()
		if outImg.Bounds().Dx() != d.width || outImg.Bounds().Dy() != d.height {
			t.Errorf("(AsImage) Wrong size: %d, %d, expected %d, %d",
				outImg.Bounds().Dx, outImg.Bounds().Dy, d.width, d.height)
		}

		if err := SaveImageAsPng(outImg, "testout/gray-"+d.fn+".png"); err != nil {
			t.Error(err)
			continue
		}
		
		var rgbImg RGBImage
		rgbImg.SetImage(img)
		if rgbImg.Width != d.width || rgbImg.Height != d.height {
			t.Errorf("(RGBImage) Wrong size: %d, %d, expected %d, %d",
				rgbImg.Width, rgbImg.Height, d.width, d.height)
		}
		
		outImg = rgbImg.AsImage()
		if outImg.Bounds().Dx() != d.width || outImg.Bounds().Dy() != d.height {
			t.Errorf("(AsImage) Wrong size: %d, %d, expected %d, %d",
				outImg.Bounds().Dx, outImg.Bounds().Dy, d.width, d.height)
		}

		if err := SaveImageAsPng(outImg, "testout/rgb-"+d.fn+".png"); err != nil {
			t.Error(err)
			continue
		}
	}
}
