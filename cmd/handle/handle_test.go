package handle

import (
	"reflect"
	"testing"
)

func TestHandleCsvFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleCsvFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleCsvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
