package db

import (
	"reflect"
	"testing"
)

func TestNewDBContext(t *testing.T) {
	type args struct {
		host     string
		user     string
		pass     string
		database string
		driver   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "tank-mysql",
			wantErr: false,
			args:    args{"192.168.115.209", "mvd", "mvd", "mvd", "mysql"},
			want:    true,
		},
		{
			name:    "sqlserver",
			wantErr: false,
			args:    args{"192.168.115.204", "php", "php", "Gambrinus-Kelterei", "sqlserver"},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDBContext(tt.args.host, tt.args.user, tt.args.pass, tt.args.database, tt.args.driver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDBContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_got := false
			if got.C != nil {
				_got = true
			}
			if !reflect.DeepEqual(_got, tt.want) {
				t.Errorf("NewDBContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
