---
weight: 5
bookCollapseSection: true
keywords: ["n-quads","n-triples","triplet","rdf"]
title: "Knowledge-Graph Representations"
---

# Knowledge-Graph Representations

In the previous sections we described, how can we make statements of facts by the application of resources, predicates and literal values. These statements are named _triples_.

We also introduced the concept of IRI, that makes possible to write unambiguous statements in standardized format, that are generally applicable. This way we have a simple solution to define the nodes and edges of the knowledge graph, including their labels, by the use of a plain text editor. We can store these graph descriptions in plain text files, that we can load into a graph database, as well as we can export the graphs from the database into text files.

The graph database has its own internal representation to store and index the nodes and edges. When we store the knowledge representation in text files, we have several options to choose a standard format for this. In the following sub-sections, we summarize the most widespread formats we can choose from, when we work with {{< cayley >}}.

The most fundamental, base framework we use to describe resources is called RDF. The RDF is the acronym of [Resource Description Framework](https://en.wikipedia.org/wiki/Resource_Description_Framework) expression.

The RDF standard defines several file-serialization formats. The most frequently used formats are text based formats, that also have their corresponding mime-types.

The next table briefly summarizes the nuts-and-bolts of these formats:

{{< hint info >}}
__Note:__

This table lists these format only for the sake of completeness. In the first approximation to this topic, it is not needed to dive into the details of all these formats and the complete semantic technology stack. We focus on the practical aspects.

In order to use {{< cayley >}} we only need to know `n-quads`, `n-triples` formats. The `Turtle` format is also very useful and we will use it in some of our examples.
{{< /hint >}}

| format | file extension | mime-type | when to use? |
| --- | --- | --- | --- |
| [n-triples](https://www.w3.org/TR/n-triples/) | `.nt` | `application/n-triples` | In case we need acceptable performance, and want to ensure high level of compatibility. |
| [n-quads](https://www.w3.org/TR/n-quads/) | `.nq` | `application/n-quads` | The same as with n-triples |
| [Turtle](https://www.w3.org/TR/turtle/) | `.ttl` | `application/x-turtle` | In case we need to read and write the content manually. |
| [N3](https://www.w3.org/TeamSubmission/n3/) | `.n3` | `text/n3` | in case we need to apply RDF rules. |
| [RDF/XML](https://www.w3.org/TR/rdf-syntax-grammar/) | `.xml` | `application/rdf+xml` | If we have to use XML. |
| [JSON-LD](https://www.w3.org/2018/jsonld-cg-reports/json-ld/) | `.jsonld` | `application/ld+json` | When we want to provide a JSON API, or we need to connect to a JSON-LD API, and there is no high performance need. |

## n-triples, n-quads

These two formats are nearly identical. In fact, the `n-triples` are the statements that we introduced in the previous sections. In the text file, each statement is placed a separate line that is closed by the `.` character. For example:

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" .
<Luke-Skywalker-from-StarWars> <has-age> 23 .
```

The `n-quads` format differs only slightly from the `n-triples` format. The `n-quads` statements may contain a fourth element, that is a string that called (graph-) label. this label can identify the graph the statement belongs to. This way it is possible to store more than one graphs in the database, that are not fully connected with each other, that makes the performance optimization easier in case of indexing and querying.

The following code fragment contains `n-quads` statements, that are extended with the `"a-new-hope"` label:

```txt
<Luke-Skywalker-from-StarWars> <is-a> <Person> "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-given-name> "Luke" "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-family-name> "Skywalker" "a-new-hope" .
<Luke-Skywalker-from-StarWars> <has-age> 23 "a-new-hope" .
```

This label refers to the film episode in which the characters appeared first time.

The next code fragment describes the Dantooine planet in `n-triples` format. This example is made for the sake of comparison with the other formats that we will describe below. The same planet will be represented by the several formats in the upcoming examples.

```txt
<https://swapi.co/resource/planet/25> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://swapi.co/vocabulary/Planet> .
<https://swapi.co/resource/planet/25> <http://www.w3.org/2000/01/rdf-schema#label> "Dantooine"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <http://www.ontotext.com/business-object/type> "Planet" .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/climate> "temperate"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/desc> "None"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/diameter> "9830"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/gravity> "1 standard"^^<http://www.w3.org/2001/XMLSchema#string> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/orbitalPeriod> "378"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/population> "1000"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/rotationPeriod> "25"^^<http://www.w3.org/2001/XMLSchema#integer> .
<https://swapi.co/resource/planet/25> <https://swapi.co/vocabulary/terrain> "oceans, savannas, mountains, grasslands"^^<http://www.w3.org/2001/XMLSchema#string> .
```

## Turtle

The Turtle is a very popular format. It looks very similar to the `n-quads` format, but it is much less verbose. Moreover we can define the prefixes and namespaces in the head part of the file, then we can use the prefixed version of IRIs in the main part of the file, which makes the knowledge description much more dense, and eye-friendly.

When we define the statements to a specific subject, we need to refer to the subject only once, then we only have to list the predicate/object pairs, that we close with a `;` character, then we need to close the whole statement-set with the final `.` character.

Among these features, the Turtle format offers other facilitation and convenience services, so this format is really perfect to use for describing knowledge in case we need to manage manually the content.

The Dantooine planet in Turtle format:

```txt
@prefix bo: <http://www.ontotext.com/business-object/> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix voc: <https://swapi.co/vocabulary/> .
@prefix xml: <http://www.w3.org/XML/1998/namespace> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<https://swapi.co/resource/planet/25> a voc:Planet ;
    rdfs:label "Dantooine"^^xsd:string ;
    bo:type "Planet" ;
    voc:climate "temperate"^^xsd:string ;
    voc:desc "None"^^xsd:string ;
    voc:diameter 9830 ;
    voc:gravity "1 standard"^^xsd:string ;
    voc:orbitalPeriod 378 ;
    voc:population 1000 ;
    voc:rotationPeriod 25 ;
    voc:terrain "oceans, savannas, mountains, grasslands"^^xsd:string .
```

## N3

The N3 format (Notation3 in longer form) was worked out by Tim-Berners Lee, who also created the original RDF/XML format.

The N3 format is suitable to make complete semantic specifications. It is equivalent to the XML/RDF format, but it is much more readable than the XML content, so much transparent to human eyes. The N3 format is very similar to the Turtle. in fact the Turtle format is a subset of the N3 format.

the Dantooine planet in N3 format:

```txt
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix ns0: <http://www.ontotext.com/business-object/> .
@prefix ns1: <https://swapi.co/vocabulary/> .

<https://swapi.co/resource/planet/25>
  a <https://swapi.co/vocabulary/Planet> ;
  rdfs:label "Dantooine"^^xsd:string ;
  ns0:type "Planet" ;
  ns1:climate "temperate"^^xsd:string ;
  ns1:desc "None"^^xsd:string ;
  ns1:diameter 9830 ;
  ns1:gravity "1 standard"^^xsd:string ;
  ns1:orbitalPeriod 378 ;
  ns1:population 1000 ;
  ns1:rotationPeriod 25 ;
  ns1:terrain "oceans, savannas, mountains, grasslands"^^xsd:string .
```

## RDF/XML

The RDF/XML is the serialization format the was primarily created, and which is able to completely satisfy all the expectations against the knowledge representation on the semantic web.

Unfortunately the RDF/XML format is rather verbose, so it is not really well readable for human beings. We use it in cases we must use XML format.

The Dantooine planet in XML/RDF format:

```xml
<?xml version="1.0" encoding="utf-8" ?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
         xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
         xmlns:ns0="http://www.ontotext.com/business-object/"
         xmlns:ns1="https://swapi.co/vocabulary/">

  <rdf:Description rdf:about="https://swapi.co/resource/planet/25">
    <rdf:type rdf:resource="https://swapi.co/vocabulary/Planet"/>
    <rdfs:label rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Dantooine</rdfs:label>
    <ns0:type>Planet</ns0:type>
    <ns1:climate rdf:datatype="http://www.w3.org/2001/XMLSchema#string">temperate</ns1:climate>
    <ns1:desc rdf:datatype="http://www.w3.org/2001/XMLSchema#string">None</ns1:desc>
    <ns1:diameter rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">9830</ns1:diameter>
    <ns1:gravity rdf:datatype="http://www.w3.org/2001/XMLSchema#string">1 standard</ns1:gravity>
    <ns1:orbitalPeriod rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">378</ns1:orbitalPeriod>
    <ns1:population rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1000</ns1:population>
    <ns1:rotationPeriod rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">25</ns1:rotationPeriod>
    <ns1:terrain rdf:datatype="http://www.w3.org/2001/XMLSchema#string">oceans, savannas, mountains, grasslands</ns1:terrain>
  </rdf:Description>

</rdf:RDF>
```

## JSON-LD

The JSON-LD representations are documents that conform with the JSON standard. The JSON representation format is widely used by the REST APIs, as the main format of the requests and responses. It is well readable for humans too. Unfortunately the size of this format is farther big, relative to the other formats, and bigger files may consume too much computational resources for processing, so it is not the most perfect format if we need high throughput, and processing speed.

The Dantooine planet in JSON-LD format:

```json
[
    {
        "@id": "https://swapi.co/resource/planet/25",
        "@type": [
            "https://swapi.co/vocabulary/Planet"
        ],
        "http://www.ontotext.com/business-object/type": [
            {
                "@value": "Planet"
            }
        ],
        "http://www.w3.org/2000/01/rdf-schema#label": [
            {
                "@value": "Dantooine"
            }
        ],
        "https://swapi.co/vocabulary/climate": [
            {
                "@value": "temperate"
            }
        ],
        "https://swapi.co/vocabulary/desc": [
            {
                "@value": "None"
            }
        ],
        "https://swapi.co/vocabulary/diameter": [
            {
                "@value": 9830
            }
        ],
        "https://swapi.co/vocabulary/gravity": [
            {
                "@value": "1 standard"
            }
        ],
        "https://swapi.co/vocabulary/orbitalPeriod": [
            {
                "@value": 378
            }
        ],
        "https://swapi.co/vocabulary/population": [
            {
                "@value": 1000
            }
        ],
        "https://swapi.co/vocabulary/rotationPeriod": [
            {
                "@value": 25
            }
        ],
        "https://swapi.co/vocabulary/terrain": [
            {
                "@value": "oceans, savannas, mountains, grasslands"
            }
        ]
    },
    {
        "@id": "https://swapi.co/vocabulary/Planet"
    }
]
```
