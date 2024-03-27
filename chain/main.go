package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}

type baseRuleChain struct {
	next RuleChain
}

// baseChain的具体方法由继承类进行实现
func (b *baseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("not implement")
}

// 返回内置的next节点，所有继承类可复用
func (b *baseRuleChain) Next() RuleChain {
	return b.next
}

// 判断next 节点是否非空，非空则执行next节点，可复用，进行相邻节点的串联
func (b *baseRuleChain) applyNext(ctx context.Context, params map[string]interface{}) error {
	if b.Next() != nil {
		return b.Next().Apply(ctx, params)
	}
	return nil
}

/* 对具体的规则节点进行定义声明 */

// 分别执行了：
// 当前节点的校验逻辑
// 倘若当前节点的校验逻辑未通过，则抛出错误，终止流程
// 倘若当前节点校验通过，执行后继节点的校验逻辑
// 针对后继节点的响应结果可以进行一定的后处理工作

type CheckTokenRule struct {
	baseRuleChain
}

func NewCheckTokenRule(next RuleChain) RuleChain {
	return &CheckTokenRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckTokenRule) Apply(ctx context.Context, parms map[string]interface{}) error {
	// 校验 token 是否合法
	token, _ := parms["token"].(string)
	if token != "myToken" {
		return fmt.Errorf("invalid token: %s", token)
	}

	if err := c.applyNext(ctx, parms); err != nil {
		fmt.Println("check token rule err post process...")
		return err
	}
	fmt.Println("check token rule common post process...")
	return nil
}

type CheckAgeRule struct {
	baseRuleChain
}

func NewCheckAgeRule(next RuleChain) RuleChain {
	return &CheckAgeRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAgeRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验 age 是否合法
	age, _ := params["age"].(int)
	if age < 18 {
		return fmt.Errorf("invalid age: %d", age)
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check age rule err post process...")
		return err
	}

	fmt.Println("check age rule common post process...")
	return nil
}

type CheckAuthorizedStatus struct {
	baseRuleChain
}

func NewCheckAuthorizedStatus(next RuleChain) RuleChain {
	return &CheckAuthorizedStatus{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAuthorizedStatus) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验是否已认证
	if authorized, _ := params["authorized"].(bool); !authorized {
		return errors.New("not authorized yet")
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check authorized status rule err post process...")
		return err
	}

	fmt.Println("check authorized statuse rule common post process...")
	return nil
}

func main() {
	// ctx := context.Background()
	checkAuthorizedRule := NewCheckAuthorizedStatus(nil)
	checkAgeRule := NewCheckAgeRule(checkAuthorizedRule)
	checkTokenRule := NewCheckTokenRule(checkAgeRule)

	if err := checkTokenRule.Apply(context.Background(), map[string]interface{}{
		"token": "myToken",
		"age":   1,
	}); err != nil {
		// 校验未通过，终止发奖流程
		log.Println(err)
		return
	}

	// 通过前置校验流程，进行奖励发放
	// sendReward(ctx, map[string]interface{}{
	// 	"token": "myToken",
	// 	"age":   1,
	// })
}
