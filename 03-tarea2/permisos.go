package main

import "fmt"

type Permiso uint8

const (
	Ejecutar Permiso = 1 << iota
	Escribir
	Leer
)

// describir(Leer|Ejecutar) == "r-x"

func describir(p Permiso) string {
	resultado := "-"

	if p&Leer != 0 {
		resultado = "r"
	}

	if p&Escribir != 0 {
		resultado += "w"
	} else {
		resultado += "-"
	}

	if p&Ejecutar != 0 {
		resultado += "x"
	} else {
		resultado += "-"
	}

	return resultado
}

func main() {
	usuario := Leer | Escribir | Ejecutar
	grupo := Leer | Ejecutar
	otros := Leer

	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))
	fmt.Printf("numérico: %d%d%d\n", usuario, grupo, otros)

	// Conceder permiso de escritura a otros
	otros |= Escribir

	// Revocar permiso de ejecución al grupo
	grupo &^= Ejecutar

	// Mostrar los permisos actualizados
	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))

}
