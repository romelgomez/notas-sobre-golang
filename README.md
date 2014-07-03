Notas sobre GoLang
==================

1. Sinopsis
-----------

El lenguaje de programación Go es un proyecto open source para hacer programas más productivos. 

Go es expresivo, conciso, limpio y eficiente. Sus mecanismos de concurrencia hace fácil escribir programas que obtienen el máximo rendimiento de las máquinas en red y multinucleos. 


2. Instalación personalizada
----------------------------

   1. **Sistema Operativo:** Ubuntu 14.04 LTS
   
   2. **IDE:** IntelliJ IDEA

   2. **Locación Personalizada:** La distribución binaria de Go asume que ella será instalada en /usr/local/go , sin embargo para este esquema las herramientas de Go serán instaladas en una locación diferente.
   
   3. **Sistema de control de versiónes:** github.com 

   4. **Esquema del sistema de archivos:** Un esquema personal creado con el fin de contar con un espacio de trabajo integro, que contenga en la medida de lo posible, cada elemento relacionado con el desarrollo de proyectos.

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


   5. #### Ejemplo de las variables de entorno en: $HOME/.bashrc ####

        ```
        # JAVA JDK SET
        export JAVA_HOME=$HOME/workspace/configuration/languages/java/jdk/jdk1.8.0_05
        export JDK_HOME=$HOME/workspace/configuration/languages/java/jdk/jdk1.8.0_05
        
        # IDEA
        export PATH=$PATH:$HOME/workspace/configuration/programs/idea/idea-IU-135.909/bin
        
        # GoLang SET
        export GOROOT=$HOME/workspace/configuration/languages/go
        export PATH=$PATH:$GOROOT/bin
        
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
