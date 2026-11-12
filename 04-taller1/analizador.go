
package main

import (
    "fmt"
    "strings"
    "unicode"
    "unicode/utf8"
)

func main() {
    texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

    fmt.Println("Texto:", texto)

fmt.Println("Bytes:", len(texto))
fmt.Println("Runas:", utf8.RuneCountInString(texto))

palabras := strings.Fields(texto)
fmt.Println("Palabras:", len(palabras))

mayusculas := 0
digitos := 0

for _, caracter := range texto {
    if unicode.IsUpper(caracter) {
        mayusculas++
    }

    if unicode.IsDigit(caracter) {
        digitos++
    }
}

fmt.Println("Mayúsculas:", mayusculas)
fmt.Println("Dígitos:", digitos)

conTilde := 0

for _, caracter := range texto {
    if strings.ContainsRune("áéíóúÁÉÍÓÚñÑ", caracter) {
        conTilde++
    }
}

fmt.Println("Con tilde/ñ:", conTilde)

    _ = utf8.RuneCountInString
}
