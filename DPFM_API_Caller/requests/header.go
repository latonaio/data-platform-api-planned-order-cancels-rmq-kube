package requests

type Header struct {
	PlannedOrder int   `json:"PlannedOrder"`
	IsCancelled  *bool `json:"IsCancelled"`
}
