package main

import (
	"fmt"
	"time"
)

// Exibe métricas gerais do sistema
func ExibirMetricas() {
	fmt.Printf("Total de produtos cadastrados: %d\n", len(Produtos))
	fmt.Printf("Número de pedidos já encerrados: %d\n", NumeroPedidosExpedidos())
	fmt.Printf("Número de pedidos em andamento: %d\n", NumeroPedidosAndamento())
	fmt.Printf("Faturamento total até o momento: R$%.2f\n", FaturamentoTotal())
	fmt.Printf("Tempo médio de expedição dos pedidos: %s\n", TempoMedioExpedicao())
}

// Calcula o faturamento total do sistema
func FaturamentoTotal() float64 {
	var total float64
	for _, pedido := range Pedidos {
		if !pedido.ExpedidoEm.IsZero() { // Considera apenas pedidos expedidos
			total += pedido.ValorTotal
		}
	}
	return total
}

// Retorna o número de pedidos que já foram expedidos
func NumeroPedidosExpedidos() int {
	count := 0
	for _, pedido := range Pedidos {
		if !pedido.ExpedidoEm.IsZero() {
			count++
		}
	}
	return count
}

// Retorna o número de pedidos que ainda estão em andamento
func NumeroPedidosAndamento() int {
	count := 0
	for _, pedido := range Pedidos {
		if pedido.ExpedidoEm.IsZero() {
			count++
		}
	}
	return count
}

// Calcula o tempo médio de expedição dos pedidos
func TempoMedioExpedicao() time.Duration {
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
