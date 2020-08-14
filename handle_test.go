package llbl

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("non query", func(t *testing.T) {
		var w writer
		defer w.AssertExpectations(t)

		Handle(&w, &dns.Msg{
			MsgHdr: dns.MsgHdr{
				Opcode: dns.OpcodeNotify,
			},
		})
	})

	t.Run("IPv4", func(t *testing.T) {
		var w writer
		w.On("WriteMsg", mock.AnythingOfType("*dns.Msg")).Return(nil).Once()
		defer w.AssertExpectations(t)

		Handle(&w, &dns.Msg{
			MsgHdr: dns.MsgHdr{
				Opcode: dns.OpcodeQuery,
			},
			Question: []dns.Question{
				{
					Name:  "foo.localhost",
					Qtype: dns.TypeA,
				},
			},
		})

		assert.Len(t, w.written.Answer, 1)
		assert.Equal(t, "foo.localhost	0	IN	A	127.0.0.1", w.written.Answer[0].String())
	})

	t.Run("IPv6", func(t *testing.T) {
		var w writer
		w.On("WriteMsg", mock.AnythingOfType("*dns.Msg")).Return(nil).Once()
		defer w.AssertExpectations(t)

		Handle(&w, &dns.Msg{
			MsgHdr: dns.MsgHdr{
				Opcode: dns.OpcodeQuery,
			},
			Question: []dns.Question{
				{
					Name:  "foo.localhost",
					Qtype: dns.TypeAAAA,
				},
			},
		})

		assert.Len(t, w.written.Answer, 1)
		assert.Equal(t, "foo.localhost	0	IN	AAAA	::1", w.written.Answer[0].String())
	})
}

type writer struct {
	mock.Mock
	written *dns.Msg
}

func (w *writer) LocalAddr() net.Addr {
	args := w.Called()
	return args.Get(0).(net.Addr)
}

func (w *writer) RemoteAddr() net.Addr {
	args := w.Called()
	return args.Get(0).(net.Addr)
}

func (w *writer) WriteMsg(m *dns.Msg) error {
	args := w.Called(m)
	w.written = m
	return args.Error(0)
}

func (w *writer) Write(b []byte) (int, error) {
	args := w.Called(b)
	return args.Int(0), args.Error(1)
}

func (w *writer) Close() error {
	args := w.Called()
	return args.Error(0)
}

func (w *writer) TsigStatus() error {
	args := w.Called()
	return args.Error(0)
}

func (w *writer) TsigTimersOnly(b bool) {
	w.Called(b)
}

func (w *writer) Hijack() {
	w.Called()
}
