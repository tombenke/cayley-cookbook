// Package graphml provides an encoder for GraphML format
package yed_graphml

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/cayleygraph/quad"
)

func init() {
	quad.RegisterFormat(quad.Format{
		Name:   "graphml",
		Ext:    []string{".graphml"},
		Mime:   []string{"application/xml"},
		Writer: func(w io.Writer) quad.WriteCloser { return NewWriter(w) },
	})
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

type Writer struct {
	w       io.Writer
	written bool
	err     error

	nodes map[string]int
	cur   int
}

func (w *Writer) writeNode(s string, color string) int {
	if w.err != nil {
		return -1
	}
	i, ok := w.nodes[s]
	if ok {
		return i
	}
	i = w.cur
	w.cur++
	w.nodes[s] = i
	_, w.err = fmt.Fprintf(w.w, "\t\t<node id=\"n%d\"><data key=\"d0\">", i)
	if w.err != nil {
		return -1
	}
	_, w.err = fmt.Fprintf(w.w, `
            <y:ShapeNode>  
              <y:Geometry x="170.5" y="-15.0" width="%d.0" height="30.0"/>  
              <y:Fill color="%s" transparent="false"/>  
              <y:BorderStyle type="line" width="1.0" color="#000000"/>  
              <y:NodeLabel>`, len(s)*10, color)
	if w.err != nil {
		return -1
	}
	if w.err = xml.EscapeText(w.w, []byte(s)); w.err != nil {
		return -1
	}
	_, w.err = fmt.Fprintf(w.w, `</y:NodeLabel>
              <y:Shape type="ellipse"/>  
            </y:ShapeNode>  
    `)
	if w.err != nil {
		return -1
	}
	if _, w.err = w.w.Write([]byte("</data></node>\n")); w.err != nil {
		return -1
	}
	return i
}

func getColorByValue(v quad.Value) string {
	switch v.(type) {
	case quad.IRI:
		// IRI
		return "#FFFF99"

	case quad.BNode:
		//Blank Node
		return "#E1E1E1"

	default:
		// Literal value
		return "#CCFFCC"
	}
}

func (w *Writer) WriteQuad(q quad.Quad) error {
	if w.err != nil {
		return w.err
	} else if !q.IsValid() {
		return quad.ErrInvalid
	}
	if !w.written {
		if _, err := w.w.Write([]byte(header)); err != nil {
			return err
		}
		w.written = true
		w.nodes = make(map[string]int)
	}
	s := w.writeNode(q.Subject.String(), getColorByValue(q.Subject))
	o := w.writeNode(q.Object.String(), getColorByValue(q.Object))
	if w.err != nil {
		return w.err
	}
	_, w.err = fmt.Fprintf(w.w, "\t\t<edge source=\"n%d\" target=\"n%d\"><data key=\"d1\">", s, o)
	if w.err != nil {
		return w.err
	}
	_, w.err = fmt.Fprintf(w.w, `
    <y:PolyLineEdge>  
          <y:Path sx="0.0" sy="-15.0" tx="29.5" ty="0.0">  
            <y:Point x="425.0" y="0.0"/>  
          </y:Path>  
          <y:LineStyle type="line" width="1.0" color="#000000"/>  
          <y:Arrows source="none" target="standard"/>  
          <y:EdgeLabel alignment="center" backgroundColor="#FFFFFF" configuration="AutoFlippingLabel" distance="2.0" fontFamily="Dialog" fontSize="12" fontStyle="plain" hasLineColor="false" height="17.96875" horizontalTextPosition="center" iconTextGap="4" modelName="centered" preferredPlacement="anywhere" ratio="0.5" textColor="#000000" verticalTextPosition="bottom" visible="true">`)
	if w.err != nil {
		return w.err
	}
	if w.err = xml.EscapeText(w.w, []byte(q.Predicate.String())); w.err != nil {
		return w.err
	}
	_, w.err = fmt.Fprintf(w.w, `</y:EdgeLabel>  
          <y:BendStyle smoothed="false"/>  
        </y:PolyLineEdge>  
 `)
	if w.err != nil {
		return w.err
	}
	_, w.err = w.w.Write([]byte("</data></edge>\n"))
	return w.err
}

func (w *Writer) WriteQuads(buf []quad.Quad) (int, error) {
	for i, q := range buf {
		if err := w.WriteQuad(q); err != nil {
			return i, err
		}
	}
	return len(buf), nil
}

func (w *Writer) Close() error {
	if w.err != nil {
		return w.err
	}
	if !w.written {
		if _, w.err = w.w.Write([]byte(header)); w.err != nil {
			return w.err
		}
	}
	if _, w.err = w.w.Write([]byte(footer)); w.err != nil {
		return w.err
	}
	w.err = fmt.Errorf("closed")
	return nil
}

const header = `<?xml version="1.0" encoding="UTF-8"?>
<graphml xmlns="http://graphml.graphdrawing.org/xmlns"   
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  
         xsi:schemaLocation="http://graphml.graphdrawing.org/xmlns 
           http://www.yworks.com/xml/schema/graphml/1.1/ygraphml.xsd"  
         xmlns:y="http://www.yworks.com/xml/graphml">

    <key id="d0" for="node" yfiles.type="nodegraphics"/>  
    <key id="d1" for="edge" yfiles.type="edgegraphics"/>  
	<graph id="G" edgedefault="directed">
`
const footer = "\t</graph>\n</graphml>\n"
