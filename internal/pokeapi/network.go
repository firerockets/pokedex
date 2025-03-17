package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocation(url string) (Resource, error) {
	if url == "" {
		url = baseURL + "/location-area/"
	}

	data, err := c.handleRequest(url)
	if err != nil {
		return Resource{}, err
	}

	var resource Resource
	err = json.Unmarshal(data, &resource)

	if err != nil {
		return Resource{}, err
	}

	return resource, nil
}

func (c *Client) ExploreLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	data, err := c.handleRequest(url)
	if err != nil {
		return LocationArea{}, err
	}

	var location LocationArea
	err = json.Unmarshal(data, &location)

	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	data, err := c.handleRequest(url)

	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (c *Client) handleRequest(url string) ([]byte, error) {

	data, ok := c.cache.Get(url)

	if ok {
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("server responded with error: %w", err)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading from response: %w", err)
	}

	c.cache.Add(url, data)

	return data, nil
}
