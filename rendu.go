package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("197c8a6a", "hCNsZ4oaAFXdUKWP")
	client := nexmo.NewClient(http.DefaultClient, auth)
	insight, _, err := client.Insight.GetBasicInsight(nexmo.BasicInsightRequest{
		Number: os.Args[1],
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Le nom du pays d'origine:", insight.CountryName)
	fmt.Println("Numéro local:", insight.NationalFormatNumber)
	fmt.Println("Numéro international:", insight.InternationalFormatNumber)
}


//  func third()  {
	// 	fmt.Println("message : ")
	//     jsonData := map[string]
	//     jsonValue, _ := json.Marshal(jsonData)
	//     response, err = http.Post("https://rest.nexmo.com/sms/{json}", os.Args[2] "application/json", bytes.NewBuffer(jsonValue))
	//     if err != nil {
	//         fmt.Printf("Api request failed with error %s\n", err)
	//     } else {
	//         data, _ := ioutil.ReadAll(response.Body)
	//         fmt.Println(string(data))
	//     }
	// }



// func second()  {
// 	fmt.Println("Appel en cours......")
//     jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
//     jsonValue, _ := json.Marshal(jsonData)
//     response, err = http.Post("https://api.nexmo.com/v1/calls", os.Args[1] "application/json", bytes.NewBuffer(jsonValue))
//     if err != nil {
//         fmt.Printf("The HTTP request failed with error %s\n", err)
//     } else {
//         data, _ := ioutil.ReadAll(response.Body)
//         fmt.Println(string(data))
//     }
// }