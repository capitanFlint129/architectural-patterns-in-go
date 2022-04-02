package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	parseStringTestCaseName                 = "Parse string"
	parseStringWithoutNewLineTestCaseName   = "Parse string without new line on end"
	parseEmptyStringTestCaseName            = "Parse empty string"
	parseStringWithoutDelimiterTestCaseName = "Parse string without delimiter"
	parseTrimWhitespacesTestCaseName        = "Parse string with whitespaces"
)

type inputData struct {
	delimiter   string
	inputString string
}

type expectedResult struct {
	parsedString []string
}

func Test_Parser(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: parseStringTestCaseName,
			inputData: inputData{
				delimiter:   "|",
				inputString: "cmd_1 | cmd_2\n",
			},
			expectedResult: expectedResult{
				parsedString: []string{"cmd_1", "cmd_2"},
			},
		},
		{
			testCaseName: parseStringWithoutNewLineTestCaseName,
			inputData: inputData{
				delimiter:   "|",
				inputString: "cmd_1 | cmd_2",
			},
			expectedResult: expectedResult{
				parsedString: []string{"cmd_1", "cmd_2"},
			},
		},
		{
			testCaseName: parseEmptyStringTestCaseName,
			inputData: inputData{
				delimiter:   "|",
				inputString: "",
			},
			expectedResult: expectedResult{
				parsedString: []string{},
			},
		},
		{
			testCaseName: parseStringWithoutDelimiterTestCaseName,
			inputData: inputData{
				delimiter:   "|",
				inputString: "cmd_1\n",
			},
			expectedResult: expectedResult{
				parsedString: []string{"cmd_1"},
			},
		},
		{
			testCaseName: parseTrimWhitespacesTestCaseName,
			inputData: inputData{
				delimiter:   "|",
				inputString: "  \t cmd_1   \n",
			},
			expectedResult: expectedResult{
				parsedString: []string{"cmd_1"},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			parser := NewParser(testData.inputData.delimiter)
			result := parser.Parse(testData.inputData.inputString)

			assert.Equal(t, testData.expectedResult.parsedString, result)
		})
	}
}
