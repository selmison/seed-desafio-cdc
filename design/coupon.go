package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var ApplyCouponDTO = Type("ApplyCouponDTO", func() {
	Description("Apply Coupon Type")
	Attribute("code", String)
	Attribute("purchase_id", String)
	Required("code", "purchase_id")
})

var CouponDTO = Type("CouponDTO", func() {
	Description("Coupon Type")
	Attribute("id", String)
	Attribute("code", String)
	Attribute("discount", Float32, func() {
		Minimum(0)
	})
	Attribute("validity", String)
	Required("id", "code", "discount", "validity")
})

var CreateCouponDTO = Type("CreateCouponDTO", func() {
	Description("New Coupon Type")
	Attribute("code", String)
	Attribute("discount", Float32, func() {
		Minimum(0)
	})
	Attribute("validity", String)
	Required("code", "discount", "validity")
})
