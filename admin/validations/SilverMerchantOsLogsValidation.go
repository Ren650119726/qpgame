package validations

import (
	"qpgame/common/mvc"
)

type SilverMerchantOsLogsValidation struct{}

// 添加/修改动作数据验证
func (self SilverMerchantOsLogsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("merchant_id", "银商编号不能为空").
		Validate()
}