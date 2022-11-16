package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// all jso fileds set for the zoho data Format
type Product struct {
	Id                          primitive.ObjectID   `json:"id" bson:"_id",omitempty`
	ZohoItemId                  *string              `json:"item_id" bson:"zohoItemId" validate:"required,min=2,max=100",omitempty`
	NAME                        *string              `json:"name" bson:"name" ,omitempty`
	PimcoreId                   *string              `json:"sku" bson:"pimcoreId" ,omitempty`
	Brand                       *string              `json:"brand" bson:"brand" ,omitempty`
	Manufacturer                *string              `json:"manufacturer" bson:"manufacturer" ,omitempty`
	CategoryId                  *string              `json:"category_id" bson:"categoryId" ,omitempty`
	CategoryName                *string              `json:"category_name" bson:"category_name" ,omitempty`
	HsnOrSac                    *string              `json:"hsn_or_sac" bson:"hsnOrSac" ,omitempty`
	ImageName                   *string              `json:"image_name" bson:"imageName" ,omitempty`
	ImageType                   *string              `json:"image_type" bson:"imageType" ,omitempty`
	Status                      *string              `json:"status" bson:"status" ,omitempty`
	Source                      *string              `json:"source" bson:"source" ,omitempty`
	IsLinkedWithZohoCrm         *bool                `json:"is_linked_with_zohocrm" bson:"isLinkedWithZohoCrm" ,omitempty`
	ZohoCrmProductId            *string              `json:"zcrm_product_id" bson:"zohoCrmProductId" ,omitempty`
	CrmOwnerId                  *string              `json:"crm_owner_id" bson:"crmOwnerId" ,omitempty`
	Unit                        *string              `json:"unit" bson:"unit" ,omitempty`
	Description                 *string              `json:"description" bson:"description" ,omitempty`
	ItemTaxPreferences          []ItemTaxPreferences `json:"item_tax_preferences" bson:"itemTaxPreferences" ,omitempty`
	Rate                        *float64             `json:"rate" bson:"rate" ,omitempty`
	AccountId                   *string              `json:"account_id" bson:"accountId" ,omitempty`
	AccountName                 *string              `json:"account_name" bson:"accountName" ,omitempty`
	TaxId                       *string              `json:"tax_id" bson:"taxId" ,omitempty`
	TaxName                     *string              `json:"tax_name" bson:"taxName" ,omitempty`
	TaxPercentage               *float64             `json:"tax_percentage" bson:"taxPercentage" ,omitempty`
	TaxType                     *string              `json:"tax_type" bson:"taxType" ,omitempty`
	IsDefaultTaxApplied         *bool                `json:"is_default_tax_applied" bson:"isDefaultTaxApplied" ,omitempty`
	IsTaxable                   *bool                `json:"is_taxable" bson:"isTaxable" ,omitempty`
	TaxExemptionId              *string              `json:"tax_exemption_id" bson:"taxExemptionId" ,omitempty`
	TaxExemptionCode            *string              `json:"tax_exemption_code" bson:"taxExemptionCode" ,omitempty`
	TaxabilityType              *string              `json:"taxability_type" bson:"taxabilityType" ,omitempty`
	Documents                   []interface{}        `json:"documents" bson:"documents" ,omitempty`
	PurchaseDescription         *string              `json:"purchase_description" bson:"purchaseDescription" ,omitempty`
	PriceBookRate               *float64             `json:"pricebook_rate" bson:"priceBookRate" ,omitempty`
	SalesRate                   *float64             `json:"sales_rate" bson:"salesRate" ,omitempty`
	PurchaseRate                *float64             `json:"purchase_rate" bson:"purchaseRate" ,omitempty`
	PurchaseAccountId           *string              `json:"purchase_account_id" bson:"purchaseAccountId" ,omitempty`
	PurchaseAccountName         *string              `json:"purchase_account_name" bson:"purchaseAccountName" ,omitempty`
	InventoryAccountId          *string              `json:"inventory_account_id" bson:"inventoryAccountId" ,omitempty`
	InventoryAccountName        *string              `json:"inventory_account_name" bson:"inventoryAccountName" ,omitempty`
	Tags                        []string             `json:"tags" bson:"tags" ,omitempty`
	ItemType                    *string              `json:"item_type" bson:"itemType" ,omitempty`
	ProductType                 *string              `json:"product_type" bson:"productType" ,omitempty`
	IsReturnable                *bool                `json:"is_returnable" bson:"isReturnable" ,omitempty`
	ReorderLevel                *string              `json:"reorder_level" bson:"reorderLevel" ,omitempty`
	MinimumOrderQuantity        *string              `json:"minimum_order_quantity" bson:"minimumOrderQuantity" ,omitempty`
	MaximumOrderQuantity        *string              `json:"maximum_order_quantity" bson:"maximumOrderQuantity" ,omitempty`
	InitialStock                *int64               `json:"initial_stock" bson:"initialStock" ,omitempty`
	InitialStockRate            *float64             `json:"initial_stock_rate" bson:"initialStockRate" ,omitempty`
	TotalInitialStock           *int64               `json:"total_initial_stock" bson:"totalInitialStock" ,omitempty`
	StockOnHand                 *float64             `json:"stock_on_hand" bson:"stockOnHand" ,omitempty`
	VendorId                    *string              `json:"vendor_id" bson:"vendorId" ,omitempty`
	VendorName                  *string              `json:"vendor_name" bson:"vendorName" ,omitempty`
	AssetValue                  *string              `json:"asset_value" bson:"assetValue" ,omitempty`
	AvailableStock              *float64             `json:"available_stock" bson:"availableStock" ,omitempty`
	ActualAvailableStock        *float64             `json:"actual_available_stock" bson:"actualAvailableStock" ,omitempty`
	CommittedStock              *float64             `json:"committed_stock" bson:"committedStock" ,omitempty`
	ActualCommittedStock        *float64             `json:"actual_committed_stock" bson:"actualCommittedStock" ,omitempty`
	AvailableForSaleStock       *float64             `json:"available_for_sale_stock" bson:"availableForSaleStock" ,omitempty`
	ActualAvailableForSaleStock *float64             `json:"actual_available_for_sale_stock" bson:"actualAvailableForSaleStock" ,omitempty`
	TrackBatchNumber            *bool                `json:"track_batch_number" bson:"trackBatchNumber" ,omitempty`
	IsComboProduct              *bool                `json:"is_combo_product" bson:"isComboProduct" ,omitempty`
	IsAdvancedTrackingMissing   *bool                `json:"is_advanced_tracking_missing" bson:"isAdvancedTrackingMissing" ,omitempty`
	CustomFields                []CustomFields       `json:"custom_fields" bson:"customFields" ,omitempty`
	SalesChannels               []interface{}        `json:"sales_channels" bson:"salesChannels" ,omitempty`
	Warehouses                  []Warehouses         `json:"warehouses" bson:"warehouses" ,omitempty`
	Branches                    []Branches           `json:"branches" bson:"branches" ,omitempty`
	PreferredVendors            []interface{}        `json:"preferred_vendors" bson:"preferredVendors" ,omitempty`
	PackageDetails              PackageDetails       `json:"package_details" bson:"packageDetails" ,omitempty`
	Product_id                  string               `json:"product_id" bson:"product_id" ,omitempty`
	ZohoCreatedTime             *string              `json:"created_time" bson:"zohoCreatedTime"`
	ZohoLastModifiedTime        *string              `json:"last_modified_time" bson:"zohoLastModifiedTime"`
	Created_at                  time.Time            `json:"created_at" bson:"created_at"`
	Updated_at                  time.Time            `json:"updated_at" bson:"updated_at"`
}

type PackageDetails struct {
	Length        *string `json:"length" bson:"length" ,omitempty`
	Width         *string `json:"width" bson:"width" ,omitempty`
	Height        *string `json:"height" bson:"height" ,omitempty`
	Weight        *string `json:"weight" bson:"weight" ,omitempty`
	WeightUnit    *string `json:"weight_unit" bson:"weightUnit" ,omitempty`
	DimensionUnit *string `json:"dimension_unit" bson:"dimensionUnit" ,omitempty`
}

type ItemTaxPreferences struct {
	TaxSpecification *string  `json:"tax_specification" bson:"taxSpecification" ,omitempty`
	TaxType          *int64   `json:"tax_type" bson:"taxType" ,omitempty`
	TaxName          *string  `json:"tax_name" bson:"taxName" ,omitempty`
	TaxPercentage    *float64 `json:"tax_percentage" bson:"taxPercentage" ,omitempty`
	TaxId            *string  `json:"tax_id" bson:"taxId" ,omitempty`
}

type CustomFields struct {
	TaxSpecification *string     `json:"field_id" bson:"taxSpecification" ,omitempty`
	TaxType          *string     `json:"customfield_id" bson:"taxType" ,omitempty`
	ShowInStore      *bool       `json:"show_in_store" bson:"showInStore" ,omitempty`
	ShowInPortal     *bool       `json:"show_in_portal" bson:"showInPortal" ,omitempty`
	IsActive         *bool       `json:"is_active" bson:"isActive" ,omitempty`
	Index            *int64      `json:"index" bson:"index" ,omitempty`
	Label            *string     `json:"label" bson:"label" ,omitempty`
	ShowOnPdf        *bool       `json:"show_on_pdf" bson:"showOnPdf" ,omitempty`
	EditOnPortal     *bool       `json:"edit_on_portal" bson:"editOnPortal" ,omitempty`
	EditOnStore      *bool       `json:"edit_on_store" bson:"editOnStore" ,omitempty`
	ApiName          *string     `json:"api_name" bson:"apiName" ,omitempty`
	ShowInAllPdf     *bool       `json:"show_in_all_pdf" bson:"showInAllPdf" ,omitempty`
	ValueFormatted   *string     `json:"value_formatted" bson:"valueFormatted" ,omitempty`
	SearchEntity     *string     `json:"search_entity" bson:"searchEntity" ,omitempty`
	DataType         *string     `json:"data_type" bson:"dataType" ,omitempty`
	PlaceHolder      *string     `json:"placeholder" bson:"placeHolder" ,omitempty`
	Value            interface{} `json:"value" bson:"value" ,omitempty`
	IsDependentField *bool       `json:"is_dependent_field" bson:"isDependentField" ,omitempty`
}

