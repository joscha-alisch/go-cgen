package parser

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"modernc.org/cc/v3"
	"testing"
)

type parseArgs struct {
	Config       *cc.Config
	IncludePaths []string
	SysIncludes  []string
	Sources      []cc.Source
}

type parseReturns struct {
	Ast *cc.AST
	Err error
}

type hostConfigArgs struct {
	Cpp  string
	Opts []string
}

type hostConfigReturns struct {
	Predefined      string
	IncludePaths    []string
	SysIncludePaths []string
	Err             error
}

func TestParser(t *testing.T) {
	tests := []struct {
		desc              string
		config            Config
		headers           []string
		hostConfigReturns hostConfigReturns
		expectedArgs      parseArgs
		expectedHostArgs  hostConfigArgs
		expectedErr       error
	}{
		{
			desc:              "just a header",
			config:            Config{},
			headers:           s("header.h"),
			hostConfigReturns: hostConfigReturns{},
			expectedHostArgs:  hostConfigArgs{},
			expectedArgs: parseArgs{
				Config: &cc.Config{},
				Sources: []cc.Source{
					{Name: "header.h"},
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(tt *testing.T) {
			p := New(test.config)

			pArgs := &parseArgs{}
			pReturns := parseReturns{
				Ast: &cc.AST{},
				Err: errors.New("parse error"),
			}

			parse = fakeParse(pArgs, pReturns)
			hArgs := &hostConfigArgs{}
			hostConfig = fakeHostConfig(hArgs, hostConfigReturns{})

			ast, err := p.Parse(test.headers...)
			if checkErr(err, test.expectedErr, tt) {
				return
			}

			if ast != pReturns.Ast {
				tt.Error("did not get the expected AST back\n")
			}

			if err != pReturns.Err {
				tt.Error("did not get the expected parsing error back\n")
			}

			if !cmp.Equal(*hArgs, test.expectedHostArgs) {
				tt.Errorf("hostConfig was not called with correct arguments:\n%s\n", cmp.Diff(test.expectedHostArgs, hArgs))
			}

			spew.Config.DisablePointerAddresses = true
			pArgsDump := spew.Sdump(*pArgs)
			expectedArgsDump := spew.Sdump(test.expectedArgs)

			if !cmp.Equal(pArgsDump, expectedArgsDump) {
				tt.Errorf("parse was not called with correct arguments:\n%s\n", cmp.Diff(expectedArgsDump, pArgsDump))
			}
		})
	}

}

func checkErr(err, expected error, t *testing.T) bool {
	if err == nil && expected != nil {
		t.Errorf("did not get an error but expected: %s", expected.Error())
		return true
	} else if err != nil && expected != nil && err.Error() != expected.Error() {
		t.Errorf("expected error: %s\n     got error: %s\n", expected.Error(), err.Error())
		return true
	}
	return false
}

func s(s ...string) []string {
	return s
}

func fakeParse(args *parseArgs, returns parseReturns) func(cfg *cc.Config, includePaths, sysIncludePaths []string, sources []cc.Source) (*cc.AST, error) {
	return func(cfg *cc.Config, includePaths, sysIncludePaths []string, sources []cc.Source) (*cc.AST, error) {
		args.Config = cfg
		args.IncludePaths = includePaths
		args.SysIncludes = sysIncludePaths
		args.Sources = sources

		return returns.Ast, returns.Err
	}
}

func fakeHostConfig(args *hostConfigArgs, returns hostConfigReturns) func(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
	return func(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
		args.Cpp = cpp
		args.Opts = opts

		return returns.Predefined, returns.IncludePaths, returns.SysIncludePaths, returns.Err
	}
}
