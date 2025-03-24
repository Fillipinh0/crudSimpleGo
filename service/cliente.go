package service

import (
	"fmt"
	"trabalhoproz/database"
)

// Nome Telefone Renda data_nascimento CEP
type Ativo struct {
	ID          int
	Nome        string
	Localizacao string
	Descricao   string
	Status      string
}

const insereComando string = "INSERT INTO ativo (nome_ativo, localização_ativo, descrição_ativo, status_ativo) VALUES (?, ?, ?, ?)"

const selectComando string = "SELECT * FROM ativo"
const selectPorNome string = "SELECT * FROM ativo ORDER BY nome_ativo ASC"

const selectPorStatus string = "SELECT * FROM ativo WHERE status_ativo = ?"

func InsereAtivo(ativo Ativo) {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	fmt.Printf("Inserindo ativo: %+v\n", ativo)

	resultado, err := conexao.Exec(insereComando, ativo.Nome, ativo.Localizacao, ativo.Descricao, ativo.Status)
	if err != nil {
		fmt.Println("Erro ao inserir ativo:", err)
		return
	}

	id, _ := resultado.LastInsertId()
	fmt.Printf("Ativo inserido com sucesso! ID: %d\n", id)
}

func SelecionarAtivo() []Ativo {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	resultado, erro := conexao.Query(selectComando)
	if erro != nil {
		fmt.Println("Erro ao executar o SELECT:", erro)
		return nil
	}
	defer resultado.Close()

	var listaAtivo []Ativo

	for resultado.Next() {
		var ativo Ativo
		erro := resultado.Scan(&ativo.ID, &ativo.Nome, &ativo.Localizacao, &ativo.Descricao, &ativo.Status)
		if erro != nil {
			fmt.Println("Erro ao ler a linha:", erro)
			continue
		}

		fmt.Printf("Ativo encontrado: %+v\n", ativo)
		listaAtivo = append(listaAtivo, ativo)
	}

	return listaAtivo
}

func SelecionarAtivoName() []Ativo {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	resultado, erro := conexao.Query(selectPorNome)
	if erro != nil {
		fmt.Println("Erro ao executar o SELECT:", erro)
		return nil
	}
	defer resultado.Close()

	var listaAtivo []Ativo

	for resultado.Next() {
		var ativo Ativo
		erro := resultado.Scan(&ativo.ID, &ativo.Nome, &ativo.Localizacao, &ativo.Descricao, &ativo.Status)
		if erro != nil {
			fmt.Println("Erro ao ler a linha:", erro)
			continue
		}

		fmt.Printf("Ativo encontrado: %+v\n", ativo)
		listaAtivo = append(listaAtivo, ativo)
	}

	return listaAtivo
}

func FiltrarAtivos(nome string, status string) []Ativo {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	var query string
	var args []interface{}

	// Construir a query dinamicamente com base nos filtros
	if nome != "" && status != "" {
		query = "SELECT * FROM ativo WHERE nome_ativo LIKE ? AND status_ativo = ?"
		args = append(args, "%"+nome+"%", status)
	} else if nome != "" {
		query = "SELECT * FROM ativo WHERE nome_ativo LIKE ?"
		args = append(args, "%"+nome+"%")
	} else if status != "" {
		query = "SELECT * FROM ativo WHERE status_ativo = ?"
		args = append(args, status)
	} else {
		query = selectComando // Sem filtros, retorna todos os registros
	}

	// Executar a consulta
	resultado, erro := conexao.Query(query, args...)
	if erro != nil {
		fmt.Println("Erro ao executar o SELECT:", erro)
		return nil
	}
	defer resultado.Close()

	var listaAtivo []Ativo
	for resultado.Next() {
		var ativo Ativo
		erro := resultado.Scan(&ativo.ID, &ativo.Nome, &ativo.Localizacao, &ativo.Descricao, &ativo.Status)
		if erro != nil {
			fmt.Println("Erro ao ler a linha:", erro)
			continue
		}
		listaAtivo = append(listaAtivo, ativo)
	}

	return listaAtivo
}
