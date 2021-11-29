package decimal

import (
	"reflect"
	"testing"
)

func TestNewFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Valid amount",
			args: args{
				s: "1234,56",
			},
			want:    1234.56,
			wantErr: false,
		},
		{
			name: "negative amount",
			args: args{
				s: "-1234,56",
			},
			want:    -1234.56,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AsMajorUnits(), tt.want) {
				t.Errorf("NewFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
