package requests

type ItemComponent struct {
	PlannedOrder       int   `json:"PlannedOrder"`
	PlannedOrderItem   int   `json:"PlannedOrderItem"`
	BillOfMaterial     int   `json:"BillOfMaterial"`
	BillOfMaterialItem int   `json:"BillOfMaterialItem"`
	IsCancelled        *bool `json:"IsCancelled"`
}
