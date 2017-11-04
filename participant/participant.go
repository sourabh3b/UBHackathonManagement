package participant

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Participant - data model
type Participant struct {
	TeamName                          string   `bson:"_id" json:"teamName"`
	Password                          string   `bson:"password" json:"password"`
	ProjectInfo                       string   `bson:"projectInfo" json:"projectInfo"`
	TeamPlayers                       []string `bson:"teamPlayers" json:"teamPlayers"`
	University                        string   `bson:"university" json:"university"`
	SoftwareOrProgrammingLanguageUsed string   `bson:"softwareOrProgrammingLanguageUsed" json:"softwareOrProgrammingLanguageUsed"`
	HardwareUsed                      string   `bson:"hardwareUsed" json:"hardwareUsed"`
}

//GetParticipant - handler to get expenses
func GetParticipant(teamName string) (Participant,error) {
	participantObject := Participant{}
	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return participantObject, errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	// Collection Expense
	err = session.DB("UBHacking").C("Participant").Find(bson.M{"_id": teamName}).One(&participantObject)
	if err != nil {
		fmt.Println("Unable to find participantObject by ID", err.Error())
		return participantObject, errors.New("Unable to find participantObject by ID " + err.Error())
	}

	return participantObject, err
}