package db

import (
	schema "gomongo/pkg/schema"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// NoSQL - exercicio 1 & 2
func (db *Database) SeedStudents() (*mongo.InsertManyResult, error) {

	disciplinas := []schema.Disciplina{
		{
			Codigo:    "ICSB60",
			Nome:      "NoSQL Banco de Dados Não Relacionais",
			Professor: "Leandro",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{"CE-202", "CB-201"},
			Horarios:  []string{"T2", "T3"},
		},
		{
			Codigo:    "CSB30",
			Nome:      "Introdução A Banco de Dados",
			Professor: "Leandro",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{},
			Horarios:  []string{},
		},
		{
			Codigo:    "ICSV30",
			Nome:      "Processamento Digital de Imagens",
			Professor: "Bogdan",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{},
			Horarios:  []string{},
		},
		{
			Codigo:    "ICSV40",
			Nome:      "Computação Gráfica",
			Professor: "Dutra",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{},
			Horarios:  []string{},
		},
		{
			Codigo:    "ICSE30",
			Nome:      "Engenharia de Software",
			Professor: "Adolfo Gustavo",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{"CB-201"},
			Horarios:  []string{"N1", "N2", "N3", "N4"},
		},
		{
			Codigo: "MA112",
			Nome: "Álgebra Linear",
			Professor: "Luiz",
			Registro: time.Now(),
			Nota: nil,
			Salas: []string{},
			Horarios: []string{},
		},
	}

	zdisciplinas := append(disciplinas,
		schema.Disciplina{
			Codigo:    "CS101",
			Nome:      "Introdução a Lógica para computação",
			Professor: "Luders",
			Registro:  time.Now(),
			Nota:      nil,
			Salas:     []string{"CE-302"},
			Horarios:  []string{"T2", "T3"},
		},
	)
	alunos := []schema.Aluno{
		{
			Nome:        "zé",
			Matricula:   "2023008",
			Curso:       "Engenharia de Software",
			Email:       "ze@alunos.utfpr.edu.br",
			Registro:    time.Now(),
			Disciplinas: zdisciplinas,
		},
		{
			Nome:        "joao",
			Matricula:   "12010223",
			Curso:       "Engenharia Eletrônica",
			Email:       "joao@alunos.utfpr.edu.br",
			Registro:    time.Now(),
			Disciplinas: disciplinas,
		},
		{
			Nome:        "ana",
			Matricula:   "12010221",
			Curso:       "Engenharia Mecatrônica",
			Email:       "ana@alunos.utfpr.edu.br",
			Registro:    time.Now(),
			Disciplinas: disciplinas,
		},
		{
			Nome:        "lucas",
			Matricula:   "12010222",
			Curso:       "Ciência da Computação",
			Email:       "lucas@alunos.utfpr.edu.br",
			Registro:    time.Now(),
			Disciplinas: disciplinas,
		},
	}

	coll := db.Client.Database(db.Name).Collection("alunos")
	result, err := coll.InsertMany(db.Ctx, alunos)
	if err != nil {
		return nil, err
	}
	log.Printf("result alunos: %v\n", result)
	return result, nil
}

func (db *Database) CleanupStudents() {
	coll := db.Client.Database(db.Name).Collection("alunos")
	deleteResult, err := coll.DeleteMany(db.Ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	log.Printf("students deleted: %v\n", deleteResult.DeletedCount)
}
