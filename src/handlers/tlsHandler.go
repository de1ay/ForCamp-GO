package handlers

import (
	"forcamp/conf"
	"crypto/tls"
	"github.com/gorilla/mux"
	"net/http"
)

// Read TLS certificates to TlsConfig
func getTlsConfig() *tls.Config {
	TlsConfig := &tls.Config{}
	TlsConfig.Certificates = make([]tls.Certificate, 3)
	TlsConfig.Certificates[0], _ = tls.LoadX509KeyPair(conf.API_SITE_TLS, conf.API_SITE_TLS_KEY)
	TlsConfig.Certificates[1], _ = tls.LoadX509KeyPair(conf.MAIN_SITE_TLS, conf.MAIN_SITE_TLS_KEY)
	TlsConfig.Certificates[2], _ = tls.LoadX509KeyPair(conf.WWW_MAIN_SITE_TLS, conf.WWW_MAIN_SITE_TLS_KEY)
	TlsConfig.BuildNameToCertificate()
	return TlsConfig
}

func HandleTLS(router *mux.Router) {
	TlsConfig := getTlsConfig()
	Server := http.Server{
		TLSConfig: TlsConfig,
		Handler: router,
	}
	TLSListener, _ := tls.Listen("tcp", conf.TLS_PORT, TlsConfig)
	Server.Serve(TLSListener)
}