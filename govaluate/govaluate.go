package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to "true", the bool value.
	fmt.Println(result)

	expression, err = govaluate.NewEvaluableExpression("foo > 0")
	if err != nil {
		fmt.Println(err)
		return
	}
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1

	result, err = expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to "false", the bool value.
	fmt.Println(result)

	expression, err = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	if err != nil {
		fmt.Println(err)
		return
	}
	parameters = make(map[string]interface{}, 8)
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80

	result, err = expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to "false", the bool value.
	fmt.Println(result)

	expression, err = govaluate.NewEvaluableExpression("http_response_body == 'service is ok'")
	if err != nil {
		fmt.Println(err)
		return
	}
	parameters = make(map[string]interface{}, 8)
	parameters["http_response_body"] = "service is ok"

	result, err = expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to "true", the bool value.
	fmt.Println(result)

	expression, err = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	if err != nil {
		fmt.Println(err)
		return
	}

	parameters = make(map[string]interface{}, 8)
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512

	result, err = expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to "50.0", the float64 value.
	fmt.Println(result)

	expression, err = govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err = expression.Evaluate(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result is now set to true
	fmt.Println(result)

}
