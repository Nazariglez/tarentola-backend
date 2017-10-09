// Created by nazarigonzalez on 9/10/17.

package rolemodel

func GetList() ([]Role, error) {
	roles := []Role{}
	if err := db.Find(&roles).Error; err != nil {
		return []Role{}, err
	}

	return roles, nil
}
