# google sheet

Used to pull data and then code gne from

URL = https://docs.google.com/spreadsheets/d/16eeYgh8dus50fISokKK8TMVWLR8A18Er-p5dBcO0FYw/edit?usp=drive_web&ouid=108808652634038924311
spreadsheetId = 16eeYgh8dus50fISokKK8TMVWLR8A18Er-p5dBcO0FYw


shtsrv := //authenticated Sheets Service

//values is a [] []interface{}

values, err := shtsrv.Spreadsheets.Values.Get("spreadsheetID", "sheetName").Do()

//handle the error

//printthe first value from the first row

fmt.Println(values[0][0].(string))