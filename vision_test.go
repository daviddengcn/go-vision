package vision

import (
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

	for _, d := range datasets {
		img, err := ImageFromFile("testdata/" + d.fn)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("Image loaded: %v", d.fn)
		
		var m GrayImage
		if err := m.SetImage(img); err != nil {
			t.Error(err)
			continue
		}

		if m.Width != d.width || m.Height != d.height {
			t.Errorf("(GrayImage) Wrong size: %d, %d, expected %d, %d",
				m.Width, m.Height, d.width, d.height)
		}
		img = m.AsImage()
		if img.Bounds().Dx() != d.width || img.Bounds().Dy() != d.height {
			t.Errorf("(AsImage) Wrong size: %d, %d, expected %d, %d", 
				img.Bounds().Dx, img.Bounds().Dy, d.width, d.height)
		}
		
		if err := SaveImageAsPng(img, d.fn + ".png"); err != nil {
			t.Error(err)
			continue
		}
	}
}
