package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	database "gomongo/internal/db"
	"gomongo/pkg/schema"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ex 1 & 2
func exec1_2(db database.Database) {
	_, err := db.SeedStudents()
	if err != nil {
		panic(err)
	}
}

// ex 3
func exec3(db database.Database) {

	filter := bson.M{"disciplinas.codigo": "CS101"}
	cursor, err := db.Collection("alunos").Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.Ctx)
	for cursor.Next(db.Ctx) {
		var aluno schema.Aluno
		if err := cursor.Decode(&aluno); err != nil {
			log.Fatal(err)
		}
		log.Printf("alunos %v\n", aluno)
	}
}

// ex 4
func exec4(db database.Database) {
	filter := bson.M{"matricula": "2023001"}
	update := bson.M{"$set": bson.M{"curso": "Engenharia Eletrônica"}}
	result, err := db.Collection("alunos").UpdateOne(db.Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("update result: %v\n", result)
}

// ex 5
func exec5(db database.Database) {
	opts := options.Find().SetProjection(bson.M{"nome": 1, "matricula": 1, "_id": 0})
	cursor, err := db.Collection("alunos").Find(db.Ctx, bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.Ctx)
	for cursor.Next(db.Ctx) {
		aluno := struct {
			ID bson.ObjectID
			Nome string
			Matricula  string
		}{}
		if err := cursor.Decode(&aluno); err != nil {
			log.Fatal(err)
		}
		log.Printf("aluno(id, nome, matricula) %v\n", aluno)
	}
}

// ex 6
func exec6(db database.Database) {
	filter := bson.M{"matricula": "2023015"}
	update := bson.M{"$pull": bson.M{"disciplinas": bson.M{"codigo": "PH101"}}}

	result, err := db.Collection("alunos").UpdateOne(db.Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result update: %v\n", result)
}

// ex 7
func exec7(db database.Database) {
	filter := bson.M{
		"$or": []bson.M{
			{"curso": "Ciência da Computação"},
			{"curso": "Engenharia de Software"},
		},
	}
	
	cursor, err := db.Collection("alunos").Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.Ctx)
	for cursor.Next(db.Ctx) {
		var aluno schema.Aluno
		if err := cursor.Decode(&aluno); err != nil {
			log.Fatal(err)
		}
		log.Printf("aluno %v\n", aluno)
	}
}

// exec8
func exec8(db database.Database) {
	filter := bson.M{"matricula": "2023008", "disciplinas.codigo": "MA112"}
	update := bson.M{"$set": bson.M{"disciplinas.$.nota": 9.5}}

	result, err := db.Collection("alunos").UpdateOne(db.Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Nota atualizada: %d\n", result.ModifiedCount)
}

// exec9
func exec9(db database.Database) {
	filter := bson.M{"disciplinas": bson.M{"$elemMatch": bson.M{"nota": bson.M{"$gte": 9.0}}}}

	cursor, err := db.Collection("alunos").Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(db.Ctx) {
		var aluno schema.Aluno
		if err := cursor.Decode(&aluno); err != nil {
			log.Fatal(err)
		}
		log.Printf("aluno %v\n", aluno)
	}
	defer cursor.Close(db.Ctx)	
}

// ex 11
func exec11(db database.Database) {
	filter := bson.M{"curso": "Ciência da Computação"}
	update := bson.M{"$push": bson.M{
		"disciplinas": schema.Disciplina{
			Codigo: "CS405", 
			Nome: "Inteligência Artificial", 
			Professor: "Tacla",
			Registro: time.Now(),
			Salas: []string{}, 
			Horarios: []string{},
		},
	}}

	result, err := db.Collection("alunos").UpdateMany(db.Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Documentos atualizados: %d\n", result.ModifiedCount)
}

func main() {

	db := database.New("universidade")
	db.Connect()
	exec1_2(db)	  
	exec3(db)
	exec4(db)
	exec5(db)
	exec6(db)
	exec7(db)
	exec8(db)
	exec9(db)
	exec11(db)
	// db.CleanupStudents()

	log.Println("send SIGINT/SIGTERM or press CTRL+C to stop....")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	db.Disconnect()
}
