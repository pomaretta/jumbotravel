package invoice

import (
	"bytes"
	"fmt"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Creator struct {
	Invoice *Invoice
	maroto  pdf.Maroto
}

type InvoiceResult struct {
	Invoice *Invoice
	Buffer  *bytes.Buffer
}

func New(invoice *Invoice) *Creator {
	return &Creator{
		Invoice: invoice,
		maroto:  pdf.NewMaroto(consts.Portrait, consts.A4),
	}
}

func (c *Creator) createHeader() (err error) {
	c.maroto.RegisterHeader(func() {
		c.maroto.Row(20, func() {
			c.maroto.Col(4, func() {
				err = c.maroto.FileImage("assets/invoice/logo.png", props.Rect{
					Center:  true,
					Percent: 100,
				})
			})
			c.maroto.ColSpace(5)
			c.maroto.Col(4, func() {
				if c.Invoice.SignatureUrl != "" {
					c.maroto.QrCode(c.Invoice.SignatureUrl, props.Rect{
						Center:  true,
						Percent: 100,
					})
				}
			})
		})
	})
	return
}

func (c *Creator) createInformation() (err error) {

	m := c.maroto

	topMargin := 8
	levelSpacing := 6
	colWidth := 2
	colSpace := 1

	m.Row(45, func() {
		m.Col(uint(colWidth), func() {
			m.Text("ASSISTANT", props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
			})
			m.Text(c.Invoice.Assistant.Name, props.Text{
				Top: float64(topMargin + levelSpacing),
			})
			m.Text(c.Invoice.Assistant.DNI, props.Text{
				Top:  float64(topMargin + levelSpacing*2),
				Size: 8,
			})
			m.Text(c.Invoice.Assistant.Email, props.Text{
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
			m.Text(c.Invoice.Provider.Name, props.Text{
				Top: float64(topMargin + levelSpacing),
			})
			m.Text(c.Invoice.Provider.DNI, props.Text{
				Top:  float64(topMargin + levelSpacing*2),
				Size: 8,
			})
			m.Text(c.Invoice.Provider.Airport, props.Text{
				Top:  float64(topMargin + levelSpacing*3),
				Size: 8,
			})
			m.Text(c.Invoice.Provider.Email, props.Text{
				Top:  float64(topMargin + levelSpacing*4),
				Size: 8,
			})
		})
		m.ColSpace(uint(colSpace))
		m.Col(uint(colWidth), func() {
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
			m.Text(fmt.Sprintf("%d", c.Invoice.Id), props.Text{
				Top:   float64(topMargin),
				Style: consts.Bold,
				Align: consts.Right,
				Color: blue,
			})
			m.Text(c.Invoice.Date.Format("02/01/2006"), props.Text{
				Top: float64(topMargin + levelSpacing),
				// Style: consts.Bold,
				Align: consts.Right,
			})
			m.Text(fmt.Sprintf("%d", len(c.Invoice.Bookings)), props.Text{
				Top: float64(topMargin + levelSpacing*2),
				// Style: consts.Bold,
				Align: consts.Right,
			})
			m.Text(fmt.Sprintf("%.2f€", c.Invoice.Total()), props.Text{
				Top: float64(topMargin + levelSpacing*3),
				// Style: consts.Bold,
				Align: consts.Right,
			})
		})
	})

	return
}

func getTableHeaders() []string {
	return []string{
		"Reference ID",
		"Product",
		"Quantity",
		"Price",
	}
}

func (c *Creator) createTable() (err error) {

	m := c.maroto
	headers := getTableHeaders()
	content := c.Invoice.GetContent()

	m.SetBackgroundColor(blue)
	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Bookings", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: white,
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(white)
	m.TableList(headers, content, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &gray,
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
			m.Text(fmt.Sprintf("%0.2f EUR€", c.Invoice.Total()), props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	return
}

func (c *Creator) Create() (*InvoiceResult, error) {

	c.maroto.SetPageMargins(10, 15, 10)

	// Header
	err := c.createHeader()
	if err != nil {
		return nil, err
	}

	// Information
	err = c.createInformation()
	if err != nil {
		return nil, err
	}

	// Table
	err = c.createTable()
	if err != nil {
		return nil, err
	}

	buffer, err := c.maroto.Output()
	if err != nil {
		return nil, err
	}

	return &InvoiceResult{
		Invoice: c.Invoice,
		Buffer:  &buffer,
	}, nil
}
