package pokeapi

type Pokemon struct {
	ID                     int64         `json:"id"`
	Name                   string        `json:"name"`
	BaseExperience         int64         `json:"base_experience"`        
}