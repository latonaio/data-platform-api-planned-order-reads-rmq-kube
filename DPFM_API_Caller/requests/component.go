package requests

type Component struct {
	PlannedOrder                           int      `json:"PlannedOrder"`
	BOMItem                                int      `json:"BOMItem"`
	BOMItemDescription                     *string  `json:"BOMItemDescription"`
	BillOfMaterialCategory                 *string  `json:"BillOfMaterialCategory"`
	BillOfMaterialItemNumber               *int     `json:"BillOfMaterialItemNumber"`
	BillOfMaterialInternalID               *int     `json:"BillOfMaterialInternalID"`
	Reservation                            *int     `json:"Reservation"`
	ReservationItem                        *int     `json:"ReservationItem"`
	ComponentProduct                       *string  `json:"ComponentProduct"`
	ComponentProductRequirementDate        *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequiredQuantity       *float32 `json:"ComponentProductRequiredQuantity"`
	BaseUnit                               *string  `json:"BaseUnit"`
	WithdrawnQuantity                      *string  `json:"WithdrawnQuantity"`
	ComponentScrapInPercent                *string  `json:"ComponentScrapInPercent"`
	QuantityIsFixed                        *bool    `json:"QuantityIsFixed"`
	ComponentWithdrawnPlantBusinessPartner *string  `json:"ComponentWithdrawnPlantBusinessPartner"`
	ComponentWithdrawnPlant                *string  `json:"ComponentWithdrawnPlant"`
	ComponentWithdrawnStorageLocation      *string  `json:"ComponentWithdrawnStorageLocation"`
	MRPController                          *string  `json:"MRPController"`
	LastChangeDateTime                     *string  `json:"LastChangeDateTime"`
}
