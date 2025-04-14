package rutax_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zaffka/rutax"
)

func TestParseID(t *testing.T) {
	type args struct {
		taxID string
	}
	tests := []struct {
		name    string
		args    args
		want    rutax.ID
		wantErr error
	}{
		{
			name: "incorrect base fmt",
			args: args{
				taxID: "123.45678901",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrIDIncorrect,
		},
		{
			name: "incorrect length 11",
			args: args{
				taxID: "12345678901",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrIDIncorrect,
		},
		{
			name: "incorrect length 9",
			args: args{
				taxID: "123456789",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrIDIncorrect,
		},
		{
			name: "incorrect length 13",
			args: args{
				taxID: "1234567890123",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrIDIncorrect,
		},
		{
			name: "is legal entity 10",
			args: args{
				taxID: "9729219090",
			},
			want:    rutax.ID{Num: "9729219090", IsLegal: true},
			wantErr: nil,
		},
		{
			name: "wrong checksum for legal entity",
			args: args{
				taxID: "9729219095",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrChecksumFailed,
		},
		{
			name: "is private person 12",
			args: args{
				taxID: "535397320195",
			},
			want:    rutax.ID{Num: "535397320195"},
			wantErr: nil,
		},
		{
			name: "wrong checksum in position 11 for private person",
			args: args{
				taxID: "535397320185",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrChecksumFailed,
		},
		{
			name: "wrong checksum in position 12 for private person",
			args: args{
				taxID: "535397320196",
			},
			want:    rutax.ID{},
			wantErr: rutax.ErrChecksumFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rutax.ParseID(tt.args.taxID)
			require.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
