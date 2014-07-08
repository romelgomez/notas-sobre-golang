Cómo escribir código Go
------------------------------------
 
**Original:** [How to Write Go Code] (http://golang.org/doc/code.html)


1. Introducción
2. Organización del código
  1. Espacios de trabajo
  2. La variable de entorno GOPATH
  3. Rutas de paquetes
  4. Su primer programa
  5. Su primer libreria
  6. Nombre de los paquetes
3. Pruebas
4. Paquetes remotos
5. Que sigue
6. Obtener Ayuda


1. Introducción
---------------

Este documento demuestra el desarrollo de una simple paquete Go y presenta la herramienta Go (go tool), la forma estándar de extraer, construir, y instalar paquetes Go y comandos. 

La herramienta Go requiere que organice su código en una manera específica, Por favor lee este documento atentamente. Esta explica de forma simple como empezar a trabajar con su instalación Go

Una similar explicación esta disponible como un [screencast, en ingles] (http://www.youtube.com/watch?v=XCsL89YtqCs).

2. Organización del código
--------------------------

1. **Espacios de trabajo:**
 
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
    
  Los comandos y librerías se construyen a partir de diferentes tipos de paquetes de código. Discutiremos esta distinción [luego](#PackageNames).

2. **La variable de entorno GOPATH**

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
  
3. **Rutas de paquete**
  
  Los paquetes de la librería estándar se mencionan como cortas rutas tales como "fmt" y "net/http". Para sus propios paquetes, debe escoger una ruta base que no colisione con futuras incorporaciones a la biblioteca estándar u otras bibliotecas externas. 
    
  Si mantienes su código en un repositorio de código fuente en alguna parte, entonces debería usar el root de tal repositorio de código fuente como su ruta base. Por ejemplo, si tienes una cuenta GitHub en github.com/user, esta deberia ser su ruta base. 
    
  Note que no necesita publicar su código en un repositorio remoto antes de poder construirlo. Es solo un buen hábito organizar su código como si usted algun dia lo publicará. En la práctica  puede	 escoger cualquier nombre de ruta arbitrario, siempre que sea único en la biblioteca estándar y en la mayoría del ecosistema Go. 
    
  Usaremos github.com/user como nuestra ruta base. Crea un directorio dentro su espacio de trabajo para mantener el código fuente.
```   
$ mkdir -p $GOPATH/src/github.com/user
```
   
4. **Su primer programa** 
   
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