type Warehouses struct {
	WarehouseId                          *string     `json:"warehouse_id" bson:"warehouseId" ,omitempty`
	WarehouseName                        *string     `json:"warehouse_name" bson:"warehouseName" ,omitempty`
	Status                               *string     `json:"status" bson:"status" ,omitempty`
	IsPrimary                            *bool       `json:"is_primary" bson:"IsPrimary" ,omitempty`
	WarehouseStockOnHand                 *float64    `json:"warehouse_stock_on_hand" bson:"warehouseStockOnHand" ,omitempty`
	InitialStock                         *int64      `json:"initial_stock" bson:"initialStock" ,omitempty`
	InitialStockRate                     *float64    `json:"initial_stock_rate" bson:"initialStockRate" ,omitempty`
	WarehouseAvailableStock              *float64    `json:"warehouse_available_stock" bson:"warehouseAvailableStock" ,omitempty`
	Batches                              interface{} `json:"batches" bson:"batches" ,omitempty`
	SalesChannels                        interface{} `json:"sales_channels" bson:"salesChannels" ,omitempty`
	IsFbaWarehouse                       *bool       `json:"is_fba_warehouse" bson:"isFbaWarehouse" ,omitempty`
	WarehouseActualAvailableStock        *float64    `json:"warehouse_actual_available_stock" bson:"warehouseActualAvailableStock" ,omitempty`
	WarehouseCommittedStock              *float64    `json:"warehouse_committed_stock" bson:"WarehouseCommittedStock" ,omitempty`
	WarehouseActualCommittedStock        *float64    `json:"warehouse_actual_committed_stock" bson:"warehouseActualCommittedStock" ,omitempty`
	WarehouseAvailableForSaleStock       *float64    `json:"warehouse_available_for_sale_stock" bson:"warehouseAvailableForSaleStock" ,omitempty`
	WarehouseActualAvailableForSaleStock *float64    `json:"warehouse_actual_available_for_sale_stock" bson:"warehouseActualAvailableForSaleStock" ,omitempty`
	AdvancedTrackingMissingQuantity      *float64    `json:"advanced_tracking_missing_quantity" bson:"advancedTrackingMissingQuantity" ,omitempty`
}

type Branches struct {
	BranchId                          *string  `json:"branch_id" bson:"branchId" ,omitempty`
	BranchName                        *string  `json:"branch_name" bson:"branchName" ,omitempty`
	BranchAvailableStock              *float64 `json:"branch_available_stock" bson:"branchAvailableStock" ,omitempty`
	BranchAssetValue                  *int64   `json:"branch_asset_value" bson:"branchAssetValue" ,omitempty`
	BranchActualAvailableStock        *float64 `json:"branch_actual_available_stock" bson:"branchActualAvailableStock" ,omitempty`
	BranchCommittedStock              *float64 `json:"branch_committed_stock" bson:"branchCommittedStock" ,omitempty`
	BranchActualCommittedStock        *float64 `json:"branch_actual_committed_stock" bson:"branchActualCommittedStock" ,omitempty`
	BranchAvailableForSaleStock       *float64 `json:"branch_available_for_sale_stock" bson:"branchAvailableForSaleStock" ,omitempty`
	BranchActualAvailableForSaleStock *float64 `json:"branch_actual_available_for_sale_stock" bson:"branchActualAvailableForSaleStock" ,omitempty`
}
