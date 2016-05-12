<html>
	<head>
		<title>Cigol Search</title>
	</head>
	<body>
		<form action="/search" method="post"> 
			<input type="text" name="q" id="q">
			<input type="submit" value="查询">
		</form>
		<script src="/static/jquery.min.js"></script>
		<script>
			$(document).ready(function(){
				$('#q')[0].focus()
			})
		</script>
	</body>
</html>