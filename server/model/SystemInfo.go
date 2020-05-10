package model

type SystemInfo struct {
	Arch        string `json:"arch"`
	Os          string `json:"os"`
	Cpu         string `json:"cpu"`
	Environment string `json:"environment"`
	Database    string `json:"database"`
	Memory      string `json:"memory"`
}
