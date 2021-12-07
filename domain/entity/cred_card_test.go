package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCrediCard(t *testing.T) {
	type args struct {
		number          string
		name            string
		expirationMonth int
		expirationYear  int
		cvv             int
	}

	tests := []struct {
		name    string
		args    args
		want    *CreditCard
		wantErr bool
	}{
		{
			name: "Cria um cartão de credito com todas campos corretos",
			args: args{
				number:          "4193523830170205",
				name:            "José da Silva",
				expirationMonth: 10,
				expirationYear:  2029,
				cvv:             123,
			},
			want: &CreditCard{
				number:          "4193523830170205",
				name:            "José da Silva",
				expirationMonth: 10,
				expirationYear:  2029,
				cvv:             123,
			},
			wantErr: false,
		},
		{
			name: "Cria um cartão de credito o número inválido",
			args: args{
				number:          "419352383017020",
				name:            "José da Silva",
				expirationMonth: 10,
				expirationYear:  2029,
				cvv:             123,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Cria um cartão de credito o mês inválido",
			args: args{
				number:          "419352383017020",
				name:            "José da Silva",
				expirationMonth: 0,
				expirationYear:  2029,
				cvv:             123,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Cria um cartão de credito o ano inválido",
			args: args{
				number:          "419352383017020",
				name:            "José da Silva",
				expirationMonth: 10,
				expirationYear:  2020,
				cvv:             123,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, err := NewCrediCard(tt.args.number, tt.args.name, tt.args.expirationMonth, tt.args.expirationYear, tt.args.cvv)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, cc)

		})
	}
}
