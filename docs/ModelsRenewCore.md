# ModelsRenewCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 单位为分 | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**CnsName** | **string** |  | 
**CodeUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | **int32** |  | 
**Fuses** | Pointer to **int32** |  | [optional] 
**H5Url** | Pointer to **string** |  | [optional] 
**QrCodeWidth** | Pointer to **string** | 二维码宽度。 只有alipay，且 trade type 为 h5 模式有效，qr pay mode 为4 时有效； 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**QrPayMode** | Pointer to **string** | 支付二维码模式。 只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**RefundState** | Pointer to **string** |  | [optional] 
**ReturnUrl** | Pointer to **string** | 付款成功后的跳转链接。只有alipay，且 trade type 为 h5 模式有效; 用法参考 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene&#x3D;22 | [optional] 
**TimeExpire** | Pointer to **string** |  | [optional] 
**TradeNo** | Pointer to **string** |  | [optional] 
**TradeProvider** | Pointer to **string** |  | [optional] 
**TradeState** | Pointer to **string** |  | [optional] 
**TradeType** | Pointer to **string** |  | [optional] 
**TxError** | Pointer to **string** |  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**TxState** | Pointer to **int32** |  | [optional] 
**WapUrl** | Pointer to **string** |  | [optional] 
**WrapperExpiry** | **int32** |  | 

## Methods

### NewModelsRenewCore

`func NewModelsRenewCore(cnsName string, duration int32, wrapperExpiry int32, ) *ModelsRenewCore`

NewModelsRenewCore instantiates a new ModelsRenewCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRenewCoreWithDefaults

`func NewModelsRenewCoreWithDefaults() *ModelsRenewCore`

NewModelsRenewCoreWithDefaults instantiates a new ModelsRenewCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsRenewCore) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsRenewCore) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsRenewCore) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsRenewCore) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppName

`func (o *ModelsRenewCore) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelsRenewCore) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelsRenewCore) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelsRenewCore) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCnsName

`func (o *ModelsRenewCore) GetCnsName() string`

GetCnsName returns the CnsName field if non-nil, zero value otherwise.

### GetCnsNameOk

`func (o *ModelsRenewCore) GetCnsNameOk() (*string, bool)`

GetCnsNameOk returns a tuple with the CnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsName

`func (o *ModelsRenewCore) SetCnsName(v string)`

SetCnsName sets CnsName field to given value.


### GetCodeUrl

`func (o *ModelsRenewCore) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *ModelsRenewCore) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *ModelsRenewCore) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *ModelsRenewCore) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsRenewCore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsRenewCore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsRenewCore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsRenewCore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *ModelsRenewCore) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ModelsRenewCore) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ModelsRenewCore) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetFuses

`func (o *ModelsRenewCore) GetFuses() int32`

GetFuses returns the Fuses field if non-nil, zero value otherwise.

### GetFusesOk

`func (o *ModelsRenewCore) GetFusesOk() (*int32, bool)`

GetFusesOk returns a tuple with the Fuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuses

`func (o *ModelsRenewCore) SetFuses(v int32)`

SetFuses sets Fuses field to given value.

### HasFuses

`func (o *ModelsRenewCore) HasFuses() bool`

HasFuses returns a boolean if a field has been set.

### GetH5Url

`func (o *ModelsRenewCore) GetH5Url() string`

GetH5Url returns the H5Url field if non-nil, zero value otherwise.

### GetH5UrlOk

`func (o *ModelsRenewCore) GetH5UrlOk() (*string, bool)`

GetH5UrlOk returns a tuple with the H5Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH5Url

`func (o *ModelsRenewCore) SetH5Url(v string)`

SetH5Url sets H5Url field to given value.

### HasH5Url

`func (o *ModelsRenewCore) HasH5Url() bool`

HasH5Url returns a boolean if a field has been set.

### GetQrCodeWidth

`func (o *ModelsRenewCore) GetQrCodeWidth() string`

GetQrCodeWidth returns the QrCodeWidth field if non-nil, zero value otherwise.

### GetQrCodeWidthOk

`func (o *ModelsRenewCore) GetQrCodeWidthOk() (*string, bool)`

GetQrCodeWidthOk returns a tuple with the QrCodeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeWidth

`func (o *ModelsRenewCore) SetQrCodeWidth(v string)`

SetQrCodeWidth sets QrCodeWidth field to given value.

### HasQrCodeWidth

`func (o *ModelsRenewCore) HasQrCodeWidth() bool`

HasQrCodeWidth returns a boolean if a field has been set.

### GetQrPayMode

`func (o *ModelsRenewCore) GetQrPayMode() string`

GetQrPayMode returns the QrPayMode field if non-nil, zero value otherwise.

### GetQrPayModeOk

`func (o *ModelsRenewCore) GetQrPayModeOk() (*string, bool)`

GetQrPayModeOk returns a tuple with the QrPayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrPayMode

`func (o *ModelsRenewCore) SetQrPayMode(v string)`

