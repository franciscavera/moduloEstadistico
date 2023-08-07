package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/citiaps/yoinformogral-backend/mailer"
	"github.com/citiaps/yoinformogral-backend/model"
	push "github.com/citiaps/yoinformogral-backend/pushNotification"
	"github.com/citiaps/yoinformogral-backend/util"
	"github.com/citiaps/yoinformogral-backend/xlsxFile"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// var Risks = map[int]string{
// 	1: "Muy bajo",
// 	2: "Bajo",
// 	3: "Medio",
// 	4: "Alto",
// 	5: "Muy alto",
// }

// var OrderFields = map[int]string{
// 	1: "-risk",
// 	2: "risk",
// 	3: "-perceived_risk",
// 	4: "perceived_risk",
// 	5: "-global_risk",
// 	6: "global_risk",
// }

type ReportController struct {
}

func (reportController *ReportController) Routes(base *gin.RouterGroup, authNormal *jwt.GinJWTMiddleware) *gin.RouterGroup {

	// Answer - Rutas
	reportRouter := base.Group("/reports") //, middleware.SetRoles(RolAdmin, RolUser), authNormal.MiddlewareFunc())
	{

		reportRouter.POST("", reportController.Create())
		reportRouter.PUT("/:id", reportController.Update())
		reportRouter.GET("", authNormal.MiddlewareFunc(), reportController.GetAll())
		//reportRouter.GET("/:id", reportController.GetReportsFormById())
		reportRouter.GET("/:id", reportController.GetAnswersReportsByIdForm())
		
	}
	base.GET("/reports-summary", authNormal.MiddlewareFunc(), reportController.GetAggreate())
    base.GET("/reports-export", authNormal.MiddlewareFunc(), reportController.ExportReports())
	base.GET("/allAnswersReports/:id", reportController.GetAllAnswersReportsByIdForm())
	return reportRouter
}
func (reportController *ReportController) ExportReports() func(c *gin.Context) {
	return func(c *gin.Context) {
		pagination := PaginationParams{}
		err := c.ShouldBind(&pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se puedieron encontrar los parametros limit, offset", err))
			return
		}

		query := bson.M{}
		user := userModel.LoadFromContext(c)
		if pagination.User {
			query["user"] = user.ID
		}

		if user.Rol != "admin" && !pagination.User {
			query["state.section"] = user.Section
		}

		// var value []bson.M
		// if len(pagination.Categories) > 0 {
		// 	for _, category := range pagination.Categories {
		// 		value = append(value, bson.M{"categories.value": category})
		// 	}
		// 	query["$or"] = value
		// }

		if pagination.State != "" {
			query["state.label"] = pagination.State
		}
		// if pagination.Risk != 0 {
		// 	query["risk"] = pagination.Risk
		// }
		// if pagination.Perceived != 0 {
		// 	query["perceived_risk"] = pagination.Perceived
		// }
		// if pagination.Global != 0 {
		// 	query["global_risk"] = pagination.Global
		// }
		// if pagination.Report != 0 {
		// 	query["type.value"] = pagination.Report
		// }
		// if pagination.Validation != "" {
		// 	query["state.validation"] = pagination.Validation
		// }
		if !pagination.Init.IsZero() || !pagination.End.IsZero() {
			query["createdAt"] = bson.M{}
		}
		if !pagination.Init.IsZero() {
			(query["createdAt"]).(bson.M)["$gt"] = pagination.Init
		}
		if !pagination.End.IsZero() {
			(query["createdAt"]).(bson.M)["$lt"] = pagination.End
		}

		sort := "-createdAt"
		// if pagination.Order != 0 {
		// 	sort = OrderFields[pagination.Order]
		// }
		log.Printf("query %v", query)
		reports, _, err := reportModel.FindPaginate(query, pagination.Limit, pagination.Offset, sort, pagination.Count)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener la lista de reportes", err))
		}

		forms, err := formModel.FindPaginate(bson.M{}, 0, 0)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener la lista de formularios", err))
		}

		fileName, file, err := xlsxFile.GenerateXlsxReports(forms, reports)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.GetError("No se pudo crear el archivo xlsx ", err))
			return
		}
		err = file.Save(fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.GetError("No se pudo cargar el archivo ", err))
			return
		}
		defer func() {
			os.Remove(fileName)
		}()
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		downloadName := time.Now().UTC().Format("test.xlsx")
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", b)
	}
}
func (reportController *ReportController) GetAggreate() func(c *gin.Context) {
	return func(c *gin.Context) {

		pagination := PaginationParams{}
		err := c.ShouldBind(&pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se puedieron encontrar los parametros limit, offset", err))
			return
		}

		query := bson.M{}
		user := userModel.LoadFromContext(c)
		if pagination.User {
			query["user"] = user.ID
		}
		if user.Rol != "admin" && !pagination.User {
			query["state.section"] = user.Section
		}

		group := bson.M{}
		if pagination.Group == "state" {
			group["$group"] = bson.M{"_id": "$state.label", "value": bson.M{"$sum": 1}}
		}
		// } else if pagination.Group == "risk" {
		// 	group["$group"] = bson.M{"_id": "$risk", "value": bson.M{"$sum": 1}}
		// } else if pagination.Group == "perceivedRisk" {
		// 	group["$group"] = bson.M{"_id": "$perceived_risk", "value": bson.M{"$sum": 1}}
		// } else if pagination.Group == "globalRisk" {
		// 	group["$group"] = bson.M{"_id": "$global_risk", "value": bson.M{"$sum": 1}}
		// }

		data, err := reportModel.FindAggregate(query, group)
		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener la lista de perros", err))
		}
		log.Printf("data: %v", data)
		c.JSON(http.StatusOK, data)
	}
}

