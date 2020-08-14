package llbl

import (
	"net"

	"github.com/miekg/dns"
)

var ipv4loopback = net.IPv4(127, 0, 0, 1)

// Handle handles DNS query for *.localhost.
func Handle(w dns.ResponseWriter, r *dns.Msg) {
	if r.Opcode != dns.OpcodeQuery {
		return
	}

	m := dns.Msg{
		Compress: true,
	}
	m.SetReply(r)

	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			m.Answer = append(m.Answer, &dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
				},
				A: ipv4loopback,
			})
		case dns.TypeAAAA:
			m.Answer = append(m.Answer, &dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
				},
				AAAA: net.IPv6loopback,
			})
		}
	}

	_ = w.WriteMsg(&m)
}
