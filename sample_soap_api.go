package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type BrinkWebServiceRes struct {
  XMLName         xml.Name `xml:"definitions"`
  Text            string   `xml:",chardata"`
  Name            string   `xml:"name,attr"`
  TargetNamespace string   `xml:"targetNamespace,attr"`
  Wsdl            string   `xml:"wsdl,attr"`
  Wsx             string   `xml:"wsx,attr"`
  Wsu             string   `xml:"wsu,attr"`
  Wsa10           string   `xml:"wsa10,attr"`
  Wsp             string   `xml:"wsp,attr"`
  Wsap            string   `xml:"wsap,attr"`
  Msc             string   `xml:"msc,attr"`
  Soap12          string   `xml:"soap12,attr"`
  Wsa             string   `xml:"wsa,attr"`
  Wsam            string   `xml:"wsam,attr"`
  Xsd             string   `xml:"xsd,attr"`
  Tns             string   `xml:"tns,attr"`
  Soap            string   `xml:"soap,attr"`
  Wsaw            string   `xml:"wsaw,attr"`
  Soapenc         string   `xml:"soapenc,attr"`
  Types           struct {
    Text   string `xml:",chardata"`
    Schema struct {
      Text            string `xml:",chardata"`
      TargetNamespace string `xml:"targetNamespace,attr"`
      Import          []struct {
        Text           string `xml:",chardata"`
        SchemaLocation string `xml:"schemaLocation,attr"`
        Namespace      string `xml:"namespace,attr"`
      } `xml:"import"`
    } `xml:"schema"`
  } `xml:"types"`
  Message []struct {
    Text string `xml:",chardata"`
    Name string `xml:"name,attr"`
    Part struct {
      Text    string `xml:",chardata"`
      Name    string `xml:"name,attr"`
      Element string `xml:"element,attr"`
      Q1      string `xml:"q1,attr"`
      Q2      string `xml:"q2,attr"`
      Q3      string `xml:"q3,attr"`
      Q4      string `xml:"q4,attr"`
      Q5      string `xml:"q5,attr"`
      Q6      string `xml:"q6,attr"`
      Q7      string `xml:"q7,attr"`
      Q8      string `xml:"q8,attr"`
      Q9      string `xml:"q9,attr"`
      Q10     string `xml:"q10,attr"`
      Q11     string `xml:"q11,attr"`
      Q12     string `xml:"q12,attr"`
      Q13     string `xml:"q13,attr"`
      Q14     string `xml:"q14,attr"`
      Q15     string `xml:"q15,attr"`
      Q16     string `xml:"q16,attr"`
      Q17     string `xml:"q17,attr"`
      Q18     string `xml:"q18,attr"`
      Q19     string `xml:"q19,attr"`
      Q20     string `xml:"q20,attr"`
      Q21     string `xml:"q21,attr"`
      Q22     string `xml:"q22,attr"`
      Q23     string `xml:"q23,attr"`
      Q24     string `xml:"q24,attr"`
      Q25     string `xml:"q25,attr"`
      Q26     string `xml:"q26,attr"`
      Q27     string `xml:"q27,attr"`
      Q28     string `xml:"q28,attr"`
      Q29     string `xml:"q29,attr"`
      Q30     string `xml:"q30,attr"`
      Q31     string `xml:"q31,attr"`
      Q32     string `xml:"q32,attr"`
      Q33     string `xml:"q33,attr"`
      Q34     string `xml:"q34,attr"`
      Q35     string `xml:"q35,attr"`
      Q36     string `xml:"q36,attr"`
      Q37     string `xml:"q37,attr"`
      Q38     string `xml:"q38,attr"`
      Q39     string `xml:"q39,attr"`
      Q40     string `xml:"q40,attr"`
      Q41     string `xml:"q41,attr"`
      Q42     string `xml:"q42,attr"`
    } `xml:"part"`
  } `xml:"message"`
  PortType struct {
    Text      string `xml:",chardata"`
    Name      string `xml:"name,attr"`
    Operation []struct {
      Text  string `xml:",chardata"`
      Name  string `xml:"name,attr"`
      Input struct {
        Text    string `xml:",chardata"`
        Action  string `xml:"Action,attr"`
        Message string `xml:"message,attr"`
      } `xml:"input"`
      Output struct {
        Text    string `xml:",chardata"`
        Action  string `xml:"Action,attr"`
        Message string `xml:"message,attr"`
      } `xml:"output"`
      Fault []struct {
        Text    string `xml:",chardata"`
        Action  string `xml:"Action,attr"`
        Name    string `xml:"name,attr"`
        Message string `xml:"message,attr"`
      } `xml:"fault"`
    } `xml:"operation"`
  } `xml:"portType"`
  Binding struct {
    Text    string `xml:",chardata"`
    Name    string `xml:"name,attr"`
    Type    string `xml:"type,attr"`
    Binding struct {
      Text      string `xml:",chardata"`
      Transport string `xml:"transport,attr"`
    } `xml:"binding"`
    Operation []struct {
      Text      string `xml:",chardata"`
      Name      string `xml:"name,attr"`
      Operation struct {
        Text       string `xml:",chardata"`
        SoapAction string `xml:"soapAction,attr"`
        Style      string `xml:"style,attr"`
      } `xml:"operation"`
      Input struct {
        Text string `xml:",chardata"`
        Body struct {
          Text string `xml:",chardata"`
          Use  string `xml:"use,attr"`
        } `xml:"body"`
      } `xml:"input"`
      Output struct {
        Text string `xml:",chardata"`
        Body struct {
          Text string `xml:",chardata"`
          Use  string `xml:"use,attr"`
        } `xml:"body"`
      } `xml:"output"`
      Fault []struct {
        Text  string `xml:",chardata"`
        Name  string `xml:"name,attr"`
        Fault struct {
          Text string `xml:",chardata"`
          Name string `xml:"name,attr"`
          Use  string `xml:"use,attr"`
        } `xml:"fault"`
      } `xml:"fault"`
    } `xml:"operation"`
  } `xml:"binding"`
  Service struct {
    Text string `xml:",chardata"`
    Name string `xml:"name,attr"`
    Port struct {
      Text    string `xml:",chardata"`
      Name    string `xml:"name,attr"`
      Binding string `xml:"binding,attr"`
      Address struct {
        Text     string `xml:",chardata"`
        Location string `xml:"location,attr"`
      } `xml:"address"`
    } `xml:"port"`
  } `xml:"service"`
} 

func main() {

  url := "***********************************************" // api url need add here
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("locationToken", "**************************") // locationToken need to add here
  req.Header.Add("accessToken", "**************************") // accessToken need to add here

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  soapWebServiceRes := &BrinkWebServiceRes{}
 
  _ = xml.Unmarshal([]byte(body), &soapWebServiceRes)

  // soap API response 
 
 fmt.Println("soapWebServiceRes.Service.Text ",soapWebServiceRes.Service.Text)
 fmt.Println("soapWebServiceRes.Service.Name",soapWebServiceRes.Service.Name)

 fmt.Println("soapWebServiceRes.Service.Port.Text ",soapWebServiceRes.Service.Port.Text)

 fmt.Println("soapWebServiceRes.Service.Port.Name ",soapWebServiceRes.Service.Port.Name)
 fmt.Println("soapWebServiceRes.Service.Port.Binding ", soapWebServiceRes.Service.Port.Binding)
 fmt.Println("soapWebServiceRes.Service.Port.Address.Text ",soapWebServiceRes.Service.Port.Address.Text)
 fmt.Println("soapWebServiceRes.Service.Port.Address.Location ",soapWebServiceRes.Service.Port.Address.Location)

}


