package models

import "encoding/json"

type PVENodes struct {
	Data []PVENode `json:"data"`
}

type PVENode struct {
	Node      string      `json:"node"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	Uptime    int         `json:"uptime"`
	MaxDisk   int         `json:"maxdisk"`
	Disk      int         `json:"disk"`
	MaxCPU    json.Number `json:"maxcpu"`
	CPU       json.Number `json:"cpu"`
	MaxMemory int         `json:"maxmem"`
	Memory    int         `json:"mem"`
}
