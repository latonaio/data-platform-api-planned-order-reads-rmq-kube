package requests

type ItemOperationComponent struct {
	PlannedOrder                                             int      `json:"PlannedOrder"`
	PlannedOrderItem                                         int      `json:"PlannedOrderItem"`
	Operations                                               int      `json:"Operations"`
	OperationsItem                                           int      `json:"OperationsItem"`
	OperationID                                              int      `json:"OperationID"`
	BillOfMaterial                                           int      `json:"BillOfMaterial"`
	BillOfMaterialItem                                       int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID                        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID                   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionPlantBusinessPartner                           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                          *string  `json:"ProductionPlant"`
	ComponentProduct                                         *string  `json:"ComponentProduct"`
	ComponentProductBuyer                                    *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                                   *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                           *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                           *string  `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty                         *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                         *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductDeliverToPartyInOperation                *int     `json:"ComponentProductDeliverToPartyInOperation"`
	ComponentProductDeliverToPlantInOperation                *string  `json:"ComponentProductDeliverToPlantInOperation"`
	ComponentProductDeliverFromPartyInOperation              *int     `json:"ComponentProductDeliverFromPartyInOperation"`
	ComponentProductDeliverFromPlantInOperation              *string  `json:"ComponentProductDeliverFromPlantInOperation"`
	ComponentProductRequirementDateInOperation               *string  `json:"ComponentProductRequirementDateInOperation"`
	ComponentProductRequirementTimeInOperation               *string  `json:"ComponentProductRequirementTimeInOperation"`
	ComponentProductPlannedQuantityInBaseUnitInOperation     *float32 `json:"ComponentProductPlannedQuantityInBaseUnitInOperation"`
	ComponentProductPlannedQuantityInDeliveryUnitInOperation *float32 `json:"ComponentProductPlannedQuantityInDeliveryUnitInOperation"`
	ComponentProductPlannedScrapInPercentInOperation         *float32 `json:"ComponentProductPlannedScrapInPercentInOperation"`
	ComponentProductBaseUnit                                 *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                             *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductIsMarkedForBackflush                     *bool    `json:"ComponentProductIsMarkedForBackflush"`
	MRPArea                                                  *string  `json:"MRPArea"`
	MRPController                                            *string  `json:"MRPController"`
	ProductionVersion                                        *int     `json:"ProductionVersion"`
	ProductionVersionItem                                    *int     `json:"ProductionVersionItem"`
	ComponentProductReservation                              *int     `json:"ComponentProductReservation"`
	ComponentProductReservationItem                          *int     `json:"ComponentProductReservationItem"`
	PlannedOrderIsReleased                                   *bool    `json:"PlannedOrderIsReleased"`
	CreationDate                                             *string  `json:"CreationDate"`
	CreationTime                                             *string  `json:"CreationTime"`
	LastChangeDate                                           *string  `json:"LastChangeDate"`
	LastChangeTime                                           *string  `json:"LastChangeTime"`
	IsMarkedForDeletion                                      *bool    `json:"IsMarkedForDeletion"`
}
