Notas sobre GoLang
==================

1. Sinopsis
-----------

El lenguaje de programación Go es un proyecto open source para hacer programas más productivos. 

Go es expresivo, conciso, limpio y eficiente. Sus mecanismos de concurrencia hace fácil escribir programas que obtienen el máximo rendimiento de las máquinas en red y multinucleos. 


2. Sobre la instalación
-----------------------

Si consideras actualizar Go, debes remover la versión existente.

#### Linux, Mac OS X, and FreeBSD tarballs

Instalación en una ubicación personalizada

Las distribuciones binarias de Go asumen que serán instaladas en `/usr/local/go` (o `c:\Go` bajo Windows), pero es posible instalar la herramientas Go en una locación diferente. En este caso debes establecer la variable de entorno GOROOT para apuntar al directorio donde este fue instalado.

Por ejemplo, si instalaste Go en el directorio `$HOME`, debes añadir los siguientes comandos a `$HOME/.profile` o a `$HOME/.bashrc`:

```
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```

Nota: `GOROOT` debe ser establecido solo cuando se instala en una locación personalizada.

**Mi sistema de archivos:** Este esquema hace ver como tengo estructurado mi espacio de trabajo para desarrollar proyectos.

```
  .
  └── home
      └── romelgomez
          └── workspace/
              ├── configuration
              │   ├── languages
              │   │   ├── go
              │   │   ├── java
              │   │   │   ├── jdk
              │   │   │   └── library
              │   │   └── javascript
              │   │       └── node
              │   ├── programs
              │   │   └── idea
              │   └── sdks
              │       └── app-engine
              │           ├── go
              │           └── java
              ├── documentation
              │   ├── angular-js
              │   ├── go
              │   └── java
              └── projects
                  ├── otro-proyecto
                  └── golang
                      ├── bin
                      ├── pkg
                      └── src
                          └── github.com
                              └── jakecoffman
                                 └── go-angular-tutorial
                              └── romelgomez
                                 ├── notas-sobre-golang
                                 └── otro-proyecto-golang
```

`$HOME` es una variable que esta definida por el sistema y su valor puede ser consultado en una terminal:

```
romelgomez@romelgomez:~$ $HOME
bash: /home/romelgomez: Is a directory
```

Por lo que `$HOME` = `/home/romelgomez`

Las varias variables de entorno para la instalación en una ubicación personalizada buscan definir:

  1. Donde estará ubicada la distribución de Go y sus comandos compilados (/bin).
  2. Donde estará nuestro espacio de trabajo GoLang y sus comandos compilados (/bin).

Para mi sistema de archivos personalizado las variables de entorno quedarían de la siguiente forma:

1.
```
export GOROOT=$HOME/workspace/configuration/languages/go
export PATH=$PATH:$GOROOT/bin
```
2.
```
export GOPATH=$HOME/workspace/projects/golang
export PATH=$PATH:$GOPATH/bin
```

**Probando su instalación:** Compruebe que Go esté instalado correctamente creando un simple programa, de la siguiente forma:

  Cree el archivo `hola.go` y coloque el siguiente programa en el: 

```
package main
import "fmt"
func main() {
    fmt.Printf("Hola Mundo")
}
```
  Luego ejecute el programa de la siguiente forma: 

  ```
$ go run hola.go
  ```

  Si ve el mensaje "Hola Mundo", significa que su instalación Go funciona. 

  ```
hello, world
  ```

3. Ejemplos:
------------

* [Hola Mundo](/ejemplos/hola-mundo/hola-mundo.md) - Imprimiendo el clásico "hola mundo"

4. Tutoriales de Golang:
------------------------

- **En Español**
  - [Un tour por Go - Español](http://go-tour-es.appspot.com/) - Un tutorial simple e interactivo que introduce el contexto básico del lenguaje.
  - [Introducción a la programación con Go](http://golang-esp-man.blogspot.com/2014/05/introducion-la-programacion-con-go.html)

- **En Ingles**
  - [Go a través de ejemplos](https://gobyexample.com) - Un tutorial simple que introduce el contexto básico del lenguaje a través de ejemplos.
  - [Una introducción a la programación en Go](http://www.golang-book.com/) - Un libro extenso sobre GoLang
  
5. Traducciones
---------------
 
* [Cómo escribir código Go](/traducciones/como-escribir-codigo-go.md) - El original en ingles: http://golang.org/doc/code.html