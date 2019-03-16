{{define "page"}}

検索ワードの書き方

Google検索に対応しています。

<h3>あいまい検索</h3>
「isear」を検索したい場合、「ISEAR」、「Isear」などでも検索できます。(正規表現検索時は除く)<br>
以下のものに対応しています。<br>
<ul>
	<li>大文字/小文字</li>
	<li>半角/全角(数字、記号、スペース含む)</li>
	<li>ひらがな/カタカナ</li>
</ul>

<h3>複数ワードの検索</h3>
<p>ワードをスペース区切りにすることで複数のワードを検索できます。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>isear chrome</td>
		<td>[isear] [chrome]</td>
	</tr>
</table>

<h4>スペースを検索ワードに含める</h4>
<p></p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>"isear chrome"</td>
		<td>[isear chrome]</td>
	</tr>
</table>

<h3>OR,カッコの判定</h3>
<p>Google検索でOR検索を行う場合など、カッコを使う場合があります。isearでは通常検索結果から自動的に除外します。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>isear (chrome OR firefox)</td>
		<td>[isear] [chrome] [firefox]</td>
	</tr>
</table>
<h4>検索ワードに含める</h4>
<p>"(ダブルクォーテーション)や'(シングルクォーテーション)で囲うことでOR,カッコも検索できます。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>"index()"</td>
		<td>[index()]</td>
	</tr>
	<tr>
		<td>"OR"</td>
		<td>[OR]</td>
	</tr>
</table>

<h3>マイナス検索</h3>
<p>Google検索では検索結果から除外したい単語はマイナスをつけて検索をしますが、isearでも検索ワードとして認識しません。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>isear -chrome</td>
		<td>[isear]</td>
	</tr>
</table>

<h3>正規表現検索(上級者向け)</h3>
<p>@RE:をつけることで、そのワードを正規表現として検索をすることが出来ます。(設定画面から正規表現を有効にしてください)</p>
<p>\wと\Wのように、大文字小文字で挙動が違う正規表現があるためあいまい検索は適応されません。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>@RE:regexp</td>
		<td>[regexp]</td>
	</tr>
</table>
<p>@RE:の後にスペースを入れると、その後の全ての単語が正規表現として検索されます。</p>
<table class="table">
	<tr>
		<th>検索テキスト</th>
		<th>生成されるボタン</th>
	</tr>
	<tr>
		<td>@RE: regexp1 regexp2 regexp3</td>
		<td>[regexp1] [regexp2] [regexp3]</td>
	</tr>
</table>

{{end}}
