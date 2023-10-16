package main

import (
	"fmt"
)

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Valor     float64
}

var limiteProdutos = 3 // Indicar o limite de produtos
var Produtos []Produto
var proximoIDProduto = 1

// Adiciona um novo produto à lista de produtos
func adicionarProduto(nome, descricao string, valor float64) {
	produto := Produto{
		ID:        proximoIDProduto,
		Nome:      nome,
		Descricao: descricao,
		Valor:     valor,
	}
	Produtos = append(Produtos, produto)
	fmt.Printf("Produto '%s' adicionado com sucesso. ID: %d\n", produto.Nome, produto.ID)
	proximoIDProduto++
}

// Remove um produto da lista de produtos pelo ID
func removerProduto(id int) {
	for indice, produto := range Produtos {
		if produto.ID == id {
			Produtos = append(Produtos[:indice], Produtos[indice+1:]...)
			fmt.Printf("Produto '%s' removido com sucesso!\n", produto.Nome)
			return
		}
		if indice == len(Produtos)-1 && produto.ID != id {
			fmt.Println("Produto não encontrado")
			return
		}
	}
}

// Busca um produto pela ID e retorna
func buscarProdutoPorID(id int) *Produto {
	for _, produto := range Produtos {
		if produto.ID == id {
			return &produto
		}
	}
	return nil
}

// Exibe todos os produtos cadastrados
func exibirProdutos() {
	if len(Produtos) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}
	for _, produto := range Produtos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Valor)
	}
}

// Adiciona uma lista de produtos, até o limite de produtos
func adicionarProdutosEmLote(produtosLote []Produto) {
	for _, produto := range produtosLote {
		if len(Produtos) < limiteProdutos {
			adicionarProduto(produto.Nome, produto.Descricao, produto.Valor)
		} else {
			fmt.Println("Limite de produtos atingido.")
			return
		}
	}
}
