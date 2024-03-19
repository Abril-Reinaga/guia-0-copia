package estructurasinterfaces

import (
	"fmt"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	ToString() string
}

type Punto struct {
	X float64
	Y float64
}

// Pre: El parámetro x y el parámetro y no deben ser nulos.
// Post: Devuelve un punto con coordenadas x e y.
func NewPunto(x, y float64) *Punto {
	return &Punto{x, y}
}

// Pre: El punto no debe ser nulo.
// Post: Devuelve la representación en cadena del punto.
func (p *Punto) ToString() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

// Agregar al módulo figuras visto en el taller un círculo (representado por un pto que indica el centro y su radio (entero). El círculo debe implementar los métodos de la interface figura.

type Circulo struct {
	// ?
}

// Pre: El parámetro pto no debe ser nulo.
// Post: Devuelve un círculo con centro en pto y radio r.
func NewCirculo(pto *Punto, r float64) *Circulo {
	return nil
}

// Pre: El círculo no debe ser nulo.
// Post: Devuelve el área del círculo.
func (c *Circulo) Area() float64 {
	return 0
}

// Pre: El círculo no debe ser nulo.
// Post: Devuelve el perímetro del círculo.
func (c *Circulo) Perimetro() float64 {
	return 0
}

// Pre: El círculo no debe ser nulo.
// Post: Devuelve la representación en cadena del círculo.
func (c *Circulo) ToString() string {
	return ""
}
