package response

type GetRGBHistogramResponse struct {
	Red 	[]float64 `json:"red_hist"`
	Green 	[]float64 `json:"green_hist"`
	Blue 	[]float64 `json:"blue_hist"`
}