
package main

import "fmt"

type Estudiante struct {
    Nombre string
    nota   float64 // minúscula privada
}

func (e Estudiante) Saludar() string {
    return "Hola, soy " + e.Nombre
}

func main() {
    e := Estudiante{Nombre: "Jerry Alberto Serna Cevallos", nota: 9.5}

    fmt.Println(e.Saludar())
}
