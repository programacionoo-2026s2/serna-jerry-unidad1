# Bitácora de trabajo - Unidad 1

**Estudiante:** Jerry Alberto Serna Cevallos
**Asignatura:** Programación Orientada a Objetos

## Registro de sesiones

En este archivo se documentarán las sesiones de trabajo de cada tarea y taller de la Unidad 1.

Cada registro incluirá:
- Fecha y actividad realizada.
- Qué hice.
- Qué problema encontré.
- Cómo lo resolví.
- Próximo paso.

## Sesiones


### Sesión - 18 de noviembre de 2026

**Actividad:** Desarrollo de la tarea 03-tarea3: Conversor de almacenamiento por línea de comandos.

**Qué hice:**
- Creé la carpeta 03-tarea3.
- Inicialicé el módulo de Go con `go mod init tarea3`.
- Desarrollé el programa main.go para convertir B, KB, MB y GB.
- Implementé validaciones de argumentos, números y unidades.
- Utilicé códigos de salida 0, 1, 2 y 3.
- Compilé el programa y realicé ocho pruebas.
- Documenté el funcionamiento en el README.md de la tarea.

**Problema encontrado:**
Al intentar ejecutar `go mod init tarea3`, PowerShell mostró un error indicando que el comando `go` no se reconocía.

**Cómo lo resolví:**
Configuré temporalmente la variable PATH para incluir la carpeta donde estaba instalado Go. Después ejecuté `go version` para comprobar que funcionaba correctamente.

**Próximo paso:**
Completar los commits pendientes en sesiones distintas, verificar los archivos del repositorio y subir la tarea a GitHub.


## 03-tarea4: Del monolito a los paquetes

**Fecha:** 22 de noviembre de 2026

### Trabajo realizado
- Creé el módulo `tarea4` y reutilicé el conversor de la tarea anterior como versión monolítica.
- Separé la lógica de conversión en el paquete `internal/almacenamiento`.
- Implementé las constantes exportadas `B`, `KB`, `MB` y `GB`.
- Creé las funciones exportadas `Bytes` y `Convertir`, y la función privada `redondear`.
- Actualicé `main.go` para importar el nuevo paquete.
- Ejecuté los ocho casos oficiales y obtuve los resultados y códigos de salida esperados.

### Error encontrado y solución
Intenté utilizar `almacenamiento.redondear` desde `main.go`, pero el compilador mostró:

`undefined: almacenamiento.redondear`

Comprendí que las funciones que comienzan con minúscula no se pueden utilizar desde otros paquetes. Eliminé la llamada temporal y comprobé que el programa volviera a compilar correctamente.

### Aprendizaje
Aprendí a organizar un programa de Go en paquetes, utilizar rutas de importación, distinguir identificadores exportados y privados, y comprobar que una reorganización del código no cambie su funcionamiento.

## 03-tarea5: Analizador de texto

**Actividad realizada:** Desarrollo de un analizador de texto en Go.

**Avances:**
- Creación del módulo y archivo main.go.
- Implementación del conteo de bytes, runas, palabras, letras, dígitos y mayúsculas.
- Validación del límite de 42 runas.
- Ejecución exitosa de las 8 pruebas oficiales.
- Elaboración del README con instrucciones, resultados y respuestas.

**Problema encontrado:**
Inicialmente creé la carpeta 03-tarea5 fuera del repositorio principal.

**Solución aplicada:**
Moví la carpeta al repositorio serna-jerry-unidad1 y comprobé que el programa funcionara correctamente desde su ubicación definitiva.

**Aprendizaje:**
Comprendí la diferencia entre bytes y runas en UTF-8 y aprendí a contar palabras correctamente, incluso cuando existen espacios consecutivos.
## 04-taller1: FizzBuzz con switch sin condición

**Actividad realizada:** Desarrollo de un programa FizzBuzz en Go utilizando un switch sin condición.

**Avances:**
- Creé el módulo de Go y el archivo main.go.
- Implementé un bucle for para recorrer los números del 1 al 30.
- Utilicé un switch sin condición, sin if/else.
- Implementé contadores para Fizz, Buzz y FizzBuzz.
- Verifiqué los resultados: Fizz=8, Buzz=4 y FizzBuzz=2.
- Realicé un experimento cambiando el orden de los case.
- Documenté los resultados y respondí las preguntas del taller.

**Problema encontrado:**
Al modificar el orden de los case, eliminé accidentalmente una llave de cierre y el programa mostró un error de sintaxis.

**Solución aplicada:**
Revisé la estructura del código, corregí las llaves y ejecuté nuevamente el programa. Después del experimento restauré el orden original.

**Aprendizaje:**
Comprendí que Go ejecuta el primer case verdadero de un switch y que el orden de las condiciones afecta el resultado. También aprendí que no es necesario escribir break al finalizar cada case.
## 05-taller1: Estadísticas variádicas
Creé el módulo, declaré la función variádica y comprobé el error centinela ErrSinDatos.

**Avances:**
- Implementé el cálculo de la nota menor, mayor y promedio.
- Utilicé las funciones integradas min y max.
- Comprobé el funcionamiento con valores individuales y slices.
- Verifiqué el error centinela mediante errors.Is.

**Aprendizaje:**
Comprendí cómo funcionan las funciones variádicas y los resultados con nombre. También aprendí que los tres puntos permiten pasar los elementos de un slice como argumentos individuales.

**Pruebas:**
Las tres líneas de salida coinciden con los resultados esperados.
