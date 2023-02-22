# ModelsRegisterCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | 单位为分 | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**CodeUrl** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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

## Methods

### NewModelsRegisterCore

`func NewModelsRegisterCore() *ModelsRegisterCore`

NewModelsRegisterCore instantiates a new ModelsRegisterCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRegisterCoreWithDefaults

`func NewModelsRegisterCoreWithDefaults() *ModelsRegisterCore`

NewModelsRegisterCoreWithDefaults instantiates a new ModelsRegisterCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ModelsRegisterCore) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelsRegisterCore) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelsRegisterCore) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelsRegisterCore) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAppName

`func (o *ModelsRegisterCore) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelsRegisterCore) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelsRegisterCore) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelsRegisterCore) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCodeUrl

`func (o *ModelsRegisterCore) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *ModelsRegisterCore) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *ModelsRegisterCore) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *ModelsRegisterCore) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetCommitHash

`func (o *ModelsRegisterCore) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ModelsRegisterCore) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ModelsRegisterCore) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ModelsRegisterCore) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsRegisterCore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsRegisterCore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsRegisterCore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsRegisterCore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetH5Url

`func (o *ModelsRegisterCore) GetH5Url() string`

GetH5Url returns the H5Url field if non-nil, zero value otherwise.

### GetH5UrlOk

`func (o *ModelsRegisterCore) GetH5UrlOk() (*string, bool)`

GetH5UrlOk returns a tuple with the H5Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH5Url

`func (o *ModelsRegisterCore) SetH5Url(v string)`

SetH5Url sets H5Url field to given value.

### HasH5Url

`func (o *ModelsRegisterCore) HasH5Url() bool`

HasH5Url returns a boolean if a field has been set.

### GetQrCodeWidth

`func (o *ModelsRegisterCore) GetQrCodeWidth() string`

GetQrCodeWidth returns the QrCodeWidth field if non-nil, zero value otherwise.

### GetQrCodeWidthOk

`func (o *ModelsRegisterCore) GetQrCodeWidthOk() (*string, bool)`

GetQrCodeWidthOk returns a tuple with the QrCodeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeWidth

`func (o *ModelsRegisterCore) SetQrCodeWidth(v string)`

SetQrCodeWidth sets QrCodeWidth field to given value.

### HasQrCodeWidth

`func (o *ModelsRegisterCore) HasQrCodeWidth() bool`

HasQrCodeWidth returns a boolean if a field has been set.

### GetQrPayMode

`func (o *ModelsRegisterCore) GetQrPayMode() string`

GetQrPayMode returns the QrPayMode field if non-nil, zero value otherwise.

### GetQrPayModeOk

`func (o *ModelsRegisterCore) GetQrPayModeOk() (*string, bool)`

GetQrPayModeOk returns a tuple with the QrPayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrPayMode

`func (o *ModelsRegisterCore) SetQrPayMode(v string)`

SetQrPayMode sets QrPayMode field to given value.

### HasQrPayMode

`func (o *ModelsRegisterCore) HasQrPayMode() bool`

HasQrPayMode returns a boolean if a field has been set.

### GetRefundState

`func (o *ModelsRegisterCore) GetRefundState() string`

GetRefundState returns the RefundState field if non-nil, zero value otherwise.

### GetRefundStateOk

`func (o *ModelsRegisterCore) GetRefundStateOk() (*string, bool)`

GetRefundStateOk returns a tuple with the RefundState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundState

`func (o *ModelsRegisterCore) SetRefundState(v string)`

SetRefundState sets RefundState field to given value.

### HasRefundState

`func (o *ModelsRegisterCore) HasRefundState() bool`

HasRefundState returns a boolean if a field has been set.

### GetReturnUrl

`func (o *ModelsRegisterCore) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *ModelsRegisterCore) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *ModelsRegisterCore) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *ModelsRegisterCore) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTimeExpire

`func (o *ModelsRegisterCore) GetTimeExpire() string`

GetTimeExpire returns the TimeExpire field if non-nil, zero value otherwise.

### GetTimeExpireOk

`func (o *ModelsRegisterCore) GetTimeExpireOk() (*string, bool)`

GetTimeExpireOk returns a tuple with the TimeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExpire

`func (o *ModelsRegisterCore) SetTimeExpire(v string)`

SetTimeExpire sets TimeExpire field to given value.

### HasTimeExpire

`func (o *ModelsRegisterCore) HasTimeExpire() bool`

HasTimeExpire returns a boolean if a field has been set.

### GetTradeNo

`func (o *ModelsRegisterCore) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ModelsRegisterCore) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ModelsRegisterCore) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ModelsRegisterCore) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetTradeProvider

`func (o *ModelsRegisterCore) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ModelsRegisterCore) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ModelsRegisterCore) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ModelsRegisterCore) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeState

`func (o *ModelsRegisterCore) GetTradeState() string`

GetTradeState returns the TradeState field if non-nil, zero value otherwise.

### GetTradeStateOk

`func (o *ModelsRegisterCore) GetTradeStateOk() (*string, bool)`

GetTradeStateOk returns a tuple with the TradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeState

`func (o *ModelsRegisterCore) SetTradeState(v string)`

SetTradeState sets TradeState field to given value.

### HasTradeState

`func (o *ModelsRegisterCore) HasTradeState() bool`

HasTradeState returns a boolean if a field has been set.

### GetTradeType

`func (o *ModelsRegisterCore) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ModelsRegisterCore) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ModelsRegisterCore) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *ModelsRegisterCore) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.

### GetTxError

`func (o *ModelsRegisterCore) GetTxError() string`

GetTxError returns the TxError field if non-nil, zero value otherwise.

### GetTxErrorOk

`func (o *ModelsRegisterCore) GetTxErrorOk() (*string, bool)`

GetTxErrorOk returns a tuple with the TxError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxError

`func (o *ModelsRegisterCore) SetTxError(v string)`

SetTxError sets TxError field to given value.

### HasTxError

`func (o *ModelsRegisterCore) HasTxError() bool`

HasTxError returns a boolean if a field has been set.

### GetTxHash

`func (o *ModelsRegisterCore) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ModelsRegisterCore) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ModelsRegisterCore) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ModelsRegisterCore) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetTxState

`func (o *ModelsRegisterCore) GetTxState() int32`

GetTxState returns the TxState field if non-nil, zero value otherwise.

### GetTxStateOk

`func (o *ModelsRegisterCore) GetTxStateOk() (*int32, bool)`

GetTxStateOk returns a tuple with the TxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxState

`func (o *ModelsRegisterCore) SetTxState(v int32)`

SetTxState sets TxState field to given value.

### HasTxState

`func (o *ModelsRegisterCore) HasTxState() bool`

HasTxState returns a boolean if a field has been set.

### GetWapUrl

`func (o *ModelsRegisterCore) GetWapUrl() string`

GetWapUrl returns the WapUrl field if non-nil, zero value otherwise.

### GetWapUrlOk

`func (o *ModelsRegisterCore) GetWapUrlOk() (*string, bool)`

GetWapUrlOk returns a tuple with the WapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapUrl

`func (o *ModelsRegisterCore) SetWapUrl(v string)`

SetWapUrl sets WapUrl field to given value.

### HasWapUrl

`func (o *ModelsRegisterCore) HasWapUrl() bool`

HasWapUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