func (reportController *ReportController) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		pagination := PaginationParams{}
		err := c.ShouldBind(&pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se puedieron encontrar los parametros limit, offset", err))
			return
		}

		query := bson.M{}
		user := userModel.LoadFromContext(c)
		if pagination.User {
			query["user"] = user.ID
		}

		if pagination.State != "" {
			query["state.label"] = pagination.State
		}
		// if pagination.Validation != "" {
		// 	query["state.validation"] = pagination.Validation
		// }
		// if pagination.Risk != 0 {
		// 	query["risk"] = pagination.Risk
		// }
		// if pagination.Perceived != 0 {
		// 	query["perceived_risk"] = pagination.Perceived
		// }
		// if pagination.Global != 0 {
		// 	query["global_risk"] = pagination.Global
		// }
		if pagination.Report != 0 {
			query["type.value"] = pagination.Report
		}
		if !pagination.Init.IsZero() || !pagination.End.IsZero() {
			query["createdAt"] = bson.M{}
		}
		if !pagination.Init.IsZero() {
			(query["createdAt"]).(bson.M)["$gt"] = pagination.Init
		}
		if !pagination.End.IsZero() {
			(query["createdAt"]).(bson.M)["$lt"] = pagination.End
		}
		sort := "-createdAt"
		// if pagination.Order != 0 {
		// 	sort = OrderFields[pagination.Order]
		// }
		var value []bson.M

		if len(pagination.Categories) > 0 {
			for _, category := range pagination.Categories {
				value = append(value, bson.M{"categories.value": category})
			}
			query["$or"] = value
		}
		log.Printf("query %v", query)
		page, total, err := reportModel.FindPaginate(query, pagination.Limit, pagination.Offset, sort, pagination.Count)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener la lista de reportes", err))
			return
		}

		c.Header("Pagination-Count", fmt.Sprintf("%d", total))

		c.JSON(http.StatusOK, page)
	}
}

func (reportController *ReportController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		var isAnon bool
		user := userModel.LoadFromContext(c)
		if user == nil {
			isAnon = true
		}
		var report model.Report
		err := c.Bind(&report)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo decodificar json", err))
			return
		}
		if !isAnon {
			report.UserPayload = *user
		}
		err = reportModel.Create(&report)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo crear el report", err))
			return
		}
		subscriptions, err := subscriptionModel.Find(bson.M{"rol": "admin"})
		log.Printf("suibscripciones", subscriptions)
		json_message := fmt.Sprintf(
			"{ \"message\": \"Se ha creado un nuevo reporte\",\"type\":\"response\"}")

		if err != nil {
			log.Printf("Error al obtener subscripción: %v, usuario: %v", err, user.ID)
			return
		}

		log.Printf("enviando notifficación por nuevo reporte")
		push.SendNotifications(subscriptions, json_message)

		c.JSON(http.StatusOK, report)
	}
}

