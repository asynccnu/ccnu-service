package service

import "testing"

func Test_loginCCNU(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "test1", args: args{
			username: "2023214441",
			password: "HUAshi0417",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loginCCNU(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginCCNU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loginCCNU() got = %v, want %v", got, tt.want)
			}
		})
	}
}
