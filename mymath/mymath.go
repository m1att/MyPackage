package mymath

import "math"

// Sqrt - возвращает квадратный корень
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Abs - возвращает абсолютное значение
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Max - возвращает максимальное значение
func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// Yn - функция Бесселя второго рода
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
