package users

type userData struct {
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
	Otros              []userOtherData
	Ubicacion          userAddress
}
type userAddress struct {
	Provincia string
	Canton    string
	Parroquia string
	Direccion string
}

type userOtherData struct {
	name        string
	description string
}
