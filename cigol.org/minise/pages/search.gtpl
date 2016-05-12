<html>
	<body>
		<form action="/search" method="post">
			<input type="text" name="q" value="{{.Q}}" id="q">
			<input type="submit" value="查询">
		</form>
		<ul>
			{{with .Hits}}
			{{range .}}
			<li>
				<h3><a href="{{.Url}}" target="_blank">{{.Title}}</a></h3>
				<p style="font-size:12px;margin:5px 0 5px;">{{.Description}}</p>
				<p>{{.Content}}</p>
			</li>
			{{end}}
			{{end}}
		</ul>
		<script src="/static/jquery.min.js"></script>
		<script>
			$(document).ready(function(){
				$('#q')[0].focus()
			})
		</script>
	</body>
</html>