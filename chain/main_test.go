package main

import (
	"context"
	"testing"
)

func Test_RuleChainV2(t *testing.T) {
	checkAuthorizedRule := NewCheckAuthorizedStatus(nil)
	checkAgeRule := NewCheckAgeRule(checkAuthorizedRule)
	checkTokenRule := NewCheckTokenRule(checkAgeRule)

	if err := checkTokenRule.Apply(context.Background(), map[string]interface{}{
		"token": "myToken",
		"age":   1,
	}); err != nil {
		// 校验未通过，终止发奖流程
		t.Error(err)
		return
	}

	// // 通过前置校验流程，进行奖励发放
	// sendReward(ctx, params)
}
