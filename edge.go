// Graph Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Licensed under:
// The GNU Lesser General Public License, version 3.0 (LGPL-3.0)

package graph

import (
    "sync"
    "math"
    "errors"
)

// Edge represents an edge object.
type Edge struct {
    ParentNode  *Node
    ChildNode   *Node
    Distance    float64
    Properties  map[string]string
    lock        sync.RWMutex
}

// NewEdge creates a new edge object.
func NewEdge() (e *Edge) {
    e = new(Edge)
    e.Properties = make(map[string]string)
    e.lock = sync.RWMutex{}
    return
}

// AddProperty adds a property to the given edge.
func (e *Edge) AddProperty(key string, value string) {
    if e == nil {
        return
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    e.Properties[key] = value
}

// RemProperty removes a property from the given edge.
func (e *Edge) RemProperty(key string) {
    if e == nil {
        return
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    delete(e.Properties, key)
}

// GetProperty returns a property from the given edge.
func (e *Edge) GetProperty(key string) string {
    if e == nil {
        return ""
    }
    
    e.lock.RLock()
    defer e.lock.RUnlock()
    
    return e.Properties[key]
}

// HasProperty returns true if given key exists.
func (e *Edge) HasProperty(key string) bool {
    if e == nil {
        return false
    }
    
    e.lock.RLock()
    defer e.lock.RUnlock()
    
    if len(e.Properties[key]) > 0 {
        return true
    }
    
    return false
}

// Link connects an edge to both an input and output
// node.
func (e *Edge) Link(input *Node, output *Node) error {
    
    if e == nil || input == nil || output == nil {
        return errors.New("Missing Edge and/or Nodes")
    }
    
    e.linkTo(input, true)
    e.linkTo(output, false)
    
    output.parent = input
    
    return nil
}

// Unlink removes a connection from both an input and
// output node, then removes it.
func (e *Edge) Unlink() {
    
    if e == nil {
        return
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    input := e.ParentNode
    output := e.ChildNode
    
    if input != nil {
        input.removeEdge(e)
    }
    
    if output != nil {
        output.removeEdge(e)
    }
    
    output.parent = nil
    
    e = nil
}

// SetDistance sets the distance of an edge.
func (e *Edge) SetDistance(v float64) {
    
    if e == nil {
        return
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    if math.IsNaN(v) {
        e.Distance = 0
        return
    }
    
    e.Distance = math.Abs(v)
}

// SetWeight sets the weighted distance of an edge.
func (e *Edge) SetWeight(v float64) {
    
    if e == nil {
        return
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    if v == 0 || math.IsNaN(v) {
        e.Distance = 0
        return
    }
    
    // weight is 1 / distance
    e.Distance = math.Abs(1 / v)
}

func (e *Edge) linkTo(n *Node, isInput bool) error {
    
    if e == nil || n == nil {
        return errors.New("Missing Edge and/or Node")
    }
    
    e.lock.Lock()
    defer e.lock.Unlock()
    
    if isInput == true {
        e.ParentNode = n
    } else {
        e.ChildNode = n
    }
    
    n.addEdge(e)
    
    return nil
}