/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    barcode
	@Date    2022/2/10 10:55 上午:00
	@Desc
*/

package model

import (
	. "goa.design/goa/v3/dsl"
) //nolint:revive

var BarCode = Type("BarCode", func() {
	Description("Barcode describes the generate barcode")
})

var BarCodeRsp = Type("BarCodeRsp", func() {
	Description("BarCodeRsp describes the response of generate barcode")
	Field(1, "barcode", String, "barcode")

	Required("barcode")
})

var GenerateTokenReq = Type("GenerateTokenReq", func() {
	Description("GenerateTokenReq describes the generate token")

	Field(1, "id", Int64, "user_id", func() {
		Example(1)
		Minimum(1)
	})
	Field(1, "tenant_id", Int64, "tenantID", func() {
		Example(1)
		Minimum(1)
	})

	Required("id", "tenant_id")
})

var GenerateTokenRsp = Type("GenerateTokenRsp", func() {
	Description("GenerateTokenRsp describes the response of generate token")
	Field(1, "token", String, "token")

	Required("token")
})