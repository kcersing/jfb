package jufenbei

//聚分呗 4.0

var Token = ""
var Sign = ""
var Key = ""

// NotifyUrl 消息通知
var NotifyUrl = "/jufenbei/notice/url" //本地

const (
	URL = "https://jiny.dock.xin/op/alipay"

	// MerchantList 商户列表
	MerchantList = URL + "/merchant/list"
	// MerchantClaim 认领/关联商家
	MerchantClaim = URL + "/merchant/claim"

	// ProductAdd 创建商品
	ProductAdd = URL + "/merchant/product/add"
	// ProductEdit 编辑商品
	ProductEdit = URL + "/merchant/product/edit"

	// OrderCreate 创建订单
	OrderCreate = URL + "/order/create"
	// OrderClose 解约订单
	OrderClose = URL + "/order/action/close"
	// OrderExpire 订单延期
	OrderExpire = URL + "/order/action/expire"

	// SetNotifyUrl 设置通知地址
	SetNotifyUrl = URL + "/set/notice/url"
	// GetNotifyUrl 获取通知地址
	GetNotifyUrl = URL + "/get/notice/url"
)
