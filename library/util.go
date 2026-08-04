package library

import "strings"

type Animal struct {
	Name  string
	Sound string
}

func GetAnimalSound(animal Animal) string {
	return animal.Sound
}

//varadic func
func PrintFamily(family ...string) string {

	return strings.Join(family, ", ")

}

type PersegiPanjang struct {
	Panjang int8
	Lebar   int8
}

//predefined function
func HitungPersegi(s int) (l int, k int) {
	l = s * s
	k = 4 * s
	return
}

//explicit return
func HitungPersegi2(s int) (int, int) {
	l := s * s
	k := 4 * s
	return l, k
}

func (pp PersegiPanjang) Luas() int8 {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int8 {
	return 2 * (pp.Panjang + pp.Lebar)
}
