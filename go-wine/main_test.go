package wine

import (
	"testing"
)

func TestGetLocalVersion(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
	}{
		// TODO: Add test cases.
		{
			name: "wine command",
			args: args{"wine"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLocalVersion(tt.args.cmd)
			if err !=nil{
				t.Errorf("Failed %v: %v", t.Name(),err)
			}else{
				t.Logf("Name: %v No: %v", got.Name, got.No)
			}
		})
	}
}
