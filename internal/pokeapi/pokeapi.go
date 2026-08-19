package pokeapi

type APIResourceList struct {
	Count 		int
	Next 		string
	Previous 	string
	Results 	[]NamedAPIResource
}

type NamedAPIResource struct {
	Name 	string
	Url 	string
}

type LocationArea struct {
	ID 						int
	Name 					string
	GameIndex 				string
	EncounterMethodRates 	[]EncounterMethodRate
	Location 				NamedAPIResource
	Names 					[]Name
	Pokemon_encounters 		[]PokemonEncounter
}

type Name struct {
	Name 		string
	Language 	NamedAPIResource
}

type Pokemon struct {
	Name string
	Base_experience int
	Height int
	Weight int
	Stats []PokemonStat
	Types []PokemonType
}

type PokemonStat struct {
	Stat NamedAPIResource
	Base_stat int
}

type PokemonType struct {
	Type NamedAPIResource
}

type EncounterMethodRate struct {
	encounterMethod 	NamedAPIResource
	versionDetails 		[]EncounterVersionDetails
}

type EncounterVersionDetails struct {
	rate 	int
	version NamedAPIResource
}

type PokemonEncounter struct {
	Pokemon 			NamedAPIResource
	versionDetails 	[]VersionEncounterDetail
}

type VersionEncounterDetail struct {
	version 			NamedAPIResource
	maxChance 			int
	encounterDetails 	[]Encounter
}

type Encounter struct {
	minLevel 			int
	maxLevel 			int
	conditionValues 	NamedAPIResource
	chance 				int
	method 				NamedAPIResource
}