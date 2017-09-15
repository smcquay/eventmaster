package eventmaster

// Configuration struct for the eventmaster
type Config struct {
	Port int `long:"port" default:"50052" description:"Port for EventMaster gRPC + HTTP API"` // What port for the API to listen on

	CassandraServiceName string `long:"cassandra_servicename" description:"name of cassandra service to talk to"`

	CassandraPort string `long:"cassandra_port" default:"9201" description:"port of cassandra service"`

	RsyslogServer bool `short:"r" long:"rsyslog_server" description:"Flag to start TCP rsyslog server"`

	RsyslogPort int `long:"rsyslog_port" default:"50053" description:"Port for rsyslog clients to send logs to"`

	PromPort int `long:"prom_port" default:"9000" description:"Port for Prometheus client"`

	CAFile string `long:"ca_file" description:"PEM encoded CA's certificate file path"`

	CertFile string `long:"cert_file" description:"PEM encoded certificate file path"`

	KeyFile string `long:"key_file" description:"PEM encoded private key file path"`

	ConfigFile string `short:"c" long:"config" description:"location of configuration file"`

	StaticFiles string `short:"s" long:"static" description:"location of static files to use (instead of embedded files)"`

	Templates string `short:"t" long:"templates" description:"location of template files to use (instead of embedded)"`
}
