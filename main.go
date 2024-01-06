package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	global  int
	session int
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.Use(authUserMiddlware())

	e.GET("/", layoutHandler)
	e.GET("/workflows", workflowsHander)
	e.GET("/schedules", schedulesHander)
	e.GET("/archive", archiveHandler)

	e.Logger.Fatal(e.Start(":9999"))
}

type Workflow struct {
	Id         int
	Name       string
	Completion bool
}

type Schedule struct {
	Id   int
	Name string
}

type Archive struct {
	Id       int
	Name     string
	Creation string
}

type User struct {
	Name string `json:""`
}

func GetUser(ctx context.Context) *User {
	if user, ok := ctx.Value("user").(*User); ok {
		return user
	}

	return nil
}

func authUserMiddlware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// typically we'd be checking for an auth header or cookie we set.
			// but lets pretend
			if "jimmy" == "jimmy" {

			}
			ctx := context.WithValue(
				c.Request().Context(),
				"user",
				&User{Name: "jimmy"},
			)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}

}

func layoutHandler(c echo.Context) error {
	return Layout().Render(c.Request().Context(), c.Response())
}

func workflowsHander(c echo.Context) error {
	return Workflows([]Workflow{
		{
			Id:         1,
			Name:       "CreateMachine",
			Completion: true,
		},
		{
			Id:         2,
			Name:       "SolveEarthsHunger",
			Completion: false,
		},
		{
			Id:         3,
			Name:       "MakeABaby",
			Completion: true,
		},
	}).Render(c.Request().Context(), c.Response())
}

func schedulesHander(c echo.Context) error {
	return Schedules([]Schedule{
		{
			Id:   0,
			Name: "JimmyShitTime",
		},
		{
			Id:   1,
			Name: "RebootMachines",
		},
	}).Render(c.Request().Context(), c.Response())
}

func archiveHandler(c echo.Context) error {
	return templateArchive(Archive{
		Id:       0,
		Name:     "MyArchive",
		Creation: "Today",
	}).Render(c.Request().Context(), c.Response())
}
