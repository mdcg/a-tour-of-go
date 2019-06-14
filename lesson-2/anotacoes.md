# For

- Go te apenas uma estrutura de laço, o `for`
- O laço `for` básico tem três componentes separados por ponto e vírgula:

1. A instrução inicial: executada antes da primeira iteração;
2. A expressão de condição: avaliada antes de cada iteração;
3. A pós-instrução: executada no final de cada iteração;

- A instrução inicial e a pós-instrução são opcionais.

# If/Else

- A instrução `if` é muito parecida com a declaração do laço `for`, os ( ) são opcionais mas { } são obrigatórios;
- A instrução `if` pode começar com uma breve declaração antes de executar a condição, contudo, elas só serão válidas até o escopo final do `if`;
- As variáveis declaradas dentro de uma instrução de `if` curta também estão disponíveis dentro dos blocos `else`

# Switch

- Um instrução `switch` é uma forma mais curta de escreer uma sequência de declarações `if - else`.
- O switch do Go executa apenas o caso selecionado, não todos os que seguem. De fato, a declaração `break` que é necessária no final de cada caso naquelas linguagens é fornecido automaticamente no Go;
- Switch cases avaliam casos de cima para baixo, parando quando um caso for bem-sucedido;
- Switch sem condição é o mesmo que `switch true`, assim é uma boa maneira de escrever longas cadeias de if-then-else.

# Defer

- A declaração `defer` adia a execução de uma função até o final do retorno da função;
- Os argumentos das chamadas adiadas são avaliados imediatamente, mas a função não é chamada até o retorno da função;
- Caso exista mais de um `defer`, elas são empurradas por uma pilha. Quando a função retorna, as chamadas de adiamaento são executadas na ordem em que a última a entrar é a primeira a sair.