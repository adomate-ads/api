package models

type PreRegistration struct {
	ID        uint             `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Domain    string           `json:"domain" gorm:"type:varchar(128)" example:"example.com"`
	Locations []PreRegLocation `json:"locations" gorm:"type:varchar(128)"`
	Services  []PreRegService  `json:"services" gorm:"type:varchar(128)"`
	Budget    uint             `json:"budget" gorm:"type:integer" example:"100"`
}

type PreRegLocation struct {
	ID                uint   `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	PreRegistrationID uint   `json:"pre_registration_id" gorm:"type:integer" example:"1"`
	Location          string `json:"location" gorm:"type:varchar(128)" example:"Houston, TX"`
}

type PreRegService struct {
	ID                uint   `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	PreRegistrationID uint   `json:"pre_registration_id" gorm:"type:integer" example:"1"`
	Service           string `json:"service" gorm:"type:varchar(128)" example:"Dental Services"`
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
	if err := DB.Preload("Locations").Preload("Services").Find(&preRegistrations).Error; err != nil {
		return nil, err
	}
	return preRegistrations, nil
}

func GetPreRegistration(id uint) (*PreRegistration, error) {
	var preRegistration PreRegistration
	err := DB.Where("id = ?", id).Preload("Locations").Preload("Services").First(&preRegistration).Error
	if err != nil {
		return nil, err
	}
	return &preRegistration, nil
}

func GetPreRegistrationByDomain(domain string) (*PreRegistration, error) {
	var preRegistration PreRegistration
	err := DB.Where("domain = ?", domain).Preload("Locations").Preload("Services").First(&preRegistration).Error
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

func (loc *PreRegLocation) CreatePreRegLocation() error {
	err := DB.Create(&loc).Error
	if err != nil {
		return err
	}
	return nil
}

func (loc *PreRegLocation) DeletePreRegLocation() error {
	err := DB.Delete(&loc).Error
	if err != nil {
		return err
	}
	return nil
}

func (svc *PreRegService) CreatePreRegService() error {
	err := DB.Create(&svc).Error
	if err != nil {
		return err
	}
	return nil
}

func (svc *PreRegService) DeletePreRegService() error {
	err := DB.Delete(&svc).Error
	if err != nil {
		return err
	}
	return nil
}
