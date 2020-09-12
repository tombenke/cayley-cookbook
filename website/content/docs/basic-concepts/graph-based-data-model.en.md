---
weight: 1
bookCollapseSection: true
title: "Graph Based Data Model"
---

# Graph Based Data Model

When we build a knowledge-graph, then we build a conceptual model of a part of the real World.

A conceptual model is a model of a subject area or area of knowledge, sometimes called a domain.

A domain model typically represents:
- the primary entities (the things of the domain),
- the relationships among entities,
- the attributes and attribute values (sometimes called properties and property values) of the entities and the relationships,
- and sometimes rules that associate entities, relationships, and attributes (or all three) in more complicated ways.

The graph databases and semantic technologies can be very efficiently applied to solve problems, when we have to work with flexible, frequently changing, non-structured data models, and we have to analyze these data from different aspects. Therefore these technologies are very useful to describe and manage conceptual models.

The [graphs](https://hu.wikipedia.org/wiki/Gr%C3%A1f) are made of two basic elements:
- nodes, (or vertices).
- edges, that connect vertices.

We usually use graphs to represent knowledge. The nodes of the graphs are typically represent the things of the World, and the edges represent the associations among the things.

The edges are directed, and both the nodes and edges can have labels.

the Figure 1. shows a very simple graph, which contains two nodes, one with the "Luke" label, and another one with the label of "Leia", moreover there is an edge between them, with the label of "knows".

{{< figure src="/cayley-cookbook/simple-graph.png" title="figure 1: A simple graph" >}}

Some graph databases allows us to give more than one labels to nodes and edges in the form of property-value maps. These kind of graphs are called property-graphs. in case of {{< cayley >}}, our graphs will contain only one label per node and per edge.
