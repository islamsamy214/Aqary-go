package models

import "web-app/providers"

type Profile struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

func (p *Profile) GetProfile() (Profile, error) {
	db := (&providers.Sql{}).Init()

	var profile Profile
	err := db.QueryRow("SELECT id, first_name, last_name, address FROM profiles WHERE id = $1", p.ID).Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Address)
	if err != nil {
		return profile, err
	}

	return profile, nil
}
