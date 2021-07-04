package configbackend

type grpcOptions struct {
}

func GetDefGrpcOptions(name string) map[string]interface{} {

	g := make(map[string]interface{})

	g["ssl-target-name-override"] = name
	g["keep-alive-time"] = "0s"
	g["keep-alive-timeout"] = "20s"
	g["keep-alive-permit"] = false
	g["fail-fast"] = false
	g["allow-insecure"] = false
	return g
}

func SetTlsCaCertPath(m *map[string]interface{}, path string) {
	mi := *m

	tls := make(map[string]interface{})
	tls["path"] = path

	mi["tlsCACerts"] = tls
}
