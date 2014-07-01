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

Con el comando ` go run ` procesa el programa, lo compila en un ejecutable, el cual es almacenado en un directorio temporal, para luego ser ejecutado.  

`$ go run hola-mundo.go`

Ejecutar la anterior línea de código en el terminal, debería imprimir la cadena: Hola Mundo

Analicemos su estrutura: 

La primera linea: 

`package main`




