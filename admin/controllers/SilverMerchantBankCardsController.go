package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var silverMerchantBankCards = models.SilverMerchantBankCards{}
var silverMerchantBankCardsValidation = validations.SilverMerchantBankCardsValidation{}

type SilverMerchantBankCardsController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_bank_cards 银商银行卡列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商银行卡列表</strong><br />
 * 业务描述: 银商银行卡列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantBankCards
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page           页数
 * @apiParam (客户端请求参数) {int}     page_size      每页记录数
 * @apiParam (客户端请求参数) {string}  card_number    卡号
 * @apiParam (客户端请求参数) {string}  address        银行卡地址
 * @apiParam (客户端请求参数) {string}  bank_name      银行名称
 * @apiParam (客户端请求参数) {string}  name           姓名
 * @apiParam (客户端请求参数) {int} 	  merchant_id    银商用户编号
 * @apiParam (客户端请求参数) {string}  created_start  添加时间/开始
 * @apiParam (客户端请求参数) {string}  created_end    添加时间/结束
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}    	code            200, 成功
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}	     id            记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 address       银行卡地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 bank_name     银行名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 card_number   卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 created       添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}     merchant_id   银商编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}  merchant_name 银商名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}  name          用户姓名
 * @apiSuccess (data-rows每个子对象字段说明) {string}  remark        备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}     status        状态：0失效，1正常
 * @apiSuccess (data-rows每个子对象字段说明) {string}  updated       更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "address": "兰州七里河土门墩支行",
 *                "bank_name": "中国工商银行",
 *                "card_number": "6222022213146561684",
 *                "created": "2019-06-11 19:32:32",
 *                "id": "2",
 *                "merchant_id": "2",
 *                "merchant_name": "",
 *                "name": "陈海冰",
 *                "remark": "",
 *                "status": "1",
 *                "updated": "2019-06-11 19:32:32"
 *            }
 *        ],
 *        "page": 1,
 *        "page_count": 1,
 *        "total_rows": 2,
 *        "page_size": 20
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 0
 *}
 */
func (self *SilverMerchantBankCardsController) Index(ctx *Context) {
	index(ctx, &silverMerchantBankCards)
}
