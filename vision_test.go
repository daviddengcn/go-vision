package vision

import (
	"os"
	"testing"

	"github.com/golangplus/testing/assert"
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
				outImg.Bounds().Dx(), outImg.Bounds().Dy(), d.width, d.height)
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
				outImg.Bounds().Dx(), outImg.Bounds().Dy(), d.width, d.height)
		}

		if err := SaveImageAsPng(outImg, "testout/rgb-"+d.fn+".png"); err != nil {
			t.Error(err)
			continue
		}
	}
}

func TestResize(t *testing.T) {
	var gray GrayImage
	gray.Resize(Size{123, 456})
	assert.Equal(t, "Size", gray.Size, Size{123, 456})
	assert.Equal(t, "len(Pixels)", len(gray.Pixels), 123*456)

	var ig IntGrayImage
	ig.Resize(Size{123, 555})
	assert.Equal(t, "Size", ig.Size, Size{123, 555})
	assert.Equal(t, "len(Pixels)", len(ig.Pixels), 123*555)

	var rgb RGBImage
	rgb.Resize(Size{123, 456})
	assert.Equal(t, "Size", rgb.Size, Size{123, 456})
	assert.Equal(t, "len(Pixels)", len(rgb.Pixels), 123*456)
}

func TestFill(t *testing.T) {
	var gray GrayImage
	gray.Resize(Size{456, 123})
	gray.Fill(123)
	assert.Equal(t, "[0,0]", gray.Pixels[0], byte(123))

	var ig IntGrayImage
	ig.Resize(Size{555, 123})
	ig.Fill(12345)
	assert.Equal(t, "[0,0]", ig.Pixels[0], int(12345))

	var rgb RGBImage
	rgb.Resize(Size{456, 123})
	rgb.Fill(RGB{1, 2, 3})
	assert.Equal(t, "[0,0]", rgb.Pixels[0], RGB{1, 2, 3})
}
