package presenter

// Presenter é a inteface que deve ser implementada para transforma dos dados
// de output da camada de usercase e entregar em um formato especifico
// dependendo de quem ira receber-lo
type Presenter interface {
	// aprensentar os dados transformado pelo bind
	Show() ([]byte, error)
	// recebe os dados de output do usecase e transformar
	Bind(interface{}) error
}
