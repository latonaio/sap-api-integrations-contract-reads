# sap-api-integrations-contract-reads
sap-api-integrations-contract-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 契約データを取得するマイクロサービスです。    
sap-api-integrations-contract-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-contract-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/contract/overview


## 動作環境  
sap-api-integrations-contract-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-contract-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-contract-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/contract/overview   
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-contract-reads には、次の API をコールするためのリソースが含まれています。  

* ContractCollection（契約 - 契約）※契約の詳細データを取得するために、ToContractExternalPriceComponent、ToContractParty、ToContractItem、ToContractItemExternalPriceComponent、と合わせて利用されます。
* ToContractExternalPriceComponent（契約 - 契約外部価格要素 ※To）
* ToContractParty（契約 - 契約先 ※To）
* ToContractItem（契約 - 契約明細 ※To）
* ContractItemCollection（契約 - 契約明細）※契約の詳細データを取得するために、ToContractItemExternalPriceComponent、と合わせて利用されます。
* ToContractItemExternalPriceComponent（契約 - 契約明細外部価格要素 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-contract-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ContractCollection.ID（契約ID）
* inoutSDC.ContractCollection.ContractExternalPriceComponent.ContractItem.ItemID（契約明細ID）
* inoutSDC.ContractCollection.Name（契約名称）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ContractCollection" が指定されています。

