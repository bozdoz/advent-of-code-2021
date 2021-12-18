package main

import "math"

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Vector[T Numeric] struct {
	x, y T
}

func (this *Vector[T]) add(vec Vector[T]) {
	this.x += vec.x
	this.y += vec.y
}

func (this *Vector[T]) subtract(vec Vector[T]) {
	this.x -= vec.x
	this.y -= vec.y
}

func (this *Vector[T]) divideBy(num T) {
	this.x /= num
	this.y /= num
}

func (this *Vector[T]) multiplyBy(num T) {
	this.x *= num
	this.y *= num
}

func (this *Vector[T]) isEqualTo(vec Vector[T]) bool {
	return this.x == vec.x && this.y == vec.y
}

func (this *Vector[T]) angleRadians() float64 {
	return math.Atan2(float64(this.y), float64(this.x))
}

func (this *Vector[T]) angleDegrees() float64 {
	radian := this.angleRadians()
	degree := math.Round(radian * 180 / math.Pi)

	if degree < 0 {
		degree += 360
	}

	return degree
}

func (this *Vector[T]) lengthSquared() T {
	return this.x*this.x + this.y*this.y
}

func (this *Vector[T]) length() float64 {
	return math.Sqrt(float64(this.lengthSquared()))
}
