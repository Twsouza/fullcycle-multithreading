package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const (
	BASE_URL = "http://viacep.com.br/ws/%s/json/"
)

func (b *Viacep) Get(ctx context.Context, cep string) (string, error) {
	url := fmt.Sprintf(BASE_URL, cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("Viacep.Get: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Viacep.Get: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Viacep.Get: response status code is not OK: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(b)
	if err != nil {
		return "", fmt.Errorf("Viacep.Get: %w", err)
	}

	return fmt.Sprintf("ViaCEP: %s, %s, %s, %s, %s", b.Logradouro, b.Complemento, b.Bairro, b.Localidade, b.Uf), nil
}
