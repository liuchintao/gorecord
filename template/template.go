package main

import (
	"os"
	"text/template"
)

type orderData struct {
	ExtraHeader string
	Name        string
	OrderNumber int
	ShipDate    string
}

var header = `
    WidgetCo, Ltd.
    463 Shoe Factory Rd.
    Hamford, VT 20202
	{{ .ExtraHeader }}
`
var footer = `
    Thank you for your business,
    WidgetCo Order Fulfillment Department
    Ph: 818-555-0123 Email: orders@widgetco.com
`
var thanks = `
    {{ template "header" }}
    Dear {{ .Name }},
    Thank you for your order! Your order number is {{ .OrderNumber }} and will be shipped on {{ .ShipDate }}.
    {{ template "footer" }}
`

func main() {
	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("thanks").Parse(thanks)
	ordersToThank := []orderData{
		{"EXTRA-1", "Sleve McDichael", 17104, "2018-10-10"},
		{"EXTRA-2", "Bobson Dugnutt", 17106, "2018-10-12"},
	}
	for _, data := range ordersToThank {
		t.ExecuteTemplate(os.Stdout, "thanks", data)
	}
}
