package grule

import (
	"context"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type (
	gruleExecutorResource interface {
		Execute(dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error
		ExecuteWithContext(ctx context.Context, dataCtx ast.IDataContext, knowledge *ast.KnowledgeBase) error
	}

	gruleContextResource interface {
		ResetVariableChangeCount()
		IncrementVariableChangeCount()
		HasVariableChange() bool

		Add(key string, obj interface{}) error
		AddJSON(key string, JSON []byte) error
		Get(key string) model.ValueNode
		GetKeys() []string

		Retract(key string)
		IsRetracted(key string) bool
		Complete()
		IsComplete() bool
		Retracted() []string
		Reset()
	}
	gruleBuilderResource interface {
		BuildRuleFromResource(name string, version string, resource pkg.Resource) error
		BuildRuleFromResources(name string, version string, resource []pkg.Resource) error
	}
)
