package clientes

type Cliente interface {
	Nome() string
	Cpf() string
	Profissao() string
}
