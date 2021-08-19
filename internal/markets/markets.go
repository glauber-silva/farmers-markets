package markets

/*
TODO: implement a Search resource that uses at least one of the fields: Distrito, Regiao5, NomeFeira e Bairro
*/

import "github.com/jinzhu/gorm"

// Service - Struct for farmers market service
type Service struct {
	DB *gorm.DB
}

// Market - Define market structure
type Market struct {
	gorm.Model
	Long       int64
	Lat        int64
	Setcens    string
	Areap      string
	Coddist    int
	Distrito   string
	Codsubpref int
	Subprefe   string
	Regiao5    string
	Reigao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referência string
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
func (s *Service) GetMarket(ID int) (Market, error) {
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
func (s *Service) UpdateMarket(ID int, newMarket Market) (Market, error) {
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
func (s *Service) DeleteMarket(ID int64) error {
	if r := s.DB.Delete(&Market{}, ID); r.Error != nil {
		return r.Error
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
