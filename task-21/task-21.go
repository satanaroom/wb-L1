package main

/*Реализовать паттерн «адаптер» на любом примере.*/

import "fmt"

// Смысл работы паттерна Адаптер в том, что если существующий Тип и его интерфейс не совместим
// с кодом какой-то системы, то, необходимо не изменять код этого Типа, а написать для него адаптер
// Минусы паттерна:
// 1. сокрытие деталей от клиента;
// 2. увеличение количества кода

// +++ Обявляем пакет для имитации подключаемого функционала, работающего только с XML данными
type AnalyticalDataService interface {
	SendXMLData()
}

type XMLData struct {
	Data string
}

func (doc *XMLData) SendXMLData() {
	fmt.Println("Отправка XML данных", doc.Data)
}
// ---

// Обявляем Адаптер, в который инжектим структуру JSON данных
type JSONDataAdapter struct {
	jsonData *JSONData
}

// Метод адаптера, который конвертирует данные из JSON в XML и отправляет сконвертированные данные
func (adapter *JSONDataAdapter) SendXMLData() {
	adapter.jsonData.ConvertToXML()
	fmt.Println("Отправка XML данных", adapter.jsonData.Data)
}

type JSONData struct {
	Data string
}

// Метод для конвертирования данных
func (doc *JSONData) ConvertToXML() {
	doc.Data = "<xml>" + doc.Data + "</xml>"
	fmt.Println("data was converted by xml.Marshal(&doc.Data)")
}

// +++ Обявляем клиента для работы с методом отправки данных в формате XML
type client struct {
}

func (c *client) SendXMLDataToService(data AnalyticalDataService) {
	data.SendXMLData()
}
// ---

func main() {
	// Инициалиализируем клиента
	client := &client{}
	// Инициалиализируем данные в формате XML
	xmlDoc := &XMLData{
		Data: "<xml>Some XML data</xml>",
	}
	// Вызваем метод клиента, в который подаем данные в формате XML
	client.SendXMLDataToService(xmlDoc)
	jsonDoc := &JSONData{
		Data: "Some JSON data",
	}
	// Инициалиализируем адаптер
	jsonDocAdapter := &JSONDataAdapter{
		jsonData: jsonDoc,
	}
	// Вызваем метод клиента, в который подаем данные в формате JSON, конвертированные в XML
	client.SendXMLDataToService(jsonDocAdapter)
}
