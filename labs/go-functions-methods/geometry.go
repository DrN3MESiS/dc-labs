// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Point x,y
type Point struct{ x, y float64 }

//Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//Distance same thing, but as a method of the Point type
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//!-point
func (p *Point) setX(x float64) {
	p.x = x
}

func (p *Point) setY(y float64) {
	p.y = y
}

//X ...
func (p *Point) X() float64 {
	return p.x
}

//Y ...
func (p *Point) Y() float64 {
	return p.y
}

func ccw(p, q, r Point) bool {
	d := (q.Y()-p.Y())*(r.X()-q.X()) - (q.X()-p.X())*(r.Y()-q.Y())
	if d == 0 {
		return true
	}

	if d < 0 {
		return false
	}

	if d > 0 {
		return true
	}

	return true
}

func genRandom() float64 {
	min := -100
	max := 100
	val := rand.Intn(max-min) + min
	return float64(val)
}

func perimiter(polygon []Point) float64 {
	var per float64 = 0.0
	for i := 0; i < len(polygon); i++ {
		if i != len(polygon)-1 {
			p1 := polygon[i]
			p2 := polygon[i+1]
			per += math.Floor(math.Sqrt(math.Pow((p1.X()-p2.X()), 2) + math.Pow((p1.Y()-p2.Y()), 2)))
		}
	}
	return per
}

//Generator Point Generation Function
func Generator(sides int) ([]Point, error) {
	rand.Seed(time.Now().UnixNano())
	points := []Point{}
	for i := 0; i < sides; i++ {
		p := Point{}
		p.setX(genRandom())
		p.setY(genRandom())
		points = append(points, p)

	}
	return points, nil
}

func main() {
	args := os.Args[1:]
	unparsedSides := args[0]
	fmt.Println("- Generating a [" + unparsedSides + "] sides figure")

	sides, err := strconv.Atoi(unparsedSides)
	if err != nil {
		panic(err)
	}

	fmt.Println("- Figure's Vertices")
	points, err := Generator(sides)

	poly := []Point{}

	for _, p := range points {
		fmt.Printf("\t- (%f, %f)\n", p.X(), p.Y())
		for _, q := range points {
			if p == q {
				continue
			}

			valid := true

			for _, r := range points {
				if p != r && q != r {
					if !ccw(p, q, r) {
						valid = false
					}
				}
			}

			if valid {
				poly = append(poly, q)
			}
		}
	}

	//No utilice las maneras mas eficientes para realizar todos los calculos
	fmt.Println("- Figure's Perimeter")
	final := perimiter(poly)
	fmt.Printf("\t- %f", final)
}
