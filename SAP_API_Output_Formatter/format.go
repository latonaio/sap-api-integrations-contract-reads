package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-contract-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToContractCollection(raw []byte, l *logger.Logger) ([]ContractCollection, error) {
	pm := &responses.ContractCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	contractCollection := make([]ContractCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		contractCollection = append(contractCollection, ContractCollection{
			ObjectID:                           data.ObjectID,
			ID:                                 data.ID,
			BuyerID:                            data.BuyerID,
			MainBusinessProcessVariantTypeCode: data.MainBusinessProcessVariantTypeCode,
			ProcessingTypeCode:                 data.ProcessingTypeCode,
			ContractDeterminationCoveredObjectsRelevanceCode: data.ContractDeterminationCoveredObjectsRelevanceCode,
			TemplateIndicator:                          data.TemplateIndicator,
			Name:                                       data.Name,
			CreationDateTime:                           data.CreationDateTime,
			CreationIdentityUUID:                       data.CreationIdentityUUID,
			LastChangeDateTime:                         data.LastChangeDateTime,
			LastChangeIdentityUUID:                     data.LastChangeIdentityUUID,
			Date:                                       data.Date,
			ValidityStartDate:                          data.ValidityStartDate,
			ValidityEndDate:                            data.ValidityEndDate,
			TimeZoneCode:                               data.TimeZoneCode,
			DocumentLanguageCode:                       data.DocumentLanguageCode,
			IncludeBusinessPartnerHierarchyIndicator:   data.IncludeBusinessPartnerHierarchyIndicator,
			IncludeReleaseAuthorisedPartiesIndicator:   data.IncludeReleaseAuthorisedPartiesIndicator,
			ServiceLevelObjectiveID:                    data.ServiceLevelObjectiveID,
			ServiceLevelObjectiveUUID:                  data.ServiceLevelObjectiveUUID,
			ServiceLevelObjectiveName:                  data.ServiceLevelObjectiveName,
			CancellationReasonCode:                     data.CancellationReasonCode,
			InvoiceScheduleStartDate:                   data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:                     data.InvoiceScheduleEndDate,
			InvoiceScheduleInvoicingInAdvanceIndicator: data.InvoiceScheduleInvoicingInAdvanceIndicator,
			InvoiceScheduleHorizonDateCalculationFunctionCode:     data.InvoiceScheduleHorizonDateCalculationFunctionCode,
			InvoiceScheduleNextInvoiceDateCalculationFunctionCode: data.InvoiceScheduleNextInvoiceDateCalculationFunctionCode,
			ContractLifeCycleStatusCode:                           data.ContractLifeCycleStatusCode,
			ConsistencyStatusCode:                                 data.ConsistencyStatusCode,
			ReplicationProcessingStatusCode:                       data.ReplicationProcessingStatusCode,
			ApprovalStatusCode:                                    data.ApprovalStatusCode,
			ExternalContractLifeCycleStatusCode:                   data.ExternalContractLifeCycleStatusCode,
			ExternalContractReferenceProcessingStatusCode:         data.ExternalContractReferenceProcessingStatusCode,
			ExternalContractInvoiceProcessingStatusCode:           data.ExternalContractInvoiceProcessingStatusCode,
			BuyerPartyID:                               data.BuyerPartyID,
			BuyerPartyUUID:                             data.BuyerPartyUUID,
			BuyerPartyName:                             data.BuyerPartyName,
			BuyerPartyMainContactPartyID:               data.BuyerPartyMainContactPartyID,
			BuyerPartyMainContactPartyUUID:             data.BuyerPartyMainContactPartyUUID,
			BuyerPartyMainContactPartyName:             data.BuyerPartyMainContactPartyName,
			ProductRecipientPartyID:                    data.ProductRecipientPartyID,
			ProductRecipientPartyUUID:                  data.ProductRecipientPartyUUID,
			ProductRecipientPartyName:                  data.ProductRecipientPartyName,
			BillToPartyID:                              data.BillToPartyID,
			BillToPartyUUID:                            data.BillToPartyUUID,
			BillToPartyName:                            data.BillToPartyName,
			PayerPartyID:                               data.PayerPartyID,
			PayerPartyUUID:                             data.PayerPartyUUID,
			PayerPartyName:                             data.PayerPartyName,
			SellerPartyID:                              data.SellerPartyID,
			SellerPartyUUID:                            data.SellerPartyUUID,
			SellerPartyName:                            data.SellerPartyName,
			AdministratorPartyID:                       data.AdministratorPartyID,
			AdministratorPartyEmployeeID:               data.AdministratorPartyEmployeeID,
			AdministratorPartyUUID:                     data.AdministratorPartyUUID,
			AdministratorPartyName:                     data.AdministratorPartyName,
			EmployeeResponsiblePartyID:                 data.EmployeeResponsiblePartyID,
			EmployeeResponsiblePartyEmployeeID:         data.EmployeeResponsiblePartyEmployeeID,
			EmployeeResponsiblePartyUUID:               data.EmployeeResponsiblePartyUUID,
			EmployeeResponsiblePartyName:               data.EmployeeResponsiblePartyName,
			ContractingUnitPartyID:                     data.ContractingUnitPartyID,
			ContractingUnitPartyUUID:                   data.ContractingUnitPartyUUID,
			ContractingUnitPartyName:                   data.ContractingUnitPartyName,
			SalesUnitPartyID:                           data.SalesUnitPartyID,
			SalesUnitPartyUUID:                         data.SalesUnitPartyUUID,
			SalesUnitPartyName:                         data.SalesUnitPartyName,
			SalesOrganisationID:                        data.SalesOrganisationID,
			SalesOrganisationUUID:                      data.SalesOrganisationUUID,
			SalesOrganisationName:                      data.SalesOrganisationName,
			DistributionChannelCode:                    data.DistributionChannelCode,
			DivisionCode:                               data.DivisionCode,
			SalesOfficeID:                              data.SalesOfficeID,
			SalesOfficeUUID:                            data.SalesOfficeUUID,
			SalesOfficeName:                            data.SalesOfficeName,
			SalesGroupID:                               data.SalesGroupID,
			SalesGroupUUID:                             data.SalesGroupUUID,
			SalesGroupName:                             data.SalesGroupName,
			TerritoryID:                                data.TerritoryID,
			TerritoryUUID:                              data.TerritoryUUID,
			TerritoryName:                              data.TerritoryName,
			DeliveryPriorityCode:                       data.DeliveryPriorityCode,
			IncotermsClassificationCode:                data.IncotermsClassificationCode,
			IncotermsTransferLocationName:              data.IncotermsTransferLocationName,
			CurrencyCode:                               data.CurrencyCode,
			CashDiscountTermsCode:                      data.CashDiscountTermsCode,
			PriceDateTime:                              data.PriceDateTime,
			ExternalPriceDocumentCalculationStatusCode: data.ExternalPriceDocumentCalculationStatusCode,
			ExternalPriceDocumentPricingProcedureCode:  data.ExternalPriceDocumentPricingProcedureCode,
			GrossAmount:                                data.GrossAmount,
			GrossAmountCurrencyCode:                    data.GrossAmountCurrencyCode,
			NetAmount:                                  data.NetAmount,
			NetAmountCurrencyCode:                      data.NetAmountCurrencyCode,
			InvoiceScheduleNetAmount:                   data.InvoiceScheduleNetAmount,
			InvoiceScheduleNetAmountCurrencyCode:       data.InvoiceScheduleNetAmountCurrencyCode,
			RequestExternalPricing:                     data.RequestExternalPricing,
			Transfer:                                   data.Transfer,
			EntityLastChangedOn:                        data.EntityLastChangedOn,
			ToContractExternalPriceComponent:           data.ContractExternalPriceComponent.Deferred.URI,
			ToContractItem:                             data.ContractItem.Deferred.URI,
			ToContractParty:                            data.ContractParty.Deferred.URI,
		})
	}

	return contractCollection, nil
}

