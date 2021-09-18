package request

type GetImagePaletteRequest struct {
	CommonImageAnalysisRequest
	ColorsCount int `schema:"colors_count"`
}
