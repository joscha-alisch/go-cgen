package extractor

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"modernc.org/cc/v3"
)

func TestExtractor(t *testing.T) {
	tests := []struct {
		desc        string
		header      string
		expected    Definition
		expectedErr error
	}{
		{
			desc:   "enum",
			header: "enum.h",
			expected: Definition{
				Identifiers: Identifiers{
					"MyEnum":  "MyEnum",
					"Option1": "Option1",
					"Option2": "Option2",
				},
				Enums: []Enum{
					{"MyEnum", []EnumOption{
						{"Option1", Value{ValueTypeLiteral, 1}},
						{"Option2", Value{ValueTypeLiteral, 2}},
					}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(tt *testing.T) {
			ast := parseHeader(test.header)
			e := New()
			def, err := e.Extract(ast)
			if checkErr(err, test.expectedErr, tt) {
				return
			}

			if !cmp.Equal(test.expected, def) {
				tt.Errorf("\ndiff: \n%s", cmp.Diff(test.expected, def))
			}
		})
	}
}

func parseHeader(header string) *cc.AST {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	testFilePath := filepath.Join(dir, "testdata", header)

	ast, err := cc.Parse(&cc.Config{}, nil, nil, []cc.Source{{Name: testFilePath}})
	if err != nil {
		panic(fmt.Errorf("could not parse header '%s': %s", header, err.Error()))
	}
	return ast
}

func checkErr(err, expected error, t *testing.T) bool {
	if err == nil && expected != nil {
		t.Errorf("did not get an error but expected: %s", expected.Error())
		return true
	} else if err != nil && expected == nil {
		t.Errorf("did not expect an error, but got: %s", err.Error())
		return true
	} else if err != nil && expected != nil && err.Error() != expected.Error() {
		t.Errorf("expected error: %s\n     got error: %s\n", expected.Error(), err.Error())
		return true
	}
	return false
}
