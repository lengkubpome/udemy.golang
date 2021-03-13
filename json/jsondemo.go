package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	// วิธีที่ 1 Slice of pipe
	// body, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	return
	// }
	// dataStruct := Users{}
	// v := &dataStruct
	// json.Unmarshal(body, v)

	//วิธีที่ 1 dataStruct --> std.Output
	// result, err := json.MarshalIndent(dataStruct, "", "    ")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(result))

	// วิธีที่ 2 json.Decoder --> dataStruct แปลงข้อมูล json ให้อยู่ในรูป struct ที่เรารู้จัก
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Users{}
	dataStruct2 := Users{}
	jsonDecoder.Decode(&dataStruct)
	jsonDecoder.Decode(&dataStruct2)
	resp.Body.Close()

	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Kiattiphong Saetang" //เปลี่ยนชื่อตัวเองก่อนแปลงคืน

	//วิธีที่ 2 Encoder --> json
	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct)

}
