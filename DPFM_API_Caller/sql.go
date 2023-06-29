package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var component *dpfm_api_output_formatter.Component
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Component":
			func() {
				component = c.Component(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:    header,
		Component: component,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	plannedOrder := input.Header.PlannedOrder

	rows, err := c.db.Query(
		`SELECT PlannedOrder, PlannedOrderType, Product, DeliverFromParty, DeliverToParty, IssuingPlant, 
		IssuingPlantStorageLocation, ReceivingPlant, ReceivingPlantStorageLocation, ProductionPlantBusinessPartner, 
		ProductionPlant, ProductionPlantStorageLocation, MRPArea, BaseUnit, PlannedOrderQuantityInBaseUnit, 
		PlannedOrderPlannedScrapQtyInBaseUnit, PlannedOrderIssuingUnit, PlannedOrderReceivingUnit, PlannedOrderIssuingQuantity, 
		PlannedOrderReceivingQuantity, PlannedOrderPlannedStartDate, PlannedOrderPlannedStartTime, PlannedOrderPlannedEndDate, 
		PlannedOrderPlannedEndTime, LastChangeDateTime, OrderID, OrderItem, Buyer, Seller, Project, Reservation, 
		PlannedOrderLongText, MRPController, PlannedOrderIsFixed, PlannedOrderBOMFixed, LastScheduledDate, 
		ScheduledBasicEndDate, ScheduledBasicEndTime, ScheduledBasicStartDate, ScheduledBasicStartTime, SchedulingType
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data
		WHERE PlannedOrder = ?;`, plannedOrder,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Component(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Component {
	plannedOrder := input.Header.PlannedOrder
	bOMItem := input.Header.Component.BOMItem

	rows, err := c.db.Query(
		`SELECT PlannedOrder, BOMItem, BOMItemDescription, BillOfMaterialCategory, BillOfMaterialItemNumber, BillOfMaterialInternalID, 
		Reservation, ReservationItem, ComponentProduct, ComponentProductRequirementDate, ComponentProductRequiredQuantity, BaseUnit, 
		WithdrawnQuantity, ComponentScrapInPercent, QuantityIsFixed, ComponentWithdrawnPlantBusinessPartner, ComponentWithdrawnPlant, 
		ComponentWithdrawnStorageLocation, MRPController, LastChangeDateTime
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_component_data
		WHERE (PlannedOrder, BOMItem) = (?, ?);`, plannedOrder, bOMItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToComponent(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
