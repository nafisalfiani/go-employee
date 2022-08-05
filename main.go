package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/nafisalfiani/go-employee/docs/swagger"
	"github.com/nafisalfiani/go-employee/src/model"
	"github.com/nafisalfiani/go-employee/src/repository"
)

// @title Employee Service
// @version 1.0
// @description API Collection for Employee Service
// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com
// @host localhost:3000
// @basePath /
func main() {
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowMethods: "GET,POST,DELETE,PATCH",
		},
	))

	// swagger path
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", hello)
	app.Post("/employee", CreateEmployee)
	app.Get("/employee/:id", GetEmployee)
	app.Get("/employee", GetEmployees)
	app.Patch("/employee", UpdateEmployee)
	app.Delete("/employee/:id", DeleteEmployee)

	app.Listen(":3000")
}

// HealthCheck
// @Summary Check server status
// @Description Will return hello message if the server is up
// @Tags Hello
// @Produce json
// @Success 200 {object} model.HTTPResp
// @Router / [get]
func hello(c *fiber.Ctx) error {
	return c.JSON(model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Hello",
			Body:  "Welcome to employee service",
		},
	})
}

// Create Employee
// @Summary      Create new employee
// @Description  This API will create new employee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        employee body model.Employee true "Employee Data"
// @Success      200  {object}  model.HTTPResp
// @Failure      400  {object}  model.HTTPResp
// @Router       /employee [post]
func CreateEmployee(c *fiber.Ctx) error {
	resp := model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Create Employee",
		},
	}

	employee := model.Employee{}
	if err := json.Unmarshal(c.Body(), &employee); err != nil {
		resp.Message.Body = "Failed to create employee"
		resp.Error = err.Error()
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	err := repository.CreateEmployee(c.Context(), employee)
	if err != nil {
		resp.Message.Body = "Failed to create employee"
		resp.Error = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Message.Body = "Employee created"
	return c.JSON(resp)
}

// Update Employee
// @Summary      Update existing employee
// @Description  This API will update existing employee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        employee body model.Employee true "Employee Data"
// @Success      200  {object}  model.HTTPResp
// @Failure      400  {object}  model.HTTPResp
// @Router       /employee [patch]
func UpdateEmployee(c *fiber.Ctx) error {
	resp := model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Update Employee",
		},
	}

	employee := model.Employee{}
	if err := json.Unmarshal(c.Body(), &employee); err != nil {
		resp.Message.Body = "Failed to update employee"
		resp.Error = err.Error()
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	err := repository.UpdateEmployee(c.Context(), employee)
	if err != nil {
		resp.Message.Body = "Failed to update employee"
		resp.Error = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Message.Body = "Employee updated"
	return c.JSON(resp)
}

// Get Employee
// @Summary      Get existing employee by id
// @Description  This API will get existing employee with given id
// @Tags         Employee
// @Produce      json
// @Param        id path int true "Employee Id"
// @Success      200  {object}  model.HTTPResp
// @Failure      400  {object}  model.HTTPResp
// @Router       /employee/{id} [get]
func GetEmployee(c *fiber.Ctx) error {
	resp := model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Get Employee",
		},
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		resp.Message.Body = "Failed to get employee"
		resp.Error = err.Error()
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	employee, err := repository.GetEmployee(c.Context(), id)
	if err != nil {
		resp.Message.Body = "Failed to get employee"
		resp.Error = err.Error()
		return c.Status(http.StatusNotFound).JSON(resp)
	}

	resp.Message.Body = "Get employee successful"
	resp.Data = employee
	return c.JSON(resp)
}

// Get Employee List
// @Summary      Get all existing employee
// @Description  This API will get all existing employee
// @Tags         Employee
// @Produce      json
// @Param        page query int false "page count"
// @Param        page_size query int false "page size"
// @Success      200  {object}  model.HTTPResp
// @Router       /employee [get]
func GetEmployees(c *fiber.Ctx) error {
	resp := model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Get Employees",
		},
	}

	page, _ := strconv.Atoi(c.Query("page", "-1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	employees, pg, err := repository.GetEmployees(c.Context(), page, pageSize)
	if err != nil {
		resp.Message.Body = "Failed to get employees"
		resp.Error = err.Error()
		return c.Status(http.StatusNotFound).JSON(resp)
	}

	resp.Message.Body = "Get employees successful"
	resp.Data = employees
	resp.Pagination = pg
	return c.JSON(resp)
}

// Delete Employee
// @Summary      Delete existing employee by id
// @Description  This API will delete employee data with given id
// @Tags         Employee
// @Produce      json
// @Param        id path int true "Employee Id"
// @Success      200  {object}  model.HTTPResp
// @Failure      400  {object}  model.HTTPResp
// @Router       /employee/{id} [delete]
func DeleteEmployee(c *fiber.Ctx) error {
	resp := model.HTTPResp{
		Message: model.HTTPMessage{
			Title: "Delete Employee",
		},
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		resp.Message.Body = "Failed to delete employee"
		resp.Error = err.Error()
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	if err := repository.DeleteEmployee(c.Context(), id); err != nil {
		resp.Message.Body = "Failed to delete employee"
		resp.Error = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp.Message.Body = "Employee deleted"
	return c.JSON(resp)
}
