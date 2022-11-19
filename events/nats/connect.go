package nats

import (
	"crypto/tls"
	"fmt"
	"github.com/halilylm/ticketing-common/logger"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Options struct {
	ClusterID string
	ClientID  string
	Servers   []string
	TLSConfig *tls.Config
	Logger    logger.Logger
}

func ConnectToStan(options Options, handler stan.ConnectionLostHandler) (stan.Conn, error) {
	natsOptions := nats.GetDefaultOptions()
	if options.TLSConfig != nil {
		natsOptions.Secure = true
		natsOptions.TLSConfig = options.TLSConfig
	}
	natsOptions.Servers = options.Servers
	conn, err := natsOptions.Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to nats at %v", options.Servers)
	}
	return stan.Connect(options.ClusterID, options.ClientID, stan.NatsConn(conn), stan.SetConnectionLostHandler(handler))
}
