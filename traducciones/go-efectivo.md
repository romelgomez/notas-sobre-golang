## Go Efectivo

1. Introducción
  1. Ejemplos
2. Formato
3. Comentarios
4. Nombres
  1. Nombres de paquetes
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

Go es un nuevo lenguaje. Aunque incorpora ideas de lenguajes existentes, tiene propiedades inusuales que hace los programas Go efectivos, diferentes en carácter de programas escritos en sus relativos. Una simple traducción de un programa en JAVA o C++ a GO es poco probable que produzca un resultado satisfactorio - los programas JAVA están escritos en JAVA, no en Go. Por otra parte, pensar en el problema desde una perspectiva Go podría producir un programa exitoso pero muy diferente. En otras palabras, para escribir Go bien, es importante entender sus propiedades y modismos. Es también importante conocer las convenciones establecidas para programar en Go, tal como el nombrar, el formato, la construcción del programa, y así sucesivamente, de modo que los programas que escribas sean fáciles de entender por otros programadores.

Este documento da consejos para escribir claro e idiomático código Go. Se aumenta la especificación del lenguaje, el tour por Go, y como escribir código go, todos los cuales debe leer primero.  

Ejemplos

La fuente de paquetes Go esta destinada a servir no sólo como librería del núcleo, sino también como ejemplos de cómo utilizar el lenguaje. Además, muchos paquetes contienen funcionales ejemplos ejecutables independientes que puedes ejecutar directamente desde el sitio web golang.org, tal como este, (si es necesario, haga clic en la palabra "Example" para abrirlo). Si tienes una pregunta sobre cómo abordar un problema o cómo se podría implementar algo, la documentación, código y ejemplos, la librería puede ofrecer respuestas, ideas y antecedentes.

## Formato

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

Algunos detalles de formato que permanecen. De forma muy resumida:

Sangría
  Nosotros usamos tabs para el sangrado y gofmt los emite de forma predeterminada. Utilice espacios sólo si es necesario.

Longitud de la línea
  Go no tiene límite de longitud de línea, No te preocupes por desbordar una tarjeta perforada. Si una línea la considera demasiado larga, envuélvala y aplique sangría con un tabulador extra.

Paréntesis
  Go necesita menos paréntesis que C y Java: estructuras de control (if, for, switch) no tienen paréntesis en su sintaxis. Además, la jerarquía de prioridad de los operadores es más corta y más clara, por lo que

```
x<<8 + y<<16
```

Significa que los espacios implican, a diferencia de los otros lenguajes.

## Comentarios

Go proporciona bloques de comentarios estilo C `/* */` y líneas de comentarios estilo C++ `//`. Las líneas de comentarios son la norma; los bloques de comentarios aparecen mayormente como comentarios de paquetes, pero son útiles dentro de una expresión o para desactivar grandes áreas de código.

El programa—y el servidor web—`godoc` procesa archivos fuentes para extraer la documentación sobre el contenido del paquete. Los comentarios que se presentan ante las declaraciones de nivel superior, sin nuevas líneas intermedias, se extraen junto con la declaración para servir como texto explicativo para el artículo. La naturaleza y el estilo de estos comentarios, determina la calidad de la documentación que `godoc` produce.

Cada paquete debe tener un comentario de paquete, un bloque de comentario precedente a la cláusula del paquete. Para los paquetes de varios archivos, el comentario del paquete sólo tiene que estar presente en un archivo, y cualquiera que lo haga. El comentario de paquete debe introducir el paquete y proporcionar información relevante de el paquete como un todo. Aparecerá de primero en la página godoc y debe establecer la documentación detallada que sigue.

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

Los comentarios no necesitan formato extra tal como las banderas de estrellas. La salida generada puede incluso no ser presentada en una fuente de ancho fijo, por lo que no dependen de la separación para la alineación —godoc, como `gofmt`, se encarga de eso. Los comentarios no son interpretados en texto plano, de modo que en HTML y otras anotaciones tales como _esto_ serán reproducidas palabra por palabra y no deben ser usadas. Un ajuste
que `godoc` hace es desplegar texto con sangría en una fuente de ancho fijo, adecuados para fragmentos de programas. El comentario de paquete para el `paquete fmt` usa este para buenos resultados.

Dependiendo del contexto, godoc puede que incluso no reforme los comentarios, así que para asegúrese de que se vean derechos: utilice una ortografía correcta, puntuación, y estructura de oración, líneas de plegado largas, y así sucesivamente.

Dentro de un paquete, cualquier comentario que precede inmediatamente una declaración de alto nivel sirve como un comentario documental para tal declaración. Cada (capitalizado) nombre exportado en un programa debe tener un comentario documental.

Los comentarios de documentación trabajan mejor como frases completas, cuales permiten una amplia variedad de presentaciones automatizadas. La primera frase debe ser un resumen de una frase que comienza con el nombre que se declara.

```
// Compile parses a regular expression and returns, if successful, a Regexp
// object that can be used to match against text.
func Compile(str string) (regexp *Regexp, err error) {
```

Si el comentario siempre comienza con el nombre, la salida de `godoc` puede ser provechosamente ejecutada través de `grep`. Imagine que no puedas recordar el nombre “Compile” pero mirando la función de análisis de expresiones regulares, así usted podrá ejecutar el comando.

```
$ godoc regexp | grep parse
```

Si todos los comentarios documentales en el paquete comienza con: “Esta función”  (“This function...”) grep no le ayudará a recordar el nombre. Pero debido a que el paquete comienza cada comentario documental con el nombre, usted vería algo como esto, que le recuerde la palabra que estás buscando.

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

## Nombres

Los nombres son tan importantes en el Go como en cualquier otro idioma. Incluso tienen un efecto semántico: la visibilidad de un nombre fuera de un paquete depende de si el primer carácter está en mayúscula. Por lo tanto, vale la pena pasar un poco de tiempo hablando de convenciones de nombres en los programas Go.

### Nombres de Paquetes

Cuando un paquete es importado, el nombre del paquete se convierte en un descriptor de acceso de los contenidos. Después

```
import "bytes"
```

El paquete importado puede hablar de bytes. Buffer. Es útil que todos los que usan el paquete puedan utilizar el mismo nombre para hacer referencia a su contenido, lo que implica que el nombre del paquete debe ser bueno: corto, conciso y sugerente.

Por convención, a los paquetes se les da nombre de una sola palabra, minúsculas; no debería haber ninguna necesidad de guiones o mixedCaps.
