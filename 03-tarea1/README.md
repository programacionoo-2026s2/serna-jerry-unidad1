
# Tarea de clase 1 - Programa suma.go

## 1. ¿Qué imprime el programa original?

El programa original imprime:

total: 0

Aunque debería sumar los números del 1 al 3, el resultado es 0.

## 2. ¿Por qué sucede? (Ámbito)

Esto sucede porque dentro del ciclo for se utiliza :=, que declara una nueva variable llamada total.

Esta nueva variable tiene un ámbito limitado al bloque del ciclo for. Por eso, la variable total que fue declarada fuera del ciclo no cambia y mantiene su valor inicial de 0.

## 3. Corrección del programa

Para corregir el problema, se eliminó el carácter : de la siguiente línea:

Antes:
total := total + i

Después:
total = total + i

Con este cambio, el programa modifica la variable original en lugar de crear una nueva.

El resultado corregido es:

total: 6

## 4. Análisis con go vet

Primero ejecuté el comando:

go vet ./...

Este comando mostró un error porque la carpeta no contiene un módulo de Go (go.mod).

Después ejecuté:

go vet suma.go

Este segundo comando terminó sin mostrar advertencias.

Esto significa que go vet no detectó problemas en el archivo analizado. Sin embargo, el analizador shadow, que permite identificar variables que ocultan a otras variables del mismo nombre, no está incluido en los análisis predeterminados de go vet.

## 5. Investigación sobre el analizador shadow

El analizador shadow es una herramienta de análisis estático de Go que permite detectar posibles casos de ocultamiento de variables.

Esto sucede cuando se declara una variable dentro de un ámbito más pequeño con el mismo nombre que otra variable existente.

En el programa original suma.go, la instrucción:

total := total + i

crea una nueva variable total dentro del ciclo for, ocultando la variable declarada fuera del ciclo.

Esto provoca que la variable original mantenga su valor de 0.

El analizador shadow puede ayudar a identificar este tipo de problemas, aunque no forma parte de los análisis que go vet ejecuta de manera predeterminada.

Para corregir el problema se reemplazó := por =, permitiendo actualizar la variable original y obtener el resultado correcto de 6.

Fuente:
https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/shadow
