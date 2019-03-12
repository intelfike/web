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
	<nav class="navbar">
		<ul class="list-inline m-0">
			<li class="list-inline-item m-1"><a href="/_top/index">トップ</a></li>
			<li class="list-inline-item m-1"><a data-toggle="collapse" href="#usage-links" role="button" aria-expanded="false">使い方</a></li>
			<li class="list-inline-item m-1"><a href="/_install/index">インストール</a></li>
			<li class="list-inline-item m-1"><a href="/_support/index">サポート</a></li>
			<li class="list-inline-item m-1"><a href="/_thanks/index">Thanks</a></li>
		</ul>
	</nav>
	<div id="usage-links" class="collapse">
		<nav class="navbar">
			<ul class="list-group m-0">
				<li class="list-item m-1"><a href="/_usage/functions">機能</a></li>
				<li class="list-item m-1"><a href="/_usage/settings">設定</a></li>
				<li class="list-item m-1"><a href="/_usage/name">名称・用語</a></li>
				<li class="list-item m-1"><a href="/_usage/sKey">ショートカットキー</a></li>
				<li class="list-item m-1"><a href="/_usage/sWord">検索テキスト</a></li>
				<li class="list-item m-1"><a href="/_usage/commandMode">コマンドモード</a></li>
			</ul>
		</nav>
	</div>
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

	<!-- V4 version -->
	<!-- <script type="text/javascript" src="https://cdn.jsdelivr.net/npm/bootstrap.native@2.0.15/dist/bootstrap-native-v4.min.js"></script> -->
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://unpkg.com/popper.js@1.12.6/dist/umd/popper.js" integrity="sha384-fA23ZRQ3G/J53mElWqVJEGJzU0sTs+SvzG8fXVWP+kJQ1lwFAOkcUOysnlKJC33U" crossorigin="anonymous"></script>
	<script src="https://unpkg.com/bootstrap-material-design@4.1.1/dist/js/bootstrap-material-design.js" integrity="sha384-CauSuKpEqAFajSpkdjv3z9t8E7RlpJ1UP0lKM/+NdtSarroVKu069AlsRPKkFBz9" crossorigin="anonymous"></script>
	<script>$(document).ready(function() { $('body').bootstrapMaterialDesign(); });</script>
</body>
</html>