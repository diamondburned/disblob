package svg

import (
	"strings"
	"testing"

	"github.com/hexops/autogold"
)

func TestInline(t *testing.T) {
	var builder strings.Builder

	if err := Inline(&builder, strings.NewReader(regularSVG)); err != nil {
		t.Fatal(err)
	}

	wantMinify.Equal(t, builder.String())
}

var wantMinify = autogold.Want("minify", "data:image/svg+xml;base64,PHN2ZyB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyIgeG1sbnM6c3ZnPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczpzb2RpcG9kaT0iaHR0cDovL3NvZGlwb2RpLnNvdXJjZWZvcmdlLm5ldC9EVEQvc29kaXBvZGktMC5kdGQiIHhtbG5zOmlua3NjYXBlPSJodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy9uYW1lc3BhY2VzL2lua3NjYXBlIiBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IDAgMCAxMjggMTI4IiBpZD0ic3ZnMTYiIGlua3NjYXBlOnZlcnNpb249IjAuOTIuMyAoMjQwNTU0NiwgMjAxOC0wMy0xMSkiIHNvZGlwb2RpOmRvY25hbWU9ImVtb2ppX3UxZjUyNC5zdmciPjxzb2RpcG9kaTpuYW1lZHZpZXcgcGFnZWNvbG9yPSIjZmZmZmZmIiBib3JkZXJjb2xvcj0iIzY2NjY2NiIgYm9yZGVyb3BhY2l0eT0iMSIgb2JqZWN0dG9sZXJhbmNlPSIxMCIgZ3JpZHRvbGVyYW5jZT0iMTAiIGd1aWRldG9sZXJhbmNlPSIxMCIgaW5rc2NhcGU6cGFnZW9wYWNpdHk9IjAiIGlua3NjYXBlOnBhZ2VzaGFkb3c9IjIiIGlua3NjYXBlOndpbmRvdy13aWR0aD0iMTkyMCIgaW5rc2NhcGU6d2luZG93LWhlaWdodD0iMTA0NiIgaWQ9Im5hbWVkdmlldzEzIiBzaG93Z3JpZD0iZmFsc2UiIGlua3NjYXBlOnpvb209IjUuMjE0OTEyNSIgaW5rc2NhcGU6Y3g9IjY0LjQ5MTM2OSIgaW5rc2NhcGU6Y3k9IjY0LjQ0MTAzNiIgaW5rc2NhcGU6d2luZG93LXg9Ii0xMSIgaW5rc2NhcGU6d2luZG93LXk9Ii0xMSIgaW5rc2NhcGU6d2luZG93LW1heGltaXplZD0iMSIgaW5rc2NhcGU6Y3VycmVudC1sYXllcj0ic3ZnMTYiIGlua3NjYXBlOnNuYXAtYmJveD0idHJ1ZSIgaW5rc2NhcGU6c25hcC1iYm94LW1pZHBvaW50cz0idHJ1ZSIgLz48ZGVmcyBpZD0iZGVmczIwIiAvPjxwYXRoIHN0eWxlPSJvcGFjaXR5OjE7ZmlsbDojM2U5OWZmO2ZpbGwtb3BhY2l0eToxO3N0cm9rZTpub25lO3N0cm9rZS13aWR0aDoxO3N0cm9rZS1saW5lY2FwOmJ1dHQ7c3Ryb2tlLWxpbmVqb2luOm1pdGVyO3N0cm9rZS1taXRlcmxpbWl0OjQ7c3Ryb2tlLWRhc2hhcnJheTpub25lO3N0cm9rZS1kYXNob2Zmc2V0OjA7c3Ryb2tlLW9wYWNpdHk6MSIgZD0iTTcuNDAwMDAwMS4wQzI2LjI2NjY2Ny4wIDQ0LjU1ODA2Ljk1ODc4OSA2My40MjQ3MjcuOTU4Nzg5IDgyLjI5MTM5My45NTg3ODkgMTAxLjczMzMzLjAgMTIwLjYuMGM0LjA5OTYuMCA3LjQgMy4zMDA0IDcuNCA3LjQuMCAxOC44NjY2NjctLjU3NTI3IDMwLjQ0NjUzOC0uNTc1MjcgNDkuMzEzMjA1LjAgMTguODY2NjY3LjU3NTI3IDQ1LjAyMDEyNS41NzUyNyA2My44ODY3OTUuMCA0LjA5OTYtMy4zMDA0IDcuNC03LjQgNy40LTE4Ljg2NjY3LjAtMzUuMjQwNDgyLTEuOTE3NTgtNTQuMTA3MTQ5LTEuOTE3NThDNDcuNjI2MTg0IDEyNi4wODI0MiAyNi4yNjY2NjcgMTI4IDcuNDAwMDAwMSAxMjggMy4zMDA0IDEyOCAwIDEyNC42OTk2LjAgMTIwLjZjMC0xOC44NjY2Ny45NTg3ODktMzguNjkyMTIyLjk1ODc4OS01Ny41NTg3ODlDLjk1ODc4OSA0NC4xNzQ1NDUuMCAyNi4yNjY2NjcuMCA3LjQuMCAzLjMwMDQgMy4zMDA0LjAgNy40MDAwMDAxLjB6IiBpZD0icmVjdDg0MCIgLz48cGF0aCBzdHlsZT0iZmlsbDojMjA4OWZmO2ZpbGwtb3BhY2l0eToxO3N0cm9rZTpub25lO3N0cm9rZS13aWR0aDoxcHg7c3Ryb2tlLWxpbmVjYXA6YnV0dDtzdHJva2UtbGluZWpvaW46bWl0ZXI7c3Ryb2tlLW9wYWNpdHk6MSIgZD0iTTEyNS4xNDI1OCAxLjE3NzczNDVDODcuMjAxODc3IDQ2LjA4Mzk0NCAzMi4wNTMwNyA3Ny40MjYwNDYuODcxMDk0IDEyNC4xOTUzMmMxLjM1NTM0MyAxLjYwNjYgMy4zODIxMjMgMi42MjY5NSA1LjY1ODIwMjYgMi42MjY5NSAxOC44NjY2Njc0LjAgNDAuMjI1MTMwNC0xLjkxNzk3IDU5LjA5MTc5NzQtMS45MTc5NyAxOC44NjY2NjYuMCAzNS4yNDA3NTYgMS45MTc5NyA1NC4xMDc0MjYgMS45MTc5NyA0LjA5OTYuMCA3LjQwMDM5LTMuMzAwNzkgNy40MDAzOS03LjQwMDM5LjAtMTguODY2NjctLjU3NjE3LTQ1LjAyMDA1Ni0uNTc2MTctNjMuODg2NzI0LjAtMTguODY2NjY2LjU3NjE3LTMwLjQ0NTgzMi41NzYxNy00OS4zMTI0OTk1LjAtMS45NTU0MDctLjc1NjAzLTMuNzI0NTc1LTEuOTg2MzMtNS4wNDQ5MjJ6IiBpZD0icGF0aDg1NSIgLz48dGV4dCBzdHlsZT0iZm9udC1zdHlsZTpub3JtYWw7Zm9udC12YXJpYW50Om5vcm1hbDtmb250LXdlaWdodDo0MDA7Zm9udC1zdHJldGNoOm5vcm1hbDtmb250LXNpemU6NTQuNDQxMjE5MzNweDtsaW5lLWhlaWdodDoxLjI1O2ZvbnQtZmFtaWx5OmNvbWljIG5ldWU7Zm9udC12YXJpYW50LWxpZ2F0dXJlczpub3JtYWw7Zm9udC12YXJpYW50LWNhcHM6bm9ybWFsO2ZvbnQtdmFyaWFudC1udW1lcmljOm5vcm1hbDtmb250LWZlYXR1cmUtc2V0dGluZ3M6bm9ybWFsO3RleHQtYWxpZ246c3RhcnQ7bGV0dGVyLXNwYWNpbmc6MDt3b3JkLXNwYWNpbmc6MDt3cml0aW5nLW1vZGU6bHItdGI7dGV4dC1hbmNob3I6c3RhcnQ7ZmlsbDojZmZmO2ZpbGwtb3BhY2l0eToxO3N0cm9rZTpub25lO3N0cm9rZS13aWR0aDoyLjkxNjQ5Mzg5IiB4PSIxMi41MzYyODciIHk9IjgyLjcwMDU2MiIgaWQ9InRleHQ4MzciPjx0c3BhbiBzb2RpcG9kaTpyb2xlPSJsaW5lIiBpZD0idHNwYW44MzUiIHg9IjEyLjUzNjI4NyIgeT0iODIuNzAwNTYyIiBzdHlsZT0iZm9udC1zdHlsZTpub3JtYWw7Zm9udC12YXJpYW50Om5vcm1hbDtmb250LXdlaWdodDo3MDA7Zm9udC1zdHJldGNoOm5vcm1hbDtmb250LWZhbWlseTpjb21pYyBuZXVlO2ZpbGw6I2ZmZjtzdHJva2Utd2lkdGg6Mi45MTY0OTM4OSI+YTwvdHNwYW4+PC90ZXh0Pjx0ZXh0IGlkPSJ0ZXh0ODQxIiB5PSI4Mi42NTc3NjgiIHg9IjQ5LjAwMTcwOSIgc3R5bGU9ImZvbnQtc3R5bGU6bm9ybWFsO2ZvbnQtdmFyaWFudDpub3JtYWw7Zm9udC13ZWlnaHQ6NDAwO2ZvbnQtc3RyZXRjaDpub3JtYWw7Zm9udC1zaXplOjU1LjQ0NjU1NjA5cHg7bGluZS1oZWlnaHQ6MS4yNTtmb250LWZhbWlseTpjb21pYyBuZXVlO2ZvbnQtdmFyaWFudC1saWdhdHVyZXM6bm9ybWFsO2ZvbnQtdmFyaWFudC1jYXBzOm5vcm1hbDtmb250LXZhcmlhbnQtbnVtZXJpYzpub3JtYWw7Zm9udC1mZWF0dXJlLXNldHRpbmdzOm5vcm1hbDt0ZXh0LWFsaWduOnN0YXJ0O2xldHRlci1zcGFjaW5nOjA7d29yZC1zcGFjaW5nOjA7d3JpdGluZy1tb2RlOmxyLXRiO3RleHQtYW5jaG9yOnN0YXJ0O2ZpbGw6I2ZmZjtmaWxsLW9wYWNpdHk6MTtzdHJva2U6bm9uZTtzdHJva2Utd2lkdGg6Mi45NzAzNTEyMiI+PHRzcGFuIHN0eWxlPSJmb250LXN0eWxlOm5vcm1hbDtmb250LXZhcmlhbnQ6bm9ybWFsO2ZvbnQtd2VpZ2h0OjcwMDtmb250LXN0cmV0Y2g6bm9ybWFsO2ZvbnQtZmFtaWx5OmNvbWljIG5ldWU7ZmlsbDojZmZmO3N0cm9rZS13aWR0aDoyLjk3MDM1MTIyIiB5PSI4Mi42NTc3NjgiIHg9IjQ5LjAwMTcwOSIgaWQ9InRzcGFuODM5IiBzb2RpcG9kaTpyb2xlPSJsaW5lIj5iPC90c3Bhbj48L3RleHQ+PHRleHQgc3R5bGU9ImZvbnQtc3R5bGU6bm9ybWFsO2ZvbnQtdmFyaWFudDpub3JtYWw7Zm9udC13ZWlnaHQ6NDAwO2ZvbnQtc3RyZXRjaDpub3JtYWw7Zm9udC1zaXplOjU5LjAwNDM5MDcycHg7bGluZS1oZWlnaHQ6MS4yNTtmb250LWZhbWlseTpjb21pYyBuZXVlO2ZvbnQtdmFyaWFudC1saWdhdHVyZXM6bm9ybWFsO2ZvbnQtdmFyaWFudC1jYXBzOm5vcm1hbDtmb250LXZhcmlhbnQtbnVtZXJpYzpub3JtYWw7Zm9udC1mZWF0dXJlLXNldHRpbmdzOm5vcm1hbDt0ZXh0LWFsaWduOnN0YXJ0O2xldHRlci1zcGFjaW5nOjA7d29yZC1zcGFjaW5nOjA7d3JpdGluZy1tb2RlOmxyLXRiO3RleHQtYW5jaG9yOnN0YXJ0O2ZpbGw6I2ZmZjtmaWxsLW9wYWNpdHk6MTtzdHJva2U6bm9uZTtzdHJva2Utd2lkdGg6My4xNjA5NDk0NyIgeD0iODUuNTI2OTM5IiB5PSI4NC4xMjA0OTkiIGlkPSJ0ZXh0ODQ1Ij48dHNwYW4gc29kaXBvZGk6cm9sZT0ibGluZSIgaWQ9InRzcGFuODQzIiB4PSI4NS41MjY5MzkiIHk9Ijg0LjEyMDQ5OSIgc3R5bGU9ImZvbnQtc3R5bGU6bm9ybWFsO2ZvbnQtdmFyaWFudDpub3JtYWw7Zm9udC13ZWlnaHQ6NzAwO2ZvbnQtc3RyZXRjaDpub3JtYWw7Zm9udC1mYW1pbHk6Y29taWMgbmV1ZTtmaWxsOiNmZmY7c3Ryb2tlLXdpZHRoOjMuMTYwOTQ5NDciPmM8L3RzcGFuPjwvdGV4dD48L3N2Zz4=")

