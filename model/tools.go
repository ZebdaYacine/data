package model

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateCaseString(value string) string {
	cases := []string{"upper", "lower", "char_lower", "char_upper", "first_char_lower", "fisrt_char_upper"}
	index := rand.Intn(len(value))
	switch cases[rand.Intn(len(cases))] {
	case "upper":
		return strings.ToUpper(value)
	case "lower":
		return strings.ToLower(value)
	case "char_lower":
		return strings.ToLower(value[:index]) + strings.ToLower(value[index:index+1]) + strings.ToLower(value[index+1:])
	case "char_upper":
		return strings.ToUpper(value[:index]) + strings.ToUpper(value[index:index+1]) + strings.ToUpper(value[index+1:])
	case "first_char_lower":
		return strings.ToLower(value[:1]) + value[1:]
	case "fisrt_char_upper":
		return strings.ToUpper(value[:1]) + value[1:]
	default:
		return ""
	}
}

func generateDataFormat() string {
	min := 1945
	max := 2005
	cases := []string{"yyyy-mm-dd", "yyyy/mm/dd", "yyyy\\mm\\dd", "dd\\mm\\yyyy", "dd-mm-yyyy", "dd/mm/yyyy"}
	day := rand.Intn(28)
	mounth := rand.Intn(12)
	year := rand.Intn(max-min+1) + min
	switch cases[rand.Intn(len(cases))] {
	case "yyyy-mm-dd":
		{
			return strconv.Itoa(year) + "-" + strconv.Itoa(mounth) + "-" + strconv.Itoa(day)
		}
	case "dd-mm-yyyy":
		{
			return strconv.Itoa(day) + "-" + strconv.Itoa(mounth) + "-" + strconv.Itoa(year)
		}
	case "yyyy/mm/dd":
		{
			return strconv.Itoa(year) + "/" + strconv.Itoa(mounth) + "/" + strconv.Itoa(day)
		}
	case "dd/mm/yyyy":
		{
			return strconv.Itoa(day) + "/" + strconv.Itoa(mounth) + "/" + strconv.Itoa(year)
		}
	case "yyyy\\mm\\dd":
		{
			return strconv.Itoa(year) + "\\" + strconv.Itoa(mounth) + "\\" + strconv.Itoa(day)
		}
	case "dd\\mm\\yyyy":
		{
			return strconv.Itoa(year) + "\\" + strconv.Itoa(mounth) + "\\" + strconv.Itoa(day)
		}
	default:
		return ""
	}
}

func generateMobileNumber() string {
	cases := []string{"06", "05", "07"}
	mobilis := rand.Intn(79000000-50000000+1) + 50000000
	ooredoo := rand.Intn(90000000-60000000+1) + 60000000
	djizzy := rand.Intn(90000000-70000000+1) + 70000000
	switch cases[rand.Intn(len(cases))] {
	case "06":
		{
			return "06" + strconv.Itoa(mobilis)
		}
	case "05":
		{
			return "05" + strconv.Itoa(ooredoo)
		}
	case "07":
		{
			return "07" + strconv.Itoa(djizzy)
		}

	default:
		return ""
	}
}

func GenerateRandomAssure() Assure {
	selectedWilaya := Wilayas[rand.Intn(len(Wilayas))]
	Fname := BaseNames[rand.Intn(len(BaseNames))]
	Lname := BaseLastNames[rand.Intn(len(BaseLastNames))]
	return Assure{
		NomPrenom:  generateCaseString(Fname) + " " + generateCaseString(Lname),
		DateN:      generateDataFormat(),
		LieuN:      selectedWilaya.Communes[rand.Intn(len(selectedWilaya.Communes))],
		Genre:      Genres[rand.Intn(len(Genres))],
		LieuRes:    selectedWilaya.Communes[rand.Intn(len(selectedWilaya.Communes))],
		Commune:    selectedWilaya.Communes[rand.Intn(len(selectedWilaya.Communes))],
		Wilaya:     selectedWilaya.Name,
		Region:     selectedWilaya.Region,
		NumTel:     generateMobileNumber(),
		Profession: Professions[rand.Intn(len(Professions))],
		Secteur:    Secteurs[rand.Intn(len(Secteurs))],
		Revenu:     strconv.Itoa(rand.Intn(1000000000-100000+1) + 100000),
		EtatCivil:  EtatCivils[rand.Intn(len(EtatCivils))],
	}
}

func GenerateRandomMaladie(maladie Maladie, c string) Maladie {
	return Maladie{
		CodeMaladie: generateCaseString(maladie.CodeMaladie),
		Libelle:     generateCaseString(maladie.Libelle),
		Taux:        generateCaseString(maladie.Taux) + c,
		//TypeMaladie: generateCaseString(LibellesMaladies[rand.Intn(len(LibellesMaladies))]),
	}
}

func GenerateRandomMidecament(m Medicament, c string) Medicament {
	return Medicament{
		Nom:     generateCaseString(m.Nom),
		Famille: generateCaseString(m.Famille),
		TR:      m.TR + c,
		Marque:  generateCaseString(m.Marque),
	}
}

func GenerateRandomPrestation() Prestation {
	return Prestation{
		Libelle:  generateCaseString(Libelles[rand.Intn(len(Libelles))]),
		TypePres: generateCaseString(Types[rand.Intn(len(Types))]),
	}
}

func GenerateRandomAssuranceMaladiePrestation(a int, p int) AssuranceMaladiePrestation {
	selectedWilaya := Wilayas[rand.Intn(len(Wilayas))]
	return AssuranceMaladiePrestation{
		AssureID:       rand.Intn(a),
		MaladieID:      rand.Intn(29),
		PrestationID:   rand.Intn(p),
		PrixPrest:      strconv.FormatInt(rand.Int63n(100000), 10),
		Commune:        selectedWilaya.Communes[rand.Intn(len(selectedWilaya.Communes))],
		Wilaya:         selectedWilaya.Name,
		Region:         selectedWilaya.Region,
		DatePrestation: generateDataFormat(),
	}
}

func GenerateRandomAssuranceMaladieMidecament(a int) AssuranceMaladieMedicament {
	selectedWilaya := Wilayas[rand.Intn(len(Wilayas))]
	return AssuranceMaladieMedicament{
		AssureID:     rand.Intn(a),
		MaladieID:    rand.Intn(29),
		MedicamentID: rand.Intn(81),
		PrixMedic:    strconv.FormatInt(rand.Int63n(10000), 10),
		Commune:      selectedWilaya.Communes[rand.Intn(len(selectedWilaya.Communes))],
		Wilaya:       selectedWilaya.Name,
		Region:       selectedWilaya.Region,
		Date:         generateDataFormat(),
	}
}

func GenerateRandomMaladieMidecament() MaladieMedicament {
	return MaladieMedicament{
		MaladieID:    rand.Intn(29),
		MedicamentID: rand.Intn(81),
	}
}
