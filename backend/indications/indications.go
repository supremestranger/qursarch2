package indications

import "backend/db"

func GetIndications() ([]string, error) {
	rows, err := db.DB.Query("SELECT Name FROM Indications")
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
