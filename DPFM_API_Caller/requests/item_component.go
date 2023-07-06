package requests

type ItemComponent struct {
	PlannedOrder                                   int      `json:"PlannedOrder"`
	PlannedOrderItem                               int      `json:"PlannedOrderItem"`
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                      *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID              *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID         *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID        *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	ProductionPlantBusinessPartner                 *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                *string  `json:"ProductionPlant"`
	ComponentProduct                               *string  `json:"ComponentProduct"`
	ComponentProductBuyer                          *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                         *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                 *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                 *string  `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty               *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant               *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductRequirementDate                *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime                *string  `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantityInBaseUnit     *float32 `json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit *float32 `json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	ComponentProductPlannedScrapInPercent          *float32 `json:"ComponentProductPlannedScrapInPercent"`
	ComponentProductBaseUnit                       *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductIsMarkedForBackflush           *bool    `json:"ComponentProductIsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	StockConfirmationBusinessPartner               *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                         *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantStorageLocation          *string  `json:"StockConfirmationPlantStorageLocation"`
	MRPArea                                        *string  `json:"MRPArea"`
	MRPController                                  *string  `json:"MRPController"`
	ProductionVersion                              *int     `json:"ProductionVersion"`
	ProductionVersionItem                          *int     `json:"ProductionVersionItem"`
	PlannedOrderIsReleased                         *bool    `json:"PlannedOrderIsReleased"`
	CreationDate                                   *string  `json:"CreationDate"`
	CreationTime                                   *string  `json:"CreationTime"`
	LastChangeDate                                 *string  `json:"LastChangeDate"`
	LastChangeTime                                 *string  `json:"LastChangeTime"`
	IsMarkedForDeletion                            *bool    `json:"IsMarkedForDeletion"`
}
