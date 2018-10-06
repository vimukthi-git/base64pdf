package main

import (
	"log"
	"os"

	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/jmoiron/jsonq"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "jsonfile, f",
			Value: "",
			Usage: "json file containing a base64 pdf in a field",
		},
		cli.StringFlag{
			Name:  "jsonpath, p",
			Value: "",
			Usage: "dot separated path to base64 string inside the json file provided. eg: root.array.0.base64field",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "output.pdf",
			Usage: "output pdf filename",
		},
	}

	app.Action = func(c *cli.Context) error {
		jsonFile := c.String("jsonfile")
		if jsonFile == "" {
			return errors.New("jsonfile is required")
		}
		jsonPath := c.String("jsonpath")
		if jsonPath == "" {
			return errors.New("jsonpath is required")
		}
		data, err := extractBase64(jsonFile, jsonPath)
		if err != nil {
			return err
		}
		return createPdf(data, c.String("output"))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func extractBase64(filename, jsonpath string) (string, error) {
	path := strings.Split(jsonpath, ".")
	data := map[string]interface{}{}
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	dec := json.NewDecoder(f)
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	return jq.String(path...)
}

func createPdf(data, outputFile string) error {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
