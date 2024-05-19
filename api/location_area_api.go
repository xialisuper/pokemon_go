package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (l LocationAreaResponse) isEqual(other LocationAreaResponse) bool {

	if l.Count != other.Count {
		return false
	}

	if &(l.Next) != &(other.Next) {
		return false
	}
	if &(l.Previous) != &(other.Previous) {
		return false
	}

	if len(l.Results) != len(other.Results) {
		return false
	}

	for i := 0; i < len(l.Results); i++ {
		if l.Results[i].Name != other.Results[i].Name {
			return false
		}
		if l.Results[i].Url != other.Results[i].Url {
			return false
		}
	}

	return true
}

// EncounterMethod represents the encounter method details
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Version represents the version details
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VersionDetail represents the details for a specific version
type VersionDetail struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

// EncounterMethodRate represents the encounter method rate details
type EncounterMethodRate struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

// Location represents the location details
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Language represents the language details
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Name represents the name details for different languages
type Name struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

// EncounterDetail represents the details of a specific encounter
type EncounterDetail struct {
	MinLevel        int             `json:"min_level"`
	MaxLevel        int             `json:"max_level"`
	ConditionValues []interface{}   `json:"condition_values"`
	Chance          int             `json:"chance"`
	Method          EncounterMethod `json:"method"`
}

// Pokemon represents the Pokemon details
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonVersionDetail represents the version-specific details for a Pokemon encounter
type PokemonVersionDetail struct {
	Version          Version           `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

// PokemonEncounter represents the encounter details for a specific Pokemon
type PokemonEncounter struct {
	Pokemon        PokemonDetail          `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

// LocationAreaEndpointResponse represents the main data structure
type LocationAreaEndpointResponse struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             Location              `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

func createRequestWithUrl(requestUrl *string) (*http.Request, error) {
	if requestUrl == nil {
		return http.NewRequest(
			http.MethodGet,
			"https://pokeapi.co/api/v2/location-area",
			nil,
		)
	}
	req, err := http.NewRequest(http.MethodGet, *requestUrl, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) GetLocationAreas(requestUrl *string) (res LocationAreaResponse, err error) {
	// make a  get request to the api and return the location areas
	// https://pokeapi.co/api/v2/location-area
	url := baseURL + "/location-area"
	if requestUrl != nil {
		url = *requestUrl
	}
	// check if the url is in cache

	cachedResponse, ok := c.cache.Get(url)

	// if it is in cache, return the cached response
	if ok {
		fmt.Println(" ------------------- cache hit ------------------")
		response := LocationAreaResponse{}
		err := json.Unmarshal(cachedResponse, &response)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return response, nil

	}
	fmt.Println("------------------- cache miss, making request to api... ------------------")
	// if it is not in cache, make a request to the api and cache the response

	req, err := createRequestWithUrl(requestUrl)

	if err != nil {

		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()
	// 读取响应的主体数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return LocationAreaResponse{}, err
	}

	// 解析json数据
	response := LocationAreaResponse{}
	if err := json.Unmarshal(body, &response); err != nil {

		return LocationAreaResponse{}, err

	}

	// 缓存响应数据
	c.cache.Add(url, body)

	return response, nil
}

func (c *Client) ExploreLocationAreaByName(name string) (res LocationAreaEndpointResponse, err error) {

	url := baseURL + "/location-area/" + name

	// check if the url is in cache

	cachedResponse, ok := c.cache.Get(url)

	// if it is in cache, return the cached response
	if ok {
		fmt.Println(" ------------------- cache hit ------------------")
		response := LocationAreaEndpointResponse{}
		err := json.Unmarshal(cachedResponse, &response)
		if err != nil {
			return LocationAreaEndpointResponse{}, err
		}
		return response, nil
	}
	fmt.Println("------------------- cache miss, making request to api... ------------------")
	// if it is not in cache, make a request to the api and cache the response

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return LocationAreaEndpointResponse{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaEndpointResponse{}, err
	}
	defer resp.Body.Close()
	// 读取响应的主体数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaEndpointResponse{}, err
	}

	// 解析json数据
	response := LocationAreaEndpointResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return LocationAreaEndpointResponse{}, err
	}

	// 缓存响应数据
	c.cache.Add(url, body)

	return response, nil

}
