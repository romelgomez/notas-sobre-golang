## Go Efectivo

1. [Introducción](#1-introducción)
  1. [Ejemplos](#1-ejemplos)
2. [Formato](#2-formato)
3. [Comentarios](3-comentarios)
4. [Nombres](4-nombres)
  1. [Nombres de paquetes](1-nombres-de-paquetes)
  2. Métodos captadores (Getters)
  3. Nombres de interfaces
  4. Tapas mixtas (MixedCaps)
5. Punto y coma
6. Estructuras de control
  1. if
  2. Redeclaración y reasignación	
  3. For
  4. Switch
  5. Type switch
7. Funciones
  1. Retorna múltiples valores
  2. Named result parameters (Parámetros de resultados Nombrados)	
  3. Diferida
8. Datos
  1. Asignación con new
  2. Constructores y literales compuestas
  3. Asignación con make
  4. Matrices
  5. Rebanadas (Slices)
  6. Rebanadas de dos dimensiones
  7. Mapas
  8. Imprimir
  9. Anexar
9. Inicialización
  1. Constantes
  2. Variables
  3. La función init
10. Métodos
  1. Punteros vs Valores
11. Interfaces y otros tipos
  1. Interfaces
  2. Conversiones
  3. Conversiones de interfaz y las afirmaciones de tipo
  4. Generalidad
  5. Interfaces y Métodos
12. Identificador en blanco
  1. Identificador en blanco en asignaciones múltiples
  2. Variables y importaciones no utilizadas
  3. Importación de efectos secundarios
  4. Comprobaciones de interfaz
13. Incorporación
14. Concurrencia
  1. Compartir comunicando
  2. Rutinas Go (GoRoutines)
  3. Canales
  4. Canales de Canales
  5. Paralelización
  6. Un tampón con fugas
15. Errores
  1. pánico
  2. recuperar
16. Un servidor web


## 1. Introducción

Go es un nuevo lenguaje. Aunque incorpora ideas de lenguajes existentes, tiene propiedades inusuales que hace efectivos los programas Go, diferentes en carácter de programas escritos en sus relativos. Una simple traducción de un programa en JAVA o C++ a GO es poco probable que produzca un resultado satisfactorio - los programas JAVA están escritos en JAVA, no en Go. Por otra parte, pensar en el problema desde una perspectiva Go podría producir un exitoso pero muy diferente programa. En otras palabras, para escribir Go bien, es importante entender sus propiedades y modismos. Es también importante conocer las convenciones establecidas para programar en Go, tal como el nombrado, el formato, la construcción del programa, y así sucesivamente, de modo que los programas que escribas sean fáciles de entender por otros programadores.

Este documento da consejos para escribir código Go claro e idiomático. Se aumenta [La especificación del lenguaje](http://golang.org/ref/spec), el [tour por Go](http://tour.golang.org/), y [Como Escribir Código Go](http://golang.org/doc/code.html), todos los cuales debe leer primero.

### 1. Ejemplos

La [fuente de paquetes Go](http://golang.org/src/pkg/) esta destinada a servir no sólo como librería del núcleo, sino también como ejemplos de cómo utilizar el lenguaje. Además, muchos paquetes contienen funcionales ejemplos ejecutables independientes que puedes ejecutar directamente desde el sitio web [golang.org](http://golang.org/), tal como [este](http://golang.org/pkg/strings/#example_Map), (si es necesario, haga clic en la palabra "Example" para abrirlo). Si tienes una pregunta sobre cómo abordar un problema o cómo se podría implementar algo, la documentación, código y ejemplos, la librería puede ofrecer respuestas, ideas y antecedentes.

## 2. Formato

Los problemas de formato son las más polémicos, pero los menos consecuentes. Las personas se pueden adaptar a diferentes estilos de formato pero es mejor si ellos no tienen que hacerlo, y menos tiempo se le dedica al tópico si todos se adhiere al mismo estilo. El problema es cómo acercarse a esta utopía sin una guía de estilo prescriptiva larga.  

Con Go tomamos un enfoque inusual y dejamos que máquina cuide de la mayoría de los problemas de formato. El programa `gofmt` (también disponible como `go fmt`, el cual opera a nivel del paquete en lugar de a nivel del archivo fuente) lee un programa Go y emite el fuente en un estilo estándar de sangría y alineación vertical, retiene y si es necesario reforma los comentarios. Si usted quiere saber cómo manejar alguna nueva situación de diseño, ejecute `run gofmt`; si la respuesta no me parece correcta, reorganize su programa (o abra una incidencia sobre gofmt), no trabaje alrededor de él.   

A modo de ejemplo, no hay necesidad de gastar tiempo haciendo coincidir los comentarios sobre los campos de una estructura, `gofmt` lo hará por usted. Teniendo en cuenta la declaración

```
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt`  alinearán las columnas:

```
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

Todo el código Go en los paquetes estándar se ha formateado con `gofmt`.

Algunos detalles de formato que faltan. De forma muy resumida:

- Sangría
  Nosotros usamos tabs para el sangrado y gofmt los emite de forma predeterminada. Utilice espacios sólo si es necesario.

- Longitud de la línea
  Go no tiene límite de longitud de línea, No te preocupes por desbordar una tarjeta perforada. Si una línea la considera demasiado larga, envuélvala y aplique sangría con un tabulador extra.

- Paréntesis
  Go necesita menos paréntesis que C y Java: estructuras de control (if, for, switch) no tienen paréntesis en su sintaxis. Además, la jerarquía de prioridad de los operadores es más corta y más clara, por lo que

```
x<<8 + y<<16
```

  Significa que los espacios implican, a diferencia de los otros lenguajes.

## 3. Comentarios

Go proporciona bloques de comentarios estilo C `/* */` y líneas de comentarios estilo C++ `//`. Las líneas de comentarios son la norma; los bloques de comentarios aparecen mayormente como comentarios de paquetes, pero son útiles dentro de una expresión o para desactivar grandes áreas de código.

El programa—y el servidor web—`godoc` procesa archivos fuentes para extraer la documentación sobre el contenido del paquete. Los comentarios que se presentan ante las declaraciones de nivel superior, sin nuevas líneas intermedias, se extraen junto con la declaración para servir como texto explicativo para el elemento. La naturaleza y el estilo de estos comentarios, determina la calidad de la documentación que `godoc` produce.

Cada paquete debe tener un comentario de paquete, un bloque de comentario precedente a la cláusula del paquete. Para los paquetes de varios archivos, el comentario del paquete sólo tiene que estar presente en un archivo, y cualquiera que lo haga. El comentario de paquete debe introducir el paquete y proporcionar información relevante del paquete como un todo. Este aparecerá de primero en la página de godoc y debe establecer la documentación detallada que sigue:

```
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

Si el paquete es simple, el comentario de paquete puede ser breve.

```
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

Los comentarios no necesitan formato extra tal como las pancartas de estrellas. La salida generada puede incluso no ser presentada en una fuente de ancho fijo, por lo que no dependen de la separación para la alineación —godoc, como `gofmt`, se encarga de eso. Los comentarios no son interpretados en texto plano, de modo que en HTML y otras anotaciones tales como _esto_ serán reproducidas palabra por palabra y no deben ser usadas. Un ajuste
que `godoc` hace es desplegar texto con sangría en una fuente de ancho fijo, adecuadas para fragmentos de programas. El comentario de paquete para el [fmt package](http://golang.org/pkg/fmt/) usa esto para buenos resultados.

Dependiendo del contexto, `godoc` puede que incluso no reforme los comentarios, así que para asegúrese de que se vean derechos: utilice una ortografía correcta, puntuación, y una estructura de oración, líneas de plegado largas, y así sucesivamente.

Dentro de un paquete, cualquier comentario que precede inmediatamente una declaración de alto nivel sirve como un comentario documental para tal declaración. Cada nombre (capitalizado) exportado en un programa debe tener un comentario documental.

Los comentarios de documentación trabajan mejor como frases completas, cuales permiten una amplia variedad de presentaciones automatizadas. La primera frase debe ser un resumen de una frase que comienza con el nombre que se declara.

```
// Compile parses a regular expression and returns, if successful, a Regexp
// object that can be used to match against text.
func Compile(str string) (regexp *Regexp, err error) {
```

Si el comentario siempre comienza con el nombre, la salida de `godoc` puede ser provechosamente ejecutada través de `grep`. Imagine que no puedas recordar el nombre “Compile” pero cuando miramos la función de análisis para expresiones regulares, así usted podrá ejecutar el comando.

```
$ godoc regexp | grep parse
```

Si todos los comentarios documentales en el paquete comienza con: “Esta función”  (“This function...”) grep no le ayudará a recordar el nombre. Pero debido a que el paquete comienza cada comentario documental con el nombre, usted vería algo como esto, que le recordara la palabra que estás buscando.

```
$ godoc regexp | grep parse
    Compile parses a regular expression and returns, if successful, a Regexp
    parsed. It simplifies safe initialization of global variables holding
    cannot be parsed. It simplifies safe initialization of global variables
$
```

Sintaxis de declaración de Go permite la agrupación de las declaraciones. Un solo comentario documental puede introducir un grupo de constantes o variables relacionadas. Dado que toda la declaración es presentada, tal comentario a menudo puede ser rutinaria.

```
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

La agrupación también puede indicar relaciones entre los elementos, tales como el hecho de que un conjunto de variables está protegido por una exclusión mutua.

```
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```

## 4. Nombres

Los nombres son tan importantes en el Go como en cualquier otro idioma. Incluso tienen un efecto semántico: la visibilidad de un nombre fuera de un paquete depende de si el primer carácter está en mayúscula. Por lo tanto, vale la pena pasar un poco de tiempo hablando de convenciones de nombres en los programas Go.

### 1. Nombres de Paquetes

Cuando un paquete es importado, el nombre del paquete se convierte en un descriptor de acceso para el contenido. Después

```
import "bytes"
```

El paquete importado puede hablar de bytes. Buffer. Es útil que todos los que usan el paquete puedan utilizar el mismo nombre para hacer referencia a su contenido, lo que implica que el nombre del paquete debe ser bueno: corto, conciso y sugerente.

Por convención, a los paquetes se les dan nombres de una sola palabra, minúsculas; no debería haber ninguna necesidad de guiones o mixedCaps. Errar en el lado de la brevedad, ya que cualquiera que use su paquete escribirá ese nombre. Y no te preocupes por colisiones. El nombre del paquete es más que el nombre predeterminado para las importaciones; no tiene que ser único a lo largo del todo el código fuente, y en el raro caso de una colisión el paquete importador puede elegir un nombre diferente para utilizar a nivel local. En cualquier caso, la confusión es raro porque el nombre del archivo en la importación determina simplemente qué paquete se está utilizando.

Otra convención es que el nombre del paquete es el nombre base de su directorio de origen; el paquete en `src/pkg/encoding/base64` es importado como `encoding/base64` pero tiene nombre `base64`, no encoding_base64 y no encodingBase64. 

El importador de un paquete utilizará el nombre para hacer referencia a su contenido, así que los nombres exportados en el paquete pueden utilizar este hecho para evitar tartamudear. (No utilice la notación import, la cual puede simplificar las pruebas que deban ejecutarse fuera del paquete que esta probando, pero por otro lado debe ser evitado). Por ejemplo, el tipo `buffered reader` en el paquete `bufio` es llamado `Reader`, no `BufReader`, porque los usuarios ven como `bufio.Reade`, el cual es un nombre conciso y claro. Por otra parte, porque las entidades importadas siempre se abordan con su nombre de paquete, `bufio.Reader` no entra en conflicto con `io.Reader`. Similarmente, la función para hacer nuevas instancias de `ring.Ring` —que es la definición de un constructor en Go— que normalmente se llama `NewRing`, pero desde que `Ring` es el único tipo exportado por el paquete, y puesto que el paquete se llama `ring`, se llama simplemente `New`, el cual los clientes del paquete ven como `ring.New`. Usa la estructura del paquete para ayudar a escoger buenos nombres. 

Otro ejemplo corto que se lee bien es `once.Do`; `once.Do(setup)` y no se mejorarse al escribir `once.DoOrWaitUntilDone(setup)`. Nombres largos no hacen automáticamente las cosas mas legibles. Una util documentación comentada puede a menudo ser más valiosa que un nombre extra largo.   

### 2. Getters 

Go no proporciona soporte automático para getters y setters. No hay nada malo con proporcionar getters y setters por si mismo, y es a menudo apropiado hacerlo, pero no es idiomático o necesario colocar `Get` en el nombre del getter. Si tienes un campo llamado `owner` (en minúsculas `lower case`, no exportado), el método getter debe ser llamado Owner (mayúsculas `upper case`, exportados), no `GetOwner`. El uso de nombres upper-case para exportar proporciona el gancho para discriminar el campo del método. Una función setter, si es necesaria, es probable que se llame `SetOwner`. Ambos nombres se leen bien en la practica:
          
```          
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

###  3 Nombres de interfaces  

Por convención, interfaces de un método son nombrados por el nombre del método mas un `-er` sufijo o similar modificación para construir un sustantivo agente `Reader`, `Writer`, `Formatter`, `CloseNotifier` etc. 

Hay un numero de tales nombres... Read, Write, Close, Flush, String, etc que tienen firmas canónicas y significados. Para evitar la confusión, no le des a sus métodos uno de esos nombres amenos que este tenga la misma firma y significado. Por el contrario, si su tipo implementa un método con el mismo significado que un método en un tipo bien conocido, dale el mismo nombre y firma; llame a su método de convertidor de cadena String no ToString.
                                                                                                                                                                                                                                                                                                                                                                                                                                       
### 4. MixedCaps

Finalmente, la convección en Go es usar `MixedCaps` o `mixedCaps` en lugar de subrayados para escribir los nombres de varias palabras.

## 5. Punto y coma

Al igual que C. Gramática formal de Go usa punto y coma para terminar declaraciones, ero a diferencia de C, esos puntos y comas no aparecen en el fuente. En cambio, el analizador léxico utiliza una regla simple para insertar un punto y coma de forma automática cuando lo revisa por lo que el texto de entrada está mayormente libre de ellos. 

La regla es la siguiente. Si el último token antes de un salto de línea es un identificador (cual incluye palabras como int y float64), un literal tan básico como un número o una cadena constante, o uno de los tokens

```
break continue fallthrough return ++ -- ) }
```

el analizador léxico siempre inserta un punto y coma después de del token. Esto podría resumirse como "Si la nueva línea viene después de un token que podría poner fin a una declaración, inserte un punto y coma".

Un punto y coma puede omitirse inmediatamente antes de una llave de cierre, por lo que una declaración como

```
go func() { for { dst <- <-src } }()
```

no necesita ningún punto y coma. Programas Go idiomáticos tienen punto y coma sólo en lugares como en cláusulas de bucle `for` para separar los elementos de inicialización, condición, y de continuación. También son necesarios para separar varias declaraciones en una línea, en caso de deba escribir el código de esa manera.  

Una de las consecuencias de las normas de introducción de punto y coma es que no se puede colocar la llave de apertura de una estructura de control (if, for, switch, o select) en la siguiente línea. Si lo hace, se le insertará un punto y coma antes de la llave, lo que podría causar efectos no deseados. Escríbelos de esta manera:

```
if i < f() {
    g()
}
```

no de esta forma:

```
if i < f()  // wrong!
{           // wrong!
    g()
}
```

## 6. Estructuras de control

La estructura de control de Go son afines a las de C, pero difieren en aspectos importantes. No existe `do` o el bucle `while`, solo un ligeramente generalizado `for`; `switch` es mas flexible; `if` y `switch` aceptar una sentencia de inicialización opcional como el de `for`; `break` y `continue` declaraciones toman una etiqueta opcional para identificar lo que se rompe o continua; y existen nuevas estructuras de control que incluyen un tipo `switch` y un multiplexor de comunicaciones de múltiples vías, `select`. La sintaxis también es ligeramente diferente: no hay paréntesis y los cuerpos deben ser siempre delimitados por llaves.    

### 1. if

En Go un simple `if` luce de la siguiente forma: 
  
```  
if x > 0 {
    return y
}  
``` 
  
Las llaves obligatorias fomentan la escritura sencilla de declaraciones if en varias líneas. Es un buen estilo para hacerlo de todos modos, especialmente cuando el cuerpo contiene una instrucción de control tal como un `return` o `break`. 

Ya que if y switch acepta una declaración de inicialización, es común ver que se utiliza para establecer una variable local.

```
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
} 
```

Este es un ejemplo de una situación común donde el código debe protegerse contra una secuencia de condiciones de error. El código se lee bien si el flujo exitoso del control corre por la página, eliminar los casos de error que puedan surgir. Como los casos de error suelen acabar en sentencias return, el código resultante no necesita else.

```
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### 2. Redeclaración y reasignación

Aparte: El último ejemplo en la sección anterior demuestra un detalle de cómo la forma de declaración corta `:=` funciona.

```
f, err := os.Open(name)
```



### 3. For
### 4. Switch
### 5. Type switch
## 7. Funciones
### 1. Retorna múltiples valores
### 2. Named result parameters (Parámetros de resultados Nombrados)	
### 3. Diferida
## 8. Datos
### 1. Asignación con new
### 2. Constructores y literales compuestas
### 3. Asignación con make
### 4. Matrices
### 5. Rebanadas (Slices)
### 6. Rebanadas de dos dimensiones
### 7. Mapas
### 8. Imprimir
### 9. Anexar
## 9. Inicialización
### 1. Constantes
### 2. Variables
### 3. La función init
## 10. Métodos
### 1. Punteros vs Valores
## 11. Interfaces y otros tipos
### 1. Interfaces
### 2. Conversiones
### 3. Conversiones de interfaz y las afirmaciones de tipo
### 4. Generalidad
### 5. Interfaces y Métodos
## 12. Identificador en blanco
### 1. Identificador en blanco en asignaciones múltiples
### 2. Variables y importaciones no utilizadas
### 3. Importación de efectos secundarios
### 4. Comprobaciones de interfaz
## 13. Incorporación
## 14. Concurrencia
### 1. Compartir comunicando
### 2. Rutinas Go (GoRoutines)
### 3. Canales
### 4. Canales de Canales
### 5. Paralelización
### 6. Un tampón con fugas
## 15. Errores
### 1. pánico
### 2. recuperar
## 16. Un servidor web