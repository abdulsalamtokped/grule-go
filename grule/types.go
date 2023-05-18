package grule

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type (
	gruleEngine struct {
		validator        *engine.GruleEngine
		dataCtx          ast.IDataContext
		knowledgeLibrary *ast.KnowledgeLibrary
		ruleBuilder      *builder.RuleBuilder
	}

	GruleOptions struct {
		MaxCycle                        uint64
		ReturnErrOnFailedRuleEvaluation bool
	}
)

var grule *gruleEngine
