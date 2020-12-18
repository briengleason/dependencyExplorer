package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Project struct {
	XMLName xml.Name `xml:"project"`
	Dependencies []Dependency `xml:"dependencies>dependency"`
	Properties Properties `xml:"properties"`
}

type Properties struct {
	Entries map[string]string
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	type entry struct {
		XMLName xml.Name
		Key     string `xml:"name,attr"`
		Value   string `xml:",chardata"`
	}
	e := entry{}
	p.Entries = map[string]string{}
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		e.Key = e.XMLName.Local
		p.Entries[e.Key] = e.Value
		fmt.Println(e.Key, e.Value)
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

type Property struct {
	XMLName xml.Name `xml:",any"`
	Value string `xml:",any"`
}

type Dependency struct {
	XMLName xml.Name `xml:"dependency"`
	GroupId string   `xml:"groupId"`
	ArtifactId string   `xml:"artifactId"`
	Version string   `xml:"version"`
	Scope string   `xml:"scope"`
}


func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("pom.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var project Project
	xml.Unmarshal(byteValue, &project)

	fmt.Println(project.Properties)

	//for i, dependencies := 0, project.Dependencies; i < len(project.Dependencies); i++ {
	//	fmt.Println("ArtifactId: " + dependencies[i].ArtifactId)
	//	fmt.Println("GroupId: " + dependencies[i].GroupId)
	//	fmt.Println("Version: " + dependencies[i].Version)
	//}

}