package gomox

import (
	"encoding/json"
	"fmt"
	"github.com/bhfonseca/go-mox/models"
	"github.com/bhfonseca/go-mox/utils"
)

func GetAllPVENodes(PVEAddress, PVEAuth string) (models.PVENodes, error) {
	response, err := utils.PVEAPIRequester(PVEAddress, "api2/json/nodes", PVEAuth, "GET", "")
	if err != nil {
		return models.PVENodes{}, err
	}

	var nodesData models.PVENodes
	err = json.Unmarshal([]byte(response), &nodesData)
	if err != nil {
		return models.PVENodes{}, fmt.Errorf("Erro ao decodificar JSON: %v", err)
	}

	return nodesData, nil
}
