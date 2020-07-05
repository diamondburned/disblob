package svg

import (
	"bytes"
	"strings"
	"testing"
)

const _HeaderSVG = `<svg xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:cc="http://creativecommons.org/ns#" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:svg="http://www.w3.org/2000/svg" xmlns="http://www.w3.org/2000/svg" xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd" xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape" width="128" height="128" style="enable-background:new 0 0 128 128" id="svg16" inkscape:version="0.92.3 (2405546, 2018-03-11)" sodipodi:docname="emoji_u1f524.svg"><body/></svg>`

const _HeaderMinified = `<svg width="128" height="128" style="enable-background:new 0 0 128 128"><body/></svg>`

func TestMinifyOpeningTag(t *testing.T) {
	var output bytes.Buffer

	if err := minifyOpeningTag(&output, strings.NewReader(_HeaderSVG)); err != nil {
		t.Fatal(err)
	}

	if s := output.String(); s != _HeaderMinified {
		t.Fatal("Mismatch minify:\n", s)
	}
}

func TestInline(t *testing.T) {
	b, err := Inline([]byte(_SVG))
	if err != nil {
		t.Fatal(err)
	}

	if s := string(b); s != _MinifiedSVG {
		t.Fatal("Mismatch minify:\n", s)
	}
}

const _SVG = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
   xmlns:dc="http://purl.org/dc/elements/1.1/"
   xmlns:cc="http://creativecommons.org/ns#"
   xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
   xmlns:svg="http://www.w3.org/2000/svg"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   width="128"
   height="128"
   style="enable-background:new 0 0 128 128;"
   version="1.1"
   id="svg16"
   inkscape:version="0.92.3 (2405546, 2018-03-11)"
   sodipodi:docname="emoji_u1f524.svg">
  <sodipodi:namedview
     pagecolor="#ffffff"
     bordercolor="#666666"
     borderopacity="1"
     objecttolerance="10"
     gridtolerance="10"
     guidetolerance="10"
     inkscape:pageopacity="0"
     inkscape:pageshadow="2"
     inkscape:window-width="1920"
     inkscape:window-height="1046"
     id="namedview13"
     showgrid="false"
     inkscape:zoom="5.2149125"
     inkscape:cx="64.491369"
     inkscape:cy="64.441036"
     inkscape:window-x="-11"
     inkscape:window-y="-11"
     inkscape:window-maximized="1"
     inkscape:current-layer="svg16"
     inkscape:snap-bbox="true"
     inkscape:snap-bbox-midpoints="true" />
  <metadata
     id="metadata22">
    <rdf:RDF>
      <cc:Work
         rdf:about="">
        <dc:format>image/svg+xml</dc:format>
        <dc:type
           rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
      </cc:Work>
    </rdf:RDF>
  </metadata>
  <defs
     id="defs20" />
  <path
     style="opacity:1;fill:#3e99ff;fill-opacity:1;stroke:none;stroke-width:1;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1;"
     d="M 7.4000001,0 C 26.266667,0 44.55806,0.958789 63.424727,0.958789 82.291393,0.958789 101.73333,0 120.6,0 c 4.0996,0 7.4,3.3004 7.4,7.4 0,18.866667 -0.57527,30.446538 -0.57527,49.313205 0,18.866667 0.57527,45.020125 0.57527,63.886795 0,4.0996 -3.3004,7.4 -7.4,7.4 -18.86667,0 -35.240482,-1.91758 -54.107149,-1.91758 C 47.626184,126.08242 26.266667,128 7.4000001,128 3.3004,128 0,124.6996 0,120.6 0,101.73333 0.958789,81.907878 0.958789,63.041211 0.958789,44.174545 0,26.266667 0,7.4 0,3.3004 3.3004,0 7.4000001,0 Z"
     id="rect840" />
  <path
     style="fill:#2089ff;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
     d="M 125.14258,1.1777345 C 87.201877,46.083944 32.05307,77.426046 0.871094,124.19532 c 1.355343,1.6066 3.382123,2.62695 5.6582026,2.62695 18.8666674,0 40.2251304,-1.91797 59.0917974,-1.91797 18.866666,0 35.240756,1.91797 54.107426,1.91797 4.0996,0 7.40039,-3.30079 7.40039,-7.40039 0,-18.86667 -0.57617,-45.020056 -0.57617,-63.886724 0,-18.866666 0.57617,-30.445832 0.57617,-49.3124995 0,-1.955407 -0.75603,-3.724575 -1.98633,-5.044922 z"
     id="path855" />
  <text
     style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:54.44121933px;line-height:1.25;font-family:'Comic Neue';font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:start;fill:#ffffff;fill-opacity:1;stroke:none;stroke-width:2.91649389"
     x="12.536287"
     y="82.700562"
     id="text837">
    <tspan
       sodipodi:role="line"
       id="tspan835"
       x="12.536287"
       y="82.700562"
       style="font-style:normal;font-variant:normal;font-weight:bold;font-stretch:normal;font-family:'Comic Neue';fill:#ffffff;stroke-width:2.91649389">a</tspan>
  </text>
  <text
     id="text841"
     y="82.657768"
     x="49.001709"
     style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:55.44655609px;line-height:1.25;font-family:'Comic Neue';font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:start;fill:#ffffff;fill-opacity:1;stroke:none;stroke-width:2.97035122">
    <tspan
       style="font-style:normal;font-variant:normal;font-weight:bold;font-stretch:normal;font-family:'Comic Neue';fill:#ffffff;stroke-width:2.97035122"
       y="82.657768"
       x="49.001709"
       id="tspan839"
       sodipodi:role="line">b</tspan>
  </text>
  <text
     style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:59.00439072px;line-height:1.25;font-family:'Comic Neue';font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:start;fill:#ffffff;fill-opacity:1;stroke:none;stroke-width:3.16094947"
     x="85.526939"
     y="84.120499"
     id="text845">
    <tspan
       sodipodi:role="line"
       id="tspan843"
       x="85.526939"
       y="84.120499"
       style="font-style:normal;font-variant:normal;font-weight:bold;font-stretch:normal;font-family:'Comic Neue';fill:#ffffff;stroke-width:3.16094947">c</tspan>
  </text>
