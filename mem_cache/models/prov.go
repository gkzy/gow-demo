package models

type Prov struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetProvData get prov data slice
func (m *Prov) GetProvData() (data []*Prov, err error) {
	data = make([]*Prov, 0)
	data = append(data, &Prov{
		ID:   51,
		Name: "四川",
	})
	data = append(data, &Prov{
		ID:   50,
		Name: "重庆",
	})
	data = append(data, &Prov{
		ID:   11,
		Name: "北京",
	})
	return
}
