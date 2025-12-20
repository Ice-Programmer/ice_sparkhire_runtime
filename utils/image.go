package utils

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
)

const (
	defaultFontPath = "/Users/chenjiahan/project/Graduation Project/ice_sparkhire_runtime/.conf/font/AlibabaPuHuiTi-3-75-SemiBold.ttf"
)

// GenerateCharAvatar generate single character image
func GenerateCharAvatar(char rune, size int) ([]byte, error) {
	// 1. create a panel
	dc := gg.NewContext(size, size)

	// 2. set background color
	dc.SetColor(GetElegantColorByChar(char))
	dc.Clear()

	// 3. load font
	fontSize := float64(size) * 0.6
	if err := dc.LoadFontFace(defaultFontPath, fontSize); err != nil {
		return nil, fmt.Errorf("load font path error: %v", err)
	}

	// 4. set font color
	textColor := color.White
	dc.SetColor(textColor)

	// 5. draw character and center
	text := string(char)
	dc.DrawStringAnchored(text, float64(size)/2, float64(size)/2, 0.5, 0.5)

	// 6. save picture
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, fmt.Errorf("encode png error: %w", err)
	}

	return buf.Bytes(), nil
}

func GetElegantColorByChar(char rune) color.RGBA {
	palette := []color.RGBA{
		{142, 154, 175, 255}, // 灰蓝色
		{184, 172, 160, 255}, // 暖灰色
		{110, 137, 139, 255}, // 灰绿色
		{214, 173, 160, 255}, // 烟粉色
		{157, 129, 137, 255}, // 藕荷色
		{125, 133, 151, 255}, // 雾霾蓝
		{163, 145, 147, 255}, // 褐灰色
		{199, 178, 153, 255}, // 亚麻色
	}

	return palette[int(char)%len(palette)]
}
