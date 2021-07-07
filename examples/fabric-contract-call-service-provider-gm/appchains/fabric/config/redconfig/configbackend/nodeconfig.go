package configbackend

type configBackend interface {
	GetNodeName() string

	SetConfig(m *map[string]interface{})
}

type configBase struct {
	NodeName string
}

func (c *configBase) GetNodeName() string {
	return c.NodeName
}
