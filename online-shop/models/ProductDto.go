package models

type ProductDto struct {
	id int `json:"id"`
	name string `json:"name"`
	price float32 `json:"value"`
}

func NewProduct(id int, name string, value float32) *ProductDto {
	p := ProductDto {
		id: id,
		name: name,
		price: value,
	}
	return &p
}

func (p *ProductDto) getId() int {
	return p.id
}

func (p *ProductDto) getName() string {
	return p.name
}

func (p *ProductDto) getPrice() float32 {
	return p.price
}