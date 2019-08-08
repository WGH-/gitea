package validation

import (
	"regexp"
	"testing"

	"github.com/go-macaron/binding"
)

func getRegexpErrorString(pattern string) string {
	// It would be unwise to rely on that regexp
	// compilation errors don't ever change across Go releases.
	_, err := regexp.Compile(pattern)
	if err != nil {
		return err.Error()
	}
	return ""
}

var regexpValidationTestCases = []validationTestCase{
	{
		description: "Empty regexp",
		data: TestForm{
			Regexp: "",
		},
		expectedErrors: binding.Errors{},
	},
	{
		description: "Valid regexp",
		data: TestForm{
			Regexp: "(master|release)",
		},
		expectedErrors: binding.Errors{},
	},

	{
		description: "Invalid regexp",
		data: TestForm{
			Regexp: "master)(",
		},
		expectedErrors: binding.Errors{
			binding.Error{
				FieldNames:     []string{"Regexp"},
				Classification: ErrRegexp,
				Message:        getRegexpErrorString("master)("),
			},
		},
	},
}

func Test_RegexpValidation(t *testing.T) {
	AddBindingRules()

	for _, testCase := range regexpValidationTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			performValidationTest(t, testCase)
		})
	}
}
