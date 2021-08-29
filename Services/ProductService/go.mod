module ProductService

go 1.16

replace ProductModel => ../../Models/ProductModel

require (
	ProductDb v0.0.0-00010101000000-000000000000
	ProductModel v0.0.0-00010101000000-000000000000
)

replace ProductDb => ../../DaoFactory/ProductDb
