package models

import . "golang-studies/alura/db"

type Produto struct {
	Nome, Descricao string
	Preco float64
	Id, Quantidade int
}

func SelectDeTodosOsProdutos() []Produto {
	db := ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER by id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	var produtos []Produto

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}