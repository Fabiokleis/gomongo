package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Disciplina struct {
	Codigo    string    `bson:"codigo"`
	Nome      string    `bson:"nome"`
	Professor string    `bson:"professor"`
	Registro  time.Time `bson:"registro"`
	Nota      *float64  `bson:"nota"`
	Salas     []string  `bson:"salas"`
	Horarios  []string  `bson:"horarios"`
}

// Aluno -> Disciplina
type Aluno struct {
	ID          bson.ObjectID `bson:"_id,omitempty"` // Maps to MongoDB's _id field
	Nome        string        `bson:"nome"`
	Matricula   string        `bson:"matricula"`
	Curso       string        `bson:"curso"`
	Email       string        `bson:"email"`
	Registro    time.Time     `bson:"registro"`
	Disciplinas []Disciplina  `bson:"disciplinas"`
}
