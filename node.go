// Graph Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Licensed under:
// The GNU Lesser General Public License, version 3.0 (LGPL-3.0)

package graph

import (
    "sync"
    "errors"
)

// Node represents a node object.
type Node struct {
    Edges       []*Edge
    Properties  map[string]string
    lock        sync.RWMutex
    index       int
	state       int
	data        int
	parent      *Node
}

// NewNode creates a new node object
func NewNode() (n *Node) {
    n = new(Node)
    n.Edges = make([]*Edge, 0)
    n.Properties = make(map[string]string)
    n.lock = sync.RWMutex{}
    n.index = 0
    n.state = 0
    n.data = 0
    n.parent = nil
    return
}

// ParentNodes returns a slice of Nodes with output
// links to node given.
func (n *Node) ParentNodes() []*Node {
    
    out := make([]*Node, 0)
    
    if n == nil {
        return out
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    for _, edge := range n.Edges {
        if edge.ChildNode == n {
            out = append(out, edge.ParentNode)
        }
    }
    
    return out
}

// ChildNodes returns a slice of Nodes with input
// links to node given.
func (n *Node) ChildNodes() []*Node {
    
    out := make([]*Node, 0)
    
    if n == nil {
        return out
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    for _, edge := range n.Edges {
        if edge.ParentNode == n {
            out = append(out, edge.ChildNode)
        }
    }
    
    return out
}

// ParentEdges returns a slice of Edges with input
// links to node given.
func (n *Node) ParentEdges() []*Edge {
    
    out := make([]*Edge, 0)
    
    if n == nil {
        return out
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    for _, edge := range n.Edges {
        if edge.ChildNode == n {
            out = append(out, edge)
        }
    }
    
    return out
}

// ChildEdges returns a slice of Edges with output
// links to node given.
func (n *Node) ChildEdges() []*Edge {
    
    out := make([]*Edge, 0)
    
    if n == nil {
        return out
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    for _, edge := range n.Edges {
        if edge.ParentNode == n {
            out = append(out, edge)
        }
    }
    
    return out
}

// AddProperty adds a property to the node given.
func (n *Node) AddProperty(key string, value string) {
    
    if n == nil {
        return
    }
    
    n.lock.Lock()
    defer n.lock.Unlock()
    
    n.Properties[key] = value
}

// RemProperty removes a property from the node given.
func (n *Node) RemProperty(key string) {
    
    if n == nil {
        return
    }
    
    n.lock.Lock()
    defer n.lock.Unlock()
    
    delete(n.Properties, key)
}

// GetProperty returns a property from the given node.
func (n *Node) GetProperty(key string) string {
    
    if n == nil {
        return ""
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    return n.Properties[key]
}

// HasProperty returns true if given propert key exists.
func (n *Node) HasProperty(key string) bool {
    
    if n == nil {
        return false
    }
    
    n.lock.RLock()
    defer n.lock.RUnlock()
    
    if len(n.Properties[key]) > 0 {
        return true
    }
    
    return false
}

// NumLinks returns the number of parent and child links
// from the given node.
func (n *Node) NumLinks() (parent int, child int) {
    
    if n ==  nil {
        return
    }
    
    for _, edge := range n.Edges {
        if edge.ParentNode == n {
            child++
        }else if edge.ChildNode == n{
            parent++
        }
    }
    
    return
}

func (n *Node) addEdge(e *Edge) error {
    
    if n == nil || e == nil {
        return errors.New("Missing Edge and/or Node")
    }
    
    n.lock.Lock()
    defer n.lock.Unlock()
    
    n.Edges = append(n.Edges, e)
    
    return nil
}

func (n *Node) removeEdge(e *Edge) error {
    
    if n == nil || e == nil {
        return errors.New("Missing Edge and/or Node")
    }
    
    n.lock.Lock()
    defer n.lock.Unlock()
    
    eI := 0
    contains := false
    
    for i, edge := range n.Edges {
        if edge == e {
            eI = i
            contains = true
            break
        }
    }
    
    if contains == false {
        return nil
    }
    
    copy(n.Edges[eI:], n.Edges[eI+1:])
    n.Edges[len(n.Edges)-1] = nil
    n.Edges = n.Edges[:len(n.Edges)-1]
    
    return nil
}