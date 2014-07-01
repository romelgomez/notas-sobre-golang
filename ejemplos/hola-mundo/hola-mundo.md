Hola Mundo
==========

Este programa imprime Hola Mundo en el terminal:  
 
```go
  package main
  import "fmt"
  func main() {
    fmt.Println("Hola Mundo")
  }
```

[/ejemplos/hola-mundo/hola-mundo.go] (/ejemplos/hola-mundo/hola-mundo.go)

El comando `go run` procesa el programa, este lo compila en un ejecutable, el cual es almacenado en un directorio temporal, para luego ser ejecutado.  

`$ go run hola-mundo.go`

Ejecutar la anterior línea de código en el terminal, debería imprimir la cadena: Hola Mundo

### Analicemos su estrutura: ###

  La primera línea: `package main` es conocida como la declaración del paquete, cada programa GO debe empezar con la declaración del paquete. Los paquetes son la forma en que GO organiza
y reutiliza el código. Hay dos tipos de programas: Ejecutables y librerías. Las aplicaciones ejecutables son los programas que pueden ser ejecutados directamente desde la terminal 
(tales como .deb en ubuntu), las librerias son coleciones de codigo que empaquetamos en conjunto para que podamos usarlas en otros programas.  

  La segunda línea  `import "fmt"`, consta de dos partes, la primera, la palabra clave `import` es como incluimos código de otros paquetes para usar en nuestros programas, la segunda, 
`"fmt"` es el paquete que incluimos, el paquete `fmt` (abreviatura de formato) implementa el formato para la entrada y salida de datos.

La tercera línea observamos la declaración de la función: 

```go
  func main() {
    fmt.Println("Hola Mundo")
  }
```
  Las funciones son los componentes básicos de un programa GO. Ellas tienen entradas, salidas y una series de pasos llamados declaraciones las cuales son ejecutadas en orden. Todas las 
funciones empiezan con la palabra clave `func` seguida del nombre de la función (`main` en este caso), un lista de cero o más "parámetros" envueltos por parientes, y un tipo de retorno
opcional y un "cuerpo" el cual está envuelto por llaves. Esta función no tiene parámetros, no retorna nada y tiene solo una declaración. El nombre `main` es especial porque esta es la 
función que es llamada cuando ejecute el programa. 
  
La cuarta y última pieza del programa es esta línea: 

```go
  fmt.Println("Hola Mundo")
```

En esta declaración accedemos a otra función dentro del paquete `fmt` llamado `Println` (Println significa imprimir en una línea). Para luego crear la cadena `Hola Mundo`.
