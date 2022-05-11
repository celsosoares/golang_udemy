package formas

import (
	"fmt"
	"math"
)

// Uma interface é um conjunto de métodos que usamos para definir um conjunto de ações.

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	r := Retangulo{10, 15}
	EscreverArea(r)

	c := Circulo{10}
	EscreverArea(c)
}
