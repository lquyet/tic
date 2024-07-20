package entity

import (
	"fmt"
	"math/rand"
)

type Product struct {
	ID          string `json:"id" faker:"uuid_digit"`
	Category    string `json:"category" faker:"oneof:food,drink,snack"`
	Price       int    `json:"price" faker:"oneof:1000,2000,3000,4000,5000,6000,7000,8000,9000,10000"`
	Image       string `json:"avatar"`
	ProductName string `json:"productName" faker:"oneof:Quantum Widget, Stellar Gadget, Apex Tool, Nano Device, Titan Gear, Quantum Bolt, Stellar Module, Apex Component, Nano Instrument, Titan Appliance, Quantum Apparatus, Stellar Contraption, Apex Mechanism, Nano Widget, Titan Device, Quantum Gizmo, Stellar Apparatus, Apex Toolset, Nano Module, Titan Contraption"`
}

type GetProductsResponse struct {
	Data       []Product  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (p *Product) GenImageURL() {
	p.Image = fmt.Sprintf("https://picsum.photos/id/%d/300/200", rand.Intn(99)+1)
}
