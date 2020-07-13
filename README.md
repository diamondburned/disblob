---


---

<h1 id="disblob">disblob</h1>
<p>A stylesheet(css) generator for use with Discord.</p>
<hr>
<p><em><strong>All original code is done by the insane but genius <a href="https://github.com/diamondburned/">diamondburned</a></strong></em></p>
<blockquote>
<p><a href="https://github.com/diamondburned/disblob">Master repo</a></p>
</blockquote>
<p><em><strong>This fork is an attempt to refine the <code>README</code> and make it easier for others to setup!</strong></em></p>
<hr>
<h1 id="status">Status</h1>

<table>
<thead>
<tr>
<th></th>
<th>Status</th>
<th></th>
</tr>
</thead>
<tbody>
<tr>
<td></td>
<td>Edits:</td>
<td></td>
</tr>
<tr>
<td></td>
<td><code>README</code> Only</td>
<td></td>
</tr>
</tbody>
</table>
<table>
<thead>
<tr>
<th></th>
<th>Fork</th>
<th></th>
</tr>
</thead>
<tbody>
<tr>
<td></td>
<td>Fork Only</td>
<td></td>
</tr>
<tr>
<td></td>
<td>(No PR)</td>
<td></td>
</tr>
</tbody>
</table>
<table>
<thead>
<tr>
<th></th>
<th>Testing</th>
<th></th>
</tr>
</thead>
<tbody>
<tr>
<td></td>
<td>-Client Only-</td>
<td></td>
</tr>
<tr>
<td></td>
<td>-  (No Tests have been done outside of the client)</td>
<td></td>
</tr>
</tbody>
</table><hr>
<ul>
<li class="task-list-item"><input type="checkbox" class="task-list-item-checkbox" checked="true" disabled=""> Client (BD(BetterBiscord)~(100% Working!)</li>
<li class="task-list-item"><input type="checkbox" class="task-list-item-checkbox" disabled=""> Client (Other css injectors)</li>
<li class="task-list-item"><input type="checkbox" class="task-list-item-checkbox" disabled=""> Browsers (Chrome, etc. (Use <a href="https://github.com/openstyles/stylus">stylus</a> to inject css)</li>
<li class="task-list-item"><input type="checkbox" class="task-list-item-checkbox" checked="true" disabled=""> Firefox Browser (Testing done with stylish(DO NOT USE(Considered spyware))</li>
</ul>
<hr>

<table>
<thead>
<tr>
<th></th>
<th>Contributors:</th>
<th></th>
</tr>
</thead>
<tbody>
<tr>
<td></td>
<td>2 (<a href="https://github.com/diamondburned/">diamondburned</a> &amp; <a href="https://github.com/ThatGeekyWeeb">ThatGeekyWeeb</a>)</td>
<td></td>
</tr>
</tbody>
</table><hr>
<h2 id="dependencies">Dependencies</h2>
<ol>
<li>(All Blobmoji deps(<a href="https://github.com/C1710/blobmoji/wiki/Build-instructions">https://github.com/C1710/blobmoji/wiki/Build-instructions</a>))</li>
<li><code>go</code></li>
<li><code>zopfli</code></li>
<li>1+ GiB of storage space</li>
</ol>
<h2 id="building">Building</h2>
<pre class=" language-sh"><code class="prism  language-sh">git submodule update --init
cd blobmoji
make -j$(nproc)
cd ..
mv blobmoji/build/renamed_flags/* blobmoji/svg/
mv blobmoji/build/resized_flags/ blobmoji/

# This will generate the stylesheet.
go run . -defpath definitions.css -datpath data.css &gt; &lt;css file&gt;
</code></pre>
<h2 id="releases">Releases</h2>
<p>My github pages link can be used to <code>@import</code> the css</p>
<pre class=" language-css"><code class="prism  language-css"><span class="token atrule"><span class="token rule">@import</span> <span class="token url">url("https://thatgeekyweeb.github.io/src/import.data.css")</span><span class="token punctuation">;</span></span>
</code></pre>
<p>Additionally there are releases <a href="https://github.com/ThatGeekyWeeb/disblob/releases">here!</a></p>
<h1 id="contributing">Contributing</h1>
<p>Any Pull Requests, Opened Issues, and forks are  Greatly welcomed!<br>
You help make disblob better by Contributing!</p>
<blockquote>
<p>Written with <a href="https://stackedit.io/">StackEdit</a>.</p>
</blockquote>