func ConvertToToContractExternalPriceComponent(raw []byte, l *logger.Logger) ([]ToContractExternalPriceComponent, error) {
	pm := &responses.ToContractExternalPriceComponent{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToContractExternalPriceComponent. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toContractExternalPriceComponent := make([]ToContractExternalPriceComponent, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toContractExternalPriceComponent = append(toContractExternalPriceComponent, ToContractExternalPriceComponent{
			ObjectID:                     data.ObjectID,
			ParentObjectID:               data.ParentObjectID,
			ContractID:                   data.ContractID,
			MajorLevelOrdinalNumberValue: data.MajorLevelOrdinalNumberValue,
			MinorLevelOrdinalNumberValue: data.MinorLevelOrdinalNumberValue,
			TypeCode:                     data.TypeCode,
			Description:                  data.Description,
			CategoryCode:                 data.CategoryCode,
			RateDecimalValue:             data.RateDecimalValue,
			RateMeasureUnitCode:          data.RateMeasureUnitCode,
			RateCurrencyCode:             data.RateCurrencyCode,
			RateBaseDecimalValue:         data.RateBaseDecimalValue,
			RateBaseMeasureUnitCode:      data.RateBaseMeasureUnitCode,
			CalculatedAmount:             data.CalculatedAmount,
			CalculatedAmountCurrencyCode: data.CalculatedAmountCurrencyCode,
			EffectiveIndicator:           data.EffectiveIndicator,
			ManuallyChangedIndicator:     data.ManuallyChangedIndicator,
			InactivityReasonCode:         data.InactivityReasonCode,
			OriginCode:                   data.OriginCode,
		})
	}

	return toContractExternalPriceComponent, nil
}

func ConvertToToContractItem(raw []byte, l *logger.Logger) ([]ToContractItem, error) {
	pm := &responses.ToContractItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToContractItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toContractItem := make([]ToContractItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toContractItem = append(toContractItem, ToContractItem{
			ObjectID:                   data.ObjectID,
			ParentObjectID:             data.ParentObjectID,
			ContractID:                 data.ContractID,
			ItemID:                     data.ItemID,
			ParentItemID:               data.ParentItemID,
			ParentItemUUID:             data.ParentItemUUID,
			ProcessingTypeCode:         data.ProcessingTypeCode,
			DeterminationRelevanceCode: data.DeterminationRelevanceCode,
			ReleaseControlCode:         data.ReleaseControlCode,
			Description:                data.Description,
			ValidityStartDate:          data.ValidityStartDate,
			ValidityEndDate:            data.ValidityEndDate,
			WarrantyGoodwillCode:       data.WarrantyGoodwillCode,
			InvoiceScheduleStartDate:   data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:     data.InvoiceScheduleEndDate,
			InvoiceScheduleInvoicingInAdvanceIndicator:            data.InvoiceScheduleInvoicingInAdvanceIndicator,
			InvoiceScheduleNextInvoiceDateCalculationFunctionCode: data.InvoiceScheduleNextInvoiceDateCalculationFunctionCode,
			InvoiceScheduleHorizonDateCalculationFunctionCode:     data.InvoiceScheduleHorizonDateCalculationFunctionCode,
			CancellationReasonCode:                                data.CancellationReasonCode,
			ContractLifeCycleStatusCode:                           data.ContractLifeCycleStatusCode,
			ConsistencyStatusCode:                                 data.ConsistencyStatusCode,
			ExternalContractLifeCycleStatusCode:                   data.ExternalContractLifeCycleStatusCode,
			ExternalContractReferenceProcessingStatusCode:         data.ExternalContractReferenceProcessingStatusCode,
			ExternalContractInvoiceProcessingStatusCode:           data.ExternalContractInvoiceProcessingStatusCode,
			ProductID:                            data.ProductID,
			ProductUUID:                          data.ProductUUID,
			ProductDescription:                   data.ProductDescription,
			Quantity:                             data.Quantity,
			QuantityUnitCode:                     data.QuantityUnitCode,
			TargetQuantity:                       data.TargetQuantity,
			TargetQuantityUnitCode:               data.TargetQuantityUnitCode,
			TotalPeriodReleaseQuantity:           data.TotalPeriodReleaseQuantity,
			TotalPeriodReleaseQuantityUnitCode:   data.TotalPeriodReleaseQuantityUnitCode,
			TotalPeriodOpenQuantity:              data.TotalPeriodOpenQuantity,
			TotalPeriodOpenQuantityUnitCode:      data.TotalPeriodOpenQuantityUnitCode,
			TargetStatusCode:                     data.TargetStatusCode,
			DeliveryPriorityCode:                 data.DeliveryPriorityCode,
			IncotermsClassificationCode:          data.IncotermsClassificationCode,
			IncotermsTransferLocationName:        data.IncotermsTransferLocationName,
			CashDiscountTermsCode:                data.CashDiscountTermsCode,
			PriceDateTime:                        data.PriceDateTime,
			NetAmount:                            data.NetAmount,
			NetAmountCurrencyCode:                data.NetAmountCurrencyCode,
			NetPriceAmount:                       data.NetPriceAmount,
			NetPriceAmountCurrencyCode:           data.NetPriceAmountCurrencyCode,
			InvoiceScheduleNetAmount:             data.InvoiceScheduleNetAmount,
			InvoiceScheduleNetAmountCurrencyCode: data.InvoiceScheduleNetAmountCurrencyCode,
			ProductRecipientPartyID:              data.ProductRecipientPartyID,
			ProductRecipientPartyUUID:            data.ProductRecipientPartyUUID,
			ProductRecipientPartyFormattedName:   data.ProductRecipientPartyFormattedName,
			BillToPartyID:                        data.BillToPartyID,
			BillToPartyUUID:                      data.BillToPartyUUID,
			BillToPartyFormattedName:             data.BillToPartyFormattedName,
			PayerPartyID:                         data.PayerPartyID,
			PayerPartyUUID:                       data.PayerPartyUUID,
			PayerPartyFormattedName:              data.PayerPartyFormattedName,
			EntityLastChangedOn:                  data.EntityLastChangedOn,
			ToContractItemExternalPriceComponent: data.ContractItemExternalPriceComponent.Deferred.URI,
		})
	}

	return toContractItem, nil
}

func ConvertToToContractItemExternalPriceComponent(raw []byte, l *logger.Logger) ([]ToContractItemExternalPriceComponent, error) {
	pm := &responses.ToContractItemExternalPriceComponent{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToContractItemExternalPriceComponent. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toContractItemExternalPriceComponent := make([]ToContractItemExternalPriceComponent, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toContractItemExternalPriceComponent = append(toContractItemExternalPriceComponent, ToContractItemExternalPriceComponent{
			ObjectID:                     data.ObjectID,
			ParentObjectID:               data.ParentObjectID,
			ContractID:                   data.ContractID,
			ContractItemID:               data.ContractID,
			MajorLevelOrdinalNumberValue: data.MajorLevelOrdinalNumberValue,
			MinorLevelOrdinalNumberValue: data.MinorLevelOrdinalNumberValue,
			TypeCode:                     data.TypeCode,
			Description:                  data.Description,
			CategoryCode:                 data.CategoryCode,
			RateDecimalValue:             data.RateDecimalValue,
			RateMeasureUnitCode:          data.RateMeasureUnitCode,
			RateCurrencyCode:             data.RateCurrencyCode,
			RateBaseDecimalValue:         data.RateBaseDecimalValue,
			RateBaseMeasureUnitCode:      data.RateBaseMeasureUnitCode,
			CalculatedAmount:             data.CalculatedAmount,
			CalculatedAmountCurrencyCode: data.CalculatedAmountCurrencyCode,
			EffectiveIndicator:           data.EffectiveIndicator,
			ManuallyChangedIndicator:     data.ManuallyChangedIndicator,
			InactivityReasonCode:         data.InactivityReasonCode,
			OriginCode:                   data.OriginCode,
		})
	}

	return toContractItemExternalPriceComponent, nil
}

func ConvertToToContractParty(raw []byte, l *logger.Logger) ([]ToContractParty, error) {
	pm := &responses.ToContractParty{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToContractParty. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toContractParty := make([]ToContractParty, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toContractParty = append(toContractParty, ToContractParty{
			ObjectID:                   data.ObjectID,
			ParentObjectID:             data.ParentObjectID,
			ContractID:                 data.ContractID,
			PartyID:                    data.PartyID,
			PartyEmployeeID:            data.PartyEmployeeID,
			PartyUUID:                  data.PartyUUID,
			PartyName:                  data.PartyName,
			PartyTypeCode:              data.PartyTypeCode,
			RoleCode:                   data.RoleCode,
			MainIndicator:              data.MainIndicator,
			CountryCode:                data.CountryCode,
			StateCode:                  data.StateCode,
			CareOfName:                 data.CareOfName,
			AddressLine1:               data.AddressLine1,
			AddressLine2:               data.AddressLine2,
			HouseNumber:                data.HouseNumber,
			Street:                     data.Street,
			AddressLine4:               data.AddressLine4,
			AddressLine5:               data.AddressLine5,
			District:                   data.District,
			City:                       data.City,
			DifferentCity:              data.DifferentCity,
			StreetPostalCode:           data.StreetPostalCode,
			County:                     data.County,
			CompanyPostalCode:          data.CompanyPostalCode,
			POBoxIndicator:             data.POBoxIndicator,
			POBox:                      data.POBox,
			POBoxPostalCode:            data.POBoxPostalCode,
			POBoxDeviatingCountryCode:  data.POBoxDeviatingCountryCode,
			POBoxDeviatingStateCode:    data.POBoxDeviatingStateCode,
			POBoxDeviatingCity:         data.POBoxDeviatingCity,
			TimeZoneCode:               data.TimeZoneCode,
			Latitude:                   data.Latitude,
			Longitude:                  data.Longitude,
			Building:                   data.Building,
			Floor:                      data.Floor,
			Room:                       data.Room,
			Phone:                      data.Phone,
			Mobile:                     data.Mobile,
			Fax:                        data.Fax,
			Email:                      data.Email,
			WebSite:                    data.WebSite,
			CorrespondenceLanguageCode: data.CorrespondenceLanguageCode,
			BestReachedByCode:          data.BestReachedByCode,
		})
	}

	return toContractParty, nil
}

func ConvertToContractItemCollection(raw []byte, l *logger.Logger) ([]ContractItemCollection, error) {
	pm := &responses.ContractItemCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractItemCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	contractItemCollection := make([]ContractItemCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		contractItemCollection = append(contractItemCollection, ContractItemCollection{
			ObjectID:                   data.ObjectID,
			ParentObjectID:             data.ParentObjectID,
			ContractID:                 data.ContractID,
			ItemID:                     data.ItemID,
			ParentItemID:               data.ParentItemID,
			ParentItemUUID:             data.ParentItemUUID,
			ProcessingTypeCode:         data.ProcessingTypeCode,
			DeterminationRelevanceCode: data.DeterminationRelevanceCode,
			ReleaseControlCode:         data.ReleaseControlCode,
			Description:                data.Description,
			ValidityStartDate:          data.ValidityStartDate,
			ValidityEndDate:            data.ValidityEndDate,
			WarrantyGoodwillCode:       data.WarrantyGoodwillCode,
			InvoiceScheduleStartDate:   data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:     data.InvoiceScheduleEndDate,
			InvoiceScheduleInvoicingInAdvanceIndicator:            data.InvoiceScheduleInvoicingInAdvanceIndicator,
			InvoiceScheduleNextInvoiceDateCalculationFunctionCode: data.InvoiceScheduleNextInvoiceDateCalculationFunctionCode,
			InvoiceScheduleHorizonDateCalculationFunctionCode:     data.InvoiceScheduleHorizonDateCalculationFunctionCode,
			CancellationReasonCode:                                data.CancellationReasonCode,
			ContractLifeCycleStatusCode:                           data.ContractLifeCycleStatusCode,
			ConsistencyStatusCode:                                 data.ConsistencyStatusCode,
			ExternalContractLifeCycleStatusCode:                   data.ExternalContractLifeCycleStatusCode,
			ExternalContractReferenceProcessingStatusCode:         data.ExternalContractReferenceProcessingStatusCode,
			ExternalContractInvoiceProcessingStatusCode:           data.ExternalContractInvoiceProcessingStatusCode,
			ProductID:                            data.ProductID,
			ProductUUID:                          data.ProductUUID,
			ProductDescription:                   data.ProductDescription,
			Quantity:                             data.Quantity,
			QuantityUnitCode:                     data.QuantityUnitCode,
			TargetQuantity:                       data.TargetQuantity,
			TargetQuantityUnitCode:               data.TargetQuantityUnitCode,
			TotalPeriodReleaseQuantity:           data.TotalPeriodReleaseQuantity,
			TotalPeriodReleaseQuantityUnitCode:   data.TotalPeriodReleaseQuantityUnitCode,
			TotalPeriodOpenQuantity:              data.TotalPeriodOpenQuantity,
			TotalPeriodOpenQuantityUnitCode:      data.TotalPeriodOpenQuantityUnitCode,
			TargetStatusCode:                     data.TargetStatusCode,
			DeliveryPriorityCode:                 data.DeliveryPriorityCode,
			IncotermsClassificationCode:          data.IncotermsClassificationCode,
			IncotermsTransferLocationName:        data.IncotermsTransferLocationName,
			CashDiscountTermsCode:                data.CashDiscountTermsCode,
			PriceDateTime:                        data.PriceDateTime,
			NetAmount:                            data.NetAmount,
			NetAmountCurrencyCode:                data.NetAmountCurrencyCode,
			NetPriceAmount:                       data.NetPriceAmount,
			NetPriceAmountCurrencyCode:           data.NetPriceAmountCurrencyCode,
			InvoiceScheduleNetAmount:             data.InvoiceScheduleNetAmount,
			InvoiceScheduleNetAmountCurrencyCode: data.InvoiceScheduleNetAmountCurrencyCode,
			ProductRecipientPartyID:              data.ProductRecipientPartyID,
			ProductRecipientPartyUUID:            data.ProductRecipientPartyUUID,
			ProductRecipientPartyFormattedName:   data.ProductRecipientPartyFormattedName,
			BillToPartyID:                        data.BillToPartyID,
			BillToPartyUUID:                      data.BillToPartyUUID,
			BillToPartyFormattedName:             data.BillToPartyFormattedName,
			PayerPartyID:                         data.PayerPartyID,
			PayerPartyUUID:                       data.PayerPartyUUID,
			PayerPartyFormattedName:              data.PayerPartyFormattedName,
			EntityLastChangedOn:                  data.EntityLastChangedOn,
			ToContractItemExternalPriceComponent: data.ContractItemExternalPriceComponent.Deferred.URI,
		})
	}

	return contractItemCollection, nil
}
