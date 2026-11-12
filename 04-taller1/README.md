
# Taller 4 - Analizador de texto

## Objetivo

Desarrollar un programa en Go que analice una cadena de texto utilizando las bibliotecas strings, unicode y unicode/utf8.

## 1. Estrategia de desarrollo

Primero creé el archivo analizador.go dentro de la carpeta 04-taller1.

Después declaré una variable llamada texto con la siguiente cadena:

"Programación en Go: rápido, simple y ¡divertido! Año 2026."

Implementé cada requisito por separado y ejecuté el programa para comprobar los resultados.

## 2. Conteo de bytes y runas

Utilicé len(texto) para contar los bytes y utf8.RuneCountInString(texto) para contar las runas.

Resultados:

- Bytes: 62
- Runas: 58

## 3. Conteo de palabras

Utilicé strings.Fields(texto) para separar las palabras y len(palabras) para contarlas.

Resultado:

- Palabras: 9

## 4. Conteo de mayúsculas y dígitos

Recorrí el texto utilizando range.

Con unicode.IsUpper identifiqué las letras mayúsculas y con unicode.IsDigit identifiqué los dígitos.

Resultados:

- Mayúsculas: 3
- Dígitos: 4

## 5. Conteo de letras con tilde o ñ

Utilicé strings.ContainsRune para comprobar si cada carácter pertenece al conjunto de vocales acentuadas o letras ñ y Ñ.

Resultado:

- Con tilde/ñ: 3

## 6. Diferencia entre bytes y runas

Los bytes y las runas no coinciden porque Go utiliza la codificación UTF-8, donde algunos caracteres ocupan más de un byte.

En el texto analizado, los caracteres ó, á, ¡ y ñ ocupan dos bytes cada uno, pero representan una sola runa.

Por eso el texto contiene 62 bytes y 58 runas.

## 7. Resultado final

Al ejecutar el comando:

go run analizador.go

Obtuve los siguientes resultados:

Bytes: 62
Runas: 58
Palabras: 9
Mayúsculas: 3
Dígitos: 4
Con tilde/ñ: 3

## Conclusión

Este taller me permitió comprender cómo analizar cadenas de texto en Go.

Aprendí a contar bytes, runas y palabras, además de identificar letras mayúsculas, dígitos y caracteres acentuados.

También comprendí la importancia de Unicode y UTF-8 para trabajar correctamente con textos que contienen caracteres especiales.
