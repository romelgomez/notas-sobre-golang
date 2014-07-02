Notas sobre GoLang
==================

Sinopsis
--------

El lenguaje de programación Go es un proyecto open source para hacer programas más productivos. 

Go es expresivo, conciso, limpio y eficiente. Sus mecanismos de concurrencia hace fácil escribir programas que obtienen el máximo rendimiento de las máquinas en red y multinucleos. 


Instalación personalizada
-------------------------

### Descrita para el siguiente entorno: ###

- **Sistema Operativo:** Ubuntu 14.04 LTS
- **IDE:** IntelliJ IDEA
- **Locación:** Personalizada, la distribución binaria de Go asume que ella será instalada en /usr/local/go , sin embargo para esquema las herramientas GO serán instaladas en una locación diferente.
- **Sistema de control de versiónes:** github.com 

```
workspace/
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
│       └── appengine
│           ├── go
│           └── java
├── documentation
│   ├── angular-js
│   ├── go
│   └── java
└── projects
    ├── current
    │   ├── angular-js
    │   │   └── github.com
    │   │       └── romelgomez
    |   |           ├── notas-sobre-angular-js
    │   │           └── otro-proyecto
    │   └── golang
    │       ├── bin
    │       ├── pkg
    │       └── src
    │           └── github.com
    │               └── romelgomez
    |                  ├── notas-sobre-angular-js  
    │                  └── otro-proyecto
    └── legacy
```


### Ejemplo de las variables de entorno en: $HOME/.bashrc ###

#### JAVA JDK SET ####
```
export JAVA_HOME=$HOME/workspace/configuration/languages/java/jdk/jdk1.8.0_05
export JDK_HOME=$HOME/workspace/configuration/languages/java/jdk/jdk1.8.0_05
```

#### IDEA ####
```
export PATH=$PATH:$HOME/workspace/configuration/programs/idea/idea-IU-135.909/bin
```

#### GoLang SET ####

```
export GOROOT=$HOME/workspace/configuration/languages/go
export PATH=$PATH:$GOROOT/bin
```

```
export GOPATH=$HOME/workspace/projects/current/golang
export PATH=$PATH:$GOPATH/bin
```



Ejemplos:
---------

* [Hola Mundo] (/ejemplos/hola-mundo/hola-mundo.md) - Imprimiendo el clásico "hola mundo"


Tutoriales de Golang:
---------------------

- **En Español**
    - [Un tour por Go - Español] (http://go-tour-es.appspot.com/) - Un tutorial simple e interactivo que introduce el contexto básico del lenguaje.

- **En Ingles**
    - [Go a través de ejemplos] (https://gobyexample.com) - Un tutorial simple que introduce el contexto básico del lenguaje a travez de ejemplos.
    - [Una introducción a la programación en Go] (http://www.golang-book.com/) - Un libro extenso sobre GoLang
