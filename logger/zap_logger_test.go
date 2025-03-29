package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		options []zap.Option
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(tt.args.options...)
			got.Infof("hello world")
		})
	}
}
