package runtime

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCallerInfo(t *testing.T) {
	wd, _ := os.Getwd()
	wd = path.Dir(wd)

	_ = SetCallerInfoFileNamePrefixToTrim(wd)

	tests := []struct {
		name string
		want CallerInfo
	}{
		{
			name: "success: returns CallerInfo",
			want: CallerInfo{
				OK:         true,
				FileName:   "runtime/runtime_test.go",
				LineNumber: 33,
				FuncName:   "github.com/hidori/go-tools/runtime.TestGetCallerInfo.func1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCallerInfo()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
