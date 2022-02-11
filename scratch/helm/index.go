package helm

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"helm.sh/helm/v3/pkg/repo"
)

func encodeIndex(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	gz := gzip.NewWriter(buf)
	if err := json.NewEncoder(gz).Encode(data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	// fmt.Println(buf.String())

	return buf.Bytes(), nil
}

func decodeIndex(data []byte) error {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("1: %s", err)
	}
	defer gz.Close()

	data, err = ioutil.ReadAll(gz)
	if err != nil {
		return fmt.Errorf("11: %s", err)
	}

	index := &repo.IndexFile{}
	if err := json.Unmarshal(data, index); err != nil {
		return fmt.Errorf("2: %s", err)
	}

	fmt.Println("----------+++---------------")
	fmt.Println(index)

	return nil
}

// func decodeIndex(data []byte) error {
// 	buf := &bytes.Buffer{}
// 	_, err := buf.Write(data)
// 	if nil != err {
// 		return fmt.Errorf("1: %s", err)
// 	}
// 	var out json.RawMessage
// 	gz, err := gzip.NewReader(buf)
// 	if nil != err {
// 		return fmt.Errorf("2: %s", err)
// 	}

// 	if err := json.NewDecoder(gz).Decode(&out); err != nil {
// 		return fmt.Errorf("3: %s", err)
// 	}

// 	if err := gz.Close(); err != nil {
// 		return err
// 	}

// 	fmt.Println("----------+++---------------")
// 	fmt.Println(string(out))

// 	// index := &repo.IndexFile{}
// 	// if err := yaml.Unmarshal([]byte(out), index); err != nil {
// 	// 	return errors.Wrapf(err, "parse out")
// 	// }

// 	return nil
// }
