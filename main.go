package main

import (
	"fmt"
)

//int satu := 1

func main() {
	// fmt.Println("angka1 = ", 1)
	// fmt.Println("angka2 = ", 2)
	// fmt.Println("angka3 = ", 3.5)

	// fmt.Println("true", true)
	// fmt.Println("false", false)

	// fmt.Println("ini string")
	// fmt.Println(len("test hitung"))
	// fmt.Println("Ibnu Zaman"[0])

	// var angka int
	// angka = 1

	// var kalimat string
	// kalimat = "Ibnu Zaman"

	// ibnu := 100
	// fmt.Println(ibnu)

	// zaman := false
	// fmt.Println(zaman)

	// fmt.Println(angka)
	// fmt.Println(kalimat)

	// MultipleVariabel()
	// TypeDeclaration()
	// AugmentedAssignments()
	// UnaryOperator()
	// Array()
	// Slice()
	// AppendInSlice()
	// MakeInSlice()
	// CopyInSlice()
	// Mapping()
	// SwitchKondisi()
	// IterationWithRange()
	// BreakAndContinue()
	// fmt.Println(FungsiParameter("Ibnu"))
	// fmt.Println(MultipleValues())

	// Multiple Values
	// namaDepan, namaBelakang := MultipleValues()
	// fmt.Println(namaDepan, namaBelakang)

	// Mengambil Value 1 dari Multiple Values
	// namaDepan, _ := MultipleValues()
	// fmt.Println(namaDepan)

	// fmt.Println(Variadic(10, 10, 10, 10))

	FungctionAsParameter("andika", Pasangan)
	pacar := Pasangan
	FungctionAsParameter("laki laki", pacar)

}

func MultipleVariabel() {
	var (
		namaDepan    = "Ibnu"
		namaBelakang = "Zaman"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}

func TypeDeclaration() {
	// type noKTP string

	type NIM int

	var nimIbnu NIM = 210018150

	// var noKtp noKTP = "321221"
	// fmt.Println(noKtp)
	// fmt.Println(noKTP("string"))

	fmt.Println(nimIbnu)
	fmt.Println("NIM Mahasiswa Informatika", NIM(2100018))
}

func AugmentedAssignments() {
	var a = 10
	a += 10
	fmt.Println(a)
}

func UnaryOperator() {
	var j = 1
	// j++
	j++
	fmt.Print(j)
}

func Array() {
	// var nama [3]string
	// nama[0] = "ibnu"
	// nama[1] = "odni"
	// nama[2] = "pimak"

	// fmt.Print(nama[0])

	var angka = [3]int{
		11, 33, 44,
	}

	fmt.Print(angka)

	var kalimat = [3]string{
		"satu", "dua", "tiga",
	}

	// fmt.Print(kalimat)
	// kalimat[2] = "ubah"
	fmt.Print(len(kalimat))
}

func Slice() {

	name := [...]int{
		0, 1, 2, 3, 4,
	}
	slice := name[2:4]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
}

func AppendInSlice() {
	data := [...]string{
		"ibnu", "januar", "teja", "kepret", "aji",
	}
	slice1 := data[2:]
	slice2 := append(slice1, "fahrir")

	fmt.Println(slice1)
	fmt.Println(slice2)
}

func MakeInSlice() {
	slice := make([]int, 2, 5)
	slice[0] = 1
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}

func CopyInSlice() {
	data := [...]string{
		"gelas", "meja", "kursi",
	}
	slice1 := data[:]
	slice2 := make([]string, len(slice1), cap(slice1))

	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func Mapping() {
	pepek := map[string]int{
		"nama":   1234,
		"alamat": 41321,
	}
	delete(pepek, "alamat")

	fmt.Println(pepek)
	// fmt.Println(pepek["alamat"])
}

func SwitchKondisi() {

	nama := "d"
	length := len(nama)
	// switch length := len(nama); length > 100 {
	// case true:
	// 	fmt.Println("pepek")
	// case false:
	// 	fmt.Println("salah")
	// 	fmt.Println(len(nama))
	// }
	switch {
	case length > 6:
		fmt.Println("bener")
	case length < 5:
		fmt.Println("kine lima")
	}
}

func IterationWithRange() {
	nama := []string{
		"ibnu",
		"zaman",
	}

	for index, names := range nama {
		fmt.Println("index", index, "=", names)
	}
}

func BreakAndContinue() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			fmt.Println("break")
			continue
		}
		if i == 9 {
			fmt.Println("break")
		}
		fmt.Println("lanjut")
	}
}

func FungsiParameter(nama string) string {
	return "namanya : " + nama
}

func MultipleValues() (string, string) {
	return "ibnu", "zaman"
}

func Variadic(angka ...int) int {
	total := 0
	for _, angka := range angka {
		total += angka
	}
	return total
}

func FungctionAsParameter(cowo string, pacar func(string) string) {
	fmt.Println(cowo, pacar(cowo))
}
func Pasangan(cowo string) string {
	if cowo == "laki laki" {
		return "gay"
	} else {
		return "normal"
	}
}
