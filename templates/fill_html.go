package templates

import (
	"os"
	"text/template"

	dbdata "devopsowy.pl/dbdoc/db_data"
)

type Elements struct {
	Tables     dbdata.Tables
	Procedures dbdata.Procedures
	Functions  dbdata.Functions
}

func FillHTML(procedures dbdata.Procedures, functions dbdata.Functions, tables dbdata.Tables) {

	var counter int64 = 1

	for i := range tables.Data {
		if i < len(tables.Data) {
			tables.Data[i].Ident = counter
			counter++
		} else {
			break
		}
	}

	for i := range functions.Data {
		if i < len(functions.Data) {
			functions.Data[i].Ident = counter
			counter++
		} else {
			break
		}
	}

	for i := range procedures.Data {
		if i < len(procedures.Data) {
			procedures.Data[i].Ident = counter
			counter++
		} else {
			break
		}
	}

	var elements Elements

	elements.Tables = tables
	elements.Functions = functions
	elements.Procedures = procedures

	tpl := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Spis Treści</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
	<style>
		#content {
			display: flex;
		}
		#toc {
			flex: 1;
			padding: 20px;
			background-color: #f0f0f0;
		}
		#info {
			flex: 3;
			padding: 20px;
		}
		.info-item {
			display: none;
		}
	</style>
	</head>
	<body>
	
	<div id="content">
		<div id="toc">
			<h2>Spis Treści</h2>
			<input type="text" id="searchInput" onkeyup="searchContent()" placeholder="Szukaj...">
			<details>
				<summary>Tabele</summary>
				<ul id="tablesList">
				{{range .Tables.Data}}
					<li><a href="#" onclick="showInfo({{.Ident}})">{{.ObjectName}}</a></li>
				{{end}}
				</ul>
			</details>
			<details>
				<summary>Funkcje</summary>
				<ul id="functionsList">
				{{range .Functions.Data}}
					<li><a href="#" onclick="showInfo({{.Ident}})">{{.ObjectName}}</a></li>
				{{end}}
				</ul>
			</details>
			<details>
				<summary>Procedury</summary>
				<ul id="proceduresList">
				{{range .Procedures.Data}}
					<li><a href="#" onclick="showInfo({{.Ident}})">{{.ObjectName}}</a></li>
				{{end}}
				</ul>
			</details>
		</div>
		<div id="info">
		{{range .Tables.Data}}
			<div id="info-{{.Ident}}" class="info-item">
				<h2>{{.ObjectName}}</h2>
				<p>To są informacje na temat {{.ObjectName}}.</p>
			<table class="table table-hover table-bordered table-100">
			  <tbody>
			  <tr>
			    <th>Name</th>
				<th>Data Type</th>
				<th>Length</th>
				<th>Nullable</th>
			  </tr>
			  {{range .Columns}}
			  <tr>
			    <th>{{.Name}}</th>
				<th>{{.DataType}}</th>
				<th>{{.MaxLength}}</th>
				<th>{{.IsNullable}}</th>
			  </tr>
			  {{end}}
			  </tbody>
			</table>
			</div>
		{{end}}
		{{range .Functions.Data}}
		<div id="info-{{.Ident}}" class="info-item">
			<h2>{{.ObjectName}}</h2>
			<p>To są informacje na temat {{.ObjectName}}.</p>
			<table class="table table-hover table-bordered table-100">
			<tbody>
			<tr>
			  <th>Name</th>
			  <th>Data Type</th>
			  <th>Length</th>
			  <th>Output</th>
			</tr>
			{{range .Parameters}}
			<tr>
			  <th>{{.Name}}</th>
			  <th>{{.DataType}}</th>
			  <th>{{.MaxLength}}</th>
			  <th>{{.IsOutput}}</th>
			</tr>
			{{end}}
			</tbody>
		  </table>
		</div>
		{{end}}
		{{range .Procedures.Data}}
		<div id="info-{{.Ident}}" class="info-item">
			<h2>{{.ObjectName}}</h2>
			<p>To są informacje na temat {{.ObjectName}}.</p>
			<table class="table table-hover table-bordered table-100">
			<tbody>
			<tr>
			  <th>Name</th>
			  <th>Data Type</th>
			  <th>Length</th>
			  <th>Nullable</th>
			</tr>
			{{range .Parameters}}
			<tr>
			  <th>{{.Name}}</th>
			  <th>{{.DataType}}</th>
			  <th>{{.MaxLength}}</th>
			  <th>{{.IsOutput}}</th>
			</tr>
			{{end}}
			</tbody>
		  </table>
		</div>
		{{end}}
		</div>
	</div>
	
	<script>
		function showInfo(infoId) {
			// Ukryj wszystkie pozycje informacji
			var infoItems = document.getElementsByClassName("info-item");
			for (var i = 0; i < infoItems.length; i++) {
				infoItems[i].style.display = "none";
			}
			// Pokaż wybraną pozycję informacji
			var selectedInfo = document.getElementById("info-" + infoId);
			if (selectedInfo) {
				selectedInfo.style.display = "block";
			}
		}
	
		function searchContent() {
			var input, filter, ul, li, a, i, txtValue;
			input = document.getElementById('searchInput');
			filter = input.value.toUpperCase();
			ul = document.getElementById("toc");
			li = ul.getElementsByTagName('li');
			for (i = 0; i < li.length; i++) {
				a = li[i].getElementsByTagName("a")[0];
				txtValue = a.textContent || a.innerText;
				if (txtValue.toUpperCase().indexOf(filter) > -1) {
					li[i].style.display = "";
				} else {
					li[i].style.display = "none";
				}
			}
		}
	</script>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.8/dist/umd/popper.min.js" integrity="sha384-I7E8VVD/ismYTF4hNIPjVp/Zjvgyol6VFvRkX/vR+Vc4jQkC+hVqc2pM8ODewa9r" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.min.js" integrity="sha384-0pUGZvbkm6XF6gxjEnlmuGrJXVbNuzT9qBBavbLwCsOGabYfZo0T0to5eqruptLy" crossorigin="anonymous"></script>
	
	</body>
	</html>`

	file, err := os.Create("templates/output.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := template.Must(template.New("HTMLTemplate").Parse(tpl))
	err = t.Execute(file, elements)
	if err != nil {
		panic(err)
	}

}
