package types

import (
	"fmt"
	"math"
)

type Vector3d struct {
	X, Y, Z int
}

func NewVector3d(x, y, z int) Vector3d {
	return Vector3d{x, y, z}
}

func (this *Vector3d) LengthSquared() int {
	return this.X*this.X + this.Y*this.Y + this.Z*this.Z
}

func (this Vector3d) Length() float64 {
	return math.Sqrt(float64(this.LengthSquared()))
}

func (this Vector3d) Add(vec Vector3d) Vector3d {
	return Vector3d{
		this.X + vec.X,
		this.Y + vec.Y,
		this.Z + vec.Z,
	}
}

func (this *Vector3d) Subtract(vec Vector3d) Vector3d {
	return Vector3d{
		this.X - vec.X,
		this.Y - vec.Y,
		this.Z - vec.Z,
	}
}

func (this *Vector3d) dotProduct(vec Vector3d) int {
	return this.X*vec.X + this.Y*vec.Y + this.Z*vec.Z
}

func (this *Vector3d) crossProduct(vec Vector3d) *Vector3d {
	return &Vector3d{
		this.Y*vec.Z - this.Z*vec.Y,
		this.Z*vec.X - this.X*vec.Z,
		this.X*vec.Y - this.Y*vec.X,
	}
}

func (this *Vector3d) distanceTo(vec2 Vector3d) float64 {
	return this.Subtract(vec2).Length()
}

func (this *Vector3d) angleBetween(vec2 Vector3d) float64 {
	return math.Acos(float64(this.dotProduct(vec2)) / (this.Length() * vec2.Length()))
}

func (this Vector3d) IsEqualTo(vec Vector3d) bool {
	return this.X == vec.X && this.Y == vec.Y && this.Z == vec.Z
}

func (this *Vector3d) ToString() string {
	return fmt.Sprint(this)
}
