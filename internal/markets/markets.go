package markets

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
