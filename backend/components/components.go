package components

import "backend/db"

func GetComponents() ([]string, error) {
	rows, err := db.DB.Query("SELECT Name FROM Components")
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		res = append(res, name)
	}

	return res, nil
}
