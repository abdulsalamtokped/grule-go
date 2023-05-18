package grule

import (
	"context"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func (g *gruleEngine) AppendDataContext(key string, data interface{}) ast.IDataContext {
	g.dataCtx.Add(key, data)

	return g.dataCtx
}

func (g *gruleEngine) ExecuteRule(ctx context.Context, rule string, opts ...string) error {
	var (
		ruleName    = "HMS Rules"
		ruleVersion = "0.1.1"
	)

	// Validate resource optional.
	if len(opts) > 0 {
		ruleName = opts[0]
	}
	if len(opts) > 1 {
		ruleVersion = opts[1]
	}

	// Generate the data builder.
	ruleExecuter := pkg.NewBytesResource([]byte(rule))
	if err := g.ruleBuilder.BuildRuleFromResource(ruleName, ruleVersion, ruleExecuter); err != nil {
		return err
	}
	// Config NewKnowledgeBaseInstance.
	knowledgeBaseInstance := g.knowledgeLibrary.NewKnowledgeBaseInstance(ruleName, ruleVersion)
	return g.validator.ExecuteWithContext(ctx, g.dataCtx, knowledgeBaseInstance)
}
