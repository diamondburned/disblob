package svg

import (
	"testing"
)

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
   width="100%"
   height="100%"
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

const _MinifiedSVG = `data:image/svg+xml;utf8,<svg xmlns:dc="http:%2F%2Fpurl.org%2Fdc%2Felements%2F1.1%2F" xmlns:cc="http:%2F%2Fcreativecommons.org%2Fns%23" xmlns:rdf="http:%2F%2Fwww.w3.org%2F1999%2F02%2F22-rdf-syntax-ns%23" xmlns:svg="http:%2F%2Fwww.w3.org%2F2000%2Fsvg" xmlns="http:%2F%2Fwww.w3.org%2F2000%2Fsvg" xmlns:sodipodi="http:%2F%2Fsodipodi.sourceforge.net%2FDTD%2Fsodipodi-0.dtd" xmlns:inkscape="http:%2F%2Fwww.inkscape.org%2Fnamespaces%2Finkscape" style="enable-background:new 0 0 128 128" id="svg16" inkscape:version="0.92.3 (2405546%2C 2018-03-11)" sodipodi:docname="emoji_u1f524.svg"><sodipodi:namedview pagecolor="%23ffffff" bordercolor="%23666666" borderopacity="1" objecttolerance="10" gridtolerance="10" guidetolerance="10" inkscape:pageopacity="0" inkscape:pageshadow="2" inkscape:window-width="1920" inkscape:window-height="1046" id="namedview13" showgrid="false" inkscape:zoom="5.2149125" inkscape:cx="64.491369" inkscape:cy="64.441036" inkscape:window-x="-11" inkscape:window-y="-11" inkscape:window-maximized="1" inkscape:current-layer="svg16" inkscape:snap-bbox="true" inkscape:snap-bbox-midpoints="true" %2F><defs id="defs20" %2F><path style="opacity:1;fill:%233e99ff;fill-opacity:1;stroke:none;stroke-width:1;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" d="M7.4000001.0C26.266667.0 44.55806.958789 63.424727.958789 82.291393.958789 101.73333.0 120.6.0c4.0996.0 7.4 3.3004 7.4 7.4.0 18.866667-.57527 30.446538-.57527 49.313205.0 18.866667.57527 45.020125.57527 63.886795.0 4.0996-3.3004 7.4-7.4 7.4-18.86667.0-35.240482-1.91758-54.107149-1.91758C47.626184 126.08242 26.266667 128 7.4000001 128 3.3004 128 0 124.6996.0 120.6c0-18.86667.958789-38.692122.958789-57.558789C.958789 44.174545.0 26.266667.0 7.4.0 3.3004 3.3004.0 7.4000001.0z" id="rect840" %2F><path style="fill:%232089ff;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1" d="M125.14258 1.1777345C87.201877 46.083944 32.05307 77.426046.871094 124.19532c1.355343 1.6066 3.382123 2.62695 5.6582026 2.62695 18.8666674.0 40.2251304-1.91797 59.0917974-1.91797 18.866666.0 35.240756 1.91797 54.107426 1.91797 4.0996.0 7.40039-3.30079 7.40039-7.40039.0-18.86667-.57617-45.020056-.57617-63.886724.0-18.866666.57617-30.445832.57617-49.3124995.0-1.955407-.75603-3.724575-1.98633-5.044922z" id="path855" %2F><text style="font-style:normal;font-variant:normal;font-weight:400;font-stretch:normal;font-size:54.44121933px;line-height:1.25;font-family:comic neue;font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0;word-spacing:0;writing-mode:lr-tb;text-anchor:start;fill:%23fff;fill-opacity:1;stroke:none;stroke-width:2.91649389" x="12.536287" y="82.700562" id="text837"><tspan sodipodi:role="line" id="tspan835" x="12.536287" y="82.700562" style="font-style:normal;font-variant:normal;font-weight:700;font-stretch:normal;font-family:comic neue;fill:%23fff;stroke-width:2.91649389">a<%2Ftspan><%2Ftext><text id="text841" y="82.657768" x="49.001709" style="font-style:normal;font-variant:normal;font-weight:400;font-stretch:normal;font-size:55.44655609px;line-height:1.25;font-family:comic neue;font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0;word-spacing:0;writing-mode:lr-tb;text-anchor:start;fill:%23fff;fill-opacity:1;stroke:none;stroke-width:2.97035122"><tspan style="font-style:normal;font-variant:normal;font-weight:700;font-stretch:normal;font-family:comic neue;fill:%23fff;stroke-width:2.97035122" y="82.657768" x="49.001709" id="tspan839" sodipodi:role="line">b<%2Ftspan><%2Ftext><text style="font-style:normal;font-variant:normal;font-weight:400;font-stretch:normal;font-size:59.00439072px;line-height:1.25;font-family:comic neue;font-variant-ligatures:normal;font-variant-caps:normal;font-variant-numeric:normal;font-feature-settings:normal;text-align:start;letter-spacing:0;word-spacing:0;writing-mode:lr-tb;text-anchor:start;fill:%23fff;fill-opacity:1;stroke:none;stroke-width:3.16094947" x="85.526939" y="84.120499" id="text845"><tspan sodipodi:role="line" id="tspan843" x="85.526939" y="84.120499" style="font-style:normal;font-variant:normal;font-weight:700;font-stretch:normal;font-family:comic neue;fill:%23fff;stroke-width:3.16094947">c<%2Ftspan><%2Ftext><%2Fsvg>`
