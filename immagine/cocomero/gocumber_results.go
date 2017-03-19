package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
)

type Urls struct {
	Head        string
	Devel30     string
	HeadFile    string
	Devel30File string
}

type report struct {
	numFailed     int
	nameFailed    []string
	oldnumFailed  int
	oldNameFailed []string
}

var cucumber = Urls{"http://m226.mgr.suse.de/workspace/manager-Head-sumaform-cucumber/last.html",
	"http://m226.mgr.suse.de/workspace/manager-3.0-sumaform30/last.html",
	"last.html",
	"last30.html"}

var rep report

//full diff
func fulldifference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

// this function compare and return elements that differe from 2 slices
// we need this for see if we have a regression

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop only 1 time first to find slice1 strings not in slice2
	for i := 0; i < 1; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
	}

	return diff
}

// this function get the latest cucumber result
func (url Urls) GetHeadOutput() {
	cmd := exec.Command("wget", url.Head, "-O", url.HeadFile)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(url.Head)
		panic(err)

	}
}

// just read the results.html and return a string
func (url Urls) ReadLastResults() string {
	data, err := ioutil.ReadFile(url.HeadFile)
	if err != nil {
		panic(err)
	}
	str := string(data)
	return str
}

// get name and number of failed steps
func (r *report) getFailedSteps(output string) {
	var nameFailed []string
	failedSteps := regexp.MustCompile("step failed")
	indexesFail := failedSteps.FindAllStringIndex(output, -1)
	for _, index := range indexesFail {
		nameFailed = append(nameFailed, strings.Replace(strings.Fields(strings.Split(output[index[0]-500:index[0]], "li id=")[1])[0], "'", "", -1))
	}
	r.numFailed = len(indexesFail)
	r.nameFailed = nameFailed
}

// dump type report into json
// this function ovveride knowfailures with new failures
func (r *report) dumpReportJson() {
	b, errj := json.Marshal(r.nameFailed)
	if errj != nil {
		log.Fatal(errj)
	}
	err := ioutil.WriteFile(".knowfailures.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

// read the json file
func (r *report) readReportJson() {
	raw, err := ioutil.ReadFile(".knowfailures.json")
	if _, err := os.Stat(".knowfailures.json"); os.IsNotExist(err) {
		empty := []byte("")
		err := ioutil.WriteFile(".knowfailures.json", empty, 0644)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal("read file fail json")
	}
	errj := json.Unmarshal(raw, &r.oldNameFailed)
	if errj != nil {
		log.Fatal("unmarshal json failed!\n", errj)
	}
}
func main() {
	cucumber.GetHeadOutput()
	output := cucumber.ReadLastResults()
	rep.getFailedSteps(output)
	rep.readReportJson()
	rep.dumpReportJson()
	// compare oldNameFailed and NameFailed
	if reflect.DeepEqual(rep.oldNameFailed, rep.nameFailed) {
		fmt.Println("no new errors")
	} else {
		fmt.Println("New regressions")
		fmt.Println(difference(rep.nameFailed, rep.oldNameFailed))
	}
}
