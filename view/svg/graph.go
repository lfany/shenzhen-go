// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

type Graph struct {
	Nodes    []Node
	Channels map[*Channel]struct{}
}

func (g *Graph) nearestPoint(x, y float64) (quad float64, pt Point) {
	quad = math.MaxFloat64
	test := func(p Point) {
		px, py := p.Pt()
		dx, dy := x-px, y-py
		if t := dx*dx + dy*dy; t < quad {
			quad, pt = t, p
		}
	}
	for _, n := range g.Nodes {
		for _, p := range n.Inputs {
			test(p)
		}
		for _, p := range n.Outputs {
			test(p)
		}
	}
	for c := range g.Channels {
		test(c)
	}
	return quad, pt
}

type jsonGraph struct {
	Nodes    []Node     `json:"nodes"`
	Channels []*Channel `json:"channels"`
}

func loadGraph() {
	resp, err := http.Get(apiEndpoint + "/graph")
	if err != nil {
		log.Printf("Querying API: %v", err)
		return
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	var g jsonGraph
	if err := d.Decode(&g); err != nil {
		log.Printf("Decoding response: %v", err)
	}
	graph.Nodes = g.Nodes
	graph.Channels = make(map[*Channel]struct{})
	for _, c := range g.Channels {
		graph.Channels[c] = struct{}{}
	}
}
