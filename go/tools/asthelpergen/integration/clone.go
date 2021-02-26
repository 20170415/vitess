/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by ASTHelperGen. DO NOT EDIT.

package integration

func CloneAST(in AST) AST {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case Bytes:
		return CloneBytes(in)
	case *Bytes:
		return CloneRefOfBytes(in)
	case InterfaceSlice:
		return CloneInterfaceSlice(in)
	case *InterfaceSlice:
		return CloneRefOfInterfaceSlice(in)
	case *Leaf:
		return CloneRefOfLeaf(in)
	case LeafSlice:
		return CloneLeafSlice(in)
	case *LeafSlice:
		return CloneRefOfLeafSlice(in)
	case *RefContainer:
		return CloneRefOfRefContainer(in)
	case *RefSliceContainer:
		return CloneRefOfRefSliceContainer(in)
	case *SubImpl:
		return CloneRefOfSubImpl(in)
	case ValueContainer:
		return CloneValueContainer(in)
	case *ValueContainer:
		return CloneRefOfValueContainer(in)
	case ValueSliceContainer:
		return CloneValueSliceContainer(in)
	case *ValueSliceContainer:
		return CloneRefOfValueSliceContainer(in)
	}
	// this should never happen
	return nil
}
func CloneSubIface(in SubIface) SubIface {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case *SubImpl:
		return CloneRefOfSubImpl(in)
	}
	// this should never happen
	return nil
}
func CloneBytes(n Bytes) Bytes {
	res := make(Bytes, len(n))
	copy(res, n)
	return res
}
func CloneRefOfBytes(n *Bytes) *Bytes {
	if n == nil {
		return nil
	}
	out := CloneBytes(*n)
	return &out
}
func CloneInterfaceSlice(n InterfaceSlice) InterfaceSlice {
	res := make(InterfaceSlice, len(n))
	for i, x := range n {
		res[i] = CloneAST(x)
	}
	return res
}
func CloneRefOfInterfaceSlice(n *InterfaceSlice) *InterfaceSlice {
	if n == nil {
		return nil
	}
	out := CloneInterfaceSlice(*n)
	return &out
}
func CloneRefOfLeaf(n *Leaf) *Leaf {
	if n == nil {
		return nil
	}
	out := CloneLeaf(*n)
	return &out
}
func CloneLeafSlice(n LeafSlice) LeafSlice {
	res := make(LeafSlice, len(n))
	for i, x := range n {
		res[i] = CloneRefOfLeaf(x)
	}
	return res
}
func CloneRefOfLeafSlice(n *LeafSlice) *LeafSlice {
	if n == nil {
		return nil
	}
	out := CloneLeafSlice(*n)
	return &out
}
func CloneRefOfRefContainer(n *RefContainer) *RefContainer {
	if n == nil {
		return nil
	}
	out := CloneRefContainer(*n)
	return &out
}
func CloneRefOfRefSliceContainer(n *RefSliceContainer) *RefSliceContainer {
	if n == nil {
		return nil
	}
	out := CloneRefSliceContainer(*n)
	return &out
}
func CloneRefOfSubImpl(n *SubImpl) *SubImpl {
	if n == nil {
		return nil
	}
	out := CloneSubImpl(*n)
	return &out
}
func CloneValueContainer(n ValueContainer) ValueContainer {
	return ValueContainer{
		ASTImplementationType: CloneRefOfLeaf(n.ASTImplementationType),
		ASTType:               CloneAST(n.ASTType),
		NotASTType:            n.NotASTType,
	}
}
func CloneRefOfValueContainer(n *ValueContainer) *ValueContainer {
	if n == nil {
		return nil
	}
	out := CloneValueContainer(*n)
	return &out
}
func CloneValueSliceContainer(n ValueSliceContainer) ValueSliceContainer {
	return ValueSliceContainer{
		ASTElements:               CloneSliceOfAST(n.ASTElements),
		ASTImplementationElements: CloneSliceOfRefOfLeaf(n.ASTImplementationElements),
		NotASTElements:            CloneSliceOfint(n.NotASTElements),
	}
}
func CloneRefOfValueSliceContainer(n *ValueSliceContainer) *ValueSliceContainer {
	if n == nil {
		return nil
	}
	out := CloneValueSliceContainer(*n)
	return &out
}
func CloneLeaf(n Leaf) Leaf {
	return Leaf{v: n.v}
}
func CloneRefContainer(n RefContainer) RefContainer {
	return RefContainer{
		ASTImplementationType: CloneRefOfLeaf(n.ASTImplementationType),
		ASTType:               CloneAST(n.ASTType),
		NotASTType:            n.NotASTType,
	}
}
func CloneRefSliceContainer(n RefSliceContainer) RefSliceContainer {
	return RefSliceContainer{
		ASTElements:               CloneSliceOfAST(n.ASTElements),
		ASTImplementationElements: CloneSliceOfRefOfLeaf(n.ASTImplementationElements),
		NotASTElements:            CloneSliceOfint(n.NotASTElements),
	}
}
func CloneSubImpl(n SubImpl) SubImpl {
	return SubImpl{inner: CloneSubIface(n.inner)}
}
func CloneSliceOfAST(n []AST) []AST {
	res := make([]AST, len(n))
	for i, x := range n {
		res[i] = CloneAST(x)
	}
	return res
}
func CloneSliceOfint(n []int) []int {
	res := make([]int, len(n))
	copy(res, n)
	return res
}
func CloneSliceOfRefOfLeaf(n []*Leaf) []*Leaf {
	res := make([]*Leaf, len(n))
	for i, x := range n {
		res[i] = CloneRefOfLeaf(x)
	}
	return res
}
