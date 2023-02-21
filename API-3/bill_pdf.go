package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	// "github.com/divrhino/fruitful-pdf/data"
)

type Temp struct {
	Rate float64 `json:"rate"`
}

var total float64 = 0.0

func generate_bill() {
	total = 0.0
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildList(m)
	buildFooter(m)
	err := m.OutputFileAndClose("pdfs/" + uname + ".pdf")
	if err != nil {
		fmt.Println("could not save Bill...", err)
		os.Exit(1)
	}
	fmt.Println("Bill generated successfully....")
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo.jpeg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}
			})
		})
	})
	m.Row(15, func() {
		m.Col(12, func() {
			m.Text("Your Favourite Food delivery Partner....", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Size:  15,
				Color: getOrangeColor(),
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text("User       : "+strings.Title(strings.ToLower(uname)), props.Text{
				Top:   1,
				Align: consts.Left,
				Color: getOrangeColor(),
				Size:  8,
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text("Address : "+address, props.Text{
				Top:   1,
				Align: consts.Left,
				Color: getOrangeColor(),
				Size:  8,
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text("Contact  : "+contact, props.Text{
				Top:   1,
				Align: consts.Left,
				Color: getOrangeColor(),
				Size:  8,
			})
		})
	})
	m.Row(7, func() {
		m.Col(12, func() {
			m.Text("Time      : "+time.Now().Format("2006-01-02 15:04:05 Monday"), props.Text{
				Top:   1,
				Align: consts.Left,
				Color: getOrangeColor(),
				Size:  8,
			})
		})
	})

}
func buildList(m pdf.Maroto) {
	tableHeadings := []string{"Item", "Price", "Quantity", "Hotel", "Net Price"}
	// contents := data.FruitList(20)
	contents := get_data()
	lightGreyColor := getGreyColor()
	m.SetBackgroundColor(getOrangeColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Order Details", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})

	})
	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 2, 2, 3, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 2, 2, 3, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightGreyColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

}

func buildFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {

		m.Row(10, func() {
			m.SetBackgroundColor(getOrangeColor())

			m.Col(10, func() {
				s := "Total amount(including 40Rupees delivery charge..)====>>  "
				m.Text(s, props.Text{
					Top:    2,
					Size:   11,
					Color:  color.NewWhite(),
					Family: consts.Courier,
					Style:  consts.Bold,
					Align:  consts.Left,
				})
			})
			m.Col(2, func() {
				s := fmt.Sprintf("%.2f", total+40.0)
				m.Text(s, props.Text{
					Top:    2,
					Size:   17,
					Color:  color.NewWhite(),
					Family: consts.Courier,
					Style:  consts.Bold,
					Align:  consts.Left,
				})
			})
		})
	})
}

func getOrangeColor() color.Color {
	return color.Color{
		Red:   232,
		Green: 116,
		Blue:  30,
	}
}
func getGreyColor() color.Color {
	return color.Color{
		Red:   242,
		Green: 242,
		Blue:  242,
	}
}

func get_data() [][]string {

	data := [][]string{}
	for v1, v2 := range cart {
		temp_data := []string{}
		for k1, k2 := range v2 {
			temp_data = append(temp_data, v1)
			f := db_data(v1, k1)
			np := calc_price(k2, f)
			temp_data = append(temp_data, fmt.Sprintf("%.2f", f))
			temp_data = append(temp_data, strconv.FormatUint(k2, 32))
			temp_data = append(temp_data, k1)
			temp_data = append(temp_data, fmt.Sprintf("%.2f", np))
			data = append(data, temp_data)
		}
	}
	return data
}

func db_data(item string, hotel string) float64 {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	fs := fmt.Sprintf("SELECT price FROM %s WHERE item = ?", hotel)
	result, err := db.Query(fs, item)
	if err != nil {
		panic(err.Error())
	}
	var f float64
	for result.Next() {
		var temp Temp
		err = result.Scan(&temp.Rate)
		if err != nil {
			panic(err.Error())
		}
		f = temp.Rate
	}
	return f
}

func calc_price(a uint64, f float64) float64 {
	total = total + (f * float64(a))
	return f * float64(a)
}
