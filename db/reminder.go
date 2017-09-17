package db

import (
	pb "github.com/baopham/reminder"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func FindReminder(id string, s *mgo.Session) (*pb.ReminderResponse, error) {
	session, c := reminderDB(s)
	defer session.Close()

	var reminder pb.ReminderResponse

	err := c.FindId(bson.ObjectIdHex(id)).One(&reminder)

	if err != nil {
		log.Println("failed to find: ", err)
		return nil, err
	}

	return &reminder, nil
}

func CreateReminder(r *pb.CreateRequest, s *mgo.Session) (string, error) {
	session, c := reminderDB(s)
	defer session.Close()

	id := bson.NewObjectId()
	_, err := c.UpsertId(id, r)

	if err != nil {
		log.Println("failed to save: ", err)
		return "", err
	}

	return id.Hex(), nil
}

func ReminderIndices() []mgo.Index {
	return []mgo.Index{
		{
			Key:    []string{"remindat"},
			Sparse: true,
		},
		{
			Key:    []string{"userid"},
			Sparse: true,
		},
		{
			Key:    []string{"tags"},
			Sparse: true,
		},
	}
}

func reminderDB(s *mgo.Session) (*mgo.Session, *mgo.Collection) {
	session := s.Copy()
	c := session.DB("gominder").C("reminders")
	return session, c
}
