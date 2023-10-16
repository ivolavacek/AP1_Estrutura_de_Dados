package main

import (
	"fmt"
	"time"
)

type Pedido struct {
	ID          int
	Entrega     bool
	Produtos    []ProdutoPedido
	ValorTotal  float64
	CriadoEm    time.Time
	ExpedidoEm  time.Time
}

type ProdutoPedido struct {
	Produto  Produto
	Quantidade int
}

var limitePedidos = 5 // Indicar limite de pedidos
var Pedidos []Pedido
var proximoIDPedido = 1
const taxaEntrega = 10.0
var Faturamento float64 = 0.0

// Adiciona um novo pedido à lista de pedidos
func adicionarPedido(entrega bool, produtosPedido []ProdutoPedido) {
	valor := calcularValorTotal(produtosPedido, entrega)
	pedido := Pedido{
		ID:          proximoIDPedido,
		Entrega:     entrega,
		Produtos:    produtosPedido,
		ValorTotal:  valor,
		CriadoEm:    time.Now(),
	}
  
  if len(Pedidos) < limitePedidos {
    Pedidos = append(Pedidos, pedido)
    fmt.Println("Pedido adicionado com sucesso!")
    proximoIDPedido++
  } else {
    fmt.Println("Limite de pedidos atingido.")
  }
}

// Calcula o valor total do pedido considerando a taxa de entrega se aplicável
func calcularValorTotal(produtosPedido []ProdutoPedido, entrega bool) float64 {
	valor := 0.0
	for _, item := range produtosPedido {
		valor += item.Produto.Valor * float64(item.Quantidade)
	}
	if entrega {
		valor += taxaEntrega
	}
	return valor
}

// Exibe todos os pedidos em andamento (não expedidos)
func exibirPedidosEmAndamento() {
  listaVazia := true
  for _, pedido := range Pedidos {
    if pedido.ExpedidoEm.IsZero() {
      listaVazia = false
      if pedido.Entrega {
        fmt.Printf("Pedido ID: %d, Delivery, Valor Total: R$%.2f, Criado em: %s\n", pedido.ID, pedido.ValorTotal, pedido.CriadoEm.Format("02/01/2006 15:04:05"))
      } else {
        fmt.Printf("Pedido ID: %d, Valor Total: R$%.2f, Criado em: %s\n", pedido.ID, pedido.ValorTotal, pedido.CriadoEm.Format("02/01/2006 15:04:05"))
      }
    }
  }
  if listaVazia {
    fmt.Println("Não há pedidos em andamento.")
  }
}

func exibirPedidosPorID(id int) *Pedido{
  for _, pedido := range Pedidos {
    if pedido.ID == id {
      return &pedido
    }
  }
  return nil
}

// Expede o primeiro pedido da lista que ainda não foi expedido
func expedirPedido() *Pedido {
	for i := range Pedidos {
		if Pedidos[i].ExpedidoEm.IsZero() { // Verifica se o pedido ainda não foi expedido
			Pedidos[i].ExpedidoEm = time.Now()
			Faturamento += Pedidos[i].ValorTotal // Adiciona o valor do pedido ao faturamento total
			fmt.Printf("Pedido ID %d expedido!\n", Pedidos[i].ID)
			return &Pedidos[i]
		}
	}
	fmt.Println("Todos os pedidos já foram expedidos.")
	return nil
}
