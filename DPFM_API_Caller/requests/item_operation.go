package requests

type ItemOperation struct {
	PlannedOrder     int   `json:"PlannedOrder"`
	PlannedOrderItem int   `json:"PlannedOrderItem"`
	Operations       int   `json:"Operations"`
	OperationsItem   int   `json:"OperationsItem"`
	OperationID      int   `json:"OperationID"`
	IsCancelled      *bool `json:"IsCancelled"`
}
