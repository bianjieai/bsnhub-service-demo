package iservice

// AddKey adds a key with the specified name and passphrase
func (s ServiceClientWrapper) AddKey(name string, passphrase string) (addr string, mnemonic string, err error) {
	return s.IritaClient.Insert(name, passphrase)
}

// DeleteKey deletes the specified key
func (s ServiceClientWrapper) DeleteKey(name string, passphrase string) error {
	return s.IritaClient.Delete(name, passphrase)
}

// ShowKey queries the given key
func (s ServiceClientWrapper) ShowKey(name string, passphrase string) (addr string, err error) {
	_, address, err := s.IritaClient.Find(name, passphrase)
	return address.String(), err
}

// ImportKey imports the specified key
func (s ServiceClientWrapper) ImportKey(name string, passphrase string, keyArmor string) (addr string, err error) {
	return s.IritaClient.Import(name, passphrase, keyArmor)
}

// ExportKey exports the given key
func (s ServiceClientWrapper) ExportKey(name string, passphrase string) (keyArmor string, err error) {
	return s.IritaClient.Export(name, passphrase)
}

// RecoverKey recover the specified key
func (s ServiceClientWrapper) RecoverKey(name string, passphrase string, mnemonic string) (addr string, err error) {
	return s.IritaClient.Recover(name, passphrase, mnemonic)
}
