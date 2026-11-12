
# Tarea de clase 2 - Permisos estilo Unix con iota

## Objetivo

Desarrollar un programa en Go que represente permisos de archivos al estilo Unix, utilizando iota y operadores a nivel de bits.

## 1. Implementación de la función describir

Primero creé el tipo Permiso utilizando uint8 y definí las constantes Leer, Escribir y Ejecutar mediante iota.

Después implementé la función describir utilizando el operador & para comprobar qué permisos están activados.

La función devuelve una cadena de tres caracteres:
- r: lectura.
- w: escritura.
- x: ejecución.
- -: permiso no concedido.

## 2. Definición de permisos

Establecí los siguientes permisos:

- Usuario: rwx (lectura, escritura y ejecución).
- Grupo: r-x (lectura y ejecución).
- Otros: r-- (solo lectura).

Al imprimir los tres permisos juntos, obtuve:

rwxr-xr--

## 3. Representación numérica

Utilicé fmt.Printf con %d para mostrar el valor numérico de cada conjunto de permisos.

El resultado fue:

numérico: 751

## 4. Modificación de permisos

Utilicé el operador |= para conceder escritura a otros:

otros |= Escribir

Después utilicé el operador &^= para revocar el permiso de ejecución del grupo:

grupo &^= Ejecutar

El resultado actualizado fue:

rwxr--rw-

## 5. Reto: valores tradicionales de chmod

Reorganicé las constantes definidas mediante iota para obtener:

- Leer = 4
- Escribir = 2
- Ejecutar = 1

Para lograrlo, declaré las constantes en el orden Ejecutar, Escribir y Leer.

De esta manera, el programa utiliza los valores tradicionales de permisos Unix y mantiene el mismo resultado.

## 6. Resultado final

Al ejecutar el comando:

go run permisos.go

Obtuve:

rwxr-xr--
numérico: 751
rwxr--rw-

## Conclusión

Esta actividad me permitió comprender cómo representar permisos Unix utilizando operadores a nivel de bits en Go.

Aprendí a comprobar permisos con &, activarlos con |= y desactivarlos con &^=.

También comprendí cómo funciona iota y cómo organizar las constantes para representar los valores utilizados por chmod.
