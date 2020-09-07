---
title: "Gráf alapú adatmodellezés"
weight: 1
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: true
# bookComments: true
---

# Gráf alapú adatmodellezés

Amikor tudásgráfot építünk, akkor a világ egy részéről építünk fel egy koncepcionális modellt.

Egy koncepcionális modell jellemzően egy meghatározott szakterületet fed le, ezért ezt a modellt domain-nek, tudás-tartománynak is nevezzük.

Egy domain a következőket képviseli:

- az elsődleges entitásokat (a dolgokat, amelyek a domain-be beletartoznak). Az entitásokra használni fogjuk a resource kifejezést is;
- az entitások között fennálló kapcsolatokat;
- az entitások és kapcsolatok tulajdonságait, más szóval attribútumait Ezeket property-knek is nevezzük
- esetleg a szabályokat, amelyek az entitások és attribútumaik kapcsolatára, azok belső tulajdonságaira vonatkoznak.

A gráf alapú adatbázis-kezelők, és a szemantikus technológiák hatékonyan alkalmazhatók olyan esetben, ahol rugalmas, gyakran változó, nem strukturált adatmodellel kell dolgozni, és sokféle keresési szempont alapján kell lekérdezéseket végrehajtanunk. Ebből adódóan jók a koncepcionális modellek leírására.

A [gráfok](https://hu.wikipedia.org/wiki/Gr%C3%A1f) két fontos alkotóelemből állnak:
- csomópontok, más néven csúcsok, angolul __node__, vagy __vertex__.
- élek, amelyek a csomópontokat összekötik, angolul __edge__.

A tudásreprezentációnál a node-okat jellemzően fogalmak, dolgok ábrázolására használjuk, az edge-eket pedig a köztük fennálló kapcsolatok, asszociációk jelölésére.

Az élek irányítottak, és mind a node-oknak, mind az edge-eknek lehetnek címkéi, angolul __label__.

Az 1. ábra egy egyszerű gráfot ábrázol, amin két csomópont (node) található a "Luke" és "Leia" címkékkel, továbbá köztük egy él (edge) a "knows" címkével

{{< figure src="/cayley-cookbook/simple-graph.png" title="1. ábra: Egyszerű Gráf" >}}

Bizonyos gráf adatbázis-kezelők megengedik, hogy a node-ok és edge-ek több tulajdonságot is hordozzanak, a {{< cayley >}} esetében számunkra ezek egyetlen értéket fognak jelenteni.

{{< button relref="basic-concepts" >}}&#9669; Alapfogalmak{{< /button >}}
{{< button relref="resources" >}}Erőforrások &#9659;{{< /button >}}

