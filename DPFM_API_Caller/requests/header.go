package requests

type Header struct {
	PlannedOrder                          int      `json:"PlannedOrder"`
	PlannedOrderType                      *string  `json:"PlannedOrderType"`
	Product                               *string  `json:"Product"`
	DeliverFromParty                      *int     `json:"DeliverFromParty"`
	DeliverToParty                        *int     `json:"DeliverToParty"`
	IssuingPlant                          *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation           *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                        *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation         *string  `json:"ReceivingPlantStorageLocation"`
	ProductionPlantBusinessPartner        *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                       *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation        *string  `json:"ProductionPlantStorageLocation"`
	MRPArea                               *string  `json:"MRPArea"`
	BaseUnit                              *string  `json:"BaseUnit"`
	PlannedOrderQuantityInBaseUnit        *float32 `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit *float32 `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderIssuingUnit               *string  `json:"PlannedOrderIssuingUnit"`
	PlannedOrderReceivingUnit             *string  `json:"PlannedOrderReceivingUnit"`
	PlannedOrderIssuingQuantity           *float32 `json:"PlannedOrderIssuingQuantity"`
	PlannedOrderReceivingQuantity         *float32 `json:"PlannedOrderReceivingQuantity"`
	PlannedOrderPlannedStartDate          *string  `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime          *string  `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate            *string  `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime            *string  `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                    *string  `json:"LastChangeDateTime"`
	OrderID                               *int     `json:"OrderID"`
	OrderItem                             *int     `json:"OrderItem"`
	Buyer                                 *int     `json:"Buyer"`
	Seller                                *int     `json:"Seller"`
	Project                               *string  `json:"Project"`
	Reservation                           *int     `json:"Reservation"`
	PlannedOrderLongText                  *string  `json:"PlannedOrderLongText"`
	MRPController                         *string  `json:"MRPController"`
	PlannedOrderIsFixed                   *bool    `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMFixed                  *bool    `json:"PlannedOrderBOMFixed"`
	LastScheduledDate                     *string  `json:"LastScheduledDate"`
	ScheduledBasicEndDate                 *string  `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                 *string  `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate               *string  `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime               *string  `json:"ScheduledBasicStartTime"`
	SchedulingType                        *string  `json:"SchedulingType"`
}
