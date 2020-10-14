package backoff_test

import (
	"errors"
	"testing"
	"time"

	"github.com/kskumgk63/backoff"
	"github.com/stretchr/testify/assert"
)

func TestCommander_Exec(t *testing.T) {
	type fields struct {
		Opts []backoff.Option
	}
	type args struct {
		f func() error
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
	}{
		{
			name: "change the timeout value",
			fields: fields{
				Opts: []backoff.Option{
					backoff.Timeout(2 * time.Second),
				},
			},
			args: args{
				f: func() error {
					return errors.New("internal server error")
				},
			},
			wantErrMsg: "Ends the exponential backoff because of timeout",
		},
		{
			name: "change output message when timeout",
			fields: fields{
				Opts: []backoff.Option{
					backoff.Timeout(2 * time.Second),
					backoff.TimeoutErrMessage("test"),
				},
			},
			args: args{
				f: func() error {
					return errors.New("internal server error")
				},
			},
			wantErrMsg: "test",
		},
		{
			name: "ignore specific error",
			fields: fields{
				Opts: []backoff.Option{
					backoff.IgnoreError(func(e error) bool {
						return e.Error() == "ignored error"
					}),
				},
			},
			args: args{
				f: func() error {
					return errors.New("ignored error")
				},
			},
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := backoff.NewCommander(tt.fields.Opts...)
			err := cmd.Exec(tt.args.f)
			if err != nil {
				assert.Equal(t, err.Error(), tt.wantErrMsg)
				return
			}
			assert.Nil(t, err)
		})
	}
}
