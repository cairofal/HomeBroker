package entity

type Asset struct {
	ID           string
	Name         string
	Marketvolume int
}

func NewAsset(id string, name string, marketvolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		Marketvolume: marketvolume,
	}
}
