package gomox

import (
	"encoding/json"
	"fmt"
	"github.com/bhfonseca/go-mox/models"
	"github.com/bhfonseca/go-mox/utils"
	"log"
)

func GetAllKMVMachines(PVEAddress, PVEAuth, node string, humanize bool) ([]models.KVMMachine, error) {
	var machines []models.KVMMachine

	response, err := utils.PVEAPIRequester(PVEAddress, fmt.Sprintf("api2/json/nodes/%s/qemu", node), PVEAuth, "GET", "")
	if err != nil {
		log.Printf("Erro ao consultar os KVM no nó %s: %v\n", node, err)
	}

	var kvmResponse models.KVMResponse
	if err := json.Unmarshal([]byte(response), &kvmResponse); err != nil {
		log.Printf("(GetAllKMVMachines) Erro ao decodificar JSON do nó %s: %v\n", node, err)
	}

	for _, vm := range kvmResponse.Data {
		if humanize {
			vm.Uptime = utils.SecondsToDays(vm.Uptime)
			vm.MaxMem = utils.BytesToGigabytes(vm.MaxMem)
			vm.Mem = utils.BytesToGigabytes(vm.Mem)
			vm.MaxDisk = utils.BytesToGigabytes(vm.MaxDisk)
			vm.Disk = utils.BytesToGigabytes(vm.Disk)

		}
		machines = append(machines, models.KVMMachine{
			Name:      vm.Name,
			PID:       vm.PID,
			VMID:      vm.VMID,
			Uptime:    vm.Uptime,
			Cpus:      vm.Cpus,
			MaxMem:    vm.MaxMem,
			Mem:       vm.Mem,
			MaxDisk:   vm.MaxDisk,
			Disk:      vm.Disk,
			DiskWrite: vm.DiskWrite,
			DiskRead:  vm.DiskRead,
		})
	}

	return machines, nil
}
