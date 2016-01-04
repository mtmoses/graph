// Graph Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Licensed under:
// The GNU Lesser General Public License, version 3.0 (LGPL-3.0)

package graph

import (
    "errors"
)

type Graph struct {
    nodes []*Node
    id    string
}

type Path struct {
	Weight float64
	Path  []Edge
}

// NewGraph generates a new graph object with empty nodes.
func NewGraph(id string) (g *Graph) {
    g = new(Graph)
    g.nodes = make([]*Node, 0)
    g.id = id
    return g
}

// NumNodes returns the number of nodes in a graph.
func (g *Graph) NumNodes() int {
    if g == nil {
        return 0
    }
    return len(g.nodes)
}

// AddNode creates a new node object and adds it to the graph.
// A pointer to the newly created node is returned.
func (g *Graph) AddNode(id string, name string) (*Node, error) {
    if g == nil {
        return nil, errors.New("Graph is nil")
    }
    if len(id) == 0 || len(name) == 0 {
        return nil, errors.New("Id and name required")
    }
    
    //build node
    n := NewNode()
    n.AddProperty("id", id)
    n.AddProperty("name", name)
    n.index = len(g.nodes)
    //add to graph
    g.nodes = append(g.nodes, n)
    
    return n, nil
}

// AddEdge creates a new edge and adds it to the input and output nodes
func (g *Graph) AddEdge(id string, name string, weight float64, in *Node, out *Node) error {
    if g == nil {
        return errors.New("Graph is nil")
    }
    if in == nil || out == nil {
        return errors.New("An edge requires two nodes")
    }
    if len(id) == 0 || len(name) == 0 {
        return errors.New("Id and name required")
    }
    
    //build edge
    e := NewEdge()
    e.AddProperty("id", id)
    e.AddProperty("name", name)
    e.SetWeight(weight)
    
    //set link and return any error
    return e.Link(in, out)
}

// GetNodeById returns a node form the graph by id
// or false if not found.
func (g *Graph) GetNodeById(id string) (*Node, bool) {
    if g == nil {
        return nil, false
    }
    if len(id) == 0 {
        return nil, false
    }
    
    var n *Node
    var isFound bool
    
    for _, node := range g.nodes {
        if node.GetProperty("id") == id {
            n = node
            isFound = true
            break
        }
    }
    
    return n, isFound
}

// GetNodeById returns a node form the graph by name
// or false if not found.
func (g *Graph) GetNodeByName(name string) (*Node, bool) {
    if g == nil {
        return nil, false
    }
    if len(name) == 0 {
        return nil, false
    }
    
    var n *Node
    var isFound bool
    
    for _, node := range g.nodes {
        if node.GetProperty("name") == name {
            n = node
            isFound = true
            break
        }
    }
    
    return n, isFound
}

// GetNodeById returns a node form the graph by a custom property
// or false if not found.
func (g *Graph) GetNodeByProperty(key string, value string) (*Node, bool) {
    if g == nil {
        return nil, false
    }
    if len(key) == 0 || len(value) == 0 {
        return nil, false
    }
    
    var n *Node
    var isFound bool
    
    for _, node := range g.nodes {
        if node.GetProperty(key) == value {
            n = node
            isFound = true
            break
        }
    }
    
    return n, isFound
}

// RemoveNode removes a node from the graph.
// Any edges connected to node are also removed.
func (g *Graph) RemoveNode(n *Node) {
    if g == nil || n == nil {
        return
    }
    
    //remove edges to avoid memory issues later on
    for _, edge := range n.Edges {
        if edge != nil {
            edge.Unlink()
        }
    }
    
    gI := 0
    contains := false
    
    //get index of node to remove
    for i, node := range g.nodes {
        if node == n {
            gI = i
            contains = true
            break
        }
    }
    
    if contains == false {
        return
    }
    
    //remove node from graph
    copy(g.nodes[gI:], g.nodes[gI+1:])
    g.nodes[len(g.nodes)-1] = nil
    g.nodes = g.nodes[:len(g.nodes)-1]
    
    //destroy node
    n = nil
    
    return
}
