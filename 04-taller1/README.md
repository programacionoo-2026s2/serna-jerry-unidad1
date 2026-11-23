# 04-taller1: FizzBuzz con switch sin condición

## Descripción

Este programa está desarrollado en Go y utiliza un bucle `for` para recorrer los números del 1 al 30.

Mediante un `switch` sin condición, comprueba si cada número es múltiplo de 15, 3 o 5.

- Si es múltiplo de 15, imprime `FizzBuzz`.
- Si es múltiplo de 3, imprime `Fizz`.
- Si es múltiplo de 5, imprime `Buzz`.
- Si no cumple ninguna condición, imprime el número.

El programa también utiliza contadores para registrar cuántas veces aparece cada palabra.

## Cómo ejecutar

Abrir una terminal dentro de la carpeta `04-taller1` y ejecutar:

```bash
go run .
```

## Resultado

Salida del programa original:

```text
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
Fizz=8 Buzz=4 FizzBuzz=2
```

### Experimento: cambiar el orden de los case

Se modificó temporalmente el orden del `switch`, colocando `case i%3 == 0` antes de `case i%15 == 0`.

La salida obtenida fue:

```text
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 Fizz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 Fizz
Fizz=10 Buzz=4 FizzBuzz=0
```

Después del experimento, se restauró el orden original para que el programa funcionara correctamente.

## Preguntas

### a) ¿Qué salida y qué conteos obtuviste al poner el case de 3 antes que el de 15, y por qué FizzBuzz dejó de aparecer?

Obtuve `Fizz=10 Buzz=4 FizzBuzz=0`.

Esto ocurrió porque Go revisa los `case` en orden y ejecuta únicamente el primero que cumple la condición. Los números 15 y 30 también son múltiplos de 3, por lo que se imprimió `Fizz` en lugar de `FizzBuzz`.

Por esta razón, el contador de Fizz aumentó de 8 a 10 y el de FizzBuzz disminuyó de 2 a 0.

### b) Cuando i vale 15, ¿cuántos case son verdaderos y cuál ejecuta Go? ¿Por qué no hace falta escribir break al final de cada case?

Cuando `i` vale 15, las tres condiciones son verdaderas porque 15 es divisible entre 15, 3 y 5.

Sin embargo, Go ejecuta únicamente el primer `case` verdadero. En el programa original ejecuta `case i%15 == 0`, por lo que imprime `FizzBuzz`.

No hace falta escribir `break` porque Go termina automáticamente el `switch` después de ejecutar el caso seleccionado, salvo que se utilice explícitamente `fallthrough`.