```
	"api_schema": "ContractCollection",
	"accepter": ["ContractCollection"],
	"contract_code": "157",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "ContractCollection",
	"accepter": ["All"],
	"contract_code": "157",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetContract(iD, itemID, name string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ContractCollection":
			func() {
				c.ContractCollection(iD)
				wg.Done()
			}()
		case "ContractItemCollection":
			func() {
				c.ContractItemCollection(itemID)
				wg.Done()
			}()
		case "ContractName":
			func() {
				c.ContractName(name)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 契約 の 契約データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ContractParty" は、/SAP_API_Output_Formatter/type.go 内 の Type ContractCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-contract-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-contract-reads/SAP_API_Caller.(*SAPAPICaller).ContractCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1EE893D6DAF66427C5B5",
			"ID": "157",
			"BuyerID": "",
			"MainBusinessProcessVariantTypeCode": "508",
			"ProcessingTypeCode": "SRCO",
			"ContractDeterminationCoveredObjectsRelevanceCode": "",
			"TemplateIndicator": false,
			"Name": "LaserJet Contract_ALPHA",
			"CreationDateTime": "2018-05-03T18:46:27+09:00",
			"CreationIdentityUUID": "00163E0C-6CDB-1ED7-AD96-621466A5043B",
			"LastChangeDateTime": "2020-04-06T22:41:50+09:00",
			"LastChangeIdentityUUID": "00163E0C-84E2-1ED7-80DF-E5C3286E25C7",
			"Date": "2016-07-04T09:00:00+09:00",
			"ValidityStartDate": "2016-07-04T09:00:00+09:00",
			"ValidityEndDate": "2021-03-31T09:00:00+09:00",
			"TimeZoneCode": "UTC",
			"DocumentLanguageCode": "EN",
			"IncludeBusinessPartnerHierarchyIndicator": false,
			"IncludeReleaseAuthorisedPartiesIndicator": false,
			"ServiceLevelObjectiveID": "",
			"ServiceLevelObjectiveUUID": 0,
			"ServiceLevelObjectiveName": "",
			"CancellationReasonCode": "",
			"InvoiceScheduleStartDate": 0,
			"InvoiceScheduleEndDate": 0,
			"InvoiceScheduleInvoicingInAdvanceIndicator": false,
			"InvoiceScheduleHorizonDateCalculationFunctionCode": "",
			"InvoiceScheduleNextInvoiceDateCalculationFunctionCode": "",
			"ContractLifeCycleStatusCode": "4",
			"ConsistencyStatusCode": "3",
			"ReplicationProcessingStatusCode": "2",
			"ApprovalStatusCode": "2",
			"ExternalContractLifeCycleStatusCode": "",
			"ExternalContractReferenceProcessingStatusCode": "",
			"ExternalContractInvoiceProcessingStatusCode": "",
			"BuyerPartyID": "1001710",
			"BuyerPartyUUID": "00163E0C-6CDB-1EE5-8DE9-4B48F9E4BBBA",
			"BuyerPartyName": "ALPHA Center",
			"BuyerPartyMainContactPartyID": "1004188",
			"BuyerPartyMainContactPartyUUID": "00163E0C-6CDB-1ED6-84D4-217462CEE486",
			"BuyerPartyMainContactPartyName": "Lou Daly",
			"ProductRecipientPartyID": "1001710",
			"ProductRecipientPartyUUID": "00163E0C-6CDB-1EE5-8DE9-4B48F9E4BBBA",
			"ProductRecipientPartyName": "ALPHA Center",
			"BillToPartyID": "1001710",
			"BillToPartyUUID": "00163E0C-6CDB-1EE5-8DE9-4B48F9E4BBBA",
			"BillToPartyName": "ALPHA Center",
			"PayerPartyID": "1001710",
			"PayerPartyUUID": "00163E0C-6CDB-1EE5-8DE9-4B48F9E4BBBA",
			"PayerPartyName": "ALPHA Center",
			"SellerPartyID": "",
			"SellerPartyUUID": 0,
			"SellerPartyName": "",
			"AdministratorPartyID": "",
			"AdministratorPartyEmployeeID": "",
			"AdministratorPartyUUID": 0,
			"AdministratorPartyName": "",
			"EmployeeResponsiblePartyID": "",
			"EmployeeResponsiblePartyEmployeeID": "",
			"EmployeeResponsiblePartyUUID": 0,
			"EmployeeResponsiblePartyName": "",
			"ContractingUnitPartyID": "",
			"ContractingUnitPartyUUID": 0,
			"ContractingUnitPartyName": "",
			"SalesUnitPartyID": "US1100",
			"SalesUnitPartyUUID": "00163E03-A070-1EE2-88B8-D4A36EBC8D9B",
			"SalesUnitPartyName": "Sales Unit US",
			"SalesOrganisationID": "US1100",
			"SalesOrganisationUUID": "00163E03-A070-1EE2-88B8-D4A36EBC8D9B",
			"SalesOrganisationName": "Sales Unit US",
			"DistributionChannelCode": "Z1",
			"DivisionCode": "00",
			"SalesOfficeID": "",
			"SalesOfficeUUID": 0,
			"SalesOfficeName": "",
			"SalesGroupID": "",
			"SalesGroupUUID": 0,
			"SalesGroupName": "",
			"TerritoryID": "",
			"TerritoryUUID": 0,
			"TerritoryName": "",
			"DeliveryPriorityCode": "",
			"IncotermsClassificationCode": "",
			"IncotermsTransferLocationName": "",
			"CurrencyCode": "USD",
			"CashDiscountTermsCode": "ZB00",
			"PriceDateTime": "2017-06-29T09:00:00+09:00",
			"ExternalPriceDocumentCalculationStatusCode": "2",
			"ExternalPriceDocumentPricingProcedureCode": "A17001",
			"GrossAmount": "62.000000",
			"GrossAmountCurrencyCode": "USD",
			"NetAmount": "62.000000",
			"NetAmountCurrencyCode": "USD",
			"InvoiceScheduleNetAmount": "2170.000000",
			"InvoiceScheduleNetAmountCurrencyCode": "USD",
			"RequestExternalPricing": false,
			"Transfer": false,
			"EntityLastChangedOn": "2020-04-06T22:41:50+09:00",
			"ContractExternalPriceComponent": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ContractCollection('00163E0C6CDB1EE893D6DAF66427C5B5')/ContractExternalPriceComponent",
			"ContractItem": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ContractCollection('00163E0C6CDB1EE893D6DAF66427C5B5')/ContractItem",
			"ContractParty": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ContractCollection('00163E0C6CDB1EE893D6DAF66427C5B5')/ContractParty"
		}
	],
	"time": "2022-04-29T13:41:16+09:00"
}

```
