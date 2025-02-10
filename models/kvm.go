package models

import "encoding/json"

type KVMResponse struct {
	Data []KVMMachine `json:"data"`
}

type KVMMachine struct {
	VMID      json.Number `json:"vmid"`
	Name      string      `json:"name"`
	Status    string      `json:"status"`
	Uptime    int         `json:"uptime"`
	PID       int         `json:"pid"`
	Cpus      json.Number `json:"cpus"`
	MaxMem    json.Number `json:"maxmem"`
	Mem       json.Number `json:"mem"`
	MaxDisk   json.Number `json:"maxdisk"`
	Disk      json.Number `json:"disk"`
	DiskWrite int         `json:"diskwrite"`
	DiskRead  int         `json:"diskread"`
}
