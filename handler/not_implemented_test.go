//go:build !evokeos && !f1

package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"multi-platform-demo/os/configuration"
)

func TestHandler_GetOSName(t *testing.T) {
	type fields struct {
		Config configuration.Configuration
	}
	type args struct {
		c   echo.Context
		rec *httptest.ResponseRecorder
	}
	type want struct {
		code int
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "gets os name",
			fields: fields{
				Config: configuration.Configuration{
					Port: "1234",
				},
			},
			args: func() args {
				c, rec := StubContextRecorder(nil)
				return args{
					c:   c,
					rec: rec,
				}
			}(),
			want: want{
				code: http.StatusOK,
				body: bytes.NewBufferString("not implemented, port: 1234"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			h := Handler{
				Config: tt.fields.Config,
			}

			// Assertions
			m := fmt.Sprintf("GetOSName(%v)", tt.args.c)
			if !tt.wantErr(t, h.GetOSName(tt.args.c), m) {
				return
			}
			assert.Equalf(t, tt.want.code, tt.args.rec.Code, m)
			assert.Equalf(t, tt.want.body, tt.args.rec.Body, m)
		})
	}
}

func TestHandler_GetOrg(t *testing.T) {
	type fields struct {
		Config configuration.Configuration
	}
	type args struct {
		c   echo.Context
		rec *httptest.ResponseRecorder
	}
	type want struct {
		code int
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "gets org",
			fields: fields{
				Config: configuration.Configuration{
					Port: "1234",
				},
			},
			args: func() args {
				c, rec := StubContextRecorder(nil)
				return args{
					c:   c,
					rec: rec,
				}
			}(),
			want: want{
				code: http.StatusOK,
				body: bytes.NewBufferString("not implemented, port: 1234"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			h := Handler{
				Config: tt.fields.Config,
			}

			// Assertions
			m := fmt.Sprintf("GetOrg(%v)", tt.args.c)
			if !tt.wantErr(t, h.GetOrg(tt.args.c), m) {
				return
			}
			assert.Equalf(t, tt.want.code, tt.args.rec.Code, m)
			assert.Equalf(t, tt.want.body, tt.args.rec.Body, m)
		})
	}
}
