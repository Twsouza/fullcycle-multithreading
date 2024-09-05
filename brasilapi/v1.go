package brasilapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

const (
	BASE_URL = "https://brasilapi.com.br/api/cep/v1/%s"
)

func (b *BrasilAPI) Get(ctx context.Context, cep string) (string, error) {
	url := fmt.Sprintf(BASE_URL, cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("BrasilAPI.Get: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("BrasilAPI.Get: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("BrasilAPI.Get: response status code is not OK: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(b)
	if err != nil {
		return "", fmt.Errorf("BrasilAPI.Get: %w", err)
	}

	return fmt.Sprintf("BrasilAPI: %s, %s, %s, %s, %s", b.Street, b.Neighborhood, b.City, b.State, b.Cep), nil
}
