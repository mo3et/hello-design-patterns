package factorymethod

// IRuleConfigParser IRuleConfigParser
// 抽象产品
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser
type jsonRuleConfigParser struct{}

// Parse JSON
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser
type yamlRuleConfigParser struct{}

// ParseYAML
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// IRuleConfigParserFactory 工厂方法接口 !!
// 抽象工厂 为了实例化对应工厂
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct{}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct{}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
// 进行判断 输入什么产品类型 就调用对应工厂 生成对应产品
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
