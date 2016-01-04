// Graph Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Licensed under:
// The GNU Lesser General Public License, version 3.0 (LGPL-3.0)

package graph

import (
    "errors"
)

const (
	dequeued = ^(1<<31 - 1)
)

// Distance returns the shrotest path between two node,
// start and finish. All edges must have a positive weight,
// otherwise this function will return an error.
func (g *Graph) Distance(start *Node, finish *Node) (distance float64, err error) {
    
    if g == nil {
        return distance, errors.New("Graph is empty or nil")
    }
    if start == nil || finish == nil {
        return distance, errors.New("Start and finish nodes required")
    }
    
    nodesBase := nodeSlice(make([]*Node, len(g.nodes)))
    copy(nodesBase, g.nodes)
	for i := range nodesBase {
		nodesBase[i].state = 1<<31 - 1
		nodesBase[i].data = i
	}
    
    start.state = 0
    nodes := &nodesBase
    nodes.Init()
    
    for len(*nodes) > 0 {
        
        current := nodes.pop()
        
        for _, edge := range current.Edges {
            weight := float64(current.state) + edge.Distance
            if weight < float64(current.state) {
                return distance, errors.New("Negative edge length")
            }
            v := edge.ChildNode
            if nodes.Contains(v) && weight < float64(v.state) {
                v.parent = current
                nodes.update(v.data, int(weight))
            }
        }
        
        if current == finish {
            //found finish so stop processing
            return distance, nil
        }
        
        if current.parent != nil {
            distance = distance + float64(current.state - current.parent.state)
            if current.parent == finish {
                //found finish so stop processing
                return distance, nil
            }
        }
        
    }
    
    return 0, errors.New("Finish node not reachable from start node")
}

// FilerPath returns the filered path based on a passed function parameter.
// The function parameter is an iteratee function that must return true for
// a path to be included.
func (g *Graph) FilterPath(paths []Path, f func(Path) bool) ([]Path, error) {
    
    filtered_path := make([]Path, 0)
    
    if g == nil {
        return paths, errors.New("Graph is empty or nil")
    }
    if len(paths) == 0 {
        return paths, errors.New("paths required")
    }
    if f == nil {
        return paths, errors.New("Function required")   
    }
    
    for i := range paths {
        current_path := paths[i]
        if f(current_path) == true {
            filtered_path = append(filtered_path, current_path)
        }
    }
    
    return filtered_path, nil
    
}

// Search returns the shortest path from the root node to every other
// node in the graph using the Dijkstra algorithm. All edges must have
// a positive weight, otherwise this function will return an error.
func (g *Graph) Search(root *Node) ([]Path, error) {
    
    paths := make([]Path, len(g.nodes))
    
    if g == nil {
        return paths, errors.New("Graph is empty or nil")
    }
    if root == nil {
        return paths, errors.New("Root node required")
    }
    
    nodesBase := nodeSlice(make([]*Node, len(g.nodes)))
    copy(nodesBase, g.nodes)
	for i := range nodesBase {
		nodesBase[i].state = 1<<31 - 1
		nodesBase[i].data = i
	}
    
    root.state = 0
    nodes := &nodesBase
    nodes.Init()
    
    for len(*nodes) > 0 {
        
        current := nodes.pop()
        
        // range over edges to get child nodes
        for _, edge := range current.Edges {
            
            //check for negative 
            weight := float64(current.state) + edge.Distance
            if weight < float64(current.state) {
                return paths, errors.New("Negative edge length")
            }
            
            v := edge.ChildNode
            
            if nodes.Contains(v) && weight < float64(v.state) {
                v.parent = current
                nodes.update(v.data, int(weight))
            }
            
        }
        
        // build path to this node
		if current.parent != nil {
        
            var path Path
            path.Weight = float64(current.state)
            path.Path = make([]Edge, 0)
            copy(path.Path, paths[current.parent.index].Path)
            
            var edge Edge
            edge.Distance = float64(current.state - current.parent.state)
            edge.ParentNode = current.parent
            edge.ChildNode = current
            path.Path = append(path.Path, edge)
            paths = append(paths, path)
        
		} else {
			paths[current.index] = Path{Weight: float64(current.state), Path: []Edge{}}
		}
        
    }
    
    return paths, nil
}
