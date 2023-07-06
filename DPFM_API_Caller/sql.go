package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-order-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
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
	var header *[]dpfm_api_output_formatter.Header
	// var headerDoc *[]dpfm_api_output_formatter.HeaderDoc
	var item *[]dpfm_api_output_formatter.Item
	var itemComponent *[]dpfm_api_output_formatter.ItemComponent
	var itemOperation *[]dpfm_api_output_formatter.ItemOperation
	var itemOperationComponent *[]dpfm_api_output_formatter.ItemOperationComponent

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()

		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemComponent":
			func() {
				itemComponent = c.ItemComponent(mtx, input, output, errs, log)
			}()
		case "ItemComponents":
			func() {
				itemComponent = c.ItemComponents(mtx, input, output, errs, log)
			}()

		case "ItemOperation":
			func() {
				itemOperation = c.ItemOperation(mtx, input, output, errs, log)
			}()
		case "ItemOperations":
			func() {
				itemOperation = c.ItemOperations(mtx, input, output, errs, log)
			}()
		case "ItemOperationComponent":
			func() {
				itemOperationComponent = c.ItemOperationComponent(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                 header,
		Item:                   item,
		ItemComponent:          itemComponent,
		ItemOperation:          itemOperation,
		ItemOperationComponent: itemOperationComponent,
	}
	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE PlannedOrder = %d ", input.Header.PlannedOrder)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data
		` + where + ` ORDER BY IsMarkedForDeletion DESC, ProductionVersion DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	where := fmt.Sprintf("WHERE PlannedOrder = %d ", input.Header.PlannedOrder)

	itemIDs := ""
	for _, v := range input.Header.Item {
		itemIDs = fmt.Sprintf("%s, %d", itemIDs, v.PlannedOrderItem)
	}

	if len(itemIDs) != 0 {
		where = fmt.Sprintf("%s\nAND PlannedOrderItem IN ( %s ) ", where, itemIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_data
		`+where+` ORDER BY PlannedOrder DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	where := fmt.Sprintf("WHERE PlannedOrder = %d ", input.Header.PlannedOrder)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_data
		` + where + ` ORDER BY PlannedOrder DESC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponent(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	var args []interface{}
	plannedOrder := input.Header.PlannedOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, plannedOrder, v.PlannedOrderItem, x.Operations, x.OperationsItem, x.BillOfMaterial, x.BillOfMaterialItem)
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_component_data
		WHERE (PlannedOrder, PlannedOrderItem, Operations, OperationsItem, BillOfMaterial, BillOfMaterialItem) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion ASC, PlannedOrder DESC, PlannedOrderItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	item := &dpfm_api_input_reader.Item{}
	itemComponent := &dpfm_api_input_reader.ItemComponent{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	if len(item.ItemComponent) > 0 {
		itemComponent = &item.ItemComponent[0]
	}

	where := fmt.Sprintf("WHERE PlannedOrder = %d", input.Header.PlannedOrder)
	if item != nil {
		if &item.PlannedOrderItem != nil {
			where = fmt.Sprintf("%s\nAND PlannedOrderItem = %d", where, item.PlannedOrderItem)
		}
	}
	if itemComponent != nil {
		if itemComponent.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemComponent.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_component_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, PlannedOrder DESC, PlannedOrderItem ASC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	item := &dpfm_api_input_reader.Item{}
	itemOperation := &dpfm_api_input_reader.ItemOperation{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	if len(item.ItemOperation) > 0 {
		itemOperation = &item.ItemOperation[0]
	}

	where := fmt.Sprintf("WHERE PlannedOrder = %d", input.Header.PlannedOrder)
	if item != nil {
		if &item.PlannedOrderItem != nil {
			where = fmt.Sprintf("%s\nAND PlannedOrderItem = %d", where, item.PlannedOrderItem)
		}
	}
	if itemOperation != nil {
		if itemOperation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemOperation.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_operations_data
		` + where + ` ORDER BY PlannedOrder DESC, OperationsItem ASC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	var args []interface{}
	plannedOrder := input.Header.PlannedOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			args = append(args, plannedOrder, v.PlannedOrderItem, w.Operations, w.OperationsItem)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_operations_data
		WHERE (PlannedOrder, PlannedOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY PlannedOrder DESC, OperationsItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperationComponent(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationComponent {
	var args []interface{}
	plannedOrder := input.Header.PlannedOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, plannedOrder, v.PlannedOrderItem, x.Operations, x.OperationsItem)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_operation_component_data
		WHERE (PlannedOrder, PlannedOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY PlannedOrder DESC, OperationsItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	return data
}
