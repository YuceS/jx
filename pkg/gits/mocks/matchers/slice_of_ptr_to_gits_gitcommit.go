// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnySliceOfPtrToGitsGitCommit() []*gits.GitCommit {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]*gits.GitCommit))(nil)).Elem()))
	var nullValue []*gits.GitCommit
	return nullValue
}

func EqSliceOfPtrToGitsGitCommit(value []*gits.GitCommit) []*gits.GitCommit {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []*gits.GitCommit
	return nullValue
}
