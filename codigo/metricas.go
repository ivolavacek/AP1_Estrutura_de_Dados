package main

import (
	"fmt"
	"time"
)

// Exibe métricas gerais do sistema
func exibirMetricas() {
  fmt.Printf("Total de produtos cadastrados: %d\n", len(Produtos))
  fmt.Printf("Número de pedidos já encerrados: %d\n", numeroPedidosExpedidos())
  fmt.Printf("Número de pedidos em andamento: %d\n", numeroPedidosAndamento())
  fmt.Printf("Faturamento total até o momento: R$%.2f\n", faturamentoTotal())
  fmt.Printf("Tempo médio de expedição dos pedidos: %s\n", tempoMedioExpedicao())
}

// Calcula o faturamento total do sistema
func faturamentoTotal() float64 {
	var total float64
	for _, pedido := range Pedidos {
		if !pedido.ExpedidoEm.IsZero() { // Considera apenas pedidos expedidos
			total += pedido.ValorTotal
		}
	}
	return total
}

// Retorna o número de pedidos que já foram expedidos
func numeroPedidosExpedidos() int {
	count := 0
	for _, pedido := range Pedidos {
		if !pedido.ExpedidoEm.IsZero() {
			count++
		}
	}
	return count
}

// Retorna o número de pedidos que ainda estão em andamento
func numeroPedidosAndamento() int {
	count := 0
	for _, pedido := range Pedidos {
		if pedido.ExpedidoEm.IsZero() {
			count++
		}
	}
	return count
}

// Calcula o tempo médio de expedição dos pedidos
func tempoMedioExpedicao() time.Duration {
  var somaTempos time.Duration
  var count int

  for _, pedido := range Pedidos {
    if !pedido.ExpedidoEm.IsZero() {
      somaTempos += pedido.ExpedidoEm.Sub(pedido.CriadoEm)
      count++
    }
  }

  if count == 0 {
    return 0
  }
  return somaTempos / time.Duration(count)
}
