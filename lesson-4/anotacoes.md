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
- Interfaces implicitas dissociam-se da definição de uma interface de sua implementação, que pode então aparecer em qualquer pacote sem pré-arranjamento.