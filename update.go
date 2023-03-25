// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package goabu

import (
	"fmt"
	"reflect"

	"github.com/abu-lang/goabu/ecarule"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type Update []Assignment

type Assignment struct {
	Resource string
	variable *ast.Variable
	Value    reflect.Value
}

func (a Assignment) String() string {
	return fmt.Sprintf("(%s,%v)", a.Resource, a.Value)
}

func appendNonempty(pool []Update, u Update) []Update {
	if len(u) == 0 {
		return pool
	}
	return append(pool, u)
}

func evalActions(actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) (Update, error) {
	res := make([]Assignment, 0)
	for _, action := range actions {
		assignment := action.Assignment
		variable := assignment.Variable
		rexpr := assignment.Expression
		rexpr = workingMemory.AddExpression(rexpr)
		exprVal, err := rexpr.Evaluate(dataContext, workingMemory)
		if err != nil {
			return nil, err
		}
		res = append(res, Assignment{
			Resource: action.Resource,
			variable: variable,
			Value:    exprVal,
		})
	}
	return res, nil
}

func condEvalActions(exp *ast.Expression, actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) (Update, error) {
	exp = workingMemory.AddExpression(exp)
	val, err := exp.Evaluate(dataContext, workingMemory)
	if err != nil {
		return nil, err
	}
	if val.Bool() {
		return evalActions(actions, dataContext, workingMemory)
	}
	return nil, nil
}

//----------------------------------LOGGER------------------------------------

// arrayMarshalerHolder is an interface for objects that can provide a [zapcore.ArrayMarshaler]
// in order to be encoded as arrays by a [*zap.Logger].
type arrayMarshalerHolder interface {
	// arrayMarshaler returns a [zapcore.ArrayMarshaler] for encoding the receiver as an array.
	arrayMarshaler() zapcore.ArrayMarshaler
}

// zapUpdate constructs [zapcore.Field] with the given [arrayMarshalerHolder].
func zapUpdate(key string, h arrayMarshalerHolder) zapcore.Field {
	return zap.Array(key, h.arrayMarshaler())
}

// zapUpdates constructs [zapcore.Field] with the given [[]arrayMarshalerHolder].
func zapUpdates[T arrayMarshalerHolder](key string, hs []T) zapcore.Field {
	return zap.Array(key, newArrayMarshaler(hs...))
}

// newArrayMarshaler creates a [zapcore.ArrayMarshaler] for encoding a [[]arrayMarshalerHolder].
func newArrayMarshaler[T arrayMarshalerHolder](hs ...T) zapcore.ArrayMarshaler {
	return zapcore.ArrayMarshalerFunc(
		func(enc zapcore.ArrayEncoder) error {
			for _, h := range hs {
				err := enc.AppendArray(h.arrayMarshaler())
				if err != nil {
					return err
				}
			}
			return nil
		})
}

// arrayMarshaler returns a [zapcore.ArrayMarshaler] for encoding the receiver as an array.
func (a Assignment) arrayMarshaler() zapcore.ArrayMarshaler {
	return zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		enc.AppendString(a.variable.Variable.Name)
		enc.AppendString(a.Resource)
		enc.AppendReflected(a.Value.Interface())
		return nil
	})
}

// arrayMarshaler returns a [zapcore.ArrayMarshaler] for encoding the receiver as an array.
func (u Update) arrayMarshaler() zapcore.ArrayMarshaler {
	return newArrayMarshaler(u...)
}
