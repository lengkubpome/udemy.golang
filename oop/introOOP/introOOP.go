package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// func Length(p *Point) float64 { //แบบเก่า
func (p *Point) Length() float64 { // สร้าง Method

	// return math.Hypot((*p).X, (*p).Y) // พรีทากอรัส
	// ระบบจะมีการจัดการ * ให้อัตโนมัติ ดังนั้นเราจะลดรูปแบบโค้ดตามบรรทัดด้านล่าง
	return math.Hypot(p.X, p.Y) // พรีทากอรัส
}

// func MoveXto(p *Point, newX float64) { //แบบเก่า
func (p *Point) MoveXto(newX float64) { // สร้าง Method
	// (*p).X = newX
	// ระบบจะมีการจัดการ * ให้อัตโนมัติ ดังนั้นเราจะลดรูปแบบโค้ดตามบรรทัดด้านล่าง
	p.X = newX
}

// *** เราจะไม่ใช้รูปแบบ Value เราจะใช้แทนค่าแบบ Pointer เท่านั้น
// func (p Point) MoveYtoByValue(newY float64) { // สร้าง Method
// 	MoveToY(&p, newY)
// }

// func MoveYto(p *Point, newY float64) { //แบบเก่า
func (p *Point) MoveYto(newY float64) { // สร้าง Method
	MoveToY(p, newY)
}

func MoveToY(p *Point, newY float64) {
	// (*p).Y = newY
	// ระบบจะมีการจัดการ * ให้อัตโนมัติ ดังนั้นเราจะลดรูปแบบโค้ดตามบรรทัดด้านล่าง
	p.Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	// fmt.Println(Length(&p)) //แบบเก่า
	fmt.Println(p.Length()) // เรียกใช้แบบ Method
	// MoveXto(&p, 7) //แบบเก่า
	p.MoveXto(7)
	// fmt.Println(Length(&p))
	fmt.Println(p.Length()) // เรียกใช้แบบ Method

	p.MoveYto(45)
	fmt.Println(p)
}
