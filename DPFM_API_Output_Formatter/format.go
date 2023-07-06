package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)
	i := 0
	for rows.Next() {
		i++
		pm := Header{}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.PlannedOrderType,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DestinationDeliverToParty,
			&pm.DestinationDeliverToPlant,
			&pm.DestinationDeliverToPlantStorageLocation,
			&pm.DepartureDeliverFromParty,
			&pm.DepartureDeliverFromPlant,
			&pm.DepartureDeliverFromPlantStorageLocation,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlant,
			&pm.OwnerProductionPlantStorageLocation,
			&pm.ProductBaseUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.PlannedOrderQuantityInBaseUnit,
			&pm.PlannedOrderQuantityInDestinationDeliveryUnit,
			&pm.PlannedOrderQuantityInDepartureDeliveryUnit,
			&pm.PlannedOrderQuantityInDestinationProductionUnit,
			&pm.PlannedOrderQuantityInDepartureProductionUnit,
			&pm.PlannedOrderDestinationDeliveryUnit,
			&pm.PlannedOrderDepartureDeliveryUnit,
			&pm.PlannedOrderDestinationProductionUnit,
			&pm.PlannedOrderDepartureProductionUnit,
			&pm.PlannedOrderPlannedScrapQtyInBaseUnit,
			&pm.PlannedOrderPlannedStartDate,
			&pm.PlannedOrderPlannedStartTime,
			&pm.PlannedOrderPlannedEndDate,
			&pm.PlannedOrderPlannedEndTime,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Projuct,
			&pm.WBSElement,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.PlannedOrderLongText,
			&pm.LastScheduledDate,
			&pm.ScheduledBasicEndDate,
			&pm.ScheduledBasicEndTime,
			&pm.ScheduledBasicStartDate,
			&pm.ScheduledBasicStartTime,
			&pm.SchedulingType,
			&pm.PlannedOrderIsReleased,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}
		header = append(header, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)
	i := 0
	for rows.Next() {
		i++
		pm := Item{}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.PrecedingProductionOrderItem,
			&pm.FollowingProductionOrderItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.PlannedOrderType,
			&pm.Product,
			&pm.Buyer,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductBaseUnit,
			&pm.ProductDeliveryUnit,
			&pm.ProductProductionUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantStorageLocation,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.PlannedOrderQuantityInBaseUnit,
			&pm.PlannedOrderQuantityInDeliveryUnit,
			&pm.PlannedOrderQuantityInProductionUnit,
			&pm.PlannedOrderPlannedScrapQtyInBaseUnit,
			&pm.PlannedOrderMinimumLotSizeQuantity,
			&pm.PlannedOrderStandardLotSizeQuantity,
			&pm.PlannedOrderMaximumLotSizeQuantity,
			&pm.PlannedOrderLotSizeRoundingQuantity,
			&pm.PlannedOrderLotSizeIsFixed,
			&pm.PlannedOrderPlannedStartDate,
			&pm.PlannedOrderPlannedStartTime,
			&pm.PlannedOrderPlannedEndDate,
			&pm.PlannedOrderPlannedEndTime,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.PlannedOrderLongText,
			&pm.LastScheduledDate,
			&pm.ScheduledBasicEndDate,
			&pm.ScheduledBasicEndTime,
			&pm.ScheduledBasicStartDate,
			&pm.ScheduledBasicStartTime,
			&pm.SchedulingType,
			&pm.PlannedOrderIsReleased,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}
		item = append(item, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}

func ConvertToItemComponent(rows *sql.Rows) (*[]ItemComponent, error) {
	defer rows.Close()
	itemComponent := make([]ItemComponent, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponent{}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequirementTime,
			&pm.ComponentProductRequiredQuantityInBaseUnit,
			&pm.ComponentProductRequiredQuantityInDeliveryUnit,
			&pm.ComponentProductPlannedScrapInPercent,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.BillOfMaterialItemText,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantStorageLocation,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.PlannedOrderIsReleased,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemComponent, err
		}
		itemComponent = append(itemComponent, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponent, nil
	}

	return &itemComponent, nil
}

func ConvertToItemOperation(rows *sql.Rows) (*[]ItemOperation, error) {
	defer rows.Close()
	itemOperation := make([]ItemOperation, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemOperation{}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.Sequence,
			&pm.SequenceText,
			&pm.OperationText,
			&pm.ProductBaseUnit,
			&pm.ProductProductionUnit,
			&pm.ProductOperationUnit,
			&pm.ProductDeliveryUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.OperationPlannedQuantityInBaseUnit,
			&pm.OperationPlannedQuantityInProductionUnit,
			&pm.OperationPlannedQuantityInOperationUnit,
			&pm.OperationPlannedQuantityInDeliveryUnit,
			&pm.OperationPlannedScrapInPercent,
			&pm.ResponsiblePlannerGroup,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.CapacityCategoryCode,
			&pm.OperationCostingRelevancyType,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.MaximumWaitDuration,
			&pm.StandardWaitDuration,
			&pm.MinimumWaitDuration,
			&pm.WaitDurationUnit,
			&pm.MaximumQueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueDurationUnit,
			&pm.MaximumMoveDuration,
			&pm.StandardMoveDuration,
			&pm.MinimumMoveDuration,
			&pm.MoveDurationUnit,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.CostElement,
			&pm.PlannedOrderIsReleased,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperation, err
		}
		itemOperation = append(itemOperation, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperation, nil
	}

	return &itemOperation, nil
}

func ConvertToItemOperationComponent(rows *sql.Rows) (*[]ItemOperationComponent, error) {
	defer rows.Close()
	itemOperationComponent := make([]ItemOperationComponent, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemOperationComponent{}
		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductDeliverToPartyInOperation,
			&pm.ComponentProductDeliverToPlantInOperation,
			&pm.ComponentProductDeliverFromPartyInOperation,
			&pm.ComponentProductDeliverFromPlantInOperation,
			&pm.ComponentProductRequirementDateInOperation,
			&pm.ComponentProductRequirementTimeInOperation,
			&pm.ComponentProductPlannedQuantityInBaseUnitInOperation,
			&pm.ComponentProductPlannedQuantityInDeliveryUnitInOperation,
			&pm.ComponentProductPlannedScrapInPercentInOperation,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.ComponentProductReservation,
			&pm.ComponentProductReservationItem,
			&pm.PlannedOrderIsReleased,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperationComponent, err
		}
		itemOperationComponent = append(itemOperationComponent, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperationComponent, nil
	}

	return &itemOperationComponent, nil
}
