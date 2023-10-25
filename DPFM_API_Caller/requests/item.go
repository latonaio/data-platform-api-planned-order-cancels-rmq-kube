package requests

type Item struct {
	PlannedOrder     int   `json:"PlannedOrder"`
	PlannedOrderItem int   `json:"PlannedOrderItem"`
	IsCancelled      *bool `json:"IsCancelled"`
}
