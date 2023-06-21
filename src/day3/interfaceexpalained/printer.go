package main

type PrinterInterface interface {
	Print()

	PageSetup()

	Preview()
}

func PrintObject(printObject PrinterInterface) {
	printObject.Print()
}

func PageSetupObject(printObject PrinterInterface) {
	printObject.PageSetup()
}

func PreviewObject(printerObject PrinterInterface) {
	printerObject.Preview()
}
