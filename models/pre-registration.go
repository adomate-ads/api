package models

type PreRegistration struct {
	ID       uint     `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Domain   string   `json:"domain" gorm:"type:varchar(128)" example:"example.com"`
	Location []string `json:"location" gorm:"type:varchar(128)" example:"Houston, TX"`
	Service  []string `json:"service" gorm:"type:varchar(128)" example:"Dental Services"`
	Budget   uint     `json:"budget" gorm:"type:integer" example:"100"`
}

func (pr *PreRegistration) CreatePreRegistration() error {
	err := DB.Create(&pr).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPreRegistrations() ([]PreRegistration, error) {
	var preRegistrations []PreRegistration
	if err := DB.Find(&preRegistrations).Error; err != nil {
		return nil, err
	}
	return preRegistrations, nil
}

func GetPreRegistration(id uint) (*PreRegistration, error) {
	var preRegistration PreRegistration
	err := DB.Where("id = ?", id).First(&preRegistration).Error
	if err != nil {
		return nil, err
	}
	return &preRegistration, nil
}

func GetPreRegistrationByDomain(domain string) (*PreRegistration, error) {
	var preRegistration PreRegistration
	err := DB.Where("domain = ?", domain).First(&preRegistration).Error
	if err != nil {
		return nil, err
	}
	return &preRegistration, nil
}

func (pr *PreRegistration) UpdatePreRegistration() error {
	err := DB.Save(&pr).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *PreRegistration) DeletePreRegistration() error {
	err := DB.Delete(&pr).Error
	if err != nil {
		return err
	}
	return nil
}

func AddLocationToPreRegistration(pr *PreRegistration, location string) error {
	pr.Location = append(pr.Location, location)
	err := pr.UpdatePreRegistration()
	if err != nil {
		return err
	}
	return nil
}

func RemoveLocationFromPreRegistration(pr *PreRegistration, location string) error {
	for i, l := range pr.Location {
		if l == location {
			pr.Location = append(pr.Location[:i], pr.Location[i+1:]...)
		}
	}
	err := pr.UpdatePreRegistration()
	if err != nil {
		return err
	}
	return nil
}

func AddServiceToPreRegistration(pr *PreRegistration, service string) error {
	pr.Service = append(pr.Service, service)
	err := pr.UpdatePreRegistration()
	if err != nil {
		return err
	}
	return nil
}

func RemoveServiceFromPreRegistration(pr *PreRegistration, service string) error {
	for i, s := range pr.Service {
		if s == service {
			pr.Service = append(pr.Service[:i], pr.Service[i+1:]...)
		}
	}
	err := pr.UpdatePreRegistration()
	if err != nil {
		return err
	}
	return nil
}
