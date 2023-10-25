package requests

type ItemOperationComponent struct {
	PlannedOrder       int   `json:"PlannedOrder"`
	PlannedOrderItem   int   `json:"PlannedOrderItem"`
	Operations         int   `json:"Operations"`
	OperationsItem     int   `json:"OperationsItem"`
	OperationID        int   `json:"OperationID"`
	BillOfMaterial     int   `json:"BillOfMaterial"`
	BillOfMaterialItem int   `json:"BillOfMaterialItem"`
	IsCancelled        *bool `json:"IsCancelled"`
}