const regularSVG = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
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

const _MinifiedSVG = `data:image/svg+xml,%3Csvg%20xmlns:dc=%22http:%2F%2Fpurl.org%2Fdc%2Felements%2F1.1%2F%22%20xmlns:cc=%22http:%2F%2Fcreativecommons.org%2Fns%23%22%20xmlns:rdf=%22http:%2F%2Fwww.w3.org%2F1999%2F02%2F22-rdf-syntax-ns%23%22%20xmlns:svg=%22http:%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns=%22http:%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns:sodipodi=%22http:%2F%2Fsodipodi.sourceforge.net%2FDTD%2Fsodipodi-0.dtd%22%20xmlns:inkscape=%22http:%2F%2Fwww.inkscape.org%2Fnamespaces%2Finkscape%22%20style=%22enable-background:new%200%200%20128%20128%22%20id=%22svg16%22%20inkscape:version=%220.92.3%20%282405546%2C%202018-03-11%29%22%20sodipodi:docname=%22emoji_u1f524.svg%22%3E%3Csodipodi:namedview%20pagecolor=%22%23ffffff%22%20bordercolor=%22%23666666%22%20borderopacity=%221%22%20objecttolerance=%2210%22%20gridtolerance=%2210%22%20guidetolerance=%2210%22%20inkscape:pageopacity=%220%22%20inkscape:pageshadow=%222%22%20inkscape:window-width=%221920%22%20inkscape:window-height=%221046%22%20id=%22namedview13%22%20showgrid=%22false%22%20inkscape:zoom=%225.2149125%22%20inkscape:cx=%2264.491369%22%20inkscape:cy=%2264.441036%22%20inkscape:window-x=%22-11%22%20inkscape:window-y=%22-11%22%20inkscape:window-maximized=%221%22%20inkscape:current-layer=%22svg16%22%20inkscape:snap-bbox=%22true%22%20inkscape:snap-bbox-midpoints=%22true%22%20%2F%3E%3Cdefs%20id=%22defs20%22%20%2F%3E%3Cpath%20style=%22opacity:1%3Bfill:%233e99ff%3Bfill-opacity:1%3Bstroke:none%3Bstroke-width:1%3Bstroke-linecap:butt%3Bstroke-linejoin:miter%3Bstroke-miterlimit:4%3Bstroke-dasharray:none%3Bstroke-dashoffset:0%3Bstroke-opacity:1%22%20d=%22M7.4000001.0C26.266667.0%2044.55806.958789%2063.424727.958789%2082.291393.958789%20101.73333.0%20120.6.0c4.0996.0%207.4%203.3004%207.4%207.4.0%2018.866667-.57527%2030.446538-.57527%2049.313205.0%2018.866667.57527%2045.020125.57527%2063.886795.0%204.0996-3.3004%207.4-7.4%207.4-18.86667.0-35.240482-1.91758-54.107149-1.91758C47.626184%20126.08242%2026.266667%20128%207.4000001%20128%203.3004%20128%200%20124.6996.0%20120.6c0-18.86667.958789-38.692122.958789-57.558789C.958789%2044.174545.0%2026.266667.0%207.4.0%203.3004%203.3004.0%207.4000001.0z%22%20id=%22rect840%22%20%2F%3E%3Cpath%20style=%22fill:%232089ff%3Bfill-opacity:1%3Bstroke:none%3Bstroke-width:1px%3Bstroke-linecap:butt%3Bstroke-linejoin:miter%3Bstroke-opacity:1%22%20d=%22M125.14258%201.1777345C87.201877%2046.083944%2032.05307%2077.426046.871094%20124.19532c1.355343%201.6066%203.382123%202.62695%205.6582026%202.62695%2018.8666674.0%2040.2251304-1.91797%2059.0917974-1.91797%2018.866666.0%2035.240756%201.91797%2054.107426%201.91797%204.0996.0%207.40039-3.30079%207.40039-7.40039.0-18.86667-.57617-45.020056-.57617-63.886724.0-18.866666.57617-30.445832.57617-49.3124995.0-1.955407-.75603-3.724575-1.98633-5.044922z%22%20id=%22path855%22%20%2F%3E%3Ctext%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:400%3Bfont-stretch:normal%3Bfont-size:54.44121933px%3Bline-height:1.25%3Bfont-family:comic%20neue%3Bfont-variant-ligatures:normal%3Bfont-variant-caps:normal%3Bfont-variant-numeric:normal%3Bfont-feature-settings:normal%3Btext-align:start%3Bletter-spacing:0%3Bword-spacing:0%3Bwriting-mode:lr-tb%3Btext-anchor:start%3Bfill:%23fff%3Bfill-opacity:1%3Bstroke:none%3Bstroke-width:2.91649389%22%20x=%2212.536287%22%20y=%2282.700562%22%20id=%22text837%22%3E%3Ctspan%20sodipodi:role=%22line%22%20id=%22tspan835%22%20x=%2212.536287%22%20y=%2282.700562%22%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:700%3Bfont-stretch:normal%3Bfont-family:comic%20neue%3Bfill:%23fff%3Bstroke-width:2.91649389%22%3Ea%3C%2Ftspan%3E%3C%2Ftext%3E%3Ctext%20id=%22text841%22%20y=%2282.657768%22%20x=%2249.001709%22%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:400%3Bfont-stretch:normal%3Bfont-size:55.44655609px%3Bline-height:1.25%3Bfont-family:comic%20neue%3Bfont-variant-ligatures:normal%3Bfont-variant-caps:normal%3Bfont-variant-numeric:normal%3Bfont-feature-settings:normal%3Btext-align:start%3Bletter-spacing:0%3Bword-spacing:0%3Bwriting-mode:lr-tb%3Btext-anchor:start%3Bfill:%23fff%3Bfill-opacity:1%3Bstroke:none%3Bstroke-width:2.97035122%22%3E%3Ctspan%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:700%3Bfont-stretch:normal%3Bfont-family:comic%20neue%3Bfill:%23fff%3Bstroke-width:2.97035122%22%20y=%2282.657768%22%20x=%2249.001709%22%20id=%22tspan839%22%20sodipodi:role=%22line%22%3Eb%3C%2Ftspan%3E%3C%2Ftext%3E%3Ctext%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:400%3Bfont-stretch:normal%3Bfont-size:59.00439072px%3Bline-height:1.25%3Bfont-family:comic%20neue%3Bfont-variant-ligatures:normal%3Bfont-variant-caps:normal%3Bfont-variant-numeric:normal%3Bfont-feature-settings:normal%3Btext-align:start%3Bletter-spacing:0%3Bword-spacing:0%3Bwriting-mode:lr-tb%3Btext-anchor:start%3Bfill:%23fff%3Bfill-opacity:1%3Bstroke:none%3Bstroke-width:3.16094947%22%20x=%2285.526939%22%20y=%2284.120499%22%20id=%22text845%22%3E%3Ctspan%20sodipodi:role=%22line%22%20id=%22tspan843%22%20x=%2285.526939%22%20y=%2284.120499%22%20style=%22font-style:normal%3Bfont-variant:normal%3Bfont-weight:700%3Bfont-stretch:normal%3Bfont-family:comic%20neue%3Bfill:%23fff%3Bstroke-width:3.16094947%22%3Ec%3C%2Ftspan%3E%3C%2Ftext%3E%3C%2Fsvg%3E`
