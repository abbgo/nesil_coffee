package serializations

type ProductCompositionForMail struct {
	ID         string `json:"id,omitempty"`
	NameTM     string `json:"name_tm,omitempty"`
	NameRU     string `json:"name_ru,omitempty"`
	NameEN     string `json:"name_en,omitempty"`
	Percentage int8   `json:"percentage,omitempty"`
}
