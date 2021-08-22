package markets

/*
TODO: implement search by Distrito, Regiao5, NomeFeira e Bairro
*/

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// Service - Struct for farmers market service
type Service struct {
	DB *gorm.DB
}

// Market - Define market structure
type Market struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Long       int64
	Lat        int64
	Setcens    string
	Areap      string
	Coddist    int
	Distrito   string
	Codsubpref int
	Subprefe   string
	Regiao5    string
	Regiao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// MarketService - An interface for the market service
type MarketService interface {
	GetMarket(ID int) (Market, error)
	GetAllMarkets() ([]Market, error)
	PostMarket(market Market) (Market, error)
	UpdateMarket(newMarket Market) (Market, error)
	DeleteMarket(ID int) error
}

// NewService - return a new market service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetMarket - retrieves market by ID from database
func (s *Service) GetMarket(ID uint) (Market, error) {
	var market Market
	if r := s.DB.First(&market, ID); r.Error != nil {
		return Market{}, r.Error
	}
	return market, nil
}

// PostMarket - insert a new market to the database
func (s *Service) PostMarket(market Market) (Market, error) {
	if r := s.DB.Save(&market); r.Error != nil {
		return Market{}, r.Error
	}
	return market, nil
}

// UpdateMarket - updates a market by ID with new market info
func (s *Service) UpdateMarket(ID uint, newMarket Market) (Market, error) {
	/*
		TODO: The field "Registro" must not be updated
	*/
	market, err := s.GetMarket(ID)
	if err != nil {
		return Market{}, err
	}

	if r := s.DB.Model(&market).Updates(newMarket); r.Error != nil {
		return Market{}, r.Error
	}
	return market, nil
}

// DeleteMarket - deletes a market by ID
func (s *Service) DeleteMarket(ID uint) error {

	if r := s.DB.Delete(&Market{}, ID); r.Error != nil {
		return r.Error
	} else if r.RowsAffected == 0 {
		return fmt.Errorf("row with id=%d cannot be deleted because it does not exist", ID)
	}
	return nil
}

// GetAllMarkets - retrieves all markets from database
func (s *Service) GetAllMarkets() ([]Market, error) {
	/*
		TODO: think about pagination
	*/
	var markets []Market
	if r := s.DB.Find(&markets); r.Error != nil {
		return markets, r.Error
	}
	return markets, nil

}
