package configbackend

const (
	_MSPID                  = "mspid"
	_CryptoPath             = "cryptoPath"
	_Peers                  = "peers"
	_CertificateAuthorities = "certificateAuthorities"
)

type OrganizationConfig struct {
	OrgName                string
	MspId                  string
	CryptoPath             string
	Peers                  []string
	CertificateAuthorities []string
}

func (o *OrganizationConfig) GetOrganizationConfig() map[string]interface{} {

	m := make(map[string]interface{})
	m[_MSPID] = o.MspId
	m[_CryptoPath] = o.CryptoPath
	m[_Peers] = o.Peers
	m[_CertificateAuthorities] = o.CertificateAuthorities

	return m
}

func (o *OrganizationConfig) SetOrganizationConfig(m *map[string]interface{}) {
	mi := *m
	mi[o.OrgName] = o.GetOrganizationConfig()
}
