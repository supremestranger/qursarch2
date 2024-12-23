package model

import (
	"backend/drugs"
	"backend/utils"
	"encoding/json"
	"net/http"
)

const DRUG_ROOT = "/drugs"

func RegisterDrugModels() {
	utils.RegisterOnGet(DRUG_ROOT, OnDrugsGet)
	utils.RegisterOnGet(DRUG_ROOT+"/{id}", OnDrugGet)
}

func OnDrugGet(rw http.ResponseWriter, req *http.Request) {
	utils.EnableCors(rw)
	id := req.PathValue("id")
	fullDrugInfo, err := drugs.GetDrug(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(fullDrugInfo)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write(res)
}

func OnDrugsGet(rw http.ResponseWriter, req *http.Request) {
	utils.EnableCors(rw)
	name := req.URL.Query().Get("name")
	needsReceipt := req.URL.Query().Get("needsReceipt") == "true"
	minAgeStr := req.URL.Query().Get("minAge")
	components := req.URL.Query()["activeComponents"] // Multiple values
	indications := req.URL.Query()["indications"]     // Multiple values
	rawDrugs, err := drugs.GetDrugs(name, needsReceipt, minAgeStr, components, indications)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	v := make([]drugs.DrugDesc, 0, len(rawDrugs))

	for _, value := range rawDrugs {
		v = append(v, value)
	}

	res, err := json.Marshal(v)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write(res)
}
