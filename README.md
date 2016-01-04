<h1>Graph<h1>
A graph library for the Go Language. Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)

<h3>License</h3>
This library is licensed under The GNU Lesser General Public License, version 3.0 (LGPL-3.0).
A copy of this license is included with the library.

<h3>Installation</h3>
First, copy the files to the src directory within your Go Root directory.
Second, build and install the library as follows:
---
go build mtmoses/graph
go install mtmoses/graph
---

<h3>Usage</h3>
Just include the graph library in your go files that access it as follows:
---
import mtmoses/graph
---

<h3>Examples</h3>
In the examples directory of the library are self-contained example files.

Create a Graph
---go
//create a graph
g := graph.NewGraph("name")
---

Create a Node
---go
node, err := g.AddNode("id", "name")
---

Insert an Edge
---go
//insert edge
g.AddEdge("id", "name", 1, nodeA, nodeB)
---

Search a Graph
---go
// Search returns the shortest path from the root node to every other
// node in the graph using the Dijkstra algorithm.
paths, err := g.Search(node)
---

Get the Distance between two Nodes
---go
// Distance returns the shrotest path between two node,
// start and finish.
distance, err := g.Distance(node_tom, node_tina)
---


