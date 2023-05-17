package main

import (
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func main() {
	dataContext := ast.NewDataContext()
	pkgs := &Package{
		Module: "Bag",
		Type:   1,
		Field:  "custom",
	}
	// Register the global function with the data context.
	dataContext.Add("Package", pkgs)
	dataContext.Add("Params", "custom")

	// Init the data-data.
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(ruleBag)))

	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err := eng1.Execute(dataContext, kb)
	if err != nil {
		fmt.Println(err)
	}
}
