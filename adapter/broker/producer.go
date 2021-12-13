package broker

// Producer é a inteface que deve ser implementada por um produce
// responsavel por emitir eventos
type Producer interface {
	Publish(msg interface{}, key []byte, topic string) error
}
