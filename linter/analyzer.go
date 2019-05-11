package linter

import (
	"fmt"
	"github.com/iamabhishek-dubey/dockerlinter/linter/rules"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"io/ioutil"
	"os"
	"text/template"
)

// Analyzer implements Analyzer.
type Analyzer struct {
	rules []*rules.Rule
}

type PageData struct {
	Text string
}

// NewAnalyzer generate a NewAnalyzer with rules to apply
func NewAnalyzer(ignoreRules []string) Analyzer {
	return newAnalyzer(ignoreRules)
}

func newAnalyzer(ignoreRules []string) Analyzer {
	var filteredRules []*rules.Rule
	for _, k := range getMakeDifference(rules.RuleKeys, ignoreRules) {
		if rule, ok := rules.Rules[k]; ok {
			filteredRules = append(filteredRules, rule)
		}
	}
	return Analyzer{rules: filteredRules}
}

// Run apply docker best practice rules to docker ast
func (a Analyzer) Run(node *parser.Node, filePath string) ([]string, error) {
	var rst []string

	f, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("create file: ", err)
	}
	f.Close()

	t, err := os.Create("temp.csv")
	if err != nil {
		fmt.Println("create file: ", err)
	}
	t.Close()

	csv_file, err := os.OpenFile("temp.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	header := "Line Number,Line,Rule Code,Description"
    csv_file.WriteString(header)

	rstChan := make(chan []string, len(a.rules))
	errChan := make(chan error, len(a.rules))

	for _, rule := range a.rules {
		go func(r *rules.Rule) {
			vrst, err := r.ValidateFunc.(func(*parser.Node) ([]rules.ValidateResult, error))(node)
			if err != nil {
				errChan <- err
			} else {
				rstChan <- rules.CreateMessage(rule, vrst, filePath)
			}
		}(rule)
		select {
		case value := <-rstChan:
			rst = append(rst, value...)
		case err := <-errChan:
			return nil, err
		}
	}

	content, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(content)

	htdata := PageData{
		Text: str,
	}

	f, err = os.Create("result.html")
	if err != nil {
		fmt.Println("create file: ", err)
	}

	tmpl := template.Must(template.New("Docker Lineter Template").Parse(htmltemplate))
	tmpl.Execute(f, htdata)
	f.Close()

	os.Remove("temp.txt")

	fmt.Println("")
	fmt.Println("The report file is generated in result.html")

	return rst, nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
// getMakeDifference is a function to create a difference set
func getMakeDifference(xs, ys []string) []string {
	if len(xs) > len(ys) {
		return makeDifference(xs, ys)
	}
	return makeDifference(ys, xs)
}

// make set difference
func makeDifference(xs, ys []string) []string {
	var set []string
	for _, c := range xs {
		if !isContain(ys, c) {
			set = append(set, c)
		}
	}
	return set
}

// isContain is a function to check if s is in xs
func isContain(xs []string, s string) bool {
	for _, x := range xs {
		if s == x {
			return true
		}
	}
	return false
}
