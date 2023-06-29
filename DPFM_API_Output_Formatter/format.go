package dpfm_api_output_formatter

import (
	"data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*Header, error) {
	pm := &requests.Header{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderType,
			&pm.Product,
			&pm.DeliverFromParty,
			&pm.DeliverToParty,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.MRPArea,
			&pm.BaseUnit,
			&pm.PlannedOrderQuantityInBaseUnit,
			&pm.PlannedOrderPlannedScrapQtyInBaseUnit,
			&pm.PlannedOrderIssuingUnit,
			&pm.PlannedOrderReceivingUnit,
			&pm.PlannedOrderIssuingQuantity,
			&pm.PlannedOrderReceivingQuantity,
			&pm.PlannedOrderPlannedStartDate,
			&pm.PlannedOrderPlannedStartTime,
			&pm.PlannedOrderPlannedEndDate,
			&pm.PlannedOrderPlannedEndTime,
			&pm.LastChangeDateTime,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Buyer,
			&pm.Seller,
			&pm.Project,
			&pm.Reservation,
			&pm.PlannedOrderLongText,
			&pm.MRPController,
			&pm.PlannedOrderIsFixed,
			&pm.PlannedOrderBOMFixed,
			&pm.LastScheduledDate,
			&pm.ScheduledBasicEndDate,
			&pm.ScheduledBasicEndTime,
			&pm.ScheduledBasicStartDate,
			&pm.ScheduledBasicStartTime,
			&pm.SchedulingType,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
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
	return header, nil
}

func ConvertToComponent(sdc *api_input_reader.SDC, rows *sql.Rows) (*Component, error) {
	pm := &requests.Component{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.BOMItem,
			&pm.BOMItemDescription,
			&pm.BillOfMaterialCategory,
			&pm.BillOfMaterialItemNumber,
			&pm.BillOfMaterialInternalID,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.ComponentProduct,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequiredQuantity,
			&pm.BaseUnit,
			&pm.WithdrawnQuantity,
			&pm.ComponentScrapInPercent,
			&pm.QuantityIsFixed,
			&pm.ComponentWithdrawnPlantBusinessPartner,
			&pm.ComponentWithdrawnPlant,
			&pm.ComponentWithdrawnStorageLocation,
			&pm.MRPController,
			&pm.LastChangeDateTime,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	component := &Component{
		PlannedOrder:                           data.PlannedOrder,
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
	return component, nil
}
