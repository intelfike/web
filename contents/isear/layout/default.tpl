<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>{{ .Title }} | isear解説</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<!-- <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous"> -->
	<link rel="stylesheet" href="https://unpkg.com/bootstrap-material-design@4.1.1/dist/css/bootstrap-material-design.min.css" integrity="sha384-wXznGJNEXNG1NFsbm0ugrLFMQPWswR3lds2VeinahP8N0zJw9VWSopbjv2x7WCvX" crossorigin="anonymous">
	<link rel="stylesheet" href="/css/default.css?_={{ .Rand }}">
</head>
<body>

<header class="header">
	<h1 class="p-2">
		<a href="/" class="normal" style="padding:10px;">
			<img src="/data/image/isear_document_title.png" height="40" alt="isearロゴ">
		</a>
	</h1>
	<nav id="global-menu" class="navbar p-0">
		<ul class="list-inline m-0">
			<a class="list-inline-item p-2 top" href="/_top/index">トップ</a>
			<a class="list-inline-item p-2 usage
				{{- if (eq .Info.Group "usage")}} collapsed{{end -}}
			" data-toggle="collapse" href="#usage-links" role="button" aria-expanded="
				{{- print (eq .Info.Group "usage") -}}
			">使い方</a>
			<a class="list-inline-item p-2 install" href="/_install/index">インストール</a>
			<a class="list-inline-item p-2 support" href="/_support/index">サポート</a>
			<a class="list-inline-item p-2 thanks" href="/_thanks/index">Thanks</a>
		</ul>
	</nav>
	{{if (eq .Info.Group "usage")}}
	<div id="usage-links" class="collapse show">
		<nav class="navbar">
			<ul class="list-inline m-0">
				<a class="list-inline-item functions" href="/_usage/functions">機能</a>
				<a class="list-inline-item settings" href="/_usage/settings">設定</a>
				<a class="list-inline-item name" href="/_usage/name">名称・用語</a>
				<a class="list-inline-item sKey" href="/_usage/sKey">ショートカットキー</a>
				<a class="list-inline-item sWord" href="/_usage/sWord">検索テキスト</a>
				<a class="list-inline-item commandMode" href="/_usage/commandMode">コマンドモード</a>
			</ul>
		</nav>
	</div>
	{{else}}
	<div id="usage-links" class="collapse">
		<nav class="navbar">
			<ul class="list-group m-0">
				<a class="m-1" href="/_usage/functions">機能</a>
				<a class="m-1" href="/_usage/settings">設定</a>
				<a class="m-1" href="/_usage/name">名称・用語</a>
				<a class="m-1" href="/_usage/sKey">ショートカットキー</a>
				<a class="m-1" href="/_usage/sWord">検索テキスト</a>
				<a class="m-1" href="/_usage/commandMode">コマンドモード</a>
			</ul>
		</nav>
	</div>
	{{end}}
</header>

<div id="middle">
	<article class="container">
		<h2>
			{{.Title}}
		</h2>
{{ template "page" . }}
	</article>
</div>

<footer calss="footer">
	&copy;2018.01 intelfike(intelfike@gmail.com)
</footer>

<script>
document.body.onload = e => {
	var gp = document.querySelector("#global-menu .{{$.Info.Group}}")
	gp.className += ' current-group'
	gp.href = '#'

	{{if (eq .Info.Group "usage") -}}
	var us = document.querySelector("#usage-links .{{$.Info.Page}}")
	us.className += ' current-page'
	us.href = '#'
	{{- end}}
}
</script>

	<!-- V4 version -->
	<!-- <script type="text/javascript" src="https://cdn.jsdelivr.net/npm/bootstrap.native@2.0.15/dist/bootstrap-native-v4.min.js"></script> -->
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://unpkg.com/popper.js@1.12.6/dist/umd/popper.js" integrity="sha384-fA23ZRQ3G/J53mElWqVJEGJzU0sTs+SvzG8fXVWP+kJQ1lwFAOkcUOysnlKJC33U" crossorigin="anonymous"></script>
	<script src="https://unpkg.com/bootstrap-material-design@4.1.1/dist/js/bootstrap-material-design.js" integrity="sha384-CauSuKpEqAFajSpkdjv3z9t8E7RlpJ1UP0lKM/+NdtSarroVKu069AlsRPKkFBz9" crossorigin="anonymous"></script>
	<script>$(document).ready(function() { $('body').bootstrapMaterialDesign(); });</script>
</body>
</html>