SetQrPayMode sets QrPayMode field to given value.

### HasQrPayMode

`func (o *ModelsRenewCore) HasQrPayMode() bool`

HasQrPayMode returns a boolean if a field has been set.

### GetRefundState

`func (o *ModelsRenewCore) GetRefundState() string`

GetRefundState returns the RefundState field if non-nil, zero value otherwise.

### GetRefundStateOk

`func (o *ModelsRenewCore) GetRefundStateOk() (*string, bool)`

GetRefundStateOk returns a tuple with the RefundState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundState

`func (o *ModelsRenewCore) SetRefundState(v string)`

SetRefundState sets RefundState field to given value.

### HasRefundState

`func (o *ModelsRenewCore) HasRefundState() bool`

HasRefundState returns a boolean if a field has been set.

### GetReturnUrl

`func (o *ModelsRenewCore) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *ModelsRenewCore) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *ModelsRenewCore) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *ModelsRenewCore) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTimeExpire

`func (o *ModelsRenewCore) GetTimeExpire() string`

GetTimeExpire returns the TimeExpire field if non-nil, zero value otherwise.

### GetTimeExpireOk

`func (o *ModelsRenewCore) GetTimeExpireOk() (*string, bool)`

GetTimeExpireOk returns a tuple with the TimeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExpire

`func (o *ModelsRenewCore) SetTimeExpire(v string)`

SetTimeExpire sets TimeExpire field to given value.

### HasTimeExpire

`func (o *ModelsRenewCore) HasTimeExpire() bool`

HasTimeExpire returns a boolean if a field has been set.

### GetTradeNo

`func (o *ModelsRenewCore) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ModelsRenewCore) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ModelsRenewCore) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ModelsRenewCore) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetTradeProvider

`func (o *ModelsRenewCore) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ModelsRenewCore) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ModelsRenewCore) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ModelsRenewCore) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeState

`func (o *ModelsRenewCore) GetTradeState() string`

GetTradeState returns the TradeState field if non-nil, zero value otherwise.

### GetTradeStateOk

`func (o *ModelsRenewCore) GetTradeStateOk() (*string, bool)`

GetTradeStateOk returns a tuple with the TradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeState

`func (o *ModelsRenewCore) SetTradeState(v string)`

SetTradeState sets TradeState field to given value.

### HasTradeState

`func (o *ModelsRenewCore) HasTradeState() bool`

HasTradeState returns a boolean if a field has been set.

### GetTradeType

`func (o *ModelsRenewCore) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ModelsRenewCore) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ModelsRenewCore) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *ModelsRenewCore) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.

### GetTxError

`func (o *ModelsRenewCore) GetTxError() string`

GetTxError returns the TxError field if non-nil, zero value otherwise.

### GetTxErrorOk

`func (o *ModelsRenewCore) GetTxErrorOk() (*string, bool)`

GetTxErrorOk returns a tuple with the TxError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxError

`func (o *ModelsRenewCore) SetTxError(v string)`

SetTxError sets TxError field to given value.

### HasTxError

`func (o *ModelsRenewCore) HasTxError() bool`

HasTxError returns a boolean if a field has been set.

### GetTxHash

`func (o *ModelsRenewCore) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ModelsRenewCore) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ModelsRenewCore) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ModelsRenewCore) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetTxState

`func (o *ModelsRenewCore) GetTxState() int32`

GetTxState returns the TxState field if non-nil, zero value otherwise.

### GetTxStateOk

`func (o *ModelsRenewCore) GetTxStateOk() (*int32, bool)`

GetTxStateOk returns a tuple with the TxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxState

`func (o *ModelsRenewCore) SetTxState(v int32)`

SetTxState sets TxState field to given value.

### HasTxState

`func (o *ModelsRenewCore) HasTxState() bool`

HasTxState returns a boolean if a field has been set.

### GetWapUrl

`func (o *ModelsRenewCore) GetWapUrl() string`

GetWapUrl returns the WapUrl field if non-nil, zero value otherwise.

### GetWapUrlOk

`func (o *ModelsRenewCore) GetWapUrlOk() (*string, bool)`

GetWapUrlOk returns a tuple with the WapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapUrl

`func (o *ModelsRenewCore) SetWapUrl(v string)`

SetWapUrl sets WapUrl field to given value.

### HasWapUrl

`func (o *ModelsRenewCore) HasWapUrl() bool`

HasWapUrl returns a boolean if a field has been set.

### GetWrapperExpiry

`func (o *ModelsRenewCore) GetWrapperExpiry() int32`

GetWrapperExpiry returns the WrapperExpiry field if non-nil, zero value otherwise.

### GetWrapperExpiryOk

`func (o *ModelsRenewCore) GetWrapperExpiryOk() (*int32, bool)`

GetWrapperExpiryOk returns a tuple with the WrapperExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperExpiry

`func (o *ModelsRenewCore) SetWrapperExpiry(v int32)`

SetWrapperExpiry sets WrapperExpiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


