package ens

type EnsName struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type EnsNameList struct {
	EnsNames []EnsName
}
