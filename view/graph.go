// Copyright 2016 Google Inc.
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

package view

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"shenzhen-go/graph"
)

// Graph handles displaying/editing a graph.
func Graph(g *graph.Graph, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s graph: %s", r.Method, r.URL)
	q := r.URL.Query()
	if _, t := q["dot"]; t {
		outputDotSrc(g, w)
		return
	}
	if _, t := q["go"]; t {
		outputGoSrc(g, w)
		return
	}
	if _, t := q["json"]; t {
		outputJSON(g, w)
		return
	}
	if _, t := q["run"]; t {
		if err := g.SaveBuildAndRun(); err != nil {
			log.Printf("Failed to save, build, run: %v", err)
			// TODO: expose error better
		}
		u := *r.URL
		u.RawQuery = ""
		http.Redirect(w, r, u.String(), http.StatusFound)
		return
	}
	if _, t := q["save"]; t {
		if err := g.SaveJSONFile(); err != nil {
			log.Printf("Failed to save JSON file: %v", err)
		}
		u := *r.URL
		u.RawQuery = ""
		http.Redirect(w, r, u.String(), http.StatusFound)
		return
	}
	if n := q["node"]; len(n) == 1 {
		Node(g, n[0], w, r)
		return
	}
	if n := q["channel"]; len(n) == 1 {
		Channel(g, n[0], w, r)
		return
	}

	var dot, svg bytes.Buffer
	if err := g.WriteDotTo(&dot); err != nil {
		log.Printf("Could not render to dot: %v", err)
		http.Error(w, "Could not render to dot", http.StatusInternalServerError)
		return
	}
	if err := dotToSVG(&svg, &dot); err != nil {
		log.Printf("Could not render dot to SVG: %v", err)
		http.Error(w, "Could not render dot to SVG", http.StatusInternalServerError)
		return
	}
	d := &struct {
		Diagram template.HTML
		Graph   *graph.Graph
	}{
		Diagram: template.HTML(svg.String()),
		Graph:   g,
	}
	if err := graphEditorTemplate.Execute(w, d); err != nil {
		log.Printf("Could not execute graph editor template: %v", err)
		http.Error(w, "Could not execute graph editor template", http.StatusInternalServerError)
		return
	}
}

func outputDotSrc(g *graph.Graph, w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "text/vnd.graphviz")
	if err := g.WriteDotTo(w); err != nil {
		log.Printf("Could not render to dot: %v", err)
		http.Error(w, "Could not render to dot", http.StatusInternalServerError)
	}
}

func outputGoSrc(g *graph.Graph, w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "text/golang")
	if err := g.WriteGoTo(w); err != nil {
		log.Printf("Could not render to Go: %v", err)
		http.Error(w, "Could not render to Go", http.StatusInternalServerError)
	}
}

func outputJSON(g *graph.Graph, w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "application/json")
	if err := g.WriteJSONTo(w); err != nil {
		log.Printf("Could not encode JSON: %v", err)
		http.Error(w, "Could not encode JSON", http.StatusInternalServerError)
		return
	}
}