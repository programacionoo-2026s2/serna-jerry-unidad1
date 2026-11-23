// Package almacenamiento proporciona funciones para convertir
// unidades de almacenamiento digital.
package almacenamiento

import "math"

// Constantes que representan la cantidad de bytes de cada unidad.
const (
	B  = 1.0
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// Bytes devuelve la cantidad de bytes correspondiente a una unidad.
// El segundo valor indica si la unidad es válida.
func Bytes(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return B, true
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	default:
		return 0, false
	}
}

// Convertir transforma un valor de una unidad de almacenamiento a otra.
// Devuelve el resultado y un booleano que indica si las unidades son válidas.
func Convertir(valor float64, origen, destino string) (float64, bool) {
	bytesOrigen, validoOrigen := Bytes(origen)
	bytesDestino, validoDestino := Bytes(destino)

	if !validoOrigen || !validoDestino {
		return 0, false
	}

	resultado := valor * bytesOrigen / bytesDestino
	return redondear(resultado), true
}

// redondear aproxima un número a dos decimales.
// Es una función privada porque comienza con minúscula.
func redondear(valor float64) float64 {
	return math.Round(valor*100) / 100
}
