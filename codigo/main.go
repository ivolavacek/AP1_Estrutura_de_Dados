package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		exibirMenu()
		escolha := obterEscolhaUsuario()
		tratarEscolha(escolha)
	}
}

func exibirMenu() {
	fmt.Println("\n======= McRonald’s =======")
	fmt.Println("1. Cadastrar produto")
	fmt.Println("2. Remover produto")
	fmt.Println("3. Buscar produto por ID")
	fmt.Println("4. Exibir todos os produtos")
	fmt.Println("5. Adicionar pedido")
	fmt.Println("6. Exibir pedidos em andamento")
	fmt.Println("7. Exibir pedidos por ID")
	fmt.Println("8. Expedir pedidos")
	fmt.Println("9. Exibir métricas")
	fmt.Println("10. Sair")
	fmt.Println("==========================")
	fmt.Println("")
	fmt.Println("Digite sua opção:")
}

func obterEscolhaUsuario() int {
	var escolha int
	fmt.Scan(&escolha)
	return escolha
}

func tratarEscolha(escolha int) {
	switch escolha {
	case 1:
		CadastrarProdutos()
	case 2:
		InterfaceRemoverProduto()
	case 3:
		BuscarProdutoPorIDInterface()
	case 4:
		ExibirProdutos()
	case 5:
		fmt.Println("Cardápio:")
		ExibirProdutos()
		fmt.Println("")
		AdicionarPedidoInterface()
	case 6:
		ExibirPedidosEmAndamento()
	case 7:
		ExibirPedidosPorIDInterface()
	case 8:
		ExpedirPedidoInterface()
	case 9:
		ExibirMetricas()
	case 10:
		fmt.Println("Obrigado por usar o sistema McRonald’s!")
		Exit()
	default:
		fmt.Println("Opção inválida. Tente novamente.")
	}
}

func CadastrarProdutos() {
	var numeroProdutos int

	fmt.Println("Quantos produtos você deseja adicionar?")
	fmt.Scan(&numeroProdutos)

	var produtosLote []Produto

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < numeroProdutos; i++ {
		var nome, descricao string
		var valor float64

		fmt.Printf("Digite o nome do produto %d:\n", i+1)
		scanner.Scan()
		nome = scanner.Text()

		fmt.Printf("Digite a descrição do produto %d:\n", i+1)
		scanner.Scan()
		descricao = scanner.Text()

		fmt.Printf("Digite o valor do produto %d:\n", i+1)
		fmt.Scan(&valor)

		produto := Produto{
			Nome:      nome,
			Descricao: descricao,
			Valor:     valor,
		}
		produtosLote = append(produtosLote, produto)
	}

	AdicionarProdutosEmLote(produtosLote)
}

func InterfaceRemoverProduto() {
	var id int

	fmt.Println("Digite o ID do produto a ser removido:")
	fmt.Scan(&id)
	RemoverProduto(id) // chamando a função do arquivo produto.go
}

func BuscarProdutoPorIDInterface() {
	var id int
	fmt.Println("Digite o ID do produto a ser buscado:")
	fmt.Scan(&id)
	produto := BuscarProdutoPorID(id)
	if produto != nil {
		fmt.Printf("\nNome: %s\nDescrição: %s\nValor: R$%.2f\n", produto.Nome, produto.Descricao, produto.Valor)
	} else {
		fmt.Println("Produto não encontrado.")
	}
}

func AdicionarPedidoInterface() {
	var entrega bool
	var opcaoEntrega string
	produtosPedido := []ProdutoPedido{} // Inicializa uma lista de produtos vazia

	for {
		var idProduto, quantidade int
		fmt.Println("Digite o ID do produto (ou digite 0 para encerrar a seleção):")
		fmt.Scan(&idProduto)

		if idProduto == 0 {
			break // Encerra a seleção de produtos
		}

		produto := BuscarProdutoPorID(idProduto)
		if produto == nil {
			fmt.Println("Produto não encontrado. Por favor, tente novamente.")
			continue // Volta para a próxima iteração
		}

		fmt.Println("Digite a quantidade do produto:")
		fmt.Scan(&quantidade)

		produtosPedido = append(produtosPedido, ProdutoPedido{
			Produto:    *produto,
			Quantidade: quantidade,
		})
	}

	fmt.Println("É um pedido de entrega? (s/n)")
	fmt.Scan(&opcaoEntrega)
	if strings.ToLower(opcaoEntrega) == "s" {
		entrega = true
	}

	AdicionarPedido(entrega, produtosPedido)
}

func ExibirPedidosPorIDInterface() {
	var id int
	fmt.Println("Digite o ID do pedido a ser buscado:")
	fmt.Scan(&id)
	pedido := ExibirPedidosPorID(id)
	if pedido != nil {
		fmt.Println("")
		for _, produto := range pedido.Produtos {
			fmt.Printf("%d %s\n", produto.Quantidade, produto.Produto.Nome)
		}
		if pedido.Entrega {
			fmt.Println("Delivery")
		}
		fmt.Printf("Valor total: R$%.2f\nCriado em: %s\n", pedido.ValorTotal, pedido.CriadoEm.Format("02/01/2006 15:04:05"))
		if !pedido.ExpedidoEm.IsZero() {
			fmt.Printf("Expedido em: %s\n", pedido.ExpedidoEm.Format("02/01/2006 15:04:05"))
		}
	} else {
		fmt.Println("Pedido não encontrado.")
	}
}

func ExpedirPedidoInterface() {
	pedidoExpedido := ExpedirPedido()
	if pedidoExpedido != nil {
		fmt.Printf("Pedido ID %d foi expedido com sucesso!\n", pedidoExpedido.ID)
	}
}

func Exit() {
	os.Exit(0)
}
