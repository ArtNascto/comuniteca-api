package global

var (
	AdminName          = "Admin"
	AdminSurname       = "Admin"
	AdminAddress       = "Rua Dr. Cristiano Altenfelder Silva"
	AdminAddressNumber = "201"
	AdminComplement    = ""
	AdminDistrict      = "Vila Carrão"
	AdminCity          = "São Paulo"
	AdminState         = "São Paulo"
	AdminZipcode       = "03448-010"
	Permissions        = []string{
		"book",
		"book.create",
		"book.update",
		"book.view",
		"user",
		"user.view",
		"user.create",
		"user.update",
		"category",
		"category.view",
		"category.update",
		"category.create",
		"role",
		"role.create",
		"role.update",
		"role.view",
		"rent",
		"rent.update",
		"rent.delete",
		"rent.create",
	}
)
