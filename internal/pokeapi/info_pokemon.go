package pokeapi

type Pokemon struct {
	ID                     int64         `json:"id"`
	Name                   string        `json:"name"`
	BaseExperience         int64         `json:"base_experience"`
	Height 				   int64		 `json:"height"`
	Weight 				   int64		 `json:"weight"`
	Stats				   []Stat		 `json:"stats"`
	Types				   []Type		 `json:"types"`
}

type Type struct {
	Slot 		int64 		`json:"slot"`
	Type 		Name 	    `json:"type"`
}

type Stat struct {
	BaseStat 		int64 	`json:"base_stat"`
	Stat 			Name 	`json:"stat"`
}

type Name struct {
	Name 		string 		`json:"name"`
	URL 		string 		`json:"url"`
}