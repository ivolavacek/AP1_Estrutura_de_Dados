# AP1_Estrutura_de_Dados

## Participantes
- André Silveira
- Ivo Lavacek
- João Araújo
- Pedro Lustosa

## Como funciona
Ao executar o arquivo main.go, aparecerá um menu com todas as funcionalidades do programa, sendo estas:
1. Cadastrar produto
2. Remover produto
3. Buscar produto por ID
4. Exibir todos os produtos
5. Adicinar pedido
6. Exibir pedidos em andamento
7. Exibir pedidos por ID
8. Expedir pedidos
9. Exibir métricas
10. Sair

Segue o funcionamento de algumas das funcionalidades do programa

### Cadastrar produtos
- Podemos cadrastrar um produto de cada vez ou em lote.
- O programa pergunta quantos produtos se deseja adicionar, então o usuário pode escolher o número de produtos a serem adicionados, servindo assim o cadastro de um produto de cada vez ou em lote numa mesma funcionalidade. 
- Em seguida, para cada produto, é pedido o nome, descrição e valor do produto.
- Caso o limite de produtos seja atingido, o programa retorna um aviso e não permite a adição de mais produtos.

### Adicionar pedido
- Ao selecionar essa opção, o programa mostra o cardápio do restaurante com ID, nome, descrição e valor de cada produto.
- Para a adição de um produto ao pedido, é requisitado o ID do produto e em seguida a quantidade.
- Podem ser adicionados vários produtos ao mesmo pedido. Para concluir o pedido deve-se digitar 0 e informar se o pedido é para delivery ou não. 
- Caso o limite de pedidos seja atingido, o programa retorna um aviso e não permite a adição de mais pedidos.

### Exibir pedidos por ID
- Ao selecionar essa opção, é requisitado o ID do pedido. Pode ser o ID de um pedido já expedido ou um ainda em andamento.
- No retorno é mostrado a quantidade e o nome do(s) produto(s) pedidos, 'Delivery' caso seja para entrega, o valor total, data de criação do pedido e data de expedição (caso o pedido já tenha sido expedido).

### Exibir métricas
- Essa opção retorna o total de produtos cadastrados, número de pedidos encerrados, número de pedidos em andamento, faturamento total (soma dos pedidos expedidos) e o tempo médio de expedição dos pedidos



------------------------------------------


**Comentários finais**
O trabalho foi feito no site 'replit.com', quando foi passado o código para o VSCode, para criar o .exe a partir do Command Prompt, algumas funcionalidades foram afetadas.
No VSCode, o arquivo main.go não rodou e o aviso de erro era que as funções definidas nos outros arquivos (metricas.go, pedido.go e produto.go) estavam 'undefined', mesmo com todos os arquivos incluidos no package 'main' e funções começadas em letras maiúsculas.
O arquivo .exe está funcionando bem, exceto no cadastro de produtos onde o programa pula o cadastro do nome dos produtos.
No replit.com, o programa funcionou perfeitamente. Segue o link [Replit](https://replit.com/@IvoLavacek/AP1-Goland-4#main.go)
