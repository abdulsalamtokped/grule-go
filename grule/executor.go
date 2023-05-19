package grule

import (
	"context"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func (g *gruleEngine) ExecuteRule(ctx context.Context, dataCtx ast.IDataContext, rule string, opts ...string) error {
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

	// Generate the rule builder.
	ruleExecuter := pkg.NewBytesResource([]byte(rule))
	if err := g.ruleBuilder.BuildRuleFromResource(ruleName, ruleVersion, ruleExecuter); err != nil {
		return err
	}
	// Config NewKnowledgeBaseInstance.
	knowledgeBaseInstance := g.knowledgeLibrary.NewKnowledgeBaseInstance(ruleName, ruleVersion)
	return g.validator.ExecuteWithContext(ctx, dataCtx, knowledgeBaseInstance)
}

func (g *gruleEngine) ExecuteRules(ctx context.Context, dataCtx ast.IDataContext, rule []string, opts ...string) error {
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

	// Generate the rule builder.
	ruleExecuter := []pkg.Resource{}
	for _, v := range rule {
		ruleExecuter = append(ruleExecuter, pkg.NewBytesResource([]byte(v)))
	}

	if err := g.ruleBuilder.BuildRuleFromResources(ruleName, ruleVersion, ruleExecuter); err != nil {
		return err
	}
	// Config NewKnowledgeBaseInstance.
	knowledgeBaseInstance := g.knowledgeLibrary.NewKnowledgeBaseInstance(ruleName, ruleVersion)
	return g.validator.ExecuteWithContext(ctx, dataCtx, knowledgeBaseInstance)
}

func (g *gruleEngine) IsRuleSucceedExecuted(dataCtx ast.IDataContext) bool {
	return dataCtx.IsComplete()
}
