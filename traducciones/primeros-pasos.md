# Primeros pasos

1. [Descargar la distribución Go](#1-descargar-la-distribución-go)
2. [Requisitos del sistema](#2-requisitos-del-sistema)
3. [Instalar las herramientas Go](#3-instalar-las-herramientas-go)
  1. [Linux, Mac OS X, y FreeBSD tarballs](#1-linux-mac-os-x-y-freebsd-tarballs)
  2. [Mac OS X paquete de instalador](#2-mac-os-x-paquete-de-instalador)
  3. [Windows](#3-windows)
4. [Compruebe la instalación](#4-compruebe-la-instalación)
5. [Configurar el entorno de trabajo](#5-configurar-el-entorno-de-trabajo)
6. [Desinstalar Go](#6-desinstalar-go)
7. [Obtener ayuda](#7-obtener-ayuda)


# 1. Descargar la distribución Go

Visite la página de descargas: https://golang.org/dl/

[Las distribuciones binarias](https://golang.org/dl/) oficiales están disponibles para los sistemas operativos FreeBSD (versión 8 y recientes) Linux, Mac OS X  (Snow Leopard y recientes), y Windows; Y para las arquitecturas de procesador 32-bit (386) y 64-bit (amd64) x86. 

Si la distribución binaria no esta disponible para su combinación de sistema operativo y arquitectura, intente [instalar desde el código fuente](http://golang.org/doc/install/source) o [instalar gccgo en lugar de gc](http://golang.org/doc/install/gccgo). 

# 2. Requisitos del sistema

El compilador `gc` soporta los siguientes sistemas operativos y arquitecturas. Por favor asegúrese que su sistema tenga estos requerimientos antes de proceder. Si su sistema operativo o arquitectura no esta en la lista, es posible que `gccgo` pueda que soporte su configuración; Para más detalles vea [Configurar y usar gccgo](http://golang.org/doc/install/gccgo)

| Sistema Operativo                       | Arquitecturas     | Notas                                                                                   |
| ----------------------------------------| ----------------- | --------------------------------------------------------------------------------------- |
| FreeBSD 8 o superior                    | amd64, 386, arm   |  Debian GNU/kFreeBSD no esta soportado; FreeBSD/ARM necesita FreeBSD 10 o superior      |
| Linux 2.6.23 o superior con glibc       | amd64, 386, arm   |   CentOS/RHEL 5.x no esta soportado; no hay una distribución binaria para ARM aún       |
| Mac OS X 10.6 o superior                | amd64, 386        |    use el gcc† que viene con Xcode‡                                                     |
| Windows XP o superior                   | amd64, 386        | usa MinGW gcc†. No hay necesidad de cygwin o msys                                       |

†gcc es requerido solo si planeas usar [cgo](http://golang.org/cmd/cgo/).

‡ Solo necesitas instalar las herramientas para la terminal para [Xcode](https://developer.apple.com/xcode/). Si ya tienes instalado Xcode 4.3+, puedes instalarlas desde la pestaña Componentes del panel de preferencias de Descargas.

# 3. Instalar las herramientas Go

Si estás actualizando desde una versión anterior de Go debes primero [remover la versión existente](#6-desinstalar-go). 

## 1. Linux, Mac OS X, y FreeBSD tarballs

[Descarga](https://golang.org/dl/) y descomprime el archivo en `/usr/local`, crea un árbol Go en `/usr/local/go`. Por ejemplo:

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

Selecciona el archivo apropiado para su instalación. Por ejemplo, si estás instalando la version Go 1.2.1 para 64-bit x86 en Linux, el archivo que desea se llama `go1.2.1.linux-amd64.tar.gz`.

(Típicamente estos comandos deben ser ejecutados como root o a través de `sudo`)

Añade `/usr/local/go/bin` a la variable de entorno PATH. Puedes hacer esto añadiendo esta línea a su `/etc/profile` (para una instalación en todo el sistema) o `$HOME/.profile`:  

`export PATH=$PATH:/usr/local/go/bin`

### Instalación en una ubicación personalizada

La distribución binaria de Go asume que será instalada en `/usr/local/go` (o `c:\Go` bajo Windows), pero es posible instalar las herramientas Go en una locación diferente. En esta caso debes establecer la variable de entorno `GOROOT` para apuntar al directorio en el cual serán instaladas. 

Por ejemplo, si instalas Go en su directorio principal debe añadir los siguientes comandos a `$HOME/.profile`:

```
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```

Nota: `GOROOT` debe ser establecido sólo cuando se instala en una locación personalizada.

## 2. Mac OS X paquete de instalador

[Descarga el paquete](https://golang.org/dl/), ábrelo, y sigue las indicaciones para instalar las herramientas Go. El paquete instala la distribución Go en `/usr/local/go`. 

El paquete debe colocar el directorio `/usr/local/go/bin` en su variable de entorno `PATH`. Puedes que requieras reiniciar cualquier sesión de Terminal abierta para que los cambios surtan efectos.

## 3. Windows

El proyecto Go proporciona dos opciones de instalación para los usuarios de Windows (además de la [instalación desde el código fuente](http://golang.org/doc/install/source)): un archivo zip que requiere que establezca algunas variables de entorno y un instalador MSI que configura su instalación automáticamente.    

### Instalador MSI

Abre el [archivo MSI](https://golang.org/dl/) y sigue las instrucciones para instalar las herramientas Go. Por defecto, el instalador coloca la distribución Go en `c:\Go`.

El instalador debería colocar el directorio `c:\Go\bin` en la variable de entorno `PATH`. Puedas que necesites reiniciar cualquier Terminal para que los cambios surtan efectos.  

### Archivo Zip

[Descarga y descomprime el archivo zip](https://golang.org/dl/) en el directorio de su elección (Nosotros sugerimos `c:\Go`). 

Si escoges un directorio distinto de `c:\Go`, debe establecer la variable de entorno `GOROOT` a su ruta seleccionada.

Añada el subdirectorio bin de su raíz Go (por ejemplo, `c:\Go\bin`) a su variable de entorno `PATH`. 

### Configurando las variables de entorno bajo Windows

Bajo Windows, puede configurar las variables de entorno a través del botón “Variables de Entorno” en la pestaña “Avanzado” del panel del control del “Sistema”. Algunas versiones de Windows proporcionan este panel de control a través de la opción “Configuración avanzada del sistema” dentro del panel de control del “Sistema”

# 4. Compruebe la instalación

Compruebe que Go esta instalado correctamente construyendo un simple programa, de la siguiente forma: 

Crea un archivo llamado `hello.go` y coloque el siguiente programa en el:

```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

Luego ejecútelo con la herramienta Go:

```
$ go run hello.go
hello, world
```

Si ve el mensaje "hello, world" significa que si instalación Go esta funcionado.  

# 5. Configurar el entorno de trabajo

Ya está casi listo. Usted sólo necesita establecer su entorno. 

Lea el documento como [escribir código Go](https://github.com/romelgomez/notas-sobre-golang/blob/master/traducciones/como-escribir-codigo-go.md), el cual proporciona instrucciones de configuración esenciales para usar las herramientas Go.  

# 6. Desinstalar Go

Para remover una existente instalación de Go de su sistema borre el directorio Go. Este esta usualmente en `/usr/local/go` bajo  Linux, Mac OS X, y FreeBSD o `c:\Go` bajo Windows.

También debe eliminar el directorio bin de Go de su variable de entorno PATH. Bajo linux y FreeBSD debes editar `/etc/profile` o `$HOME/.profile`. Si instalaste Go con el [paquete Mac OS X](#2-mac-os-x-paquete-de-instalador) entonces debes remover el archivo `/etc/paths.d/go`. Los usuarios de Windows deben leer la sección sobre [establecer variables de entorno bajo Windows](#configurando-las-variables-de-entorno-bajo-windows). 

# 7. Obtener ayuda

Para ayuda en tiempo real, pregunta a los útiles gophers en #go-nuts en el servidor IRC de [Freenode](http://freenode.net/)

La oficial lista de correo para discusiones de el lenguaje Go es [Go Nuts](http://groups.google.com/group/golang-nuts).

Reporta bugs usando el [seguidor de problemas Go](http://golang.org/issue)