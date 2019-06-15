# Ponteiros

- Um ponteiro guarda na memória o endereço de uma variável;
- O tipo *T é um ponteiro para o valor T. Seu valor zero é `nil`;

```
var p *int
```

O operador `&` gera um ponteiro para seu operando.

```
i := 42
p = &i
```

o operador `*` indica um valor subjacentedo ponteiro.

```
fmt.Println(*p) // lê i através do ponteiro p
*p = 21         // defina i através do ponteiro p
```

PS: Go não faz aritmética de ponteiros.

# Structs

- A `struct` é uma coleção de campos;
- Os valores de uma `struct` são acessados por ponto;
- Campos de structs podem ser acessado através de um ponteiro;
- Para acessar o campo X de uma struct quando tivermos o ponteiro p da struct podemos escrever (*p).X ou simplemesmente p.X;
- A struct literal indica um valor recém-alocado, ao enumerar os valores de seus campos;
- O prefixo especial `&` constrói um ponteiro para uma struct literal

# Matrizes

- O tipo [n]T é uma matriz de n valor do tipo T;

```
var a [10]int
```

- O tamanho da matriz é parte do tipo, logo, matrizes não podem ser redimensiadas.

# Slices

- Uma matriz tem um tamanho fixo. Uma slice, é dinamicamente redimensionada. Na prática, slices são muito mais comuns que matrizes;
- O tipo []T é uma slice com elementos do tipo T.
- Uma slice é formada pela especificação de dois indices, um limite menor e maior, separados por dois pontos:

```
a[low: high]
```

- Este seleciona um intervalo meio-aberto que inclui o primeiro elemento, mas exclui o último:

```
a[1:4]
```

- Uma slice não armazena todos os dados, apenas descreve uma seção de uma matriz subjacente;
- Alterando os elementos de uma slice modificada os elementos correspondentes da sua matriz subjacente;
- Outras slices que partilham desse valor também serão alteradas;
- Uma slcie literal é como um array literal, sem o comprimento;
- Essas expressões são equivalentes:

```
a[0:10]
a[:10]
a[0:]
a[:]
```

- Uma slice tem tanto um tamanho quanto uma capacidade;
- O comprimento de uma slice é o número de elementos que ela contém;
- A capacidade de uma slice é o número de elementos na matriz subjacente, contando a partir do primeiro elemento na slice;
- O valor zero de uma slice é nil;
- As slices podem ser criadas com a função `make` padrão; isto é, como você criar matrizes dinâmicamente;
- O função `make` aloca uma matriz zerada e retorna uma slice que se refere a essa matriz:

```
a := make([]int, 5)     // len(a) = 5
b := make([]int, 0, 5)  // len(b) = 0, cap(b) = 5
```

- Para adicionar novos elementos a uma slice, Go dispõe da função ``append;

```
func append(s[]T, vs ...T) []T
```

# Range

- O `range` do laço `for` itera sobre ua slice ou map;
- Ao iterar sobre uma slice, dois valores são retornados para cada iteração. O primeiro é o indice, o segundo uma cópia do elemento daquele índice;
- Você pode ignorar o índice ou valor, atribuindo o `_`;

# Maps

- Um map mapeia chaves para valores;
- o valor zero de um map é nil;
- Um map nil não tem chaves, nem chaves podem ser adicionadas;
- A função make retorna um map com um tipo determinado;
- Maps literais são como structs literais, mas as chaves são obrigatórias;
- Se o tipo de nível superior é apenas um nome do tipo, você pode omiti-lo;

# Funções valores

- Função são valores também. Elas podem ser passdas assim como outros valores;
- Elas podem ser passadas como argumentos de funções e retornar valores.

# Closures

- As funções em Go podem ser closures. Um closure é uma função valor que referencia variáveis de fora de seu corpo. A função pode acessar e atribuir nas variáveis referenciadas; nesse sentido a função é "limitada" às variáveis.