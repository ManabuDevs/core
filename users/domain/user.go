package users

type user struct {
	id                 int
	nroSocio           string
	nombres            string
	apellidos          string
	tipoIdentificacion string
	Nacionalidad       string
	FechaNacimiento    string
	Edad               int
	FechaIngreso       string
	actividadEconomica string
	genero             string
	estadoCivil        string
	Telefono           string
	Instruccion        string
	Otros              []OtherData
	Ubicacion          Address
}
type Address struct {
	Provincia string
	Canton    string
	Parroquia string
	Direccion string
}

type OtherData struct {
	name        string
	description string
}
