package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// Ability represents the ability details
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// AbilityDetail represents the details for a specific ability
type AbilityDetail struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}

// Form represents the form details
type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Version represents the version details

// GameIndex represents the game index details
type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

// Item represents the item details
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VersionDetail represents the version-specific details for an item
type ItemVersionDetail struct {
	Rarity  int     `json:"rarity"`
	Version Version `json:"version"`
}

// HeldItem represents the held item details
type HeldItem struct {
	Item           Item                `json:"item"`
	VersionDetails []ItemVersionDetail `json:"version_details"`
}

// Move represents the move details
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// MoveLearnMethod represents the move learn method details
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VersionGroup represents the version group details
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VersionGroupDetail represents the version group-specific details for a move
type VersionGroupDetail struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	VersionGroup    VersionGroup    `json:"version_group"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
}

// MoveDetail represents the details for a specific move
type MoveDetail struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

// PokemonSpecies represents the species details
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Sprite represents the sprite details
type Sprite struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// OtherSprites represents the other sprite details
type OtherSprites struct {
	DreamWorld struct {
		FrontDefault string  `json:"front_default"`
		FrontFemale  *string `json:"front_female"`
	} `json:"dream_world"`
	Home struct {
		FrontDefault     string  `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
	} `json:"home"`
	OfficialArtwork struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"official-artwork"`
	Showdown struct {
		BackDefault      string  `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        string  `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     string  `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
	} `json:"showdown"`
}

// VersionSprites represents the version-specific sprite details
type VersionSprites struct {
	RedBlue struct {
		BackDefault  string `json:"back_default"`
		BackGray     string `json:"back_gray"`
		FrontDefault string `json:"front_default"`
		FrontGray    string `json:"front_gray"`
	} `json:"red-blue"`
	Yellow struct {
		BackDefault  string `json:"back_default"`
		BackGray     string `json:"back_gray"`
		FrontDefault string `json:"front_default"`
		FrontGray    string `json:"front_gray"`
	} `json:"yellow"`
	// Add more generations as needed
}

// Versions represents the generation-specific sprite details
type Versions struct {
	GenerationI    VersionSprites `json:"generation-i"`
	GenerationII   VersionSprites `json:"generation-ii"`
	GenerationIII  VersionSprites `json:"generation-iii"`
	GenerationIV   VersionSprites `json:"generation-iv"`
	GenerationV    VersionSprites `json:"generation-v"`
	GenerationVI   VersionSprites `json:"generation-vi"`
	GenerationVII  VersionSprites `json:"generation-vii"`
	GenerationVIII VersionSprites `json:"generation-viii"`
}

// Sprites represents the sprite details
type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       *string      `json:"back_female"`
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  *string      `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      *string      `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale *string      `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	Versions         Versions     `json:"versions"`
}

// Stat represents the stat details
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// StatDetail represents the details for a specific stat
type StatDetail struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

// Type represents the type details
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// TypeDetail represents the details for a specific type
type TypeDetail struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// Generation represents the generation details
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PastTypeDetail represents the past type details for a specific generation
type PastTypeDetail struct {
	Generation Generation   `json:"generation"`
	Types      []TypeDetail `json:"types"`
}

// PokemonDetail represents the main data structure
type PokemonDetail struct {
	ID                     int             `json:"id"`
	Name                   string          `json:"name"`
	BaseExperience         int             `json:"base_experience"`
	Height                 int             `json:"height"`
	IsDefault              bool            `json:"is_default"`
	Order                  int             `json:"order"`
	Weight                 int             `json:"weight"`
	Abilities              []AbilityDetail `json:"abilities"`
	Forms                  []Form          `json:"forms"`
	GameIndices            []GameIndex     `json:"game_indices"`
	HeldItems              []HeldItem      `json:"held_items"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []MoveDetail    `json:"moves"`
	Species                PokemonSpecies  `json:"species"`
	Sprites                Sprites         `json:"sprites"`
	Cries                  struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats     []StatDetail     `json:"stats"`
	Types     []TypeDetail     `json:"types"`
	PastTypes []PastTypeDetail `json:"past_types"`
}

// CatchPokemon gets the details of a specific pokemon
func (c *Client) CatchPokemon(name string) (pokemon PokemonDetail, err error) {
	// https://pokeapi.co/api/v2/pokemon/{id or name}/
	url := baseURL + "pokemon/" + name

	cachedResponse, ok := c.cache.Get(url)
	if ok {
		fmt.Println("you have met this pokemon before using cache", name)
		err = json.Unmarshal(cachedResponse, &pokemon)
		return pokemon, err
	}

	fmt.Println("you have never met this pokemon before", name)

	req, err := createRequestWithUrl(&url)

	if err != nil {
		return pokemon, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return pokemon, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemon, err
	}

	c.cache.Add(url, body)

	return pokemon, err

}
