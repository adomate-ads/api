package models

import "time"

type BiddingStrategy struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Strategy  string    `json:"strategy" gorm:"type:varchar(128)" example:"Cost Optimize"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetBiddingStrategies() ([]BiddingStrategy, error) {
	var biddingStrategies []BiddingStrategy
	if err := DB.Find(&biddingStrategies).Error; err != nil {
		return nil, err
	}
	return biddingStrategies, nil
}

func GetBiddingStrategy(id uint) (*BiddingStrategy, error) {
	var biddingStrategy BiddingStrategy
	if err := DB.First(&biddingStrategy, id).Error; err != nil {
		return nil, err
	}
	return &biddingStrategy, nil
}

func GetBiddingStrategyByName(name string) (*BiddingStrategy, error) {
	var biddingStrategy BiddingStrategy
	if err := DB.Where("strategy = ?", name).First(&biddingStrategy).Error; err != nil {
		return nil, err
	}
	return &biddingStrategy, nil
}

func (bs *BiddingStrategy) CreateBiddingStrategy() error {
	err := DB.Create(&bs).Error
	if err != nil {
		return err
	}
	return nil
}

func (bs *BiddingStrategy) UpdateBiddingStrategy() (*BiddingStrategy, error) {
	err := DB.Save(&bs).Error
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (bs *BiddingStrategy) DeleteBiddingStrategy() error {
	err := DB.Delete(&bs).Error
	if err != nil {
		return err
	}
	return nil
}
