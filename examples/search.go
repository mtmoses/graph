// Graph Library Examples - Search
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
    fmt.Println("Search Example")
    
    //create a graph
    g := graph.NewGraph("search test")
    buildGraph(g)
    
    //Step 1: Get a Node to use as a starting pont
    node_tom, ok := g.GetNodeByName("Tom")
    if !ok {
        return
    }
    
    //Step 2: Search the Graph from choosen Node as Root
    paths, err := g.Search(node_tom)
    if err != nil {
        fmt.Println("ERROR",err)
        return
    }
    
    //Display Results
    
    // NOTE: Search determines the shortest path between all nodes in the graph
    
    //iterate over path object (Path and Weight)
    for i := range paths {
        p := paths[i]
        //get each edge withn path
        for _, e := range p.Path {
            fmt.Printf("PATH parent:%s child:%s weight:%f distance:%f \n", e.ParentNode.GetProperty("name"), e.ChildNode.GetProperty("name"), p.Weight)
        }
    }
    
}