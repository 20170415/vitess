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

func replaceInterfaceSlice(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(InterfaceSlice)[idx] = newNode.(AST)
	}
}
func replaceLeafSlice(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(LeafSlice)[idx] = newNode.(*Leaf)
	}
}
func replaceRefOfRefContainerASTType(newNode, parent AST) {
	parent.(*RefContainer).ASTType = newNode.(AST)
}
func replaceRefOfRefContainerASTImplementationType(newNode, parent AST) {
	parent.(*RefContainer).ASTImplementationType = newNode.(*Leaf)
}
func replaceRefOfRefSliceContainerASTElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(*RefSliceContainer).ASTElements[idx] = newNode.(AST)
	}
}
func replaceRefOfRefSliceContainerASTImplementationElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(*RefSliceContainer).ASTImplementationElements[idx] = newNode.(*Leaf)
	}
}
func replaceRefOfSubImplinner(newNode, parent AST) {
	parent.(*SubImpl).inner = newNode.(SubIface)
}
func replaceRefOfValueContainerASTType(newNode, parent AST) {
	parent.(*ValueContainer).ASTType = newNode.(AST)
}
func replaceRefOfValueContainerASTImplementationType(newNode, parent AST) {
	parent.(*ValueContainer).ASTImplementationType = newNode.(*Leaf)
}
func replaceValueSliceContainerASTElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(ValueSliceContainer).ASTElements[idx] = newNode.(AST)
	}
}
func replaceValueSliceContainerASTImplementationElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(ValueSliceContainer).ASTImplementationElements[idx] = newNode.(*Leaf)
	}
}
func replaceRefOfValueSliceContainerASTElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(*ValueSliceContainer).ASTElements[idx] = newNode.(AST)
	}
}
func replaceRefOfValueSliceContainerASTImplementationElements(idx int) func(AST, AST) {
	return func(newNode, container AST) {
		container.(*ValueSliceContainer).ASTImplementationElements[idx] = newNode.(*Leaf)
	}
}
func (a *application) apply(parent, node AST, replacer replacerFunc) {
	if node == nil || isNilValue(node) {
		return
	}
	saved := a.cursor
	a.cursor.replacer = replacer
	a.cursor.node = node
	a.cursor.parent = parent
	if a.pre != nil && !a.pre(&a.cursor) {
		a.cursor = saved
		return
	}
	switch n := node.(type) {
	case Bytes:
	case InterfaceSlice:
		for x, el := range n {
			a.apply(node, el, replaceInterfaceSlice(x))
		}
	case *Leaf:
	case LeafSlice:
		for x, el := range n {
			a.apply(node, el, replaceLeafSlice(x))
		}
	case *RefContainer:
		a.apply(node, n.ASTType, replaceRefOfRefContainerASTType)
		a.apply(node, n.ASTImplementationType, replaceRefOfRefContainerASTImplementationType)
	case *RefSliceContainer:
		for x, el := range n.ASTElements {
			a.apply(node, el, replaceRefOfRefSliceContainerASTElements(x))
		}
		for x, el := range n.ASTImplementationElements {
			a.apply(node, el, replaceRefOfRefSliceContainerASTImplementationElements(x))
		}
	case *SubImpl:
		a.apply(node, n.inner, replaceRefOfSubImplinner)
	case ValueContainer:
		a.apply(node, n.ASTType, replacePanic("ValueContainer ASTType"))
		a.apply(node, n.ASTImplementationType, replacePanic("ValueContainer ASTImplementationType"))
	case *ValueContainer:
		a.apply(node, n.ASTType, replaceRefOfValueContainerASTType)
		a.apply(node, n.ASTImplementationType, replaceRefOfValueContainerASTImplementationType)
	case ValueSliceContainer:
		for x, el := range n.ASTElements {
			a.apply(node, el, replaceValueSliceContainerASTElements(x))
		}
		for x, el := range n.ASTImplementationElements {
			a.apply(node, el, replaceValueSliceContainerASTImplementationElements(x))
		}
	case *ValueSliceContainer:
		for x, el := range n.ASTElements {
			a.apply(node, el, replaceRefOfValueSliceContainerASTElements(x))
		}
		for x, el := range n.ASTImplementationElements {
			a.apply(node, el, replaceRefOfValueSliceContainerASTImplementationElements(x))
		}
	}
	if a.post != nil && !a.post(&a.cursor) {
		panic(abort)
	}
	a.cursor = saved
}
