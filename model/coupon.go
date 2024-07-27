package model

// 数据库实体
type Coupon struct {
	Id          int64  `gorm:"primary_key;auto_increment"`
	Username    string `gorm:"type:varchar(20); not null"` // 用户名
	CouponName  string `gorm:"type:varchar(60); not null"` // 优惠券名称
	Amount      int64  // 最大优惠券数
	Left        int64  // 剩余优惠券数
	Stock       int64  // 面额
	Description string `gorm:"type:varchar(60)"` // 优惠券描述信息
}

type ReqCoupon struct {
	Name        string
	Amount      int64
	Description string
	Stock       int64
}

type ResCoupon struct {
	Name        string `json:"name"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

// 商家查询优惠券时，返回的数据结构
type SellerResCoupon struct {
	ResCoupon
	Amount int64 `json:"amount"`
	Left   int64 `json:"left"`
}

// 顾客查询优惠券时，返回的数据结构
type CustomerResCoupon struct {
	ResCoupon
}

func ParseSellerResCoupons(coupons []Coupon) []SellerResCoupon {
	var sellerCoupons []SellerResCoupon
	for _, coupon := range coupons {
		sellerCoupons = append(sellerCoupons,
			SellerResCoupon{ResCoupon{coupon.CouponName, coupon.Stock, coupon.Description},
				coupon.Amount, coupon.Left})
	}
	return sellerCoupons
}

func ParseCustomerResCoupons(coupons []Coupon) []CustomerResCoupon {
	var sellerCoupons []CustomerResCoupon
	for _, coupon := range coupons {
		sellerCoupons = append(sellerCoupons,
			CustomerResCoupon{ResCoupon{coupon.CouponName, coupon.Stock, coupon.Description}})
	}
	return sellerCoupons
}

//
//func (p *Coupon) SubNumberOne(msg *RobbitMqService.Message) error {
//	return p.SubProductNum(msg.ProductID, msg.UserID)
//}
//
//func (p *Coupon) SubProductNum(productID int64, userId int64) error {
//	// 创建事物
//	begin := p.sqlDb.Begin()
//	// 根据商品ID查询商品
//	var product datamodels.Product
//	if err := begin.First(&product, productID).Error; err != nil {
//		begin.Rollback()
//		return errors.New("查询订单错误：" + err.Error())
//	}
//	if product.ProductNum > 0 {
//		// 扣除商品数量
//		product.ProductNum -= 1
//		if err := begin.Save(product).Error; err != nil {
//			begin.Rollback()
//			return errors.New("扣除商品数量错误：" + err.Error())
//		}
//
//		// 创建订单
//		order := &datamodels.Order{
//			UserId:      userId,
//			ProductId:   productID,
//			OrderStatus: datamodels.OrderSuccess,
//		}
//		if err := begin.Create(order).Error; err != nil {
//			begin.Rollback()
//			return errors.New("创建订单错误：" + err.Error())
//		}
//		// 无错误则提交事物
//		begin.Commit()
//		return nil
//	} else {
//		return errors.New("商品数量不足")
//	}
//}
