package entity

type HttpResponse struct {
	Data       PrintRequest `json:"data,omitempty"`
	Message    string       `json:"message"`
	HttpStatus int          `json:"http_status"`
}

type PrintRequest struct {
	ItemName                string  `json:"item_name"`
	EstimatedWeight         float32 `json:"estimated_weight"`
	EstimatedFilamentLength float32 `json:"estimated_filament_length"`
	EstimatedDuration       int     `json:"estimated_duration"`
	FileUrl                 string  `json:"file_url"`
	Requestor               string  `json:"requestor"`
	Status                  string  `json:"status"`
}

func NewPrintRequest(itemName string, estimatedWeight float32, estimatedFilamentLength float32, estimatedDuration int, fileUrl string, requestor string, status string) *PrintRequest {
	return &PrintRequest{
		ItemName:                itemName,
		EstimatedWeight:         estimatedWeight,
		EstimatedFilamentLength: estimatedFilamentLength,
		EstimatedDuration:       estimatedDuration,
		FileUrl:                 fileUrl,
		Requestor:               requestor,
		Status:                  status,
	}
}
