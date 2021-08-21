package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	input_img_front_data, _ := os.Open("input_img/gopher.png")
	defer input_img_front_data.Close()
	input_img_back_data, _ := os.Open("input_img/sample.png")
	defer input_img_back_data.Close()

	// イメージオブジェクトを作成
	input_img_back, _, err := image.Decode(input_img_back_data)
	if err != nil {
		fmt.Println(input_img_back, err)
	}
	input_img_front, _, err := image.Decode(input_img_front_data)
	if err != nil {
		fmt.Println(input_img_front, err)
	}

	// 書き出し用イメージ作成
	output_rect := image.Rectangle{image.Pt(0, 0), input_img_back.Bounds().Size()}
	output_image_canvas := image.NewRGBA(output_rect)

	// 書き出し用イメージに書き出す。背景
	output_back_rect := image.Rectangle{image.Pt(0, 0), input_img_back.Bounds().Size()}
	draw.Draw(output_image_canvas, output_back_rect, input_img_back, image.Pt(0, 0), draw.Src)
	// 書き出し用イメージに書き出す。前面
	output_front_rect := image.Rectangle{image.Pt(0, 0), input_img_front.Bounds().Size()}
	draw.Draw(output_image_canvas, output_front_rect, input_img_front, image.Pt(0, 0), draw.Over)
	// 書き出しファイル作成
	output_image, _ := os.Create("output_result.png")
	defer output_image.Close()
	// png形式で書き出し
	png.Encode(output_image, output_image_canvas)

}
