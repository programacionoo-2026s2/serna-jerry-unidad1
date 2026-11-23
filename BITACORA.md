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
