package model

import ("github.com/globalsign/mgo/bson"
"log"
)

type Form struct {
	ID        bson.ObjectId          `json:"id" bson:"_id,omitempty"`
	Title     string                 `json:"title" bson:"title"`
	Type      map[string]interface{} `json:"type" bson:"type,omitempty"`
	Questions []interface{}          `json:"questions" bson:"questions,omitempty"`
	States    []interface{}          `json:"states" bson:"states,omitempty"`
	Characterization bson.ObjectId `json:"characterization" bson:"characterization,omitempty"`
	
}

func (formModel *Form) FindPaginate(query bson.M, limit int, offset int) ([]Form, error) {

	col, session := GetCollection(CollectionNameForm)
	defer session.Close()
	pag := []bson.M{{"$skip": offset}}
	if limit > 0 {
		pag = append(pag, bson.M{"$limit": limit})
	}
	pipeline := []bson.M{
		bson.M{"$match": query},
		// bson.M{"$facet": bson.M{
		// 	"metadata": []bson.M{{"$count": "total"}},
		// 	"data":     pag, // add projection here wish you re-shape the docs
		// }},
	}

	// pageDoc := Page{}
	pageDoc := []Form{}
	err := col.Pipe(pipeline).All(&pageDoc)

	return pageDoc, err
}

func (formModel *Form) Create(formDoc *Form) error {
	col, session := GetCollection(CollectionNameForm)
	defer session.Close()
	formDoc.ID = bson.NewObjectId()
	err := col.Insert(formDoc)

	return err
}

func (formModel *Form) FindByField(query bson.M) (*Form, error) {
	col, session := GetCollection(CollectionNameForm)
	defer session.Close()
	var formDoc Form
	err := col.Find(query).One(&formDoc)

	return &formDoc, err
}
func (formModel *Form) Get(id string) (*Form, error) {
	col, session := GetCollection(CollectionNameForm)
	defer session.Close()
	var formDoc Form
	err := col.FindId(bson.ObjectIdHex(id)).One(&formDoc)

	return &formDoc, err
}

//Si se hace esta consulta directamente en la base de datos es con un Agreggate, acá se utiliza un pipeline.
//Entrada: query 
/*Funcionamiento: Busca todas las preguntas de un formulario que el atributo "isCharacterizaton == true"*/
//Salida: Arreglo con las preguntas, cada pregunta trae su index, type,subtitle, sentence, values
func (formModel *Form) QuestionCharacterizationByIdForm(query bson.M) ([]bson.M, error) {
	col, session := GetCollection(CollectionNameForm)
	defer session.Close()

	// Definición de Pipeline (Aggregate)
	pipeline := []bson.M{
		{
			"$match": query,
		},
		bson.M{"$unwind": "$questions"},
		bson.M{
			"$match": bson.M{
				"questions.isCharacterization": true,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":       0,
				"type":      "$questions.type",
				"name":		 "$questions.name",
				"umedida":   "$questions.umedida",
				"index":     "$questions.index",
				"subtitle":  "$questions.subtitle",
				"sentence":  "$questions.sentence",
				"values": bson.M{
					"$map": bson.M{
						"input": "$questions.options",
						"as":    "option",
						"in":    "$$option.value",
					},
				},
				"texts": bson.M{
					"$map": bson.M{
						"input": "$questions.options",
						"as":    "option",
						"in":    "$$option.text",
					},
				},
			},
		},
	}

	groupForms := []bson.M{}
	err := col.Pipe(pipeline).All(&groupForms)
	log.Printf("pipe: %+v", pipeline)
	return groupForms, err
}

//Entrada: query 
/*Funcionamiento: se hace un pipeline, donde el primer paso es hacer match con la query,
la query es "query["_id"] = bson.ObjectIdHex(id)", se separan las preguntas y luego se hace match con 
todas las preguntas que isQuestion== true y que no sean de tipo text ni open, luego se agrupan y se indica que valores queremos, finalmente
se hace una proyeccion*/
//Salida: Arreglo con las preguntas, incluye su index, type y subtitle
func (formModel *Form) QuestionsAll(query bson.M) ([]bson.M, error) {
	col, session := GetCollection(CollectionNameForm)
	defer session.Close()

	// Definición de Pipeline (Aggregate)
	pipeline := []bson.M{
		{
			"$match": query,
		},
		{"$unwind": "$questions"},
		{
			"$match": bson.M{
				"questions.isQuestion": true,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":       0,
				"type":      "$questions.type",
				"name":		 "$questions.name",
				"index":     "$questions.index",
				"subtitle":  "$questions.subtitle",
				"umedida":   "$questions.umedida",
				"sentence":  "$questions.sentence",
				"values": bson.M{
					"$map": bson.M{
						"input": "$questions.options",
						"as":    "option",
						"in":    "$$option.value",
					},
				},
				"texts": bson.M{
					"$map": bson.M{
						"input": "$questions.options",
						"as":    "option",
						"in":    "$$option.text",
					},
				},
			},
		},
	}
	groupQuestions := []bson.M{}
	err := col.Pipe(pipeline).All(&groupQuestions)
	log.Printf("pipe: %+v", pipeline)
	return groupQuestions, err

}

