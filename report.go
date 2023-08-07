package model

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

type Report struct {
	ID            bson.ObjectId          `json:"id" bson:"_id,omitempty"`
	User          bson.ObjectId          `json:"user" bson:"user,omitempty"`
	Risk          int                    `json:"risk" bson:"risk"`
	GlobalRisk    int                    `json:"globalRisk" bson:"global_risk"`
	PerceivedRisk int                    `json:"perceivedRisk" bson:"perceived_risk"`
	Categories    []interface{}          `json:"categories" bson:"categories"`
	Form          bson.ObjectId          `json:"form" bson:"form,omitempty"`
	Answers       []interface{}          `json:"answers" bson:"answers,omitempty"`
	CreatedAt     time.Time              `json:"createdAt" bson:"createdAt,omitempty"`
	Type          map[string]interface{} `json:"type" bson:"type,omitempty"`
	State         interface{}            `json:"state" bson:"state,omitempty"`
	UserPayload   User                   `json:"userPayload" bson:"userPayload,omitempty"`
}

type QueryReport struct {
	ID         bson.ObjectId          `json:"id" bson:"_id,omitempty"`
	User       bson.ObjectId          `json:"user" bson:"user,omitempty"`
	Categories []interface{}          `json:"categories" bson:"categories"`
	Form       bson.ObjectId          `json:"form" bson:"form,omitempty"`
	CreatedAt  time.Time              `json:"createdAt" bson:"createdAt,omitempty"`
	Type       map[string]interface{} `json:"type" bson:"type,omitempty"`
	State      string                 `json:"state" bson:"state"`
}
type GroupByState struct {
	State string
	Count int
}
type PageReport struct {
	Metadata []map[string]int `json:"metadata" bson:"metadata,omitempty"`
	Data     []Report         `json:"data" bson:"data,omitempty"`
}

func (reportModel *Report) Create(reportDoc *Report) error {
	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	reportDoc.ID = bson.NewObjectId()
	reportDoc.CreatedAt = time.Now()
	state := make(map[string]interface{})
	state["label"] = "Recepcionada"
	reportDoc.State = state

	err := col.Insert(reportDoc)

	return err
}

func (reportModel *Report) Get(id string) (*Report, error) {
	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	var reportDoc Report
	err := col.FindId(bson.ObjectIdHex(id)).One(&reportDoc)

	return &reportDoc, err
}

func (reportModel *Report) FindAnon() ([]bson.M, error) {
	col, session := GetCollection(CollectionNameReport)
	defer session.Close()

	groupReports := []bson.M{}
	err := col.Find(bson.M{
		"user": bson.M{
			"$exists": false,
		},
	}).All(&groupReports)

	return groupReports, err
}

func (*Report) Count(query bson.M) (int, error) {
	col, session := GetCollection(CollectionNameReport)
	defer session.Close()

	groupReports := []bson.M{}
	err := col.Find(query).All(&groupReports)

	return len(groupReports), err
}

func (reportModel *Report) FindAggregate(query bson.M, group bson.M) ([]bson.M, error) {

	col, session := GetCollection(CollectionNameReport)
	defer session.Close()

	pipeline := []bson.M{
		{"$match": query},
		group,
	}
	groupReports := []bson.M{}
	err := col.Pipe(pipeline).All(&groupReports)
	log.Printf("pipe: %+v", pipeline)
	return groupReports, err
}

func (reportModel *Report) FindPaginate(query bson.M, limit int, offset int, sort string, mode bool) ([]Report, int, error) {

	col, session := GetCollection(CollectionNameReport)
	// col2, _ := GetCollection(CollectionNameReport)
	defer session.Close()
	pag := []bson.M{{"$sort": bson.M{"createdAt": -1}}}

	pag = append(pag, bson.M{"$skip": offset})

	if limit > 0 {
		pag = append(pag, bson.M{"$limit": limit})
	}
	// pipeline := []bson.M{
	// 	bson.M{"$match": query},
	// }

	// pageDoc := PageReport{}
	pageDoc := []Report{}

	if !mode {
		err := col.Find(query).Sort(sort).Limit(limit).Skip(offset).All(&pageDoc)
		if err != nil {
			return pageDoc, 0, err
		}
	}
	count, err := col.Find(query).Count()
	// err := col.Pipe(pipeline).All(&pageDoc)
	// log.Printf("pipe: %+v", pag)
	// pipeline2 := []bson.M{
	// 	bson.M{"$match": query},
	// 	bson.M{"$group": bson.M{"_id": "_id", "count": bson.M{"$sum": 1}}},
	// }
	// err = col.Pipe(pipeline2).All(&count)
	// total := count[0]["count"]
	log.Printf("count %v\n", count)
	return pageDoc, count, err
}

func (reportModel *Report) Update(id string, queryUpdate bson.M) error {
	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	log.Printf("query %v, id: %s", queryUpdate, id)
	err := col.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": queryUpdate})
	return err
}
func (reportModel *Report) Delete(id string) error {

	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Find : Obtener todos los resportes que correspondan al id del formulario
func (reportModel *Report) FindIdForm(query bson.M) ([]Report, error) {

	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	reports := []Report{}
	err := col.Find(query).All(&reports)
	return reports, err
}

//Si se hace esta consulta directamente en la base de datos es con un Agreggate, acá se utiliza un pipeline.
//Entrada: query => "query["form"] = bson.ObjectIdHex(id) donde el form debe tener el id que entra
/*Funcionamiento: Busca todos los reportes que contengan el ID del formulario, luego separa los documentos en mini documentos por answers,
luego hace el match con todas las respuestas que isCharacterization = true, luego agrupa por index de las preguntas y hace una proyeccion */
//Salida: Arreglo de arreglos, donde un arreglo grande contiene un arreglo con las respuestas de cada pregunta agrupadas por index.
func (reportModel *Report) AnswersReportsByIdForm(query bson.M) ([]bson.M, error) {

	col, session := GetCollection(CollectionNameReport)
	defer session.Close()

	//Definición de Pipeline (Agreggate)
	pipeline := []bson.M{
		{
			"$match": query,
		},
		bson.M{"$unwind": "$answers"},
		bson.M{
			"$match": bson.M{
				"answers.isCharacterization": true,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$answers.index",
				"respuestas": bson.M{
					"$push": "$answers.answer",
				},
			},
		},
		bson.M{
				"$sort": bson.M{
				  "_id": 1,
				},
		},
		bson.M{
			"$project": bson.M{
				"_id":       0,
				"index":     "$_id",
				"respuestas": 1,
			},
		},
	}

	groupReports := []bson.M{}
	err := col.Pipe(pipeline).All(&groupReports)
	log.Printf("pipe: %+v", pipeline)
	return groupReports, err
}

func (reportModel *Report) AllAnswersReportsByIdForm(query bson.M) ([]bson.M, error) {

	col, session := GetCollection(CollectionNameReport)
	defer session.Close()
	pipeline := []bson.M{
		{
			"$match": query,
		},
		bson.M{"$unwind": "$answers"},
		bson.M{
			"$group": bson.M{
				"_id": "$answers.index",
				"respuestas": bson.M{
					"$push": "$answers.answer",
				},
			},
		},
		bson.M{
				"$sort": bson.M{
				  "_id": 1,
				},
		},
		bson.M{
			"$project": bson.M{
				"_id":       0,
				"index":     "$_id",
				"respuestas": 1,
			},
		},
	}

	groupAllAnswers:= []bson.M{}
	err := col.Pipe(pipeline).All(&groupAllAnswers)
	log.Printf("pipe: %+v", pipeline)
	return groupAllAnswers, err

}
