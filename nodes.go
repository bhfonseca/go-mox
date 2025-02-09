package go_mox

import (
	"encoding/json"
	"fmt"
	"github.com/bhfonseca/go-mox/models"
	"github.com/bhfonseca/go-mox/utils"
)

func GetAllPVENodes(PVEAddress, PVEAuth string) ([]string, error) {
	response, err := utils.PVEAPIRequester(PVEAddress, "api2/json/nodes", PVEAuth, "GET", "")
	if err != nil {
		return nil, err
	}
	var nodesData models.PVENodes
	err = json.Unmarshal([]byte(response), &nodesData)
	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar JSON: %v", err)
	}

	var nodes []string
	for _, node := range nodesData.Data {
		nodes = append(nodes, node.Node)
	}
	return nodes, nil
}
