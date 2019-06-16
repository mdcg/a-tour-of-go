# Métodos

- Go não tem classes. No entanto, você pode definir método de tipos;
- O método é uma função com um argumento receptor especial;
- O receptor aparece em sua lista de argumentos entre a própria palavra-chave func e o nome do método;
- Você pode delcarar um método em tipos que não são struct também;
- Não dá para declarar receptor cujo tipo é usado em outro pacote (o qual inclui padrões tais como int);
- Você pode declarar métodos com ponteiro receptor;
- Isso significa que o tipo de receptor tem a sintaxe literal *t por algum tipo t;
- Métodos com ponteiro receptores pode modificar o valor ao qual o receptor pertence. Esses métodos muitas vezes precisam modificar seu receptor;
- Receptores de ponteiros são mais comuns do que receptores de valor;
- Com receptor de valor, o método opera sobre uma cópia original. Já o receptor de ponteiro muda o valor original.
- Go permite a indireção de ponteiros (Caso um método tenha um receptor de valor => tipo ponteiro / Método receptor ponteiro => tipo original);

# Interfaces

- Um tipo interface é definida por um conjunto de métodos;
- Um valor de tipo interface pode conter qualquer valor que implementa esses métodos;
- Um tipo implementa uma interface através da implementação dos métodos. Não há declaração explícita de intenções, não há palavra-chave "implements";
- Interfaces implicitas dissociam-se da definição de uma interface de sua implementação, que pode então aparecer em qualquer pacote sem pré-arranjamento;
- Chamando um método de um valor de interface executa o método do mesmo nome no seu tipo subjacente;
- Se o valor concreto no interior da própria interface é nulo, o método será chamado com um receptor nulo;
- Um valor de interface nil detém nem o valor nem tipo concreto;
- Chamar um método em uma interface nula resulta em erro de execução, pois não há um tipo dentro da tupla de interface para indicar qual método concreto chamar;
- O tipo de interface que especifica zero métodos é conhecido como interface vazia:

```
interface{}
```

- Uma interface vazia pode conter valores de qualquer tipo;
- Interfaces vazias são usadas pelo código que lida com valores de tipo desconhecido;

# Type assertation

- A type assertation fornece acesso ao valor concreto subjacente de um valor

```
t := i.(T)
```

- Esta declaração afirma que o valor de interface i detém o tipo concreto T e atribui o valor de T subjacente à variável t;
- Se i não detém uma T, a declaração irá desencadear um erro pânico;
- Para testar se um valor de interface possui um tipo especifico, uma type assertion pode retornar dois valores: o valor subjacente e um valor booleano que informa se a afirmação é correta.

```
t, ok := i (T)
```

# Type Switch

- Um type switch é uma construção que permite diversas type assertations em série;
- Uma type switch é como uma instrução switch regular, mas os casos de um type switch especificam os tipos (e não valores), e esses valores são comparados com o tipo de valores determinados da interface informada:

```
switch v := i.(type) {
    case T:
    // ....
    case S:
    //....
    default:
    // ....
}
```

- A declaração em um type switch tem a mesma sintaxe como uma afirmação de tipo de i.(T), mas o tipo especifico t é substituido com a palavra-chave type;
- Nisso a instrução switch testa se o valor de interface i detém um valor de tipo T ou S. Em cada um dos casos T e S, a variável v será do tipo T ou S. No caso default a variável é do mesmo tipo e valor da interface i;

# Stringers

Uma das interfaces mais ubíqua é a Stringer definida pelo pacote fmt:

```
type Stringer interface {
    String() string
}
```

A Stringer é um tipo que podemos descrever como uma string;

# Erros

- Programas Go expressam estados de erro com valores error
- O tipo error é uma interface built-in similar a fmt.Stringer:

```
type error interface {
    Error() string
}
```

- Funções frequentemente retornam um valor error e a chamada do código deve lidar com errors testando se o erro é igual a nil: 

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Println("Não conseguiu converter número: %v\n", err)
}
fmt.Println("Inteiro convertido:", i)
```

# Leitores

- O pacote io especifica a interface io.Reader, que representa o fim da leitura de um fluxo de dados;
- A biblioteca padrão contém várias implementações destas interfaces, incluindo arquivos, conexões de rede, compressores, cifras e outros.

a interface io.Reader tem o método Read:

```
func (T) Read(b []byte) (n int, err error)
```

- Read popula uma slice de bytes passados com dados e retorna o número de bytes populados e um valor de erro.

# Imagens

- O pacote image define a interface Image:

```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
