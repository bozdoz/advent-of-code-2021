package utils

import (
	"fmt"
	"math"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Vector[T Numeric] struct {
	X, Y T
}

func NewVector[T Numeric](x, y T) Vector[T] {
	return Vector[T]{x, y}
}

func (this Vector[T]) Add(vec Vector[T]) Vector[T] {
	return Vector[T]{
		this.X + vec.X,
		this.Y + vec.Y,
	}
}

func (this Vector[T]) Subtract(vec Vector[T]) Vector[T] {
	return Vector[T]{
		this.X - vec.X,
		this.Y - vec.Y,
	}
}

func (this *Vector[T]) subtractBy(vec Vector[T]) {
	this.X -= vec.X
	this.Y -= vec.Y
}

func (this *Vector[T]) divideBy(num T) {
	this.X /= num
	this.Y /= num
}

func (this *Vector[T]) multiplyBy(num T) {
	this.X *= num
	this.Y *= num
}

func (this Vector[T]) IsEqualTo(vec Vector[T]) bool {
	return this.X == vec.X && this.Y == vec.Y
}

func (this *Vector[T]) AngleRadians() float64 {
	return math.Atan2(float64(this.Y), float64(this.X))
}

func (this *Vector[T]) AngleDegrees() float64 {
	radian := this.AngleRadians()
	degree := math.Round(radian * 180 / math.Pi)

	if degree < 0 {
		degree += 360
	}

	return degree
}

func (this *Vector[T]) lengthSquared() T {
	return this.X*this.X + this.Y*this.Y
}

func (this Vector[T]) length() float64 {
	return math.Sqrt(float64(this.lengthSquared()))
}

func (this *Vector[T]) distanceTo(vec2 Vector[T]) float64 {
	return this.Subtract(vec2).length()
}

func (this *Vector[T]) ToString() string {
	return fmt.Sprint(this)
}
