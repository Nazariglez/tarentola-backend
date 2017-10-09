// Created by nazarigonzalez on 9/10/17.

package rolemodel

func GetList() []Role {
	return currentRoles
}

func GetID(name string) uint {
	var id uint = 0

	for _, r := range currentRoles {
		if name == r.Name {
			id = r.ID
			break
		}
	}

	return id
}
