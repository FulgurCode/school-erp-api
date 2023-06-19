package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	Id                       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                     string             `json:"name,omitempty" bson:"name,omitempty"`
	Class                    int                `json:"class,omitempty" bson:"class,omitempty"`
	Gender                   string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Dob                      string             `json:"dob,omitempty" bson:"dob,omitempty"`
	AadhaarNo                int                `json:"aadhaarNo,omitempty" bson:"aadhaarNo,omitempty"`
	Course                   string             `json:"course,omitempty" bson:"course,omitempty"`
	ApplicationNo            int                `json:"applicationNo,omitempty" bson:"applicationNo,omitempty"`
	AdmissionNo              int                `json:"admissionNo,omitempty" bson:"admissionNo,omitempty"`
	AdmissionCategory        string             `json:"admissionCategory,omitempty" bson:"admissionCategory,omitempty"`
	AdmissionDate            string             `json:"admissionDate,omitempty" bson:"admissionDate,omitempty"`
	Rank                     int                `json:"rank,omitempty" bson:"rank,omitempty"`
	Wgpa                     float64            `json:"wgpa,omitempty" bson:"wgpa,omitempty"`
	Religion                 string             `json:"religion,omitempty" bson:"religion,omitempty"`
	Caste                    string             `json:"caste,omitempty" bson:"caste,omitempty"`
	Obc                      bool               `json:"obc,omitempty" bson:"obc,omitempty"`
	Category                 string             `json:"category,omitempty" bson:"category,omitempty"`
	NameOfParent             string             `json:"nameOfParent,omitempty" bson:"nameOfParent,omitempty"`
	RelationshipWithGuardian string             `json:"relationshipWithGuardian,omitempty" bson:"relationshipWithGuardian,omitempty"`
	AddressOfGuardian        string             `json:"addressOfGuardian,omitempty" bson:"addressOfGuardian,omitempty"`
	OccupationOfParent       string             `json:"occupationOfParent,omitempty" bson:"occupationOfParent,omitempty"`
	Phone                    int                `json:"phone,omitempty" bson:"phone,omitempty"`
	LinguisticMinority       string             `json:"linguisticMinority,omitempty" bson:"linguisticMinority,omitempty"`
	SecondLanguage           string             `json:"secondLanguage,omitempty" bson:"secondLanguage,omitempty"`
	Status                   string             `json:"status,omitempty" bson:"status,omitempty"`
	TcNumber                 string             `json:"tcNumber,omitempty" bson:"tcNumber,omitempty"`
	TcDate                   string             `json:"tcDate,omitempty" bson:"tcDate,omitempty"`
	TcSchool                 string             `json:"tcSchool,omitempty" bson:"tcSchool,omitempty"`
	SslcNameOfBoard          string             `json:"sslcNameOfBoard,omitempty" bson:"sslcNameOfBoard,omitempty"`
	SslcPassingTime          string             `json:"sslcPassingTime,omitempty" bson:"sslcPassingTime,omitempty"`
	SslcRegisterNo           int                `json:"sslcRegisterNo,omitempty" bson:"sslcRegisterNo,omitempty"`
	Verified                 bool               `json:"verified,omitempty" bson:"verified,omitempty"`
	Confirmed                bool               `json:"confirmed,omitempty" bson:"confirmed,omitempty"`
	Import                   bool               `json:"import,omitempty" bson:"import,omitempty"`
}
