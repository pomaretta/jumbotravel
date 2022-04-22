package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   107,
		Green: 114,
		Blue:  128,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   209,
		Green: 213,
		Blue:  219,
	}
}

func getGreenColor() color.Color {
	return color.Color{
		Red:   29,
		Green: 255,
		Blue:  70,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 182,
		Blue:  215,
	}
}

func main() {

	darkGrayColor := getBlueColor()
	grayColor := getGrayColor()
	whiteColor := color.NewWhite()
	blueColor := getBlueColor()
	header := getHeader()
	contents := getContents()

	// Create MarotoPDF
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	// Set Margins
	m.SetPageMargins(10, 15, 10)

	// TODO: Header
	m.RegisterHeader(func() {
		m.Row(20, func() {

			// JumboTravel Logo
			m.Col(4, func() {
				_ = m.FileImage("assets/invoice/logo.png", props.Rect{
					Center:  true,
					Percent: 100,
				})
			})

			m.ColSpace(5)

			// Invoice QR
			m.Col(4, func() {
				m.QrCode("https://www.carlospomares.es", props.Rect{
					Center:  true,
					Percent: 100,
				})
			})
		})
	})

	// TODO: Footer
	// m.RegisterFooter(func() {
	// 	m.Row(5, func() {
	// 		m.Col(12, func() {
	// 			m.Text("JumboTravel, 2022", props.Text{
	// 				Top:   8,
	// 				Style: consts.Bold,
	// 				Size:  10,
	// 				Align: consts.Center,
	// 				Color: blueColor,
	// 			})
	// 		})
	// 	})
	// })

	// TODO: Invoice information
	m.Row(45, func() {

		topMargin := 8
		levelSpacing := 6
		colWidth := 2
		colSpace := 1

		// TODO: Agent Information
		m.Col(uint(colWidth), func() {

			m.Text("ASSISTANT", props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
			})

			m.Text("Pere Pons", props.Text{
				Top: float64(topMargin + levelSpacing),
			})

			m.Text("111111111", props.Text{
				Top:  float64(topMargin + levelSpacing*2),
				Size: 8,
			})

			m.Text("pere.pons@jumbotravel.com", props.Text{
				Top:  float64(topMargin + levelSpacing*3),
				Size: 8,
			})

		})

		m.ColSpace(uint(colSpace))

		m.Col(uint(colWidth), func() {

			m.Text("PROVIDER", props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
			})

			m.Text("Joan Piqué", props.Text{
				Top: float64(topMargin + levelSpacing),
			})

			m.Text("121212121", props.Text{
				Top:  float64(topMargin + levelSpacing*2),
				Size: 8,
			})

			m.Text("Barcelona (ES, BCN)", props.Text{
				Top:  float64(topMargin + levelSpacing*3),
				Size: 8,
			})

			m.Text("joan.pique@jumbotravel.com", props.Text{
				Top:  float64(topMargin + levelSpacing*4),
				Size: 8,
			})

		})

		m.ColSpace(uint(colSpace))

		m.Col(uint(colWidth), func() {

			// Invoice Data Headers
			m.Text("INVOICE #", props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
				Align: consts.Right,
			})

			m.Text("INVOICE DATE", props.Text{
				Top:   float64(topMargin + levelSpacing),
				Style: consts.Bold,
				Align: consts.Right,
			})

			m.Text("BOOKINGS", props.Text{
				Top:   float64(topMargin + levelSpacing*2),
				Style: consts.Bold,
				Align: consts.Right,
			})

			m.Text("TOTAL", props.Text{
				Top:   float64(topMargin + levelSpacing*3),
				Style: consts.Bold,
				Align: consts.Right,
			})

		})

		m.ColSpace(uint(colSpace))

		m.Col(uint(colWidth), func() {

			m.Text("JT000010422", props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
				Align: consts.Right,
				Color: blueColor,
			})

			m.Text("2022-04-05", props.Text{
				Top: float64(topMargin + levelSpacing),
				// Style: consts.Bold,
				Align: consts.Right,
			})

			m.Text("1", props.Text{
				Top: float64(topMargin + levelSpacing*2),
				// Style: consts.Bold,
				Align: consts.Right,
			})

			m.Text("75.5€", props.Text{
				Top: float64(topMargin + levelSpacing*3),
				// Style: consts.Bold,
				Align: consts.Right,
			})

		})

	})

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Bookings", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})

		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)

	// Invoice Items
	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	err := m.OutputFileAndClose("./invoice.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
}

func getHeader() []string {
	return []string{"Reference ID", "Product", "Quantity", "Price"}
}

func getContents() [][]string {
	return [][]string{
		{"", "Swamp", "12", "R$ 4,00"},
		{"", "Sorin, A Planeswalker", "4", "R$ 90,00"},
		{"", "Tassa", "4", "R$ 30,00"},
		{"", "Skinrender", "4", "R$ 9,00"},
		{"", "Island", "12", "R$ 4,00"},
		{"", "Mountain", "12", "R$ 4,00"},
		{"", "Plain", "12", "R$ 4,00"},
		{"", "Black Lotus", "1", "R$ 1.000,00"},
		{"", "Time Walk", "1", "R$ 1.000,00"},
		{"", "Emberclave", "4", "R$ 44,00"},
		{"", "Anax", "4", "R$ 32,00"},
		{"", "Murderous Rider", "4", "R$ 22,00"},
		{"", "Gray Merchant of Asphodel", "4", "R$ 2,00"},
		{"", "Ajani's Pridemate", "4", "R$ 2,00"},
		{"", "Renan, Chatuba", "4", "R$ 19,00"},
		{"", "Tymarett", "4", "R$ 13,00"},
		{"", "Doom Blade", "4", "R$ 5,00"},
		{"", "Dark Lord", "3", "R$ 7,00"},
		{"", "Memory of Thanatos", "3", "R$ 32,00"},
		{"", "Poring", "4", "R$ 1,00"},
		{"", "Deviling", "4", "R$ 99,00"},
		{"", "Seiya", "4", "R$ 45,00"},
		{"", "Harry Potter", "4", "R$ 62,00"},
		{"", "Goku", "4", "R$ 77,00"},
		{"", "Phreoni", "4", "R$ 22,00"},
		{"", "Katheryn High Wizard", "4", "R$ 25,00"},
		{"", "Lord Seyren", "4", "R$ 55,00"},
	}
}
