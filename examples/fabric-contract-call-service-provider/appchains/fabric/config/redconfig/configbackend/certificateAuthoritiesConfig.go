package configbackend

type CertificateAuthoritiesConfig struct {
	CAOrgName      string
	CAUrl          string
	CAName         string
	EnrollId       string
	EnrollSecret   string
	TlsCACertsPath string
}

func (o *CertificateAuthoritiesConfig) GetCertificateAuthoritiesConfig() map[string]interface{} {

	p := make(map[string]interface{})

	p["url"] = o.CAUrl

	ho := make(map[string]interface{})
	ho["verify"] = true

	p["httpOptions"] = ho

	SetTlsCaCertPath(&p, o.TlsCACertsPath)

	return p
}

func (o *CertificateAuthoritiesConfig) SetCertificateAuthoritiesConfig(m *map[string]interface{}) {
	mi := *m
	mi[o.CAOrgName] = o.GetCertificateAuthoritiesConfig()
}
