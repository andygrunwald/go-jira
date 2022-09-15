package insight

type Icon struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	URL16 string `json:"url16" validate:"required"`
	URL48 string `json:"url48" validate:"required"`
}
