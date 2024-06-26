package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	dbdata "devopsowy.pl/dbdoc/db_data"
)

type Elements struct {
	Tables     dbdata.Tables
	Procedures dbdata.Procedures
	Functions  dbdata.Functions
	Views      dbdata.Views
}

func FillHTML(procedures dbdata.Procedures, functions dbdata.Functions, tables dbdata.Tables, views dbdata.Views) {

	var counter int64 = 1

	for i := range tables.Data {
		if i < len(tables.Data) {
			tables.Data[i].Ident = counter
			counter++
		} else {
			break
		}
	}

	for i := range views.Data {
		if i < len(views.Data) {
			views.Data[i].Ident = counter
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
	elements.Views = views
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
		.normal-th th {
			font-weight: normal;
		}
		.sticky-header {
			position: sticky;
			top: 0;
			background-color: #FFFFFF; /* Kolor nagłówka */
			z-index: 1000; /* Upewnij się, że nagłówek jest nad treścią */
		}
		
		.sticky-header::after {
			content: '';
			position: absolute;
			bottom: -3px; /* Dodany margines w dół */
			left: 0;
			width: 100%;
			height: 1px;
			background-color: #ccc; /* Kolor linii oddzielającej */
			z-index: 999; /* Upewnij się, że linia jest nad treścią */
		}
		
		.sticky-header::before {
			content: '';
			position: absolute;
			top: 100%; /* Umieść przed linią */
			left: 0;
			width: 100%;
			height: 8%; /* Wysokość równa wysokości nagłówka */
			background-color: white; /* Kolor obszaru między linią a nagłówkiem */
			z-index: 998; /* Upewnij się, że obszar jest nad treścią */
		}
		table.table-fit {
  width: auto !important;
  table-layout: auto !important;
}
table.table-fit thead th,
table.table-fit tbody td,
table.table-fit tfoot th,
table.table-fit tfoot td {
  width: auto !important;
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
				<summary>Widoki</summary>
				<ul id="viewsList">
				{{range .Views.Data}}
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
				<h2 class="sticky-header">{{.ObjectName}}</h2>
				<p>To są informacje na temat {{.ObjectName}}.</p>
			<table class="table table-hover table-bordered">
			  <tbody>
			  <tr class="table-secondary">
			    <th>Name</th>
				<th>Data Type</th>
				<th>Length</th>
				<th>Nullable</th>
			  </tr>
			  {{range .Columns}}
			  <tr class="normal-th">
			    <th>{{.Name}}</th>
				<th>{{.DataType}}</th>
				<th>{{.MaxLength}}</th>
				<th>{{.IsNullable}}</th>
			  </tr>
			  {{end}}
			  </tbody>
			</table>
			<table class="table table-hover table-bordered table-100">
			<tbody>
			<tr class="table-secondary">
			  <th>Name</th>
			  <th>Columns</th>
			  <th>Type</th>
			  <th>Unique</th>
			  <th>Primary</th>
			</tr>
			{{range .Indexes}}
			<tr class="normal-th">
			  <th>{{.Name}}</th>
			  <th>{{.Columns}}</th>
			  <th>{{.Type}}</th>
			  <th>{{.IsUnique}}</th>
			  <th>{{.IsPrimary}}</th>
		    </tr>
			{{end}}
			</tbody>
			</table>
			<table class="table table-hover table-bordered table-100">
			<tbody>
			<tr class="table-secondary">
			  <th>Name</th>
			  <th>Referenced table</th>
			  <th>Columns</th>
			  <th>Referenced columns</th>
			</tr>
			{{range .ForeignKeys}}
			<tr class="normal-th">
			  <th>{{.Name}}</th>
			  <th>{{.ReferencedTableName}}</th>
			  <th>{{.ForeignKeyColumns}}</th>
			  <th>{{.ReferencedColumns}}</th>
		    </tr>
			{{end}}
			</tbody>
			</table>
			</div>
		{{end}}
		{{range .Views.Data}}
		<div id="info-{{.Ident}}" class="info-item">
			<h2>{{.ObjectName}}</h2>
			<p>To są informacje na temat {{.ObjectName}}.</p>
			<table class="table table-hover table-bordered table-100">
			<tbody>
			<tr class="table-secondary">
			  <th>Name</th>
			  <th>Data Type</th>
			</tr>
			{{range .Columns}}
			<tr class="normal-th">
			  <th>{{.Name}}</th>
			  <th>{{.DataType}}</th>
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
			<tr class="table-secondary">
			  <th>Name</th>
			  <th>Data Type</th>
			  <th>Length</th>
			  <th>Output</th>
			</tr>
			{{range .Parameters}}
			<tr class="normal-th">
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
			<tr class="table-secondary">
			  <th>Name</th>
			  <th>Data Type</th>
			  <th>Length</th>
			  <th>Nullable</th>
			</tr>
			{{range .Parameters}}
			<tr class="normal-th">
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

	// ON RELEASE MODE
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ON CODE MODE
	// file, err := os.Create("templates/output.html")
	// if err != nil {
	// 	panic(err)
	// }

	// ON RELEASE MODE
	file, err := os.Create(filepath.Dir(exePath) + "\\" + "output.html")
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
