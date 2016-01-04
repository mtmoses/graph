// Graph Library Examples - Distance
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Licensed under:
// The GNU Lesser General Public License, version 3.0 (LGPL-3.0)

package main

import (
    "fmt"
    "mtmoses/graph"
)

// just build some nodes
func buildGraph(g *graph.Graph) {
    
    //build some nodes
    //AddNode(id string, name string) (*Node, error)
    
    node_tom, _ := g.AddNode("1", "Tom")
    node_bob, _ := g.AddNode("2", "Bob")
    node_sam, _ := g.AddNode("3", "Sam")
    node_tina, _ := g.AddNode("4", "Tina")
    
    //add edges to nodes
    //(g *Graph) AddEdge(id string, name string, weight float64, in *Node, out *Node)
    g.AddEdge("1", "knows", 1, node_tom, node_bob) //Tom knows Bob
    g.AddEdge("2", "knows", 1, node_bob, node_sam) //Bob knows Sam
    g.AddEdge("3", "knows", 1, node_sam, node_tina) //Sam knows Tina
    g.AddEdge("4", "knows", 1, node_bob, node_tina) //Bob knows Tina
}

func main() {
    fmt.Println("Distance Example")
    
    //create a graph
    g := graph.NewGraph("distance test")
    buildGraph(g)
    
    //determine how far Tom is from Tina
    
    //Step 1: Get a Node to use as a starting pont
    node_tom, ok := g.GetNodeByName("Tom")
    if !ok {
        return
    }
    
    //Step 2: Get a Node as finish result
    node_tina, ok := g.GetNodeByName("Tina")
    if !ok {
        return
    }
    
    //Step 3: Get the distance between the nodes
    distance, err := g.Distance(node_tom, node_tina)
    fmt.Println("DISTANCE", distance, err)
}