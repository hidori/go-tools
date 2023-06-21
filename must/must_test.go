package must

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	type args struct {
		f func() (int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantPanic bool
	}{
		{
			name: "success: returns 1",
			args: args{
				f: func() (int, error) {
					return 1, nil
				},
			},
			want: 1,
		},
		{
			name: "fail: panics",
			args: args{
				f: func() (int, error) {
					return 0, errors.New("error!")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = Get(tt.args.f)
				})
				return
			}

			got := Get(tt.args.f)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestMust2(t *testing.T) {
	type args struct {
		f func() (int, int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		want1     int
		wantPanic bool
	}{
		{
			name: "success: returns 1, 2",
			args: args{
				f: func() (int, int, error) {
					return 1, 2, nil
				},
			},
			want:  1,
			want1: 2,
		},
		{
			name: "fail: panics",
			args: args{
				f: func() (int, int, error) {
					return 0, 0, errors.New("error!")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_, _ = Get1(tt.args.f)
				})
				return
			}

			got, got1 := Get1(tt.args.f)
			if !assert.Equal(t, tt.want, got) {
				return
			}

			if !assert.Equal(t, tt.want1, got1) {
				return
			}
		})
	}
}

func TestMust3(t *testing.T) {
	type args struct {
		f func() (int, int, int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		want1     int
		want2     int
		wantPanic bool
	}{
		{
			name: "success: returns 1, 2, 3",
			args: args{
				f: func() (int, int, int, error) {
					return 1, 2, 3, nil
				},
			},
			want:  1,
			want1: 2,
			want2: 3,
		},
		{
			name: "fail: panics",
			args: args{
				f: func() (int, int, int, error) {
					return 0, 0, 0, errors.New("error!")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_, _, _ = Get3(tt.args.f)
				})
				return
			}

			got, got1, got2 := Get3(tt.args.f)
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
