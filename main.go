package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//JSON Api Struct
type Orders struct {
	OrdersummaryBySupplierid struct {
		Supplier struct {
			Name     string `json:"name"`
			City     string `json:"city"`
			State    string `json:"state"`
			Pincode  string `json:"pincode"`
			Phoneno  string `json:"phoneNo"`
			Mobile   string `json:"mobile"`
			Username string `json:"username"`
		} `json:"supplier"`
		Orderstatus []struct {
			Orderdate string `json:"orderDate"`
			Billed    string `json:"billed"`
			Bounced   string `json:"bounced"`
			Pending   string `json:"pending"`
		} `json:"orderStatus"`
	} `json:"ordersummaryBySupplierId"`
}

func main() {

	//Api URL
	url := "http://localhost:3000/orderdata"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	//Api GET call
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	Order := Orders{}
	//Un Marshalling JSON response content to 'Order'
	jsonErr := json.Unmarshal(body, &Order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("Successfully read JSON API")

	//----------- CAN ALSO DO WITH JSON FILE INCLUDED IN THIS FOLDER,
	//----------- COMMENT ABOVE CODE AND UNCOMMENT THE CODE BELOW

	// // Open our jsonFile
	// jsonFile, err := os.Open("OrderDetails.json")
	// // Handle error then handle it
	// if err != nil {
	//     fmt.Println(err)
	// }
	// fmt.Println("Successfully Opened OrderDetails.json")
	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()
	// // read our opened jsonFile as a byte array.
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// //initialize struct
	// var Order Orders
	// // jsonFile's content into 'Order' which we defined above
	// err = json.Unmarshal(byteValue, &Order)
	// if err != nil {
	//     fmt.Println(err)
	// }

	//Create CSV File
	csvFile, err := os.Create("./OrderDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	var row []string
	row = append(row, Order.OrdersummaryBySupplierid.Supplier.City)
	row = append(row, Order.OrdersummaryBySupplierid.Supplier.Name)
	row = append(row, Order.OrdersummaryBySupplierid.Supplier.State)
	writer.Write(row)

	//Appending nested struct
	for _, emp := range Order.OrdersummaryBySupplierid.Orderstatus {
		var row1 []string
		row1 = append(row1, emp.Orderdate)
		row1 = append(row1, emp.Billed)
		row1 = append(row1, emp.Bounced)
		row1 = append(row1, emp.Pending)
		writer.Write(row1)
	}
	// remember to flush!
	writer.Flush()
}
