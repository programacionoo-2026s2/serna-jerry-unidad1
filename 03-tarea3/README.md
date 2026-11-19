
# 03-tarea3: Conversor de almacenamiento

## Descripción

Desarrollar un programa en Go que permita convertir unidades de almacenamiento desde la línea de comandos, aplicando validaciones y códigos de salida.

## Unidades disponibles

El programa permite convertir entre:

- B: Bytes
- KB: Kilobytes
- MB: Megabytes
- GB: Gigabytes

Se utiliza la equivalencia binaria:

- 1 KB = 1024 B
- 1 MB = 1024 KB
- 1 GB = 1024 MB

## Tecnologías utilizadas

- Lenguaje Go
- Visual Studio Code
- PowerShell
- Git y GitHub

## Funcionamiento del programa

El programa recibe tres argumentos:

1. Valor numérico.
2. Unidad de origen.
3. Unidad de destino.

Primero valida la cantidad de argumentos y el valor ingresado.

Después verifica que las unidades sean válidas.

Finalmente, convierte el valor a bytes y luego a la unidad solicitada.

## Cómo ejecutar

Para ejecutar el programa desde la carpeta 03-tarea3:

    go run . 2048 KB MB

Resultado:

    2048 KB = 2.00 MB

También se puede compilar:

    go build -o convertir.exe .

Y ejecutar en Windows:

    .\convertir.exe 2048 KB MB


### Verificación de códigos de salida en PowerShell

Para comprobar los códigos de salida reales del programa, primero se compila el ejecutable:

```powershell
go build -o convertir.exe .
```

Después se ejecuta una prueba y se consulta el código de salida:

```powershell
.\convertir.exe 2048 KB MB
$LASTEXITCODE
```

El código `0` indica que la conversión fue exitosa.

También se pueden comprobar los errores:

```powershell
.\convertir.exe abc KB MB
$LASTEXITCODE
```

El código `2` indica que el valor ingresado no es válido.

Los códigos utilizados son:

- `0`: ejecución correcta.
- `1`: cantidad incorrecta de argumentos.
- `2`: valor no numérico o negativo.
- `3`: unidad de almacenamiento no válida.

## Códigos de salida

- 0: Ejecución correcta.
- 1: Cantidad incorrecta de argumentos.
- 2: Valor no numérico o negativo.
- 3: Unidad de almacenamiento no válida.


## Pruebas realizadas

| N.º | Comando | Resultado esperado | Código de salida |
|---|---|---|---|
| 1 | `go run . 2048 KB MB` | 2048 KB = 2.00 MB | 0 |
| 2 | `go run . 1 gb mb` | 1 GB = 1024.00 MB | 0 |
| 3 | `go run . 1500 B KB` | 1500 B = 1.46 KB | 0 |
| 4 | `go run . 0.5 MB KB` | 0.5 MB = 512.00 KB | 0 |
| 5 | `go run . abc KB MB` | Error: valor no numérico | 2 |
| 6 | `go run . -5 KB MB` | Error: valor negativo | 2 |
| 7 | `go run . 5 TB MB` | Error: unidad no válida | 3 |
| 8 | `go run . 5 KB` | Error: cantidad incorrecta de argumentos | 1 |

Las capturas de las pruebas se encuentran en la carpeta evidencias.

## Conclusión

Esta tarea me permitió aprender a desarrollar un programa de línea de comandos en Go.

Utilicé os.Args para recibir argumentos, strconv.ParseFloat para convertir valores numéricos y strings.ToUpper para validar unidades sin distinguir entre mayúsculas y minúsculas.

También aprendí a utilizar os.Exit para manejar diferentes códigos de salida y comprobar el funcionamiento del programa mediante pruebas.

## Preguntas

### a) ¿Por qué el valor está en os.Args[1] y no en os.Args[0]?

Porque os.Args[0] contiene el nombre o la ruta del programa que se está ejecutando. Los argumentos ingresados por el usuario comienzan desde os.Args[1]. Por eso, el valor numérico se encuentra en esa posición.

### b) ¿Por qué los mensajes de error se escriben en stderr y no con fmt.Println?

Porque stderr es el canal destinado a los errores, mientras que stdout se utiliza para los resultados normales. Esto permite que otros programas o scripts distingan los resultados correctos de los mensajes de error.

### c) ¿Qué pasaría si ignoraras el error de strconv.ParseFloat?

Si ignorara el error, el programa podría continuar utilizando un valor incorrecto, como cero, cuando el usuario ingrese texto en lugar de un número. Esto produciría resultados engañosos en vez de informar que la entrada no es válida.

### d) ¿Qué ventaja tiene declarar los códigos de salida como constantes?

Las constantes permiten utilizar nombres descriptivos en lugar de números repetidos. Esto facilita la lectura, el mantenimiento del código y evita errores al utilizar los códigos de salida.
