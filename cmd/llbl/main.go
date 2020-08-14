package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ichiban/llbl"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

func main() {
	dns.HandleFunc("localhost.", llbl.Handle)

	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to listen")
		return
	}

	log.WithFields(log.Fields{
		"addr": conn.LocalAddr(),
	}).Info("listen")

	s := dns.Server{
		PacketConn: conn,
	}

	addr := conn.LocalAddr().(*net.UDPAddr)
	restore, err := llbl.Configure(addr.Port)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to configure")
		return
	}
	defer func() {
		if err := restore(); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to restore configuration")
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ActivateAndServe(); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to serve")
			return
		}
	}()

	<-sigs

	if err := s.Shutdown(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to shutdown")
		return
	}
}
