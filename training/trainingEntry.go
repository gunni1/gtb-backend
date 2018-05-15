package training

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TrainingEntry struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Location  string        `json:"location" bson:"location"`
	UserID    string        `json:"userId" bson:"userId"`
	BeginTime time.Time     `json:"beginTime" bson:"beginTime"`
	EndTime   time.Time     `json:"endTime" bson:"endTime"`
	Practices []Practice    `json:"practices" bson:"practices"`
}

type Practice struct {
	PracticeKey string `json:"practiceKey" bson:"practiceKey"`
	Reps        int64  `json:"reps" bson:"reps"`
	Duration    string `json:"duration" bson:"duration"`
}
