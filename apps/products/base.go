package products

import (
	"online-shop/apps/auth"
	infrafiber "online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Get("", handler.GetListProduct)

		productRoute.Use(infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{
			string(auth.ROLE_Admin),
		}))
		productRoute.Post("", handler.CreateProduct)
		productRoute.Get("/sku/:sku", handler.GetProductDetail)
		productRoute.Put("/:id", handler.UpdateProduct)
		productRoute.Delete("/:id", handler.DeleteProduct)
	}
}
