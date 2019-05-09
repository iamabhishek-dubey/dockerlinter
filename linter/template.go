package linter

const htmltemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<title>Dockerfile Auditor</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1"
<link rel="icon" type="image/png" href="https://iamabhishek-dubey.github.io/dockerlinter/reports/images/icons/favicon.ico"/>
<link rel="stylesheet" type="text/css" href="https://iamabhishek-dubey.github.io/dockerlinter/reports/vendor/bootstrap/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="https://iamabhishek-dubey.github.io/dockerlinter/reports/css/util.css">
<link rel="stylesheet" type="text/css" href="https://iamabhishek-dubey.github.io/dockerlinter/reports/css/main.css">
</head>
<body>
	<div class="container">
		<img src="https://iamabhishek-dubey.github.io/dockerlinter/reports/images/auditor.png">

		<h2 class="title">Dockerfile Auditor</h2>

	<table class="greyGridTable">
		<thead>
			<tr>
				<th>Line Number</th>
				<th>Line</th>
	            <th>Rule Code</th>
				<th style="width:35%">Description</th>
			</tr>
		</thead>
		<tfoot>
			<tr>
			<td></td>
			<td></td>
			<td></td>
			</tr>
		</tfoot>
		<tbody>
            {{.Text}}
		</tbody>
	</table>
</div>
</body>
</html>`
