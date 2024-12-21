package model

import (
	"backend/indications"
	"backend/utils"
	"encoding/json"
	"net/http"
)

const INDICATION_ROOT = "/indications"

func RegisterIndicationModels() {
	utils.RegisterOnGet(INDICATION_ROOT, onIndicationGet)
}

func onIndicationGet(rw http.ResponseWriter, req *http.Request) {
	utils.EnableCors(rw)
	indications, err := indications.GetIndications()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(indications)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write(json)
}
