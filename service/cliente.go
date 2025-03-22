package service

import (
	"fmt"
	"trabalhoproz/database"
)

// Nome Telefone Renda data_nascimento CEP
type Cliente struct {
	ID             int
	Nome           string
	Telefone       string
	Renda          string
	DataNascimento string
	CEP            string
}

const insereComando string = "INSERT INTO cliente (nome_cliente, telefone_cliente, renda_cliente, data_nascimento_cliente, CEP_cliente) VALUES (?, ?, ?, ?, ?)"
const selectComando string = "SELECT * FROM cliente"

func InsereCliente(cliente Cliente) {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	fmt.Printf("Inserindo cliente: %+v\n", cliente)

	resultado, err := conexao.Exec(insereComando, cliente.Nome, cliente.Telefone, cliente.Renda, cliente.DataNascimento, cliente.CEP)
	if err != nil {
		fmt.Println("Erro ao inserir cliente:", err)
		return
	}

	id, _ := resultado.LastInsertId()
	fmt.Printf("Cliente inserido com sucesso! ID: %d\n", id)
}

func SelecionarClientes() []Cliente {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	resultado, erro := conexao.Query(selectComando)
	if erro != nil {
		fmt.Println("Erro ao executar o SELECT:", erro)
		return nil
	}
	defer resultado.Close()

	var listaClientes []Cliente

	for resultado.Next() {
		var cliente Cliente
		erro := resultado.Scan(&cliente.ID, &cliente.Nome, &cliente.Telefone, &cliente.Renda, &cliente.DataNascimento, &cliente.CEP)
		if erro != nil {
			fmt.Println("Erro ao ler a linha:", erro)
			continue
		}

		fmt.Printf("Cliente encontrado: %+v\n", cliente)
		listaClientes = append(listaClientes, cliente)
	}

	return listaClientes
}
