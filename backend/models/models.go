package models

type SystemInfo struct {
	HostName   string `json:"hostname"`
	SystemOS   string `json:"systemos"`
	Processors string `json:"processors"`
	Memory     string `json:"memory"`
	HostIP     string `json:"hostip"`
	DiskUsed   string `json:"diskused"`
}
