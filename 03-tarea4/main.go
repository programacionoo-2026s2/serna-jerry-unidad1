
package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
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

	
unidadesValidas := func(unidad string) bool {
    switch unidad {
    case "B", "KB", "MB", "GB":
        return true
    default:
        return false
    }
}

	if !unidadesValidas(unidadOrigen) || !unidadesValidas(unidadDestino) {
		fmt.Fprintln(os.Stderr, "Error: unidad no válida. Use B, KB, MB o GB.")
		os.Exit(ErrUnidad)
	}

	factores := map[string]float64{
		"B":  1,
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
	}

	valorEnBytes := valor * factores[unidadOrigen]
	resultado := valorEnBytes / factores[unidadDestino]

	
fmt.Printf("%g %s = %.2f %s\n",
    valor, unidadOrigen, resultado, unidadDestino)

}
