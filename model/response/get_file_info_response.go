package response

type GetFileInfoResponse struct {
	Width string `json:"width"`
	Height string `json:"height"`
	Filetype string `json:"filetype"`
	ColorSchema string `json:"color_schema"`
	Size string `json:"size"`
}