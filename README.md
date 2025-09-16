# gomongo

NoSQL - exercícios básicos 01

foi escrito uma funcao por exercicio no arquivo [`main.go`](https://github.com/Fabiokleis/gomongo/blob/main/cmd/main.go)

exercicio 10:

updateOne faz uma alteração em campos específicos dentro de um documento existente, preservando o restante

replaceOne faz uma substituição completa, o documento original (exceto o _id) e o substitui por um documento inteiramente novo

O schema de [aluno e disciplina](https://github.com/Fabiokleis/gomongo/blob/main/pkg/schema/types.go):
```go

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
```

veja a [conexao com o banco](https://github.com/Fabiokleis/gomongo/blob/main/internal/db/connection.go) e a [seed de alunos e disciplinas](https://github.com/Fabiokleis/gomongo/blob/main/internal/db/seed.go)

## setup
mongodb docker
```shell
docker run --name my-mongo -p 27017:27017 -d mongo
```

## build & run
```shell
export MONGODB_URI=mongodb://localhost:27017
go run cmd/main.go
```
