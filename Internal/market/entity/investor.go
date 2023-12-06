package entity

type Investor struct {
	ID                    string
	Name                  string
	InvestorAssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:                    id,
		InvestorAssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assassetPosition)
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}
