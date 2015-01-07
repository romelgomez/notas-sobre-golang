# Cómo escribir código Go

1. [Introducción](#1-introducción)
2. [Organización del código](#2-organización-del-código)
  1. [Espacios de trabajo](#1-espacios-de-trabajo)
  2. [La variable de entorno GOPATH](#2-la-variable-de-entorno-gopath)
  3. [Rutas de paquetes](#3-rutas-de-paquete)
  4. [Su primer programa](#4-su-primer-programa)
  5. [Su primer librería](#5-su-primera-librería)
  6. [Nombre de los paquetes](#6-nombre-de-los-paquetes)
3. [Pruebas](#3-pruebas)
4. [Paquetes remotos](#4-paquetes-remotos)
5. [Que sigue](#5-que-sigue)
6. [Obtener Ayuda](#6-obtener-ayuda)

## 1. Introducción

Este documento demuestra el desarrollo de una simple paquete Go y presenta la herramienta Go (go tool), la forma estándar de extraer, construir, y instalar paquetes Go y comandos. 

La herramienta Go requiere que organice su código en una manera específica, Por favor lee este documento atentamente. Esta explica de forma simple como empezar a trabajar con su instalación Go

Una similar explicación esta disponible como un [screencast, en ingles](http://www.youtube.com/watch?v=XCsL89YtqCs).

## 2. Organización del código

### 1. Espacios de trabajo
 
  La herramienta go esta diseñada para trabajar con código abierto mantenido en repositorios públicos. Aunque usted no necesita publicar su código, el modelo para como el entorno está configurado trabaja de la misma forma si lo hace o no. 
       
  El código Go debe mantenerse dentro de un espacio de trabajo. Un espacio de trabajo es una jerarquía de directorios con tres directorios en su raíz:
  
    1. `src` contiene archivos de código fuente organizados en paquetes (un paquete por directorio).
    2. `pkg` contiene objetos de paquetes, y ...
    3. `bin` contiene comandos ejecutables. 
    
  La herramienta go construye paquetes fuentes y instala el binario resultante en los directorios pkg y bin. 
    
  El subdirectorio `src` típicamente contiene múltiples repositorios de control de versiones (tales como por ejemplo Git o Mercurial) que siguen el desarrollo de uno o más paquetes fuentes.
    
  Para darte una idea de como un espacio de trabajo debe lucir en la práctica, aquí un ejemplo: 

```
bin/
      streak                         # command executable
      todo                           # command executable
  pkg/
      linux_amd64/
          code.google.com/p/goauth2/
              oauth.a                # package object
          github.com/nf/todo/
              task.a                 # package object
  src/
      code.google.com/p/goauth2/
          .hg/                       # mercurial repository metadata
          oauth/
              oauth.go               # package source
              oauth_test.go          # test source
      github.com/nf/
          streak/
              .git/                  # git repository metadata
              oauth.go               # command source
              streak.go              # command source
          todo/
              .git/                  # git repository metadata
              task/
                  task.go            # package source
              todo.go                # command source  
```
  
  Este espacio de trabajo contiene tres repositorios (goauth2, streak, y todo) que comprenden dos comandos (streak y todo) y dos librerías (oauth and task).
    
  Los comandos y librerías se construyen a partir de diferentes tipos de paquetes de código. Discutiremos esta distinción [luego](http://golang.org/doc/code.html#PackageNames).

### 2. La variable de entorno GOPATH

  La variable de entorno especifica la locación de su espacio de trabajo. Esta es tal vez la única variable de entorno que necesitas establecer para desarrollar código Go.
    
  Para empezar, un directorio de espacio de trabajo y establece GOPATH en consecuencia. Su espacio de trabajo puede ser localizado donde usted desee, pero usaremos $HOME/go en este documento. Note que esta no debe ser la misma ruta que de la instalación de Go.

```
  $ mkdir $HOME/go
  $ export GOPATH=$HOME/go
```
  
  Por conveniencia, añada el subdirectorio `bin` del espacio de trabajo a su PATH 

```
  $ export PATH=$PATH:$GOPATH/bin
```  
  
### 3. Rutas de paquete
  
  Los paquetes de la librería estándar se mencionan como cortas rutas tales como "fmt" y "net/http". Para sus propios paquetes, debe escoger una ruta base que no colisione con futuras incorporaciones a la biblioteca estándar u otras bibliotecas externas. 
    
  Si mantienes su código en un repositorio de código fuente en alguna parte, entonces debería usar el root de tal repositorio de código fuente como su ruta base. Por ejemplo, si tienes una cuenta [GitHub](https://github.com/) en github.com/user, esta deberia ser su ruta base. 
    
  Note que no necesita publicar su código en un repositorio remoto antes de poder construirlo. Es solo un buen hábito organizar su código como si usted algún dia lo publicará. En la práctica  puede	 escoger cualquier nombre de ruta arbitrario, siempre que sea único en la biblioteca estándar y en la mayoría del ecosistema Go. 
    
  Usaremos github.com/user como nuestra ruta base. Crea un directorio dentro su espacio de trabajo para mantener el código fuente.

```   
$ mkdir -p $GOPATH/src/github.com/user
```
   
### 4. Su primer programa 
   
  Para compilar y ejecutar un simple programa, primero escoja la ruta del paquete (nosotros usaremos github.com/user/hello) y cree el correspondiente directorio de paquetes dentro de su espacio de trabajo:

```
$ mkdir $GOPATH/src/github.com/user/hello
```

  Luego, crea un archivo llamado hello.go dentro del directorio, que contenga el siguiente código Go.

```
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

  Ahora puedes construir y instalar el programa con la herramienta `go`:

```
go install github.com/user/hello
```

  Nota que puedes correr este comando desde cualquier parte en su sistema. La herramienta `go` encuentra el codigo fuente al buscar el paquete github.com/user/hello dentro del espacio de trabajo especificado por el GOPATH
   
  Puedes también omitir la ruta del paquete si ejecutas `go install` desde el directorio del paquete:

```
  $ cd $GOPATH/src/github.com/user/hello
  $ go install
```
 
  Este comando construye el comando hello, produciendo un binario ejecutable. Este luego instala tal binario en el directorio bin del espacio de trabajo como hello (o, bajo Windows, hello.exe). En nuestro ejemplo, que será $GOPATH/bin/hello, el cual es $HOME/go/bin/hello.
   
  La herramienta go solo imprimirá la salida cuando ocurra un error, por lo que si este comando no produce salida, significa que se han ejecutado con éxito.
   
  Puedes ahora ejecutar el programa escribiendo su ruta completa en la línea de comando:

```   
  $ $GOPATH/bin/hello
  Hello, world.
```   
  
  O como has añadido $GOPATH/bin a su PATH, solo escriba el nombre del binario: 

```
  $ hello
  Hello, world.
```

  Si estas usando un sistema de control de versiones, ahora es un excelente tiempo para inicializar un repositorio, añadir los archivos, y confirmar su primeros cambios. De nuevo, este paso es opcional: no necesitas usar un sistema de control de versiones para escribir codigo Go.  

```
  $ cd $GOPATH/src/github.com/user/hello
  $ git init
  Initialized empty Git repository in /home/user/go/src/github.com/user/hello/.git/
  $ git add hello.go
  $ git commit -m "initial commit"
  [master (root-commit) 0b4507d] initial commit
   1 file changed, 1 insertion(+)
    create mode 100644 hello.go
```

  Empujar el código a un repositorio remoto se deja como ejercicio para el lector
  
  
### 5. Su primera librería 

Vamos a escribir una librería y usarla desde el programa `hello` 

De nuevo, el primer paso es escoger la ruta del paquete (nosotros vamos a usar github.com/user/newmath) y crear el directorio del paquete.  

```
  $ mkdir $GOPATH/src/github.com/user/newmath
```

Luego, crear un archivo llamado sqrt.go en tal directorio con el siguiente contenido: 
 
```go
// Package newmath is a trivial example package.
package newmath

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
```

Ahora, prueba que el paquete compila con `go build`: 

```
$ go build github.com/user/newmath
```

O si estas trabajando el en directorio fuente del paquete, solo: 

```
$ go build
```

Esto no producirá un archivo de salida. Para hacer eso, debes usar `go install`, el cual coloca el objeto del paquete dentro del directorio pkg del espacio de trabajo. 
 
Luego de confirmar que el paquete newmath se construye, modifica su original `hello.go` (el cual esta en $GOPATH/src/github.com/user/hello) para usarlo:  

```go
package main

import (
	"fmt"
	"github.com/user/newmath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
}
```

Cada vez que la herramienta go instala un paquete o binario, también instala cualquier dependencias que tenga. Por lo que cuando instales el programa `hello`

```
$ go install github.com/user/hello
```

El paquete `newmath` será instalado también, automáticamente.

Al correr la nueva versión del programa, debes ver algún salida numérica: 
 
``` 
$ hello
Hello, world.  Sqrt(2) = 1.414213562373095
``` 

Luego de los pasos anteriores, su espacio de trabajo debe verse tal como: 
 
```  
bin/
    hello              # command executable
pkg/
    linux_amd64/       # this will reflect your OS and architecture
        github.com/user/
            newmath.a  # package object
src/
    github.com/user/
        hello/
            hello.go   # command source
        newmath/
            sqrt.go    # package source 
```

Nota que `go install` coloca el objeto `newmath.a` en un directorio dentro de `pkg/linux_amd64` que refleja su directorio de origen. Esto es para que en futuras invocaciones la herramienta go pueda encontrar el objeto del paquete y evitar recopilar el paquete innecesariamente. La parte linux_amd64 está ahí para ayudar en la compilación cruzada, y esta refejará el sistema operativo y arquitectura de su sistema. 
    
### 6. Nombre de los paquetes 

La primera declaración en un archivo fuente Go debe ser: 

```
package name
```

Donde `name` es el nombre por defecto para las importaciones. (Todos los archivos en un paquete deben usar el mismo `name`.)

Las convenciones de Go es que el nombre del paquete es el último elemento de la ruta de importación: el nombre del paquete importado como "crypto/rot13" debe ser nombrado rot13.

Los comandos ejecutables deben usar siempre el paquete `main`

No existe un requisito de que los nombres de paquetes sean únicos a través de todos los paquetes relacionados en un único binario, sólo que las rutas de importación (sus nombres de archivo completos) sean únicos.

Revisa  [Go Efectivo](http://golang.org/doc/effective_go.html#names) para aprender más sobre las convenciones de nombre.

## 3. Pruebas


Go tiene un marco de prueba ligero compuesto del comando `go test` y el paquete `testing`. 

Escribe una prueba creando un archivo con un nombre que termine en test.go que contenga funciones llamadas TestXXX con la firma `func (t *testing.T)`. El marco de prueba ejecuta cada una de tales funciones; si la función llama a una función de fallo tal como t.Error o t.Fail, la prueba se considera que ha fallado.
  
Añade una prueba al paquete newmath creando el archivo $GOPATH/src/github.com/user/newmath/sqrt_test.go conteniendo el siguiente código Go

```go
package newmath
import "testing"
func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
```

Luego ejecuta la prueba con `go test`:

```
$ go test github.com/user/newmath
ok  	github.com/user/newmath 0.165s
```

Como siempre, si estás ejecutando la herramienta `go` desde el directorio del paquete, puedes omitir la ruta del paquete: 

```
$ go test
ok  	github.com/user/newmath 0.165s
```

Ejecuta [go help test](http://golang.org/cmd/go/#hdr-Test_packages) y observa la [documentación de paquetes de pruebas](http://golang.org/pkg/testing/) para mas detalles. 

## 4. Paquetes remotos

Una ruta de importación puede describir como obtener el código fuente del paquete usando un sistema de control de versiones tal como Git o Mercural. La herramienta go usa esta propiedad para automáticamente extraer paquetes de repositorios remotos. Por ejemplo, los ejemplos descritos en este docuemento tambien se mantiene en un repositorio Mercurial amacenados en Google Code, [code.google.com/p/go.example](http://code.google.com/p/go.example). Si incluyes la URL del repositorio en la ruta de importación del paquete, go get ira a buscar, instalara, y construirá este automaticamente:
  
```
$ go get code.google.com/p/go.example/hello
$ $GOPATH/bin/hello
Hello, world.  Sqrt(2) = 1.414213562373095
```

Si el paquete especificado no está presente en un espacio de trabajo, `go get` lo colocará en el interior del primer espacio de trabajo especificado por GOPATH. (Si el paquete no existe todavía, `go get` salta buscar al remoto y se comporta igual que `go install`)

Después de emitir el anterior comando go get, el árbol de directorios del área de trabajo debe verse así:


```
bin/
    hello                 # command executable
pkg/
    linux_amd64/
        code.google.com/p/go.example/
            newmath.a     # package object
        github.com/user/
            newmath.a     # package object
src/
    code.google.com/p/go.example/
        hello/
            hello.go      # command source
        newmath/
            sqrt.go       # package source
            sqrt_test.go  # test source
    github.com/user/
        hello/
            hello.go      # command source
        newmath/
            sqrt.go       # package source
            sqrt_test.go  # test source
```

El comando `hello` almacenado en Google Code depende en el paquete `newmath` con el mismo repositorio. La importación en el archivo hello.go usa la misma convención de ruta de importación, por lo que el comando go get es capaz de localizar y instalar la dependencia del paquete, también.

```
import "code.google.com/p/go.example/newmath"
```

Esta convenciones es la forma facil de hacer que sus paquetes go estén disponibles para que otros los usen. La [Wiki Go](http://code.google.com/p/go-wiki/wiki/Projects) y [godoc.org](http://godoc.org/) proporciona una lista externa de proyectos Go.

Para más información en el uso de repositorios remotos con la herramienta go, revisa [go help importpath](http://golang.org/cmd/go/#hdr-Remote_import_paths) 

## 5. Que sigue

Suscríbete a la lista de correo [golang-announce](http://groups.google.com/group/golang-announce) para ser notificado cuando una nueva versión estable de Go está disponible.  

Revisa [Go Efectivo](http://golang.org/doc/effective_go.html) por tips en la escritura clara, e idiomática de código Go.

Toma [un tour por Go](http://tour.golang.org/) para aprender el lenguaje apropiadamente. 

Visita la [pagina de documentación](http://golang.org/doc/#articles) para un conjunto de artículos en profundidad sobre el lenguaje Go y sus librerías y herramientas. 

## 6. Obtener Ayuda

Para ayuda en tiempo real, pregunta a los útiles gophers en #go-nuts en el servidor IRC [Freenode](http://freenode.net/).

La lista oficial de discusión del lenguaje Go es [Go Nuts](http://groups.google.com/group/golang-nuts).

Reporta errores usando el [seguidor de incidencias Go](http://code.google.com/p/go/issues/list).