package runtime

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	type args struct {
		fc func() error
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "success: returns none",
			args: args{
				fc: func() error {
					return nil
				},
			},
		},
		{
			name: "fail: returns error",
			args: args{
				fc: func() error {
					return errors.New("custom error")
				},
			},
			wantErr:        true,
			wantErrMessage: "custom error",
		},
		{
			name: "fail: returns error (recovered)",
			args: args{
				fc: func() error {
					panic("panic!")
				},
			},
			wantErr:        true,
			wantErrMessage: "runtime.Recover(): panic!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Recover(tt.args.fc)
			if (err != nil) && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}
		})
	}
}

func TestRecover1(t *testing.T) {
	type args struct {
		fc func() (int, error)
	}
	tests := []struct {
		name           string
		args           args
		want           int
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "success: returns 1",
			args: args{
				fc: func() (int, error) {
					return 1, nil
				},
			},
			want: 1,
		},
		{
			name: "fail: returns error",
			args: args{
				fc: func() (int, error) {
					return 0, errors.New("custom error")
				},
			},
			wantErr:        true,
			wantErrMessage: "custom error",
		},
		{
			name: "fail: returns error (recovered)",
			args: args{
				fc: func() (int, error) {
					panic("panic!")
				},
			},
			wantErr:        true,
			wantErrMessage: "runtime.Recover(): panic!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Recover1(tt.args.fc)
			if (err != nil) && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestRecover2(t *testing.T) {
	type args struct {
		fc func() (int, int, error)
	}
	tests := []struct {
		name           string
		args           args
		want           int
		want1          int
		wantErr        bool
		wantErrMessage string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Recover2(tt.args.fc)
			if (err != nil) && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}

			if !assert.Equal(t, tt.want1, got1) {
				return
			}
		})
	}
}

func TestRecover3(t *testing.T) {
	type args struct {
		fc func() (int, int, int, error)
	}
	tests := []struct {
		name           string
		args           args
		want           int
		want1          int
		want2          int
		wantErr        bool
		wantErrMessage string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := Recover3(tt.args.fc)
			if (err != nil) && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}

			if !assert.Equal(t, tt.want1, got1) {
				return
			}

			if !assert.Equal(t, tt.want2, got2) {
				return
			}
		})
	}
}
