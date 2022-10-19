package lib

import (
	"context"
	"regexp"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

func ValiidatorInit() {
	//校验身份证
	gvalid.RegisterRule("HKIdCard", RuleHKIdCard)
}
func RuleHKIdCard(ctx context.Context, in gvalid.RuleFuncInput) error {
	isIdCard, err := regexp.MatchString(`[a-zA-Z]{1}[0-9]{6}\([0-9]{1}\)`, in.Value.String())
	if err != nil {
		return err
	}
	if !isIdCard {
		return gerror.Newf(`invalid IdCard`)
	}
	return nil
}
