package model

type Assure struct {
	NomPrenom  string
	DateN      string
	LieuN      string
	Genre      string
	LieuRes    string
	Commune    string
	Wilaya     string
	Region     string
	NumTel     string
	Profession string
	Secteur    string
	Revenu     string
	EtatCivil  string
}

type Maladie struct {
	CodeMaladie string
	Libelle     string
	Taux        string
	//TypeMaladie string
}

type Medicament struct {
	Nom     string
	Famille string
	TR      string
	Marque  string
}

type Prestation struct {
	Libelle  string
	TypePres string
}

type AssuranceMaladiePrestation struct {
	AssureID     int
	MaladieID    int
	PrestationID int
	PrixPrest    string
}

type AssuranceMaladieMedicament struct {
	AssureID     int
	MaladieID    int
	MedicamentID int
	PrixMedic    string
}

type MaladieMedicament struct {
	MaladieID    int
	MedicamentID int
}

type Wilaya struct {
	Name     string
	Region   string
	Communes []string
}
