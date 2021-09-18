package response

type GetImagePaletteResponse struct {
	Colors [][]int `json:"rgb_colors"`
}