func (reportController *ReportController) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		var report model.Report
		err := c.Bind(&report)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo convertir collection json", err))
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, util.GetError("No se encuentra parametro :id", nil))
			return
		}

		if !bson.IsObjectIdHex(id) {
			c.JSON(http.StatusInternalServerError, util.GetError("El id ingresado no es válido", nil))
			return
		}

		queryUpdate := bson.M{}
		queryUpdate["state"] = report.State
		err = reportModel.Update(id, queryUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo actualizar el reporte", err))
			return
		}

		go func(id string) {

			report, err := reportModel.Get(id)
			if err != nil {
				log.Printf("Error: report not found, %s", err.Error())
				return
			}
			user := report.UserPayload
			state := report.State.(bson.M)
			var label string
			if labelCheck, ok := state["other"].(string); ok {
				label = labelCheck
			} else {
				label = ""
			}

			if label == "Aceptado" || label == "Rechazado" {
				var response string
				if responseCheck, ok := state["response"].(string); ok {

					response = responseCheck
				} else {
					log.Printf("Respuesta no definida")
					return
				}

				subscriptions, err := subscriptionModel.Find(bson.M{"user": user.ID})

				json_message := fmt.Sprintf(
					"{ \"message\": \"Ha recibido una respuesta por su último reporte: %s\",\"type\":\"response\"}", response)

				if err != nil {
					log.Printf("Error al obtener subscripción: %v, usuario: %v", err, user.ID)
					return
				}

				log.Printf("enviando notifficación a %v, respuesta %v", id, response)
				push.SendNotifications(subscriptions, json_message)
			} else if label == "Asignado" {
				loc, err := time.LoadLocation("America/Santiago")
				if err != nil {
					fmt.Printf("%s", err.Error())
					return
				}
				date := report.CreatedAt.In(loc).Format("02/01/06 03:04:05 PM")
				var reportType string
				var inCharge bson.M
				var inChargeEmail string
				var inChargeName string

				//risk := Risks[report.GlobalRisk]
				panelUrl := os.Getenv("PANEL_URL")

				if reportTypeCheck, ok := report.Type["label"].(string); ok {

					reportType = reportTypeCheck
				} else {
					return
				}

				if inChargeFlag, ok := state["incharge"].(bson.M); ok {

					inCharge = inChargeFlag
				} else {
					return
				}

				if inChargeEmailFlag, ok := inCharge["email"].(string); ok {

					inChargeEmail = inChargeEmailFlag
				} else {
					return
				}

				if inChargeNameFlag, ok := inCharge["name"].(string); ok {

					inChargeEmail = inChargeNameFlag
				} else {
					return
				}

				templateDir := "mailer/templates/stateAssigmentTemplate.html"
				r := mailer.NewRequest([]string{inChargeEmail}, "Reporte asignado")

				err = r.SendMailSkipTLS(templateDir, map[string]string{"username": inChargeName, "date": date, "type": reportType, "panelUrl": panelUrl})
				if err != nil {
					c.JSON(http.StatusBadRequest, util.GetError("Email invalido", err))
					return
				}
				subscriptions, err := subscriptionModel.Find(bson.M{"email": inChargeEmail})
				log.Printf("suibscripciones", subscriptions)
				json_message := fmt.Sprintf(
					"{ \"message\": \"Ta han asignado un nuevo reporte\",\"type\":\"response\"}")

				if err != nil {
					log.Printf("Error al obtener subscripción: %v, usuario: %v", err, user.ID)
					return
				}

				log.Printf("enviando notifficación a usuario asignado")
				push.SendNotifications(subscriptions, json_message)
			}

		}(id)

		c.JSON(http.StatusOK, report)

	}
}
/*
func (reportController *ReportController) GetReportsFormById() func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusNotFound, util.GetError("No se encuentra parametro :id", nil))
			return
		}
		if !bson.IsObjectIdHex(id) {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo recibir parámetro", nil))
			return
		}
		query := bson.M{}
		query["form"] = bson.ObjectIdHex(id)
		group, err := reportModel.FindIdForm(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudieron obtener los reportes", nil))
			return
		}

		c.JSON(http.StatusOK, group)
	}

}*/

func (reportController *ReportController) GetAnswersReportsByIdForm() func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")
		
		if id == "" {
			c.JSON(http.StatusNotFound, util.GetError("No se encuentra parametro :id", nil))
			return
		}
		if !bson.IsObjectIdHex(id) {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo recibir parámetro", nil))
			return
		}
		query:= bson.M{}
		query["form"] = bson.ObjectIdHex(id) 
		group, err := reportModel.AnswersReportsByIdForm(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudieron obtener los reportes", nil))
			return
		}

		c.JSON(http.StatusOK , group)
	}
}

func (reportController *ReportController) GetAllAnswersReportsByIdForm() func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")
		
		if id == "" {
			c.JSON(http.StatusNotFound, util.GetError("No se encuentra parametro :id", nil))
			return
		}
		if !bson.IsObjectIdHex(id) {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo recibir parámetro", nil))
			return
		}
		query:= bson.M{}
		query["form"] = bson.ObjectIdHex(id) 
		group, err := reportModel.AllAnswersReportsByIdForm(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudieron obtener los reportes", nil))
			return
		}

		c.JSON(http.StatusOK , group)
	}

}
