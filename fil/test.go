	package fil

import (
    "encoding/json"
    "io/ioutil"
    "os"

)

// MyStruct is an example structure for this program.
type MyStruct struct {
    StructData string `json:"StructData"`
}

func Tmain() {
    filename := "myFile.json"
    err := checkFile(filename)
    if err != nil {
		print(err)
	}

    file, err := ioutil.ReadFile(filename)
    if err != nil {
		print(err)
    }

    data := []MyStruct{}

    // Here the magic happens!
    json.Unmarshal(file, &data)

    newStruct := &MyStruct{
        StructData: "peanut",
    }

    data = append(data, *newStruct)

    // Preparing the data to be marshalled and written.
    dataBytes, err := json.MarshalIndent(data, "", "")
    if err != nil {
		print(err)
    }

    err = ioutil.WriteFile(filename, dataBytes, 0644)
    if err != nil {
		print(err)
    }
}

func checkFile(filename string) error {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        _, err := os.Create(filename)
        if err != nil {
            return err
        }
    }
    return nil
}