</svg>`

const _MinifiedSVG = `<svg width="100%" height="100%" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;"><path d="M64.46,24.82C6.14,24.82 4.76,90.2 4.76,103.15C4.76,116.08 31.51,126.58 64.46,126.58C97.45,126.58 124.16,116.08 124.16,103.15C124.17,90.2 122.79,24.82 64.46,24.82Z" style="fill:rgb(180,124,228);fill-rule:nonzero;"/><path d="M83.92,40.16C84.63,38.15 86.12,36.26 88.35,34.51C90.17,33.08 92.1,31.68 93.6,29.9C98.02,24.64 100.18,16.45 99.18,9.61C98.95,8.08 97.46,5.13 98.47,3.74C100.22,1.36 104.55,3.82 105.94,5.19C110.29,9.47 112.9,15.34 114.17,21.24C115.76,28.57 115.54,36.17 113.11,43.27C111.09,49.15 106.24,54.96 99.99,56.34C94.63,57.53 88.84,53.99 85.88,49.78C83.49,46.36 82.87,43.13 83.92,40.16ZM45.01,39.32C44.3,37.32 42.83,35.43 40.58,33.67C38.75,32.24 36.86,30.84 35.34,29.06C30.93,23.8 28.77,15.62 29.77,8.78C30,7.24 31.47,4.3 30.47,2.91C28.73,0.51 24.4,2.98 23,4.36C18.66,8.64 16.05,14.51 14.76,20.4C13.19,27.72 13.4,35.33 15.84,42.43C17.85,48.31 22.72,54.12 28.96,55.51C34.31,56.7 40.1,53.15 43.04,48.95C45.44,45.52 46.07,42.3 45.01,39.32Z" style="fill:rgb(180,124,228);fill-rule:nonzero;"/><path id="a" d="M38.06,89.86C24.24,87.02 22.79,74.36 23.77,68.86L58.03,79.12C57.06,84.59 47.76,91.84 38.06,89.86Z" style="fill:rgb(47,47,47);fill-rule:nonzero;"/><clipPath id="_clip1"><path d="M38.06,89.86C24.24,87.02 22.79,74.36 23.77,68.86L58.03,79.12C57.06,84.59 47.76,91.84 38.06,89.86Z"/></clipPath><g clip-path="url(#_clip1)"><path d="M39.47,84.11C30.3,82.32 29.48,73.07 30.82,69.43L52.29,76.23C52.16,79.85 45.94,85.38 39.47,84.11Z" style="fill:rgb(251,184,23);fill-rule:nonzero;"/></g><path id="b" d="M90.87,89.86C104.69,87.02 106.14,74.36 105.16,68.86L70.9,79.11C71.87,84.59 81.18,91.84 90.87,89.86Z" style="fill:rgb(47,47,47);fill-rule:nonzero;"/><clipPath id="_clip2"><path d="M90.87,89.86C104.69,87.02 106.14,74.36 105.16,68.86L70.9,79.11C71.87,84.59 81.18,91.84 90.87,89.86Z"/></clipPath><g clip-path="url(#_clip2)"><path d="M89.45,84.11C98.62,82.32 99.44,73.07 98.1,69.43L76.64,76.23C76.76,79.85 82.98,85.38 89.45,84.11Z" style="fill:rgb(251,184,23);fill-rule:nonzero;"/></g><path d="M63.18,98.3C63.81,98.3 64.49,98.33 65.16,98.36C81,99.31 87,107.9 87.25,108.26C88.15,109.56 87.79,111.32 86.51,112.2C85.22,113.08 83.48,112.74 82.59,111.47C82.36,111.15 77.66,104.78 64.82,104.03C52.02,103.29 46.44,111.65 46.37,111.73C45.54,113.04 43.77,113.42 42.46,112.57C41.657,112.052 41.17,111.159 41.17,110.203C41.17,109.655 41.33,109.119 41.63,108.66C41.91,108.2 48.49,98.3 63.18,98.3Z" style="fill:rgb(47,47,47);fill-rule:nonzero;"/></svg>`
