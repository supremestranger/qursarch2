package model

import (
	"backend/components"
	"backend/utils"
	"encoding/json"
	"net/http"
)

const COMPONENT_ROOT = "/components"

func RegisterComponentModels() {
	utils.RegisterOnGet(COMPONENT_ROOT, onComponentGet)
}

func onComponentGet(rw http.ResponseWriter, req *http.Request) {
	utils.EnableCors(rw)
	components, err := components.GetComponents()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(components)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write(json)
}
