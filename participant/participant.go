package participant

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//Participant - data model
type Participant struct {
	TeamName                          string   `bson:"teamName" json:"teamName"`
	Password                          string   `bson:"password" json:"password"`
	ProjectInfo                       string   `bson:"projectInfo" json:"projectInfo"`
	TeamPlayers                       []string `bson:"teamPlayers" json:"teamPlayers"`
	University                        string   `bson:"university" json:"university"`
	SoftwareOrProgrammingLanguageUsed string   `bson:"softwareOrProgrammingLanguageUsed" json:"softwareOrProgrammingLanguageUsed"`
	HardwareUsed                      string   `bson:"hardwareUsed" json:"hardwareUsed"`
}
