# Convert-JSON-API-to-CSV-Golang
Here you will find a easy program to convert a JSON response API to a CSV File

## Packages required
No need to install any package explictly. Just copy paste the follwoing import in main() of main.go.

```golang
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
```

## Reading from a JSON response API

```golang
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
```
## Reading from a JSON response File

```golang

	// Open our jsonFile
	jsonFile, err := os.Open("OrderDetails.json")
  
	// Handle error then handle it
	if err != nil {
	    fmt.Println(err)
	}
	fmt.Println("Successfully Opened OrderDetails.json")
	
  // defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	
  // read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	
  //initialize struct
	var Order Orders
	
  // jsonFile's content into 'Order' which we defined above
	err = json.Unmarshal(byteValue, &Order)
	
  if err != nil {
	    fmt.Println(err)
	}
```
  
  
## Creating Writing to a CSV File

```golang

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
	
  ```

### Note
You can change the delimeter '|' of .csv format by ctrl + Click on writer.Write() function


#### Connect
[Shubham Snehi](https://in.linkedin.com/in/shubham-snehi-a62bb5189)
