package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"trabalhoproz/service"
)

/*
	func verClientes(resposta http.ResponseWriter, requisicao *http.Request) {
		// Buscar a lista de clientes do banco de dados
		listaClientes := service.SelecionarClientes()

		// Carregar o template
		pagina, erro := template.ParseFiles("template/index.html")
		if erro != nil {
			fmt.Println("Erro ao carregar o template:", erro)
			return
		}

		// Renderizar o template com a lista de clientes
		pagina.Execute(resposta, listaClientes)
	}
*/
func abreIndex(res http.ResponseWriter, req *http.Request) {
	// Buscar a lista de clientes do banco de dados
	listaClientes := service.SelecionarAtivo()

	// Carregar o template
	pagina, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar o template com a lista de clientes
	pagina.Execute(res, listaClientes)
}

func apiAtivos(res http.ResponseWriter, req *http.Request) {
	// Configurar o cabeçalho para JSON
	res.Header().Set("Content-Type", "application/json")

	// Buscar a lista de ativos
	listaAtivos := service.SelecionarAtivo()

	// Converter a lista para JSON
	err := json.NewEncoder(res).Encode(listaAtivos)
	if err != nil {
		http.Error(res, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}
}

func cadastraCliente(res http.ResponseWriter, req *http.Request) {

	// Verifica se o método é POST
	if req.Method == http.MethodPost {
		var ativo service.Ativo

		ativo.Nome = req.FormValue("nome")
		ativo.Localizacao = req.FormValue("localizacao")
		ativo.Descricao = req.FormValue("descricao")
		ativo.Status = req.FormValue("status")

		service.InsereAtivo(ativo)
		fmt.Printf("Ativo: %v\n", ativo)

		// Redireciona para a página inicial após cadastrar
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return // Garante que nada mais será executado após o redirecionamento
	}

	// Se não for POST, renderiza a página normalmente
	pagina, erro := template.ParseFiles("template/index.html")
	if erro != nil {
		fmt.Println(erro)
		return
	}
	pagina.Execute(res, nil)

}

func verAtivoName(res http.ResponseWriter, req *http.Request) {
	// Buscar os ativos ordenados por nome
	listaAtivos := service.SelecionarAtivoName()

	// Carregar o template
	pagina, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar o template com os dados
	pagina.Execute(res, listaAtivos)
}

func filtrarAtivos(res http.ResponseWriter, req *http.Request) {
	nome := req.FormValue("nome")
	status := req.FormValue("status_ativo") // isso é tipo a arvore dom do JS

	listaAtivos := service.FiltrarAtivos(nome, status)

	// Carregar o template
	pagina, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar o template com os dados filtrados
	pagina.Execute(res, listaAtivos)
}

func main() {
	// Servir arquivos estáticos (CSS, imagens, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Criando os EndPoints
	http.HandleFunc("/api/ativos", apiAtivos)
	http.HandleFunc("/", abreIndex)
	http.HandleFunc("/cadastraCliente", cadastraCliente)
	http.HandleFunc("/verAtivoName", verAtivoName)
	http.HandleFunc("/filtrarAtivos", filtrarAtivos)

	// Exibir mensagem antes de iniciar o servidor
	fmt.Println("Iniciando Servidor na porta 8080...")

	// Iniciando o servidor
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Println("Servidor com Problemas:", erro)
	}
}
