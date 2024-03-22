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
			</div>
		{{end}}
		{{range .Functions.Data}}
		<div id="info-{{.Ident}}" class="info-item">
			<h2>{{.ObjectName}}</h2>
			<p>To są informacje na temat {{.ObjectName}}.</p>
		</div>
		{{end}}
		{{range .Procedures.Data}}
		<div id="info-{{.Ident}}" class="info-item">
			<h2>{{.ObjectName}}</h2>
			<p>To są informacje na temat {{.ObjectName}}.</p>
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
