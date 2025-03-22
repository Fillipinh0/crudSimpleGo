package main

import (
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
	listaClientes := service.SelecionarClientes()

	// Carregar o template
	pagina, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar o template com a lista de clientes
	pagina.Execute(res, listaClientes)
}

func cadastraCliente(res http.ResponseWriter, req *http.Request) {

	// Verifica se o método é POST
	if req.Method == http.MethodPost {
		var cliente service.Cliente

		cliente.Nome = req.FormValue("nome")
		cliente.Telefone = req.FormValue("telefone")
		cliente.Renda = req.FormValue("renda")
		cliente.DataNascimento = req.FormValue("data_nascimento")
		cliente.CEP = req.FormValue("cep")

		service.InsereCliente(cliente)
		fmt.Printf("Cliente: %v\n", cliente)

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

func main() {
	// Servir arquivos estáticos (CSS, imagens, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Criando os EndPoints
	http.HandleFunc("/", abreIndex)
	http.HandleFunc("/cadastraCliente", cadastraCliente)

	// Exibir mensagem antes de iniciar o servidor
	fmt.Println("Iniciando Servidor na porta 8080...")

	// Iniciando o servidor
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Println("Servidor com Problemas:", erro)
	}
}
