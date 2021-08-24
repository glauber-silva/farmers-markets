package markets

/*
TODO: implement search by Distrito, Regiao5, NomeFeira e Bairro
*/

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/url"
	"time"
)

// Service - Struct for farmers market service
type Service struct {
	DB *gorm.DB
}

// Market - Define market structure
type Market struct {
	ID         uint `gorm:"primaryKey;autoIncrement" ,csv:"ID"`
	Long       int64 `csv:"LONG"`
	Lat        int64 `csv:"LAT"`
	Setcens    string `csv:"SETCENS"`
	Areap      string `csv:"AREAP"`
	Coddist    int `csv:"CODDIST"`
	Distrito   string `csv:"DISTRITO"`
	Codsubpref int `csv:"CODSUBPREF"`
	Subprefe   string `csv:"SUBPREFE"`
	Regiao5    string `csv:"REGIAO5"`
	Regiao8    string `csv:"REGIAO8"`
	NomeFeira  string `csv:"NOME_FEIRA"`
	Registro   string `csv:"REGISTRO"`
	Logradouro string `csv:"LOGRADOURO"`
	Numero     string `csv:"NUMERO"`
	Bairro     string `csv:"BAIRRO"`
	Referencia string `csv:"REFERENCIA"`
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
	SearchMarkets(opts struct{}) ([]Market, error)
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
	market, err := s.GetMarket(ID)
	if err != nil {
		return Market{}, err
	}

	if market.Registro == newMarket.Registro{
		if r := s.DB.Model(&market).Updates(newMarket); r.Error != nil {
			return Market{}, r.Error
		}
	} else {
		return Market{}, errors.New("Registro must not be updated")
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
	var markets []Market
	if r := s.DB.Find(&markets); r.Error != nil {
		return markets, r.Error
	}
	return markets, nil

}

// SearchMarkets - retrieves all marktes from database with the information provided
func (s *Service) SearchMarkets(q url.Values) ([]Market, error){
	var markets []Market
	if r := s.DB.Where("distrito LIKE upper(?)", q.Get("distrito")).
		Or("regiao5 LIKE INITCAP(?)", q.Get("regiao5")).
		Or("nome_feira LIKE upper(?)", q.Get("nome_feira")).
		Or("bairro LIKE upper(?)", q.Get("bairro")).Find(&markets); r.Error != nil {
			return markets, r.Error
	}
	return markets, nil
}