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