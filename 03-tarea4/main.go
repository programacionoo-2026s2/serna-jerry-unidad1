package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"tarea4/internal/almacenamiento"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil || math.IsNaN(valor) || math.IsInf(valor, 0) {
		fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número válido.")
		os.Exit(ErrValor)
	}

	if valor < 0 {
		fmt.Fprintln(os.Stderr, "Error: el valor no puede ser negativo.")
		os.Exit(ErrValor)
	}

	unidadOrigen := strings.ToUpper(os.Args[2])
	unidadDestino := strings.ToUpper(os.Args[3])

	_, origenValido := almacenamiento.Bytes(unidadOrigen)
	_, destinoValido := almacenamiento.Bytes(unidadDestino)

	if !origenValido || !destinoValido {
		fmt.Fprintln(os.Stderr, "Error: unidad no válida. Use B, KB, MB o GB.")
		os.Exit(ErrUnidad)
	}

	resultado, _ := almacenamiento.Convertir(valor, unidadOrigen, unidadDestino)

	fmt.Printf("%g %s = %.2f %s\n",
		valor, unidadOrigen, resultado, unidadDestino)
}
