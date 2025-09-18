# Instalación de GO

- Descargar instalador y seguir instrucciones
🔗[Pagina de Go](https://go.dev/dl/)
- Varificar instalación ´go version´
-🔗[Repositorio de Digital House](https://github.com/digitalhouse-content/go-fundamentals-bases/blob/main/functions/function/factory.go)

# Go Path (Estructura de directorios)

- GOPATH (Donde estan los proyectos)
  - bin (binarios)
  - pkg (dependencias internas y externas. Las externas se almacenan en la carpeta mod)
  -src (Se almacenan los proyectos)
- GOROOT (ejecuta el programa)

# Modules & Packages

Los paquetes son carpatas donde se van a almacenar los archivos de codigo Go incluyendo el main. Se debe evitar almacenar un paquete dentro de otro.

Los modulos es la estructura del proyecto el cual contiene todos lo paquetes que hacen que el codigo funcione.

# Variables

|Tipo de variable| Valor por defecto|
|-----------------------------------|
| Number         |        0         |
| String         |        ""        |
| Pointer        |        null      |
| Boolean        |        false     |
| Interface      |        null      |

## Declaracion de variables y asignación de valores

Para declarar una variable se debe utilizar la palabra reservada var seguido del nombre de la variable y el tipo de dato.

Ejemplo: ´var myIntVar int´

Otra forma de declararle y asignar el valor es con el nombre de la variable seguido de := y el valor de la variable.

Ejemplo: ´myVarInt := 10´

Con el & se puede acceder a la direccion de memoria donde la variable es almacenada.

# Constantes

Este tipo de dato no se puede modificar y se declara similar a las variables, solo que en vez de utilizar la palabra reservada var se usa la palabra const.

Ejemplo: ´const myIntConst int = 10´

# Bloques

Son segmentos de codigo donde se pueden declarar variables que no se necesitan durante todo el ciclo de ejecución del proyecto.
Ejemplo: ´{
  myScopeVar := 40
  fm.println("Variable scope: ", myScopeVar)
}´

# String con salto de linea

Para tener un texto con saltos de linea, a la variable en vez de usar comilla doble ("") para su asignación se usa la comilla sencilla ('')

Ejemplo: ´myStrVar := 'Este es un texto de
multiples lineas.'´

# Slices
Son similares a los arreglos, con la diferencia de que su tamaño es variable y se puede modificar durante el tiempo de ejecución. Se pueden asignar valores desde un array o variable existente y su asignación es realizada por referencia por lo que si se modifica el valor en el slice se modifica tambien en el arreglo o variable original.

# Funciones

Son secciones de código orientados a acciones que pueden repetirse mas de una vez y evitar tener que escribirlo varias veces y redicir linea de código.

Las funciones publicas son nombradas con la primera letra en mayuscula, mientras que las funciones privadas se declaran con la primera letra del nombre en minuscula.

# Llamar las funciones externas

- Desde la carpeta raiz en el CMD ejecutar el comando ´go mod init nombrepaquete´´
- Importar en el lugar donde se va a utilizar ingresando la ruta donde estan las funciones a utilizar.

# Estructuras

Bloques de codigo donde se almacenan variables y metodos correspondientes a una entidad, por ejemplo User tenfria las variables nombre, apellido, edad y atributos para obtenerlos.

# Convenciones para nombrar elementos

Existen algunas convenciones que podemos utilizar para nombrar variables, funciones y otros elementos en la programación.

1. Camel Case: las palabras se escriben juntas sin espacios y cada palabra después de la primera comienza con una letra mayúscula (excepto la primera palabra).
Ejemplo: miVariableEjemplo o calcularSalarioNeto

2. Snake Case: las palabras se escriben en mayúsculas y se separan mediante guiones bajos (underscore).
Ejemplo: mi_variable_ejemplo o calcular_salario_neto

3. Pascal Case: Es similar a Camel Case, pero en lugar de comenzar con una letra mayúscula, la primera letra de cada palabra se inicia en mayúscula.
Ejemplo: MiVariableEjemplo o CalcularSalarioNeto

4. Kebab Case:
las palabras se escriben en minúsculas y se separan mediante guiones. Aunque no es tan común en la programación, se usa en ciertos contextos, especialmente en la definición de la URLs y nombres de archivos web.
Ejemplo:
mi-variable-ejemplo o calcular-salario-neto


# Format

go fmt es una herramienta en Go (Golang) que se utiliza para formatear automáticamente el código fuente de tus programas según las convenciones de estilo establecidas por la comunidad Go. Su objetivo es mantener la coherencia y la legibilidad del código al aplicar un formato uniforme a lo largo de tu proyecto. Aquí tienes una explicación sobre cómo usar go fmt y por qué es útil:

## Uso de go fmt

Instalación de Go: Asegúrate de tener Go instalado en tu sistema. Si no lo tienes, puedes descargarlo desde el sitio oficial de Go (https://golang.org/dl/).

Ubicación del código: Abre una terminal y navega hasta el directorio que contiene tu código Go.

Ejecución de go fmt: Ejecuta el siguiente comando en la terminal:

´go fmt ./...´

El ./... indica que go fmt formateará todos los archivos en todos los subdirectorios dentro del directorio actual.
Verificación del resultado
: Una vez que
go fmt
haya terminado, podrás ver los cambios realizados en tus archivos. Los archivos que necesitaban formateo se actualizarán automáticamente con el estilo de formato Go estándar.

# Pointers

Al igual que los slices trabajan por referencia, por lo que si se realizan cambios al llamar la dirección de memoria se veran reflejados en la variable original.

# Interfaces

Sirve para instanciar distintos objetos con atributos y metodos en común.

# Errors

## Defer:

Esta palabra reservada funciona para recuperar la aplicacion de errores que se pueden generar durante la ejecución.

## Panic

Esta palabra reservada puede ser utilizada para arrojar un error cuando no se cumple la logica.

# Package

## Internos

Paquetes disponibles en el momento que se instala Go

## Time

Determina fechas y horas y realiza operaciones con ellas.

## OS

Aprovecha las funcionalidades del sistema operativo con las que se puede realizar operaciones como la lectura de archivos, configuracion de bases de datos etc...

## Log

Packete con el que se puede llevar el registro de errores que se presentan en el programa.

## Externos

Son paquetes que se requieren intalar de manera independiente para poder ser utilizados en el programa. Estos pueden estar almacenados en un repositorio como GitHub o Gitlab o ser importado de otras fuentes.

## UUID

Para instalarlo se debe utilizar el comando
´go get github.com/google/uuid´

Se utiliza para generar una identificación para diferentes usuarios o entidades que se van a utilizar en la aplicación.

## Custom

Librerias predeterminadas que se pueden subir a un repositorio y ustilizarlas en  nuevos proyectos.

# Go Routines

Permiten crear hilos o procesos para ejecutarse de manera simultanea, la cantidad de hilos que se pueden crear depende del numero de nucleos que se tengan en el equipo.

Estas Goroutines finalizan en el momento que termina el proceso main. Cuando se utilizan canales, el proceso del hilo termina cuando se le es asignado un valor al canal y este debe ser cerrado al terminar el proceso.

# Generics

Sirven para especificar los tipos de datos que se quieren recibir en una función o interfaz.