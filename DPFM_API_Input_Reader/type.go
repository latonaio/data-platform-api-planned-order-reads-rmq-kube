package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"PlannedOrder"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
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
