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
	Views      dbdata.Views
}

func FillHTML(procedures dbdata.Procedures, functions dbdata.Functions, tables dbdata.Tables, views dbdata.Views) {

	var elements Elements

	elements.Tables = tables
	elements.Views = views
	elements.Functions = functions
	elements.Procedures = procedures

	fillDetails(elements)
	fillMain(elements)

	// ON RELEASE MODE
	// exePath, err := os.Executable()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// ON CODE MODE
	// file, err := os.Create("templates/output.html")
	// if err != nil {
	// 	panic(err)
	// }

	// ON RELEASE MODE
	// file, err := os.Create(filepath.Dir(exePath) + "\\" + "output.html")
	// if err != nil {
	// 	panic(err)
	// }

	//defer file.Close()

}

func fillMain(elements Elements) {
	mainHtml, err := template.ParseFiles("templates/mainTmpl.html")
	if err != nil {
		panic(err)
	}

	filename := "main.html"

	file, err := os.Create("templates/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = mainHtml.Execute(file, elements)
	if err != nil {
		panic(err)
	}
}

func fillDetails(elements Elements) {
	// Table details
	tableDetail, err := template.ParseFiles("templates/tableTmpl.html")
	if err != nil {
		panic(err)
	}

	for _, obj := range elements.Tables.Data {
		filename := obj.ObjectName + ".html"

		file, err := os.Create("templates/details/" + filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = tableDetail.Execute(file, obj)
		if err != nil {
			panic(err)
		}
	}

	// View details
	viewDetail, err := template.ParseFiles("templates/viewTmpl.html")
	if err != nil {
		panic(err)
	}

	for _, obj := range elements.Views.Data {
		filename := obj.ObjectName + ".html"

		file, err := os.Create("templates/details/" + filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = viewDetail.Execute(file, obj)
		if err != nil {
			panic(err)
		}
	}

	// Function details
	funcionDetail, err := template.ParseFiles("templates/functionTmpl.html")
	if err != nil {
		panic(err)
	}

	for _, obj := range elements.Functions.Data {
		filename := obj.ObjectName + ".html"

		file, err := os.Create("templates/details/" + filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = funcionDetail.Execute(file, obj)
		if err != nil {
			panic(err)
		}
	}

	// Procedure details
	procDetail, err := template.ParseFiles("templates/procedureTmpl.html")
	if err != nil {
		panic(err)
	}

	for _, obj := range elements.Procedures.Data {
		filename := obj.ObjectName + ".html"

		file, err := os.Create("templates/details/" + filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = procDetail.Execute(file, obj)
		if err != nil {
			panic(err)
		}
	}
}
