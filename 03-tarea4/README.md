
# 03-tarea4: Del monolito a los paquetes

## Descripción y estructura

Este proyecto reorganiza el conversor de almacenamiento desarrollado en la tarea 03-tarea3. Inicialmente, toda la lógica estaba en `main.go` (versión monolítica). Después se separaron las operaciones de conversión en un paquete interno.

```text
03-tarea4/
├── main.go
├── go.mod
├── README.md
├── internal/
│   └── almacenamiento/
│       └── almacenamiento.go
└── evidencias/
```

- `main.go`: recibe argumentos, valida valores, muestra resultados y controla los códigos de salida.
- `go.mod`: define el módulo `tarea4`.
- `almacenamiento.go`: contiene las constantes B, KB, MB y GB, las funciones exportadas Bytes y Convertir, y la función privada redondear.
- `evidencias/`: contiene las capturas de las pruebas.

## Cómo ejecutar

Desde la carpeta `03-tarea4`:

```powershell
go run . 2048 KB MB
```

También se puede compilar y ejecutar:

```powershell
go build -o convertir.exe .
.\convertir.exe 2048 KB MB
$LASTEXITCODE
```

Los códigos de salida son:

- 0: ejecución correcta.
- 1: cantidad incorrecta de argumentos.
- 2: valor no numérico o negativo.
- 3: unidad de almacenamiento no válida.

Los mensajes de error se muestran en stderr.

## Pruebas realizadas

Se ejecutaron ocho casos con la versión modular:

| Entrada | Resultado | Código |
|---|---|---|
| 2048 KB MB | 2048 KB = 2.00 MB | 0 |
| 1 gb mb | 1 GB = 1024.00 MB | 0 |
| 1500 B KB | 1500 B = 1.46 KB | 0 |
| 0.5 MB KB | 0.5 MB = 512.00 KB | 0 |
| abc KB MB | Error: el valor debe ser un número válido. | 2 |
| -5 KB MB | Error: el valor no puede ser negativo. | 2 |
| 5 TB MB | Error: unidad no válida. Use B, KB, MB o GB. | 3 |
| 5 KB | Uso: convertir &lt;valor&gt; &lt;unidad_origen&gt; &lt;unidad_destino&gt; | 1 |

Los ocho casos produjeron los resultados esperados, conservando el comportamiento de la versión monolítica.

## Experimento del compilador

Se intentó llamar a la función privada `redondear` desde `main.go` mediante:

```go
_ = almacenamiento.redondear(3.14159)
```

Al ejecutar `go build .`, el compilador mostró:

```text
# tarea4
.\main.go:49:21: undefined: almacenamiento.redondear
```

El error ocurre porque `redondear` comienza con minúscula y no está exportada. Por eso no puede utilizarse desde otro paquete.

Después del experimento se eliminó la llamada temporal y el proyecto volvió a compilar correctamente.

## Preguntas

**a) ¿Por qué Bytes y Convertir empiezan con mayúscula y redondear con minúscula? ¿Qué ganas al dejar redondear oculta?**

En Go, los identificadores que empiezan con mayúscula son exportados y pueden utilizarse desde otros paquetes. Los que comienzan con minúscula son privados del paquete. Mantener redondear oculta protege los detalles internos y permite ofrecer una interfaz más sencilla.

**b) ¿Qué mensaje dio el compilador al usar almacenamiento.redondear desde main.go y qué regla de Go lo explica?**

El compilador mostró `undefined: almacenamiento.redondear`. Esto ocurre porque Go no permite acceder a identificadores no exportados desde otro paquete.

**c) ¿Para qué sirve que el paquete esté dentro de internal/ y qué pasaría si otro módulo intentara importarlo?**

La carpeta internal restringe quién puede importar el paquete. En este proyecto, el código del módulo tarea4 puede utilizarlo, pero otro módulo externo no puede importarlo directamente. Esto ayuda a proteger la implementación interna del proyecto.
