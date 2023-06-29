package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header    *Header    `json:"Header"`
	Component *Component `json:"Component"`
}

type PlannedOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type Header struct {
	PlannedOrder                          int       `json:"PlannedOrder"`
	PlannedOrderType                      *string   `json:"PlannedOrderType"`
	Product                               *string   `json:"Product"`
	DeliverFromParty                      *int      `json:"DeliverFromParty"`
	DeliverToParty                        *int      `json:"DeliverToParty"`
	IssuingPlant                          *string   `json:"IssuingPlant"`
	IssuingPlantStorageLocation           *string   `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                        *string   `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation         *string   `json:"ReceivingPlantStorageLocation"`
	ProductionPlantBusinessPartner        *int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                       *string   `json:"ProductionPlant"`
	ProductionPlantStorageLocation        *string   `json:"ProductionPlantStorageLocation"`
	MRPArea                               *string   `json:"MRPArea"`
	BaseUnit                              *string   `json:"BaseUnit"`
	PlannedOrderQuantityInBaseUnit        *float32  `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit *float32  `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderIssuingUnit               *string   `json:"PlannedOrderIssuingUnit"`
	PlannedOrderReceivingUnit             *string   `json:"PlannedOrderReceivingUnit"`
	PlannedOrderIssuingQuantity           *float32  `json:"PlannedOrderIssuingQuantity"`
	PlannedOrderReceivingQuantity         *float32  `json:"PlannedOrderReceivingQuantity"`
	PlannedOrderPlannedStartDate          *string   `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime          *string   `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate            *string   `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime            *string   `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                    *string   `json:"LastChangeDateTime"`
	OrderID                               *int      `json:"OrderID"`
	OrderItem                             *int      `json:"OrderItem"`
	Buyer                                 *int      `json:"Buyer"`
	Seller                                *int      `json:"Seller"`
	Project                               *string   `json:"Project"`
	Reservation                           *int      `json:"Reservation"`
	PlannedOrderLongText                  *string   `json:"PlannedOrderLongText"`
	MRPController                         *string   `json:"MRPController"`
	PlannedOrderIsFixed                   *bool     `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMFixed                  *bool     `json:"PlannedOrderBOMFixed"`
	LastScheduledDate                     *string   `json:"LastScheduledDate"`
	ScheduledBasicEndDate                 *string   `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                 *string   `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate               *string   `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime               *string   `json:"ScheduledBasicStartTime"`
	SchedulingType                        *string   `json:"SchedulingType"`
	Component                             Component `json:"Component"`
}

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
