package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	type fields struct {
		Amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "retorna um erro para uma trasação com Amount > 1000",
			fields: fields{
				Amount: 1001,
			},
			wantErr: true,
		},
		{
			name: "retorna um erro para uma trasação com Amount < 1",
			fields: fields{
				Amount: 0,
			},
			wantErr: true,
		},
		{
			name: "não retorna erros para transação com Amount > 0 e Amount <= 1000",
			fields: fields{
				Amount: 1000,
			},
			wantErr: false,
		},
		{
			name: "não retorna erros para transação com Amount > 0 e Amount <= 1000",
			fields: fields{
				Amount: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				Amount: tt.fields.Amount,
			}

			err := tr.IsValid()
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
