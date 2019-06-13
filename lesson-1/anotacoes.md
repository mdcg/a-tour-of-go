# Pacotes

- Programas começam rodando pelo pacote 'main';
- Por convenção, o nome do pacote é o mesmo que o último elemento do caminho de importação. Por exemplo, o pacote "math/rand" compreende arquivos que começam com package rand.

# Importações

- Importação consignada => 

```
import (
    "fmt"
    "math"
)
```

- Você também pode importar dessa maneira:
```
import "fmt"
import "math"
```

PS: Prefira consignada.

# Nomes exportados

- Em Go, um nome é exportado se ele começa com letra maiúscula. Por exemplo `Pizza` é um nome exportável, assim como `Pi`, que é do pacote `math`;
- `pizza` e `pi` não começam com letra maiúscula, logo eles não são exportáveis.

# Funções

- Uma função pode ter zero ou mais argumentos;
- O tipo vem vem após o nome da variável;
- Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo, você pode omitir o tipo de todos menos o último:

```
x int, y int
x, y int
```

# Resultados múltiplos

- Uma função pode retornar qualquer número de resultados.

# Valores nomeados de retorno

- Valores de retorno de Go podem ser nomeados e agirem apenas como variáveis;
- A declaração `return` sem argumentos retorna os valores atuais dos resultados. Isso é conhecido como retorno "limpo";
- Instruções de retorno limpas devem ser usadas apenas em funções curtas.

# Variáveis

- a instrução `var` declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado;

```
var c, python, java bool
```

- A declaração pode estar numa pacote ou a nível de função;
- A declaração pode incluir inicializadores, uma por variável;
- Se um inicializador está presente, o tipo pode ser omitido.

```
var i, j int = 1, 2
var c, python, java = true, false, "no!"
```

# Declaração curta de variáveis

- Dentro de uma função a instrução de atribuição curta `:=` pode ser utilizada no lugar de `var` com o tipo implicito;
- Fora de funções, não é possível usar o `:=`.

# Tipos básicos em Go

- Os tipos básicos de Go são:

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // pseudônimo para uint8

rune // pseudônimo para int32
     // representa um ponto de código Unicode

float32 float64

complex64 complex128
```

# Valores zero

- Variáveis declaradas sem um valor inicial explicitado darão seu *valor zero*:

```
0 para tipos numéricos,
false para tipos boleanos, e
"" (string vazia) para strings.
```

# Conversões de tipo

Algumas conversões numéricas:

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

ou de forma simples:

```
i := 42
f := float64(i)
u := uint(f)
```

# Inferência de tipo

- Ao declarar uma variável, sem especificar o seu tipo (usando `var` sem um tipo ou `:=`), o tipo da variável é inferida a partir do valor do lado direito:

```
var i int
j := i            // j é um int
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

# Constantes

- Constantes são declaradas como variáveis, mas com a palavra-chave `const`;
- Constantes podem ser sequências de caracteres, booleanos ou valores numéricos;
- Elas não podem ser declaradas usando :=

# Constantes numéricas

- Constantes numéricas também são valores de alta precisão;
- Uma constante sem tipo pega o tipo necessário pelo seu contexto;
