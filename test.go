package restgo

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func TestFile(file string, debug bool) error {
	testStage, err := loadTestFile(file)
	if err != nil {
		return err
	}

	return Evaluate(testStage, debug)
}

func loadTestFile(file string) (*TestStage, error) {
	fp, err := os.Open(file)
	if err != nil {
		fmt.Printf("open file error %v\n", err)
		return nil, err
	}
	defer fp.Close()

	var testStage TestStage
	d := yaml.NewDecoder(fp)
	if err := d.Decode(&testStage); err != nil {
		return nil, err
	}

	return &testStage, nil
}
