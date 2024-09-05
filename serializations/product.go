package serializations

type ProductForMail struct {
	ID            string                      `json:"id,omitempty"`
	NameTM        string                      `json:"name_tm,omitempty"`
	NameRU        string                      `json:"name_ru,omitempty"`
	NameEN        string                      `json:"name_en,omitempty"`
	DescriptionTM string                      `json:"description_tm,omitempty"`
	DescriptionRU string                      `json:"description_ru,omitempty"`
	DescriptionEN string                      `json:"description_en,omitempty"`
	Images        []string                    `json:"images,omitempty"`
	CategoryID    string                      `json:"category_id,omitempty"`
	Compositions  []ProductCompositionForMail `json:"compositions,omitempty"`
}
