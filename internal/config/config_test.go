package config

import (
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {

	tests := []struct {
		name      string
		setConfig *Config
		want      *Config
	}{
		{
			name: "config1",
			setConfig: &Config{
				Port:                 "8080",
				AccessSigningString:  "ACCESS",
				RefreshSigningString: "REFRESH",
			},
			want: &Config{
				Port:                 "8080",
				AccessSigningString:  "ACCESS",
				RefreshSigningString: "REFRESH",
			},
		},
		{
			name: "config2",
			setConfig: &Config{
				AccessSigningString:  "ACCESS",
				RefreshSigningString: "REFRESH",
			},
			want: &Config{
				Port:                 "8000",
				AccessSigningString:  "ACCESS",
				RefreshSigningString: "REFRESH",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("HTTP_PORT", tt.setConfig.Port)
			os.Setenv("SIGNING_STRING", tt.setConfig.AccessSigningString)
			os.Setenv("REFRESH_SIGNING_STRING", tt.setConfig.RefreshSigningString)
			if got := GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
