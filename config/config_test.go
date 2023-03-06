package config

import "testing"

func TestValidateConfig(t *testing.T) {
	type args struct {
		c cfg
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error url",
			args: args{
				c: cfg{
					Port:     0,
					MongoUrl: "adsasd",
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				c: cfg{
					Port:     8080,
					MongoUrl: "mongodb://localhost:27017/",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
