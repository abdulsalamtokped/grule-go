package main

import (
	"context"
	"fmt"

	"github.com/grule-go/grule"
)

func main() {
	// Data that being evaluated.
	pkgs := &Package{
		Module: "Bag",
		Type:   1,
		Field:  "custom",
	}

	// Register the global function with the data context.
	// dataContext := ast.NewDataContext()
	// dataContext.Add("Package", pkgs)
	// dataContext.Add("Module", "Bag")
	// dataContext.Add("Type", 1)

	// // Init the data-data.
	// lib := ast.NewKnowledgeLibrary()
	// ruleBuilder := builder.NewRuleBuilder(lib)
	// ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rulePlain)))
	// // ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.Ne÷wBytesResource([]byte(ruleBag)))
	// // Config.
	// kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	// eng1 := &engine.GruleEngine{MaxCycle: 1, ReturnErrOnFailedRuleEvaluation: true}

	// // See matching rules.
	// _, err := eng1.FetchMatchingRules(dataContext, kb)

	// // Execute.
	// err = eng1.Execute(dataContext, kb)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	p := grule.New(grule.GruleOptions{MaxCycle: 1, ReturnErrOnFailedRuleEvaluation: true})
	p.AppendDataContext("Package", pkgs)
	p.AppendDataContext("Module", "Bag")
	p.AppendDataContext("Type", 1)

	err := p.ExecuteRule(context.Background(), rulePlain)
	if err != nil {
		fmt.Println("err")
	}
}
