package main

import (
	"errors"
	"fmt"
	"math"
)

// エラー情報を格納する構造体を定義
type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ERROR Message:%s,Field:%s,Value:%f", e.Message, e.Field, e.Value)
}

// 図形インターフェースを定義
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 長方形と円の構造体を定義
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 長方形と円の計算用メソッドを定義
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 図形のバリデーションを行う関数を定義
func validateShape(s Shape) error {
	switch v := s.(type) {
	case Rectangle:

		if v.Width <= 0 {
			return fmt.Errorf(
				"Rectangleの検証失敗:%w",
				&ValidationError{
					Message: "幅は正の数を指定してください",
					Field:   "Rectangle.Width",
					Value:   v.Width,
				},
			)
		}
		if v.Height <= 0 {
			return fmt.Errorf(
				"Rectangleの検証失敗:%w",
				&ValidationError{
					Message: "高さは正の数を指定してください",
					Field:   "Rectangle.Height",
					Value:   v.Height,
				},
			)
		}
	case Circle:
		if v.Radius <= 0 {

			err := &ValidationError{
				Message: "半径は正の数を指定してください",
				Field:   "半径",
				Value:   v.Radius,
			}
			return fmt.Errorf("Circleの検証失敗：%w", err)
		}
	}
	return nil
}

// main関数で図形の面積・周長を計算
func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	var calcShapes []Shape

	// スライスに対して繰り返し処理を行い、バリデーションを実施
	fmt.Println("---図形のバリデーション---")
	for _, v := range shapes {
		err := validateShape(v)

		if err != nil {
			fmt.Println("バリテーションNG")
			var validationErr *ValidationError

			if errors.As(err, &validationErr) {
				fmt.Printf("%s\n", validationErr.Error())
			} else {
				fmt.Printf("その他のエラー:%s\n", err.Error())
			}

		} else {
			fmt.Println("バリテーションOK")
			calcShapes = append(calcShapes, v)
		}
	}

	// 図形の型情報（%T）・面積・周長をそれぞれ出力
	fmt.Println("---検証済み図形の一括計算---")

	for _, v := range calcShapes {
		fmt.Printf("図形:%T | 面積:%f | 周長:%f\n", v, v.Area(), v.Perimeter())
	}

}
