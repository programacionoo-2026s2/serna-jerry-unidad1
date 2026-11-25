# 05-taller1: Funciones en Go - Estadísticas variádicas

## Descripción

Programa desarrollado en Go que utiliza una función variádica para recibir cualquier cantidad de calificaciones y calcular la nota menor, la mayor y el promedio.

La función devuelve cuatro resultados con nombre y utiliza el error centinela ErrSinDatos cuando no recibe valores.

## Cómo ejecutar

Desde la carpeta 05-taller1, ejecutar:

```bash
go run .
```

El programa imprime tres líneas:

- Primera línea: estadísticas de valores individuales.
- Segunda línea: estadísticas de un slice con dos decimales.
- Tercera línea: mensaje de error cuando no se reciben datos.

## Pruebas realizadas

Resultado de la ejecución:

```text
4 9.75 7.0625 <nil>
min=7.00 max=10.00 prom=8.17
error: no se recibieron datos
```

Los resultados coinciden con la salida esperada del taller.

## Preguntas

### a) ¿Por qué los resultados se llaman menor y mayor y no min y max?

Porque min y max son funciones integradas de Go. Si utilizara esos mismos nombres para los resultados, estaría ocultando las funciones originales.

Por ejemplo, en mi código utilizo:

```go
menor = min(menor, n)
mayor = max(mayor, n)
```

Si el resultado se llamara min, Go intentaría utilizar una variable como si fuera una función y mostraría un error de compilación.

### b) ¿Qué hacen los tres puntos en estadisticas(notas...)? ¿Qué mensaje da Go si los quitas y por qué?

Los tres puntos permiten pasar los elementos del slice como argumentos individuales a la función variádica.

En mi código utilizo:

```go
mn, mx, pr, _ := estadisticas(notas...)
```

Al quitar los tres puntos, Go muestra un error como:

```text
cannot use notas (variable of type []float64) as float64 value in argument to estadisticas
```

Esto sucede porque notas es un slice de tipo []float64, mientras que la función espera recibir valores individuales de tipo float64.