package controller

import (
	"net/http"

	//"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/citiaps/yoinformogral-backend/model"
	"github.com/citiaps/yoinformogral-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type FormController struct {
}

func (formController *FormController) Routes(base *gin.RouterGroup, authNormal *jwt.GinJWTMiddleware) *gin.RouterGroup {

	formRouter := base.Group("/forms") //, middleware.SetRoles(RolAdmin, RolUser), authNormal.MiddlewareFunc())
	{

		formRouter.POST("", authNormal.MiddlewareFunc(), formController.Create())
		formRouter.GET("/:id", formController.One())
		formRouter.GET("", formController.GetAll())
		//formRouter.GET("/:id", formController.GetQuestionCharacterizationByIdForm())
		
	}
	base.GET("/question-characterization/:id", formController.GetQuestionCharacterizationByIdForm())
	base.GET("/questionsForm/:id", formController.GetQuestionsByIdForm())
	return formRouter
}

func (formController *FormController) GetAll() func(c *gin.Context) {

	return func(c *gin.Context) {
		pagination := PaginationParams{}
		err := c.ShouldBind(&pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se puedieron encontrar los parametros limit, offset", err))
			return
		}

		page, err := formModel.FindPaginate(bson.M{}, pagination.Limit, pagination.Offset)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener la lista de formularios", err))
		}

		// if len(page.Metadata) != 0 {
		// 	c.Header("Pagination-Count", fmt.Sprintf("%d", page.Metadata[0]["total"]))
		// }

		c.JSON(http.StatusOK, page)
	}
}

func (formController *FormController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {

		// Traer Usuario
		var form model.Form
		err := c.Bind(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo decodificar json", err))
			return
		}

		err = formModel.Create(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo insertar perro", err))
			return
		}

		c.JSON(http.StatusOK, form)
	}
}
//Para obtener solo un fomrulario por su id
func (formController *FormController) One() func(c *gin.Context) {
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
		group, err := formModel.Get(id)
		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se encontró el formulario", err))
			return
		}
		// Times resolved
		userQuery := c.Query("user")
		if userQuery != "" && userQuery != "undefined" {
			user, err := userModel.FindOne(bson.M{
				"_id": bson.ObjectIdHex(userQuery),
			})
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, util.GetError("Error", err))
				return
			}
			count, err := reportModel.Count(bson.M{
				"user": user.ID,
			})
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, util.GetError("Error", err))
				return
			}
			if count >= 10 {
				c.JSON(http.StatusConflict, util.GetError("Solo se puede responder", err))
				return
			}
		}

		c.JSON(http.StatusOK, group)
	}
}

/*Se recibe el parameto id (id del formulario), con eso se encuentra el formulario y se obtienen las preguntas de 
caracterizacíon
*/
func (formController *FormController) GetQuestionCharacterizationByIdForm() func(c *gin.Context) {
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
		query["_id"] = bson.ObjectIdHex(id) 
		group, err := formModel.QuestionCharacterizationByIdForm(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudieron obtener las preguntas de caracterización", nil))
			return
		}

		c.JSON(http.StatusOK , group)
	}
}
/*Se recibe el parameto id (id del formulario), con eso se encuentra el formulario y se obtienen las preguntas asociadas 
a ese formulario
*/
func (formController *FormController) GetQuestionsByIdForm() func(c *gin.Context) {
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
		query["_id"] = bson.ObjectIdHex(id) 
		group, err := formModel.QuestionsAll(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudieron obtener las preguntas de caracterización", nil))
			return
		}

		c.JSON(http.StatusOK , group)
	}
}
