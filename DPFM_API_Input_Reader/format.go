package dpfm_api_input_reader

import (
	"data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		PlannedOrder:                          data.PlannedOrder,
		PlannedOrderType:                      data.PlannedOrderType,
		Product:                               data.Product,
		DeliverFromParty:                      data.DeliverFromParty,
		DeliverToParty:                        data.DeliverToParty,
		IssuingPlant:                          data.IssuingPlant,
		IssuingPlantStorageLocation:           data.IssuingPlantStorageLocation,
		ReceivingPlant:                        data.ReceivingPlant,
		ReceivingPlantStorageLocation:         data.ReceivingPlantStorageLocation,
		ProductionPlantBusinessPartner:        data.ProductionPlantBusinessPartner,
		ProductionPlant:                       data.ProductionPlant,
		ProductionPlantStorageLocation:        data.ProductionPlantStorageLocation,
		MRPArea:                               data.MRPArea,
		BaseUnit:                              data.BaseUnit,
		PlannedOrderQuantityInBaseUnit:        data.PlannedOrderQuantityInBaseUnit,
		PlannedOrderPlannedScrapQtyInBaseUnit: data.PlannedOrderPlannedScrapQtyInBaseUnit,
		PlannedOrderIssuingUnit:               data.PlannedOrderIssuingUnit,
		PlannedOrderReceivingUnit:             data.PlannedOrderReceivingUnit,
		PlannedOrderIssuingQuantity:           data.PlannedOrderIssuingQuantity,
		PlannedOrderReceivingQuantity:         data.PlannedOrderReceivingQuantity,
		PlannedOrderPlannedStartDate:          data.PlannedOrderPlannedStartDate,
		PlannedOrderPlannedStartTime:          data.PlannedOrderPlannedStartTime,
		PlannedOrderPlannedEndDate:            data.PlannedOrderPlannedEndDate,
		PlannedOrderPlannedEndTime:            data.PlannedOrderPlannedEndTime,
		LastChangeDateTime:                    data.LastChangeDateTime,
		OrderID:                               data.OrderID,
		OrderItem:                             data.OrderItem,
		Buyer:                                 data.Buyer,
		Seller:                                data.Seller,
		Project:                               data.Project,
		Reservation:                           data.Reservation,
		PlannedOrderLongText:                  data.PlannedOrderLongText,
		MRPController:                         data.MRPController,
		PlannedOrderIsFixed:                   data.PlannedOrderIsFixed,
		PlannedOrderBOMFixed:                  data.PlannedOrderBOMFixed,
		LastScheduledDate:                     data.LastScheduledDate,
		ScheduledBasicEndDate:                 data.ScheduledBasicEndDate,
		ScheduledBasicEndTime:                 data.ScheduledBasicEndTime,
		ScheduledBasicStartDate:               data.ScheduledBasicStartDate,
		ScheduledBasicStartTime:               data.ScheduledBasicStartTime,
		SchedulingType:                        data.SchedulingType,
	}
}

func (sdc *SDC) ConvertToComponent() *requests.Component {
	dataHeader := sdc.Header
	data := sdc.Header.Component
	return &requests.Component{
		PlannedOrder:                           dataHeader.PlannedOrder,
		BOMItem:                                data.BOMItem,
		BOMItemDescription:                     data.BOMItemDescription,
		BillOfMaterialCategory:                 data.BillOfMaterialCategory,
		BillOfMaterialItemNumber:               data.BillOfMaterialItemNumber,
		BillOfMaterialInternalID:               data.BillOfMaterialInternalID,
		Reservation:                            data.Reservation,
		ReservationItem:                        data.ReservationItem,
		ComponentProduct:                       data.ComponentProduct,
		ComponentProductRequirementDate:        data.ComponentProductRequirementDate,
		ComponentProductRequiredQuantity:       data.ComponentProductRequiredQuantity,
		BaseUnit:                               data.BaseUnit,
		WithdrawnQuantity:                      data.WithdrawnQuantity,
		ComponentScrapInPercent:                data.ComponentScrapInPercent,
		QuantityIsFixed:                        data.QuantityIsFixed,
		ComponentWithdrawnPlantBusinessPartner: data.ComponentWithdrawnPlantBusinessPartner,
		ComponentWithdrawnPlant:                data.ComponentWithdrawnPlant,
		ComponentWithdrawnStorageLocation:      data.ComponentWithdrawnStorageLocation,
		MRPController:                          data.MRPController,
		LastChangeDateTime:                     data.LastChangeDateTime,
	}
}
