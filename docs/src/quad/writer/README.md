quad/writer
===========

Run examples:

```bash
    cd quad/writer
    go run export_nquads_to_stdout.go data.go
```

Convert the dot file to png, using Graphviz:

```bash
    dot -Tpng people.dot > people_dot.png
```

![people exported to DOT](./writer/people_dot.png)

Convert the graphml output to png, using the autolayout function of yEd editor.

![people exported to GraphML](./writer/people_graphml.png)
