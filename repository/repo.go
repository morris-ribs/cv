package repository

import
(
  "log"
  "cv-server-rest-go/models"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

var currentId int

var candidates models.Candidates

const SERVER = "localhost:27017"
const DBNAME = "dbcandidates"
const DOCNAME = "candidates"

// Get all candidates
func GetCandidates() models.Candidates {
  session, err := mgo.Dial(SERVER)
  defer session.Close()
  c := session.DB(DBNAME).C(DOCNAME)
  results := models.Candidates{}
  c.Find(nil).All(&results)
  if err != nil {
      log.Fatal(err)
  }
  return results
}

// Get a candidate by id
func GetCandidateById(id string) models.Candidate {
	session, err := mgo.Dial(SERVER)
  defer session.Close()
  c := session.DB(DBNAME).C(DOCNAME)

  // Grab id
  oid := bson.ObjectIdHex(id)

  result := models.Candidate{}
  err = c.FindId(oid).One(&result)
  if err != nil {
      log.Fatal(err)
  }
      
  // return empty Candidate if not found
  return result
}

// Get a candidate by key
func GetCandidateByKey(key string) models.Candidate {
	session, err := mgo.Dial(SERVER)
  defer session.Close()
  c := session.DB(DBNAME).C(DOCNAME)

  result := models.Candidate{}
  err = c.Find(bson.M{"key": key}).One(&result)
  if err != nil {
      log.Fatal(err)
  }
      
  // return empty Candidate if not found
  return result
}

// Create a candidate
func CreateCandidate(candidate models.Candidate) models.Candidate {
  session, err := mgo.Dial(SERVER)
  defer session.Close()

  candidate.Id = bson.NewObjectId()
  session.DB(DBNAME).C(DOCNAME).Insert(candidate)

  if err != nil {
      log.Fatal(err)
  }
  return candidate
}

// Delete a candidate
func DeleteCandidate(id string) string {
  session, err := mgo.Dial(SERVER)
  defer session.Close()

  // Verify id is ObjectId, otherwise bail
  if !bson.IsObjectIdHex(id) {
      return "404"
  }

  // Grab id
  oid := bson.ObjectIdHex(id)

  // Remove user
  if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
      log.Fatal(err)
  }

  // Write status
  return ""
}
