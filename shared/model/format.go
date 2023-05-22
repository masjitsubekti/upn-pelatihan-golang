package model

import (
	"time"

	"github.com/bojanz/currency"
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

const (
	// DefaultDateFormat digunakan untuk standar format input dari string ke format tanggal
	DateDMYFormat = "02/01/2006"

	// DefaultDateFormat digunakan untuk standar format input dari string ke format tanggal
	DefaultDateFormat = "2006-01-02"

	// DefaultDateFormat digunakan untuk standar format input dari string ke format tanggal dan jam
	DefaultDateTimeFormat = "2006-01-02 15:04:05"

	// DefaultTimeFormat digunakan untuk standar format input dari string ke format jam
	DefaultTimeFormat = "15:04:05"
)

// DecimalToRupiah parse decimal to rupiah format
func DecimalToRupiah(value decimal.Decimal) string {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(value)
}

// RupiahToDecimal parse rupiah format to decimal
func RupiahToDecimal(rupiah string) (value decimal.Decimal, err error) {
	locale := currency.NewLocale("id")
	formatter := currency.NewFormatter(locale)
	amount, err := formatter.Parse(rupiah, "IDR")
	if err != nil {
		return
	}
	value, err = decimal.NewFromString(amount.Number())
	if err != nil {
		return
	}
	return
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
