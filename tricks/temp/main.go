package main

import (
	"encoding/json"
	"fmt"
)

type PoaCreateB2GRequest struct {
	StartDate           string                `json:"ДатаНачалаДействия"`
	EndDate             string                `json:"ДатаОкончанияДействия"`
	RegistrationCodeSTA string                `json:"КодНОРегистрацииМЧД"`
	ActionsCodeSTA      []string              `json:"КодыНОДействияМЧД,omitempty"`
	Principal           *PoaPrincipal         `json:"СведенияОДоверителе"`
	Representative      *PoaRepresentativeB2G `json:"СведенияОПредставителе"`
}

type PoaRepresentativeB2G struct {
	Type       string           `json:"Тип"`
	Name       string           `json:"Наименование,omitempty"`
	INN        string           `json:"ИНН,omitempty"`
	OGRN       string           `json:"ОГРН,omitempty"`
	KPP        string           `json:"КПП,omitempty"`
	Individual PoaIndividualB2G `json:"ДанныеПредставителя"`
}

type PoaIndividualB2G struct {
	FIO              *PoaFIO           `json:"ФИО"`
	BirthDate        string            `json:"ДатаРождения"`
	INNFL            string            `json:"ИННФЛ"`
	OGRNIP           string            `json:"ОГРНИП,omitempty"`
	Snils            string            `json:"СНИЛС"`
	Identification   PoaIdentification `json:"ДокументУдостоверяющийЛичность"`
	EmpowermentCodes []string          `json:"СписокПолномочий"`
}

type PoaFIO struct {
	Lastname   string `json:"Фамилия"`
	Firstname  string `json:"Имя"`
	Patronymic string `json:"Отчество,omitempty"`
}

type PoaAddress struct {
	Postcode  string `json:"Индекс,omitempty"`
	Region    string `json:"КодРегиона"`
	District  string `json:"Район,omitempty"`
	City      string `json:"Город,omitempty"`
	Subregion string `json:"НасПункт,omitempty"`
	Street    string `json:"Улица,omitempty"`
	Building  string `json:"Дом,omitempty"`
	Pavilion  string `json:"Корпус,omitempty"`
	Apartment string `json:"Квартира,omitempty"`
}

type PoaManagingOrg struct {
	Name string `json:"Наименование,omitempty"`
	INN  string `json:"ИНН"`
	OGRN string `json:"ОГРН"`
	KPP  string `json:"КПП"`
}

type PoaIdentification struct {
	Type         string `json:"Тип"`
	NumberSeries string `json:"СерияНомер"`
	Unit         string `json:"КодПодразделения,omitempty"`
	IssueDate    string `json:"ДатаВыдачи"`
	IssueBy      string `json:"КемВыдан,omitempty"`
}

type PoaIndividual struct {
	FIO            *PoaFIO            `json:"ФИО,omitempty"`
	BirthDate      string             `json:"ДатаРождения,omitempty"`
	BirthPlace     string             `json:"МестоРождения,omitempty"`
	Citizenship    string             `json:"Гражданство,omitempty"`
	INNFL          string             `json:"ИННФЛ,omitempty"`
	OGRNIP         string             `json:"ОГРНИП,omitempty"`
	Position       string             `json:"Должность,omitempty"`
	Nationality    string             `json:"Национальность,omitempty"`
	Gender         string             `json:"Пол,omitempty"`
	Snils          string             `json:"СНИЛС,omitempty"`
	CountryCode    string             `json:"КодСтраныГражданства,omitempty"`
	AuthorityDoc   string             `json:"ДокументПолномочий,omitempty"`
	Identification *PoaIdentification `json:"ДокументУдостоверяющийЛичность,omitempty"`
}

type PoaPrincipal struct {
	Type               string          `json:"Тип"`
	Name               string          `json:"Наименование,omitempty"`
	INN                string          `json:"ИНН,omitempty"`
	OGRN               string          `json:"ОГРН,omitempty"`
	KPP                string          `json:"КПП,omitempty"`
	CountryCode        string          `json:"КодСтраныРегистрации,omitempty"`
	RegistrationNumber string          `json:"РегистрационныйНомер,omitempty"`
	Postal             *PoaAddress     `json:"АдресЮридический,omitempty"`
	ManagingOrg        *PoaManagingOrg `json:"СведенияОбУправляющейОрганизации,omitempty"`
	Signatory          *PoaIndividual  `json:"ДанныеПодписанта,omitempty"`
}

func main() {

	js := []byte("{\n  \"КодыНОДействияМЧД\": [\n    \"9999\"\n  ],\n  \"ДатаНачалаДействия\": \"2023-01-29\",\n  \"КодНОРегистрацииМЧД\": \"9999\",\n  \"СведенияОДоверителе\": {\n    \"ИНН\": \"9648670241\",\n    \"КПП\": \"999901581\",\n    \"Тип\": \"ЮЛ\",\n    \"ОГРН\": \"8634321388407\",\n    \"Наименование\": \"Тест_Грузоотправитель_ЭЗЗ\",\n    \"АдресЮридический\": {\n      \"Дом\": \"4\",\n      \"Город\": \"Калуга\",\n      \"Улица\": \"Тестовая\",\n      \"Индекс\": \"248000\",\n      \"КодРегиона\": \"40\"\n    },\n    \"ДанныеПодписанта\": {\n      \"Пол\": \"М\",\n      \"ФИО\": {\n        \"Имя\": \"Грузоотправитель\",\n        \"Фамилия\": \"Тест\",\n        \"Отчество\": \"ЭЗЗ\"\n      },\n      \"ИННФЛ\": \"967810566209\",\n      \"СНИЛС\": \"800-141-086 23\",\n      \"ДатаРождения\": \"2000-01-01\",\n      \"Национальность\": \"RUS\"\n    }\n  },\n  \"ДатаОкончанияДействия\": \"2024-01-29\",\n  \"СведенияОПредставителе\": {\n    \"ИНН\": \"9693940154\",\n    \"КПП\": \"999901581\",\n    \"Тип\": \"ЮЛ\",\n    \"ОГРН\": \"1908134309708\",\n    \"Наименование\": \"Тестовый ЮЛ\",\n    \"ДанныеПредставителя\": {\n      \"ФИО\": {\n        \"Имя\": \"Тест\",\n        \"Фамилия\": \"Тестов\",\n        \"Отчество\": \"Тестович\"\n      },\n      \"ИННФЛ\": \"964762216372\",\n      \"СНИЛС\": \"192-197-994 21\",\n      \"ДатаРождения\": \"2000-01-02\",\n      \"СписокПолномочий\": [\n        \"04\"\n      ],\n      \"ДокументУдостоверяющийЛичность\": {\n        \"Тип\": \"ПаспортРФ\",\n        \"КемВыдан\": \"МВД\",\n        \"ДатаВыдачи\": \"2015-01-02\",\n        \"СерияНомер\": \"2002100100\",\n        \"КодПодразделения\": \"202-100\"\n      }\n    }\n  }\n}")

	arg := new(PoaCreateB2GRequest)
	if err := json.Unmarshal(js, arg); err != nil {
		fmt.Println(err)
	}

	fmt.Println(arg.Representative.Individual.EmpowermentCodes)

}
