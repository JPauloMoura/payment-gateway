package fixture

import (
	"os"
	"testing"

	"github.com/JPauloMoura/payment-gateway/adapter/repository/fixture"
	"github.com/JPauloMoura/payment-gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

// Uma transação deve ser inserida com sucesso no banco de dados
func TestTransactionRepositoryDBInsert(t *testing.T) {
	//caminho para os scripts sql
	migration := os.DirFS("fixture/sql")

	// sobe a tabela do BD
	db := fixture.Up(migration)
	defer fixture.Down(db, migration)

	repo := NewTransactionRepositoryDB(db)

	// inseri os dados
	err := repo.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
