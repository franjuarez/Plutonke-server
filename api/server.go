package api

import (
	"example/plutonke-server/storage"
	"github.com/labstack/echo"
)

type Server struct {
	port  string
	router *echo.Echo
	store storage.Storage
}

func NewServer(port string, store storage.Storage) *Server {
	return &Server{
		port:  port,
		router: echo.New(),
		store: store,
	}
}

func (s *Server) Start() error {
	s.router.GET("/expenses", s.HandleGetAllExpenses)
	s.router.GET("/expenses/:id", s.HandleGetExpenseById)
	s.router.POST("/expenses", s.HandleAddExpense)
	s.router.PUT("/expenses/:id", s.HandleEditExpense)
	s.router.DELETE("/expenses/:id", s.HandleDeleteExpense)

	s.router.GET("/categories", s.HandleGetAllCategories)
	s.router.GET("/categories/:id", s.HandleGetCategoryById)
	s.router.POST("/categories", s.HandleAddCategory)
	s.router.PUT("/categories/:id", s.HandleEditCategory)
	s.router.DELETE("/categories/:id", s.HandleDeleteCategory)

	if err := s.router.Start(s.port); err != nil{
		return err
	}
	return nil
}
