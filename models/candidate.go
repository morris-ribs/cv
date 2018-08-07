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


type Experience struct {
  CompanyName string          `json:"companyname"`
  CompanyDescription string   `json:"companydescription"`
  Location string             `json:"location"`
	Period	 string             `json:"period"`
	Projects   Projects	        `json:"projects"`
	Technologies Names          `json:"technologies"`
}

type Experiences []Experience

type Background struct {
  SchoolName string        `json:"schoolname"`
  Location string          `json:"location"`
  Period string            `json:"period"`
  Degree string            `json:"degree"`
  Description string       `json:"description"`					
}
type Education []Background

type NameLevel struct {
    Name      string          `json:"name"`
    Level     string          `json:"level"`
}
type NameLevels []NameLevel

type FullCv struct {
  Lang            string       `json:"lang"`
  Experiences     Experiences  `json:"experiences"`
  Education       Education    `json:"education"`
  Languages       NameLevels   `json:"languages"` 
}

type FullCvs []FullCv

type Address struct {
  Adr1      string    `json:"adr1"`
  Adr2     string    `json:"adr2"`
}

type Candidate struct {
  Id        bson.ObjectId   `bson:"_id"`
  Key       string          `json:"key"`
  Name      string          `json:"name"`
  Email     string          `json:"email"`
  Nick      string          `json:"nick"`
  Phone     string          `json:"phone"`
  LinkedIn  string          `json:"linkedin"`
  Github    string          `json:"github"`
  Blog      string          `json:"blog"`
  Address   Address         `json:"address"`
  FullCvs   FullCvs         `json:"fullcvs"`
  Skills    NameLevels      `json:"skills"`
}


type Candidates []Candidate
