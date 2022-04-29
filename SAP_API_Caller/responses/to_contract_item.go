package responses

type ToContractItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                              string `json:"ObjectID"`
			ParentObjectID                                        string `json:"ParentObjectID"`
			ContractID                                            string `json:"ContractID"`
			ItemID                                                string `json:"ID"`
			ParentItemID                                          string `json:"ParentItemID"`
			ParentItemUUID                                        string `json:"ParentItemUUID"`
			ProcessingTypeCode                                    string `json:"ProcessingTypeCode"`
			DeterminationRelevanceCode                            string `json:"DeterminationRelevanceCode"`
			ReleaseControlCode                                    string `json:"ReleaseControlCode"`
			Description                                           string `json:"Description"`
			ValidityStartDate                                     string `json:"ValidityStartDate"`
			ValidityEndDate                                       string `json:"ValidityEndDate"`
			WarrantyGoodwillCode                                  string `json:"WarrantyGoodwillCode"`
			InvoiceScheduleStartDate                              string `json:"InvoiceScheduleStartDate"`
			InvoiceScheduleEndDate                                string `json:"InvoiceScheduleEndDate"`
			InvoiceScheduleInvoicingInAdvanceIndicator            bool   `json:"InvoiceScheduleInvoicingInAdvanceIndicator"`
			InvoiceScheduleNextInvoiceDateCalculationFunctionCode string `json:"InvoiceScheduleNextInvoiceDateCalculationFunctionCode"`
			InvoiceScheduleHorizonDateCalculationFunctionCode     string `json:"InvoiceScheduleHorizonDateCalculationFunctionCode"`
			CancellationReasonCode                                string `json:"CancellationReasonCode"`
			ContractLifeCycleStatusCode                           string `json:"ContractLifeCycleStatusCode"`
			ConsistencyStatusCode                                 string `json:"ConsistencyStatusCode"`
			ExternalContractLifeCycleStatusCode                   string `json:"ExternalContractLifeCycleStatusCode"`
			ExternalContractReferenceProcessingStatusCode         string `json:"ExternalContractReferenceProcessingStatusCode"`
			ExternalContractInvoiceProcessingStatusCode           string `json:"ExternalContractInvoiceProcessingStatusCode"`
			ProductID                                             string `json:"ProductID"`
			ProductUUID                                           string `json:"ProductUUID"`
			ProductDescription                                    string `json:"ProductDescription"`
			Quantity                                              string `json:"Quantity"`
			QuantityUnitCode                                      string `json:"QuantityUnitCode"`
			TargetQuantity                                        string `json:"TargetQuantity"`
			TargetQuantityUnitCode                                string `json:"TargetQuantityUnitCode"`
			TotalPeriodReleaseQuantity                            string `json:"TotalPeriodReleaseQuantity"`
			TotalPeriodReleaseQuantityUnitCode                    string `json:"TotalPeriodReleaseQuantityUnitCode"`
			TotalPeriodOpenQuantity                               string `json:"TotalPeriodOpenQuantity"`
			TotalPeriodOpenQuantityUnitCode                       string `json:"TotalPeriodOpenQuantityUnitCode"`
			TargetStatusCode                                      string `json:"TargetStatusCode"`
			DeliveryPriorityCode                                  string `json:"DeliveryPriorityCode"`
			IncotermsClassificationCode                           string `json:"IncotermsClassificationCode"`
			IncotermsTransferLocationName                         string `json:"IncotermsTransferLocationName"`
			CashDiscountTermsCode                                 string `json:"CashDiscountTermsCode"`
			PriceDateTime                                         string `json:"PriceDateTime"`
			NetAmount                                             string `json:"NetAmount"`
			NetAmountCurrencyCode                                 string `json:"NetAmountCurrencyCode"`
			NetPriceAmount                                        string `json:"NetPriceAmount"`
			NetPriceAmountCurrencyCode                            string `json:"NetPriceAmountCurrencyCode"`
			InvoiceScheduleNetAmount                              string `json:"InvoiceScheduleNetAmount"`
			InvoiceScheduleNetAmountCurrencyCode                  string `json:"InvoiceScheduleNetAmountCurrencyCode"`
			ProductRecipientPartyID                               string `json:"ProductRecipientPartyID"`
			ProductRecipientPartyUUID                             string `json:"ProductRecipientPartyUUID"`
			ProductRecipientPartyFormattedName                    string `json:"ProductRecipientPartyFormattedName"`
			BillToPartyID                                         string `json:"BillToPartyID"`
			BillToPartyUUID                                       string `json:"BillToPartyUUID"`
			BillToPartyFormattedName                              string `json:"BillToPartyFormattedName"`
			PayerPartyID                                          string `json:"PayerPartyID"`
			PayerPartyUUID                                        string `json:"PayerPartyUUID"`
			PayerPartyFormattedName                               string `json:"PayerPartyFormattedName"`
			EntityLastChangedOn                                   string `json:"EntityLastChangedOn"`
			ContractItemExternalPriceComponent                    struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ContractItemExternalPriceComponent"`
		} `json:"results"`
	} `json:"d"`
}
