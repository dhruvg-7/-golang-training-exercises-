package myerrors

func NewError(code string) string {
	return getErrorMessage(code)
}
func getErrorMessage(code string) string {
	switch code {
	case "store_i":
		return "Store Insert Error, GOT: "
	case "store_u":
		return "Store Update Error, GOT: "
	case "store_d":
		return "Store Delete Error, GOT: "
	case "store_r":
		return "Store Read Error, GOT: "
	case "store_con":
		return "Store Connection Error, GOT: "
	default:
		return "Some Error"
	}
}
