package main

import (
	"fmt"
	"time"
)

func eachFunc[T any](a []T, f func(T) bool) {
	for _, e := range a {
		if !f(e) {
			break
		}
	}
}

func main() {

	//"2023-02-01 09:30:40.840917"
	//"2023-02-01 09:30:06.690695"
	t1, _ := time.Parse("02/Jan/2006:15:04:05 -0700", "2023-02-01 09:30:40.840917")

	fmt.Println(t1)

	//t2 := time.Now()

	//fmt.Println(time.Before(time2))

	//eachFunc([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
	//	fmt.Println(v)
	//	return v < 4
	//})
	//
	//eachFunc([]string{"foo", "bar", "", "zoo"}, func(v string) bool {
	//	fmt.Println(v)
	//	return v != ""
	//})

	//eachFunc([]interface{"foo", "bar", "", "zoo"}, func(v string) bool {
	//	fmt.Println(v)
	//	return v != ""
	//})

	//lines := []string{
	//	//`\u0053`,
	//	`\u0075`,
	//	`\u006E`,
	//	`\uD83D`,
	//	"vgfg \u0000",
	//	//"{\n         \"version\": \"5.7.1\",\n         \"payload\": {\n           \"messageId\": \"40930e9b-b58b-4f7e-8c53-aae536526d68\",\n           \"type\": 5,\n           \"data\": {\n             \"directionName\": \"ПФР\",\n             \"docflowType\": 118,\n             \"docflowTypeCode\": \"18\",\n             \"docflowTypeName\": \"СЗВ-ТД\",\n             \"transactionType\": 2,\n             \"transactionTypeCode\": \"02\",\n             \"transactionTypeName\": \"УОД\",\n             \"programVersion\": \"Астрал Отчет 5.0 5.0.24.3\",\n             \"isPrimary\": false,\n             \"specOperatorUrl\": \"http://aispfr.keydisk.ru\",\n             \"documents\": [\n               {\n                 \"sourceFile\": {\n                   \"fileId\": \"764c6d67-55c4-482a-934b-94de97296d19\",\n                   \"documentId\": \"00000000-0000-0000-0000-000000000000\",\n                   \"name\": \"ПФР_050-026-004654_101000_УОД_20211119_e0598e34-0672-46ed-af89-d00b65aab64a.xml\",\n                   \"size\": 1165\n                 },\n                 \"documentId\": \"e0598e34-0672-46ed-af89-d00b65aab64a\",\n                 \"isMain\": true,\n                 \"documentTypeCode\": \"02\",\n                 \"documentTypeName\": \"УОД\",\n                 \"documentType\": 2,\n                 \"isCompressed\": true,\n                 \"isEncrypt\": false,\n                 \"isContent\": true,\n                 \"signFormat\": 5,\n                 \"content\": {\n                   \"fileId\": \"e0598e34-0672-46ed-af89-d00b65aab64a\",\n                   \"fileName\": \"ПФР_050-026-004654_101000_УОД_20211119_e0598e34-0672-46ed-af89-d00b65aab64a.xml.gz\",\n                   \"type\": \"application/xml\",\n                   \"size\": 511\n                 },\n                 \"signFiles\": []\n               }\n             ],\n             \"packageFile\": {\n               \"fileId\": \"584a8617-a104-45ae-add5-51683499a93c\",\n               \"documentId\": \"00000000-0000-0000-0000-000000000000\",\n               \"name\": \"ПФР_073-030-012251_101000_УОД_20211119_584a8617-a104-45ae-add5-51683499a93c.zip\",\n               \"size\": 6284\n             },\n             \"operator\": {\n               \"code\": \"050-026-004654\",\n               \"storageType\": 0,\n               \"cryptoProviderType\": 0,\n               \"id\": \"00000000-0000-0000-0000-000000000000\",\n               \"name\": \"АО \\"КАЛУГА АСТРАЛ\\"\"\n             },\n             \"receiver\": {\n               \"subjectId\": \"D1B09545F2CA67374409F27EB409912E10CAE544\",\n               \"certificateId\": \"c9f972b1-11b1-4228-948a-055fd4a6bad1\",\n               \"storageType\": 4,\n               \"cryptoProviderType\": 19,\n               \"id\": \"3292e18a-c727-e611-9ccb-005056a85b01\",\n               \"name\": \"Глава КФХ Маняхин В.В.\"\n             },\n             \"sender\": {\n               \"code\": \"00\",\n               \"subjectId\": \"DB1CE8224FFF07638BF15A95B451E2F13E07C5AD\",\n               \"storageType\": 0,\n               \"cryptoProviderType\": 0,\n               \"id\": \"c2a4e273-a4cc-4949-bbf3-5b2a898c437c\",\n               \"name\": \"АИС ПФР (ЭТК)\",\n               \"email\": \"Gotsutsov@101.pfr.ru\",\n               \"addresseeType\": 20\n             },\n             \"isDelayed\": false,\n             \"abonentId\": \"3292e18a-c727-e611-9ccb-005056a85b01\",\n             \"directionId\": \"3e8a6a23-d325-4c74-9bb5-e83c0bd4de4e\",\n             \"docflowId\": \"3bc8091e-8378-4992-a714-c264002e157f\",\n             \"gateId\": \"a2c54ee0-0bd4-496b-8ffb-b7c56dfa640c\",\n             \"isInbound\": true,\n             \"startedAt\": \"2021-11-18T21:35:00.313787\",\n             \"traceId\": \"584a8617-a104-45ae-add5-51683499a93c\"\n           }\n         },\n         \"payloadType\": \"RoDispatcher.Application.Services.DispatchingCenterParts.Models.MqCoreMessage\",\n         \"sender\": \"report-online-dispatcher-88d6d6dd8-94dbt\",\n         \"createdAt\": \"2021-11-18T21:35:01.6346772Z\"\n       }",
	//}
	//fmt.Println(lines)
	//
	//str := "Пояснение на требование от 20.04.2022г. № 15204.\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000 Исполнитель Демьянова И.В. \u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000 \u0000\u0000 \u0000\u0000 \u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000(3452) 29-61-58"
	//str = strings.Replace(str, "\u0000", "", -1)
	//
	//fmt.Println(str)

	//for i, v := range lines {
	//fmt.Printf("%U %U\n", i, v)
	//var err error
	//lines[i], err = strconv.Unquote(`"` + v + `"`)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//}
	//fmt.Println(lines)

	//high, low := utf16.EncodeRune('😀')
	//pairStr := string([]rune{high, low})
	//ascii := strconv.QuoteToASCII(pairStr)
	//fmt.Println(ascii)
}
