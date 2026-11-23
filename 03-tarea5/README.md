# 03-tarea5: Analizador de texto: bytes, runas y palabras

## Descripción

Este programa está desarrollado en Go y analiza textos utilizados como subtítulos de videos.

Cuenta la cantidad de bytes, runas, palabras, letras, dígitos, letras mayúsculas y runas que ocupan más de un byte en UTF-8.

También comprueba si el subtítulo supera el límite de 42 runas, para facilitar la creación de subtítulos accesibles.

## Cómo ejecutar

Desde la carpeta `03-tarea5`, ejecutar:

```bash
go run . "Canción de cuna 2026"
```

El texto debe escribirse entre comillas para que se interprete como un solo argumento.

Para compilar el programa:

```bash
go build -o analizar.exe .
```

En PowerShell se puede ejecutar:

```powershell
.\analizar.exe "Hola mundo"
```

### Códigos de salida

- **0:** ejecución correcta.
- **1:** cantidad incorrecta de argumentos.
- **2:** texto vacío o compuesto únicamente por espacios.
- **3:** el subtítulo supera el límite de 42 runas.

Los errores se muestran mediante `stderr`.

## Pruebas realizadas

| Caso | Texto o situación | Bytes | Runas | Palabras | Letras | Dígitos | Mayúsculas | Runas de más de 1 byte | Código |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | Hola mundo | 10 | 10 | 2 | 9 | 0 | 1 | 0 | 0 |
| 2 | Canción de cuna 2026 | 21 | 20 | 4 | 13 | 4 | 1 | 1 | 0 |
| 3 | ÁÉÍ ñandú | 14 | 9 | 2 | 8 | 0 | 3 | 5 | 0 |
| 4 | Dos espacios al inicio y al final: `  espacios múltiples  ` | 23 | 22 | 2 | 17 | 0 | 0 | 1 | 0 |
| 5 | Tres espacios | — | — | — | — | — | — | — | 2 |
| 6 | Este subtítulo es demasiado largo para leerse bien | 51 | 50 | 8 | 43 | 0 | 1 | 1 | 3 |
| 7 | Sin argumentos | — | — | — | — | — | — | — | 1 |
| 8 | Dos argumentos sin comillas | — | — | — | — | — | — | — | 1 |

Las ocho pruebas se realizaron con el ejecutable compilado y se verificaron sus códigos de salida.

## Preguntas

### a) ¿Por qué "Canción de cuna 2026" tiene 21 bytes y 20 runas?

Porque la letra `ó` ocupa 2 bytes en UTF-8, pero representa una sola runa. Los demás caracteres ocupan un byte cada uno.

Para comprobar el límite de subtítulos utilizaría las **runas**, porque representan los caracteres Unicode del texto y no el espacio que ocupan en memoria.

### b) En `for i, r := range texto`, ¿qué representan `i` y `r`? ¿Qué ocurre con `ñu`?

La variable `i` indica la posición inicial en bytes de cada carácter y `r` contiene la runa Unicode.

En el texto `ñu`, la primera vuelta tiene `i = 0` y `r = 'ñ'`. Como `ñ` ocupa 2 bytes, la segunda vuelta tiene `i = 2` y `r = 'u'`.

### c) ¿Por qué las palabras se cuentan con `enPalabra` en lugar de contar espacios?

Porque puede haber varios espacios consecutivos, al principio o al final del texto. Contar espacios produciría resultados incorrectos.

La variable `enPalabra` permite detectar cuándo comienza una palabra nueva. Por ejemplo, en `  espacios múltiples  ` hay varios espacios, pero solamente **2 palabras**.