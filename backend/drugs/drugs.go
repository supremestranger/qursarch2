package drugs

import (
	"backend/db"
	"strconv"
)

type PriceDesc struct {
	MedstoreName string
	Price        int
}

type DrugDesc struct {
	Id           int
	Name         string
	Description  string
	NeedsReceipt bool
	Prices       []PriceDesc
	Components   []string
	Indications  []string
}

type FullDrugInfo struct {
	Id           int
	Name         string
	Description  string
	Dosage       string
	UnitWeight   int
	Effects      string
	MinAge       int
	NeedsReceipt bool
	Prices       []PriceDesc
	Components   []string
	Indications  []string
}

func GetDrug(id string) (*FullDrugInfo, error) {
	query := `
		SELECT 
			d.Id AS DrugId,
			d.Name AS DrugName,
			d.Description AS DrugDescription,
			d.Dosage as DrugDosage,
			d.UnitWeight as DrugUnitWeight,
			d.Effects as DrugEffects,
			d.MinAge as DrugMinAge,
			d.NeedsReceipt,
			c.Name as ComponentName,
			i.Name as IndicationName,
			ms.Name AS MedstoreName,
			dm.Price
		FROM 
			DrugMap dm
		JOIN 
			Drugs d ON dm.DrugId = d.ID
		JOIN 
			Medstores ms ON dm.MedstoreId = ms.ID
		JOIN
    		DrugsComponents dc ON dc.DrugId = d.id
		JOIN
    		Components c ON c.id = dc.ComponentId
		JOIN
			DrugIndications di on di.DrugId = d.Id
		JOIN
			Indications i on i.Id = di.IndicationId
		WHERE d.Id = ($1)
	`

	var drugName, drugDescription, dosage, effects, componentName, indicationName, medstoreName string
	var drugId, minAge, unitWeight, price int
	var needsReceipt bool

	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	var fullDrugInfo FullDrugInfo
	var firstRow bool
	for rows.Next() {
		if err := rows.Scan(&drugId, &drugName, &drugDescription, &dosage, &unitWeight, &effects, &minAge, &needsReceipt, &componentName, &indicationName, &medstoreName, &price); err != nil {
			return nil, err
		}

		priceDesc := PriceDesc{
			MedstoreName: medstoreName,
			Price:        price,
		}
		prices := make([]PriceDesc, 0)
		components := make([]string, 0)
		indications := make([]string, 0)

		prices = append(prices, priceDesc)
		components = append(components, componentName)
		indications = append(indications, indicationName)

		if !firstRow {
			fullDrugInfo = FullDrugInfo{
				Id:           drugId,
				Name:         drugName,
				Description:  drugDescription,
				Dosage:       dosage,
				UnitWeight:   unitWeight,
				Effects:      effects,
				MinAge:       minAge,
				Prices:       prices,
				Components:   components,
				NeedsReceipt: needsReceipt,
				Indications:  indications,
			}
			firstRow = true
		} else {
			var updateComponents bool = true
			var updateIndications bool = true
			var updatePrices bool = true

			for _, v := range fullDrugInfo.Prices {
				if v.MedstoreName == priceDesc.MedstoreName {
					updatePrices = false
				}
			}
			for _, v := range fullDrugInfo.Components {
				if v == componentName {
					updateComponents = false
					break
				}
			}
			for _, v := range fullDrugInfo.Indications {
				if v == indicationName {
					updateIndications = false
					break
				}
			}
			if updatePrices {
				fullDrugInfo.Prices = append(fullDrugInfo.Prices, priceDesc)
			}
			if updateComponents {
				fullDrugInfo.Components = append(fullDrugInfo.Components, componentName)
			}
			if updateIndications {
				fullDrugInfo.Indications = append(fullDrugInfo.Indications, indicationName)
			}
		}

	}

	return &fullDrugInfo, nil
}

func GetDrugs(name string, needsReceiptFilter bool, minAgeStr string, components []string, indications []string) (map[string]DrugDesc, error) {
	var drugName, drugDescription, componentName, indicationName, medstoreName string
	var drugId, price int
	var needsReceipt bool
	var filterComponents string
	var filterIndications string
	if len(components) > 0 {
		filterComponents = components[0]
	}
	if len(indications) > 0 {
		filterIndications = indications[0]
	}
	_ = filterComponents
	_ = filterIndications
	query := `
		SELECT 
			d.Id AS DrugId,
			d.Name AS DrugName,
			d.Description AS DrugDescription,
			d.NeedsReceipt,
			c.Name as ComponentName,
			i.Name as IndicationName,
			ms.Name AS MedstoreName,
			dm.Price
		FROM 
			DrugMap dm
		JOIN 
			Drugs d ON dm.DrugId = d.ID
		JOIN 
			Medstores ms ON dm.MedstoreId = ms.ID
		JOIN
    		DrugsComponents dc ON dc.DrugId = d.id
		JOIN
    		Components c ON c.id = dc.ComponentId
		JOIN
			DrugIndications di on di.DrugId = d.Id
		JOIN
			Indications i on i.Id = di.IndicationId
		WHERE
		position(LOWER($1) in LOWER(d.Name)) > 0
		AND ($2 = FALSE OR d.NeedsReceipt = ($2))
		AND d.MinAge >= ($3)
		AND (COALESCE($4, '') = '' OR c.Name = ANY(string_to_array($4, ',')))
		AND (COALESCE($5, '') = '' OR i.Name = ANY(string_to_array($5, ',')))
		ORDER BY
    		d.id
	`
	minAge := 0
	if minAgeStr != "" {
		minAgeParsed, err := strconv.Atoi(minAgeStr)
		if err != nil {
			return nil, err
		}
		minAge = minAgeParsed
	}
	rows, err := db.DB.Query(query, name, needsReceiptFilter, minAge, filterComponents, filterIndications)
	if err != nil {
		return nil, err
	}
	drugs := make(map[string]DrugDesc)
	// rows.Scan(drugs)
	for rows.Next() {
		if err := rows.Scan(&drugId, &drugName, &drugDescription, &needsReceipt, &componentName, &indicationName, &medstoreName, &price); err != nil {
			return nil, err
		}

		priceDesc := PriceDesc{
			MedstoreName: medstoreName,
			Price:        price,
		}
		if val, ok := drugs[drugName]; !ok {
			prices := make([]PriceDesc, 0)
			components := make([]string, 0)
			indications := make([]string, 0)

			prices = append(prices, priceDesc)
			components = append(components, componentName)
			indications = append(indications, indicationName)
			drugs[drugName] = DrugDesc{
				Id:           drugId,
				Name:         drugName,
				Description:  drugDescription,
				Prices:       prices,
				Components:   components,
				NeedsReceipt: needsReceipt,
				Indications:  indications,
			}
		} else {
			var updateComponents bool = true
			var updateIndications bool = true
			var updatePrices bool = true

			for _, v := range val.Prices {
				if v.MedstoreName == priceDesc.MedstoreName {
					updatePrices = false
				}
			}
			for _, v := range val.Components {
				if v == componentName {
					updateComponents = false
					break
				}
			}
			for _, v := range val.Indications {
				if v == indicationName {
					updateIndications = false
					break
				}
			}
			if updatePrices {
				val.Prices = append(val.Prices, priceDesc)
			}
			if updateComponents {
				val.Components = append(val.Components, componentName)
			}
			if updateIndications {
				val.Indications = append(val.Indications, indicationName)
			}
			drugs[drugName] = val
		}
	}
	return drugs, nil
}
