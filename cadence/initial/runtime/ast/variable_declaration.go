/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import "github.com/onflow/cadence-initial/runtime/common"

type VariableDeclaration struct {
	Access            Access
	IsConstant        bool
	Identifier        Identifier
	TypeAnnotation    *TypeAnnotation
	Value             Expression
	Transfer          *Transfer
	StartPos          Position
	SecondTransfer    *Transfer
	SecondValue       Expression
	ParentIfStatement *IfStatement
}

func (d *VariableDeclaration) StartPosition() Position {
	return d.StartPos
}

func (d *VariableDeclaration) EndPosition() Position {
	return d.Value.EndPosition()
}

func (*VariableDeclaration) isIfStatementTest() {}

func (*VariableDeclaration) isDeclaration() {}

func (*VariableDeclaration) isStatement() {}

func (d *VariableDeclaration) Accept(visitor Visitor) Repr {
	return visitor.VisitVariableDeclaration(d)
}

func (d *VariableDeclaration) DeclarationIdentifier() *Identifier {
	return &d.Identifier
}

func (d *VariableDeclaration) DeclarationKind() common.DeclarationKind {
	if d.IsConstant {
		return common.DeclarationKindConstant
	}
	return common.DeclarationKindVariable
}

func (d *VariableDeclaration) DeclarationAccess() Access {
	return d.Access
}
