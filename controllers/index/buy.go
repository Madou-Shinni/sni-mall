package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "xiaomi-mall/controllers"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

type BuyController struct {
	BaseController
}

/*
提交订单执行结算
   1、获取用户信息 获取用户的收货地址信息
   2、获取购买商品的信息
   3、把订单信息放在订单表，把商品信息放在商品表
   4、删除购物车里面的选中数据
   5、跳转到支付页面
*/
func (con BuyController) DoCheckout(c *gin.Context) {
	// 1、获取用户信息 获取用户的收货地址信息
	id := GetCurrentUserId(c)
	//models.Cookie.Get(c, "userinfo", &user)

	addressResult := models.Address{}
	mysql.DB.Where("uid = ? AND default_address=1", id).Find(&addressResult)

	// 2、获取购买商品的信息
	var cartList []models.Cart
	models.Cookie.Get(c, "cartList", &cartList)
	var orderList []models.Cart
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	// 3、把订单信息放在订单表，把商品信息放在商品表
	order := models.Order{
		OrderId:     utils.GetOrderId(),
		Uid:         int(id),
		AllPrice:    allPrice,
		Phone:       addressResult.Phone,
		Name:        addressResult.Name,
		Address:     addressResult.Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(utils.GetUnix()),
	}

	err := mysql.DB.Create(&order).Error
	//增加数据成功以后可以通过  order.Id
	if err == nil {
		// 把商品信息放在商品对应的订单表
		for i := 0; i < len(orderList); i++ {
			orderItem := models.OrderItem{
				OrderId:      order.Id,
				Uid:          int(id),
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			mysql.DB.Create(&orderItem)
		}
	}

	// 4、删除购物车里面的选中数据
	var noSelectCartList []models.Cart
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	models.Cookie.Set(c, "cartList", noSelectCartList)

	c.Redirect(302, "/buy/pay")
}

// Pay 支付
func (con BuyController) Pay(c *gin.Context) {
	c.String(200, "支付页面")
}

// OrderPayStatus 获取订单状态
func (con BuyController) OrderPayStatus(c *gin.Context) {

	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}
	//获取用户信息
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)

	//获取主订单信息
	order := models.Order{}
	mysql.DB.Where("id=?", id).Find(&order)

	//判断当前数据是否合法
	if user.Id != order.Uid {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "非法请求",
		})
		return
	}

	//判断是否支付
	if order.PayStatus == 1 && order.OrderStatus == 1 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "支付成功",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "支付成功",
		})
		return
	}

}
