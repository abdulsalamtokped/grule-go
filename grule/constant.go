package grule

import "github.com/hyperjumptech/grule-rule-engine/ast"

type (
	gruleEngine struct {
		validator        gruleExecutorResource
		ruleBuilder      gruleBuilderResource
		knowledgeLibrary *ast.KnowledgeLibrary
	}

	RuleOptions struct {
		MaxCycle                        uint64
		ReturnErrOnFailedRuleEvaluation bool
	}
)

var grule *gruleEngine
