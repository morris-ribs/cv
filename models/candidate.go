package models

import
(
  "gopkg.in/mgo.v2/bson"
)

type Project struct {
    Name    string            `json:"name"`
    Description string        `json:"description"`
}

type Projects []Project

type NameStruct struct {
    Name string `json:"name"`
}
type Names []NameStruct

type Technologies struct {
    Backend Names           `json:"backend"`
    Frontend Names          `json:"frontend"`
    Database Names          `json:"database"`
    ControlVersion Names    `json:"controlversion"`
}

type Experience struct {
    CompanyName string          `json:"companyname"`
    CompanyDescription string   `json:"companydescription"`
    Location string             `json:"location"`
	Period	 string             `json:"period"`
	Projects   Projects	    `json:"projects"`
	Technologies Technologies `json:"technologies"`
}

type Experiences []Experience

type Sections struct {
    Title string            `json:"jobtitle"`
    NatTitle string         `json:"nationalitiestitle"`
    Nationalities string    `json:"nationalities"`
	IntExp  string          `json:"intexp"`
    ProfExp string          `json:"profexp"`
    Technologies string     `json:"technos"`
    Database string         `json:"database"`
    ControlVersion string   `json:"controlversion"`
    Education string        `json:"education"`
    Languages string        `json:"languages"`
    Hobbies string          `json:"hobbies"`
}

type Background struct {
    SchoolName string            `json:"schoolname"`
    Location string            `json:"location"`
    Period string            `json:"period"`
    Degree string            `json:"degree"`
    Description string            `json:"description"`					
}
type Education []Background

type Language struct {
    Name      string          `json:"name"`
    Level     string          `json:"level"`
}
type Languages []Language

type FullCv struct {
    Lang            string       `json:"lang"`
    Sections        Sections     `json:"sections"`
    Experiences     Experiences  `json:"experiences"`
    Education       Education    `json:"education"`
    Languages       Names        `json:"languages"` 
}

type FullCvs []FullCv

type Address struct {
  Adr1      string    `json:"adr1"`
  Adr2     string    `json:"adr2"`
}

type Candidate struct {
  Id        bson.ObjectId   `bson:"_id"`
  Name      string          `json:"name"`
  Email     string          `json:"email"`
  Nick      string          `json:"nick"`
  Phone     string          `json:"phone"`
  Address   Address         `json:"address"`
  FullCvs   FullCvs         `json:"fullcvs"`
}


type Candidates []Candidate
