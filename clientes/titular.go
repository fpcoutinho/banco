package clientes

type Titular struct {
	nome      string
	cpf       string
	profissao string
}

func (t *Titular) Nome() string {
	return t.nome
}

func (t *Titular) Cpf() string {
	return t.cpf
}

func (t *Titular) Profissao() string {
	return t.profissao
}

func (t *Titular) New(nome string, cpf string, profissao string) {
	t.nome = nome
	t.cpf = cpf
	t.profissao = profissao
}
