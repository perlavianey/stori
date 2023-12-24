package main

import (
	"reflect"
	"stori/database"
	"testing"
)

func TestConvertTransactions(t *testing.T) {
	type args struct {
		data [][]string
		ulid string
	}
	tests := []struct {
		name string
		args args
		want []database.Transaction
	}{
		{
			name: "test 1",
			args: args{
				data: [][]string{
					{"id", "date", "amount", "id_account"},
					{"1", "2023-10-01", "100.00", "123456789"},
				},
				ulid: "123456789",
			},
			want: []database.Transaction{
				{
					Id:        "1",
					Date:      "2023-01-01",
					Amount:    100.00,
					IdAccount: "123456789",
					Filename:  "123456789",
					Timestamp: "2023-10-01T00:00:00.000Z07:00",
				},
			},
		},
		{
			name: "test 2",
			args: args{
				data: [][]string{
					{"id", "date", "amount", "id_account"},
					{"1", "2023-01-01", "100.00", "123456789"},
					{"2", "2023-01-02", "200.00", "987654321"},
				},
				ulid: "123456789",
			},
			want: []database.Transaction{
				{
					Id:        "1",
					Date:      "2023-01-01",
					Amount:    100.00,
					IdAccount: "123456789",
					Filename:  "123456789",
					Timestamp: "2023-01-01T00:00:00.000Z07:00",
				},
				{
					Id:        "2",
					Date:      "2023-01-02",
					Amount:    200.00,
					IdAccount: "987654321",
					Filename:  "123456789",
					Timestamp: "2023-01-02T00:00:00.000Z07:00",
				},
			},
		},
		{
			name: "test 3",
			args: args{
				data: [][]string{
					{"id", "date", "amount", "id_account"},
					{"1", "2023-01-01", "100.00", "123456789"},
					{"2", "2023-01-02", "not a number", "987654321"},
				},
				ulid: "123456789",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertTransactions(tt.args.data, tt.args.ulid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
