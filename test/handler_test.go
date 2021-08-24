package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {

	client := resty.New()
	r, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, r.StatusCode())
}

func TestGetMarkets(t *testing.T) {
	client := resty.New()
	r, err := client.R().Get("http://localhost:8080/api/market")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, r.StatusCode())
}

func TestPostMarket(t *testing.T) {
	client := resty.New()
	r, err := client.R().
		SetBody(
			`{"Long": -46550164,
			"Lat": -23558733,
			"Setcens": "355030885000091",
			"Areap": "3550308005040",
			"Coddist": 87,
			"Distrito": "VILA FORMOSA",
			"Codsubpref": 26,
			"Subprefe": "ARICANDUVA-FORMOSA-CARRAO",
			"Regiao5": "Leste",
			"Regiao8": "Leste 1",
			"NomeFeira": "VILA FORMOSA",
			"Registro": "7777-6",
			"Logradouro": "RUA TESTE",
			"Numero": "S/N",
			"Bairro": "VL FORMOSA",
			"Referencia": "TRAVESSA MEIA-MEIA"
	}`).
		Post("http://localhost:8080/api/market")

	assert.NoError(t, err)
	assert.Equal(t, 200, r.StatusCode())
	// TODO: assert if is a json/dict
}
