package grule

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

func New(options RuleOptions) *gruleEngine {
	if grule == nil {
		grule = &gruleEngine{
			validator: &engine.GruleEngine{
				MaxCycle:                        options.MaxCycle,
				ReturnErrOnFailedRuleEvaluation: options.ReturnErrOnFailedRuleEvaluation,
			},
			knowledgeLibrary: ast.NewKnowledgeLibrary(),
		}
		grule.ruleBuilder = builder.NewRuleBuilder(grule.knowledgeLibrary)
	}

	return grule
}
