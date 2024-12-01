package main

import (
	"bufio"
	"data/model"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func getArgs(table *string, rows *int64, file *string) error {
	args := os.Args

	if len(args) < 7 {
		return fmt.Errorf("insufficient arguments provided")
	}
	if args[1] != "-t" {
		return fmt.Errorf("missing table name. Use -t <tablename>")
	}
	*table = args[2]

	if args[3] != "-r" {
		return fmt.Errorf("missing row count. Use -r <row_count>")
	}
	val, err := strconv.ParseInt(args[4], 10, 64)

	if err != nil {
		return fmt.Errorf("invalid value for -r flag: %v", err)
	}
	*rows = val
	if args[5] != "-f" {
		return fmt.Errorf("missing file name. Use -f <filename>")
	}
	*file = args[6]

	return nil
}

func main() {

	var table string
	var file string
	var rows int64

	err := getArgs(&table, &rows, &file)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Println("Table:", table)
	log.Println("Rows:", rows)

	err = createDataFile(file, table, int(rows))
	if err != nil {
		log.Println("Error:", err)
	}
}

func createDataFile(filename string, table string, rows int) error {
	file, err := os.Create(filename)
	percent := []string{"%", ""}
	c := ""
	format := "sql"
	numAssure := 0
	numP := 0
	d := ","
	b := "["
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter le form des ficher (json or sql):")
	format, _ = reader.ReadString('\n')
	for i := 1; i <= rows; i++ {
		var query string
		switch table {
		case "assure":
			{
				assure := model.GenerateRandomAssure()
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(assure)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf(
						"INSERT INTO assuré (id,nom_prenom, dateN, lieuN, genre, lieu_res, commune, wilaya, region, num_tel, profession, secteur, revenu, etat_civil) "+
							"VALUES (%d,'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s, '%s');\n",
						i,
						assure.NomPrenom,
						assure.DateN,
						assure.LieuN,
						assure.Genre,
						assure.LieuRes,
						assure.Commune,
						assure.Wilaya,
						assure.Region,
						assure.NumTel,
						assure.Profession,
						assure.Secteur,
						assure.Revenu,
						assure.EtatCivil,
					)
				}
			}
		case "maladie":
			{
				if i >= len(model.MaladiesChroniques) {
					break
				}
				r := rand.Intn(len(model.MaladiesChroniques))
				mld := model.MaladiesChroniques[i]
				if r/2 == 0 {
					tauxInt, _ := strconv.Atoi(mld.Taux)
					f := float64(tauxInt) / 100
					mld.Taux = strconv.FormatFloat(f, 'f', 2, 64)
					c = ""
				} else {
					index := rand.Intn(len(percent))
					c = percent[index]
				}

				maladie := model.GenerateRandomMaladie(mld, c)
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(maladie)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf(
						"INSERT INTO maladie (id,code_maladie, libelle, taux, type_maladie) "+
							"VALUES (%d,'%s', '%s', '%s');\n",
						i,
						maladie.CodeMaladie,
						maladie.Libelle,
						maladie.Taux,
						//maladie.TypeMaladie,
					)
				}
			}
		case "medicament":
			{
				if i >= len(model.Medicaments) {
					break
				}
				r := rand.Intn(len(model.Medicaments))
				md := model.Medicaments[i]
				g := r - i
				absValue := math.Abs(float64(g))
				if absValue < 10 {
					f, _ := strconv.ParseFloat(md.TR, 64)
					md.TR = strconv.FormatFloat(f*100, 'f', 0, 64)
					index := rand.Intn(len(percent))
					c = percent[index]
				} else {
					c = ""
				}

				midecament := model.GenerateRandomMidecament(md, c)
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(midecament)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf(
						"INSERT INTO medicament (id,nom, famille, TR, marque) "+
							"VALUES (%d,'%s', '%s', '%s', '%s');\n",
						i,
						midecament.Nom,
						midecament.Famille,
						midecament.TR,
						midecament.Marque,
					)
				}
			}
		case "prestation":
			{
				prestation := model.GenerateRandomPrestation()
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(prestation)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf(
						"INSERT INTO prestation (id, libelle, type_pres) "+
							"VALUES (%d,'%s', '%s');\n",
						i,
						prestation.Libelle,
						prestation.TypePres,
					)
				}
			}
		case "assurance_maladie_prestation":
			{
				if i == 1 {
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("Enter le nomber des assures:")
					numInput, _ := reader.ReadString('\n')
					numAssure, _ = strconv.Atoi(strings.TrimSpace(numInput))

					fmt.Println("Enter le nomber des prestations:")
					numInput, _ = reader.ReadString('\n')
					numP, _ = strconv.Atoi(strings.TrimSpace(numInput))
				}
				amp := model.GenerateRandomAssuranceMaladiePrestation(numAssure, numP)
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(amp)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf("INSERT INTO assurance_maladie_prestation (id, assure_id, maladie_id,prestation_id,prix_prest,commune, wilaya, region,date_prestation) "+
						"VALUES (%d,%d,%d,%d,'%s','%s','%s','%s','%s');\n",
						i,
						amp.AssureID,
						amp.MaladieID,
						amp.PrestationID,
						amp.PrixPrest,
						amp.Commune,
						amp.Wilaya,
						amp.Region,
						amp.DatePrestation,
					)
				}
			}
		case "assurance_maladie_medicament":
			{
				if i == 1 {
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("Enter le nomber des assures:")
					numInput, _ := reader.ReadString('\n')
					numAssure, _ = strconv.Atoi(strings.TrimSpace(numInput))
				}
				amm := model.GenerateRandomAssuranceMaladieMidecament(numAssure)
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(amm)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf("INSERT INTO assurance_maladie_medicament (id, assure_id, maladie_id,medicament_id,prix_medic,commune, wilaya, region,date_achat_medic) "+
						"VALUES (%d,%d,%d,%d,'%s','%s','%s','%s','%s');\n",
						i,
						amm.AssureID,
						amm.MaladieID,
						amm.MedicamentID,
						amm.PrixMedic,
						amm.Commune,
						amm.Wilaya,
						amm.Region,
						amm.Date,
					)
				}
			}
		case "maladie_medicament":
			{
				mm := model.GenerateRandomMaladieMidecament()
				if strings.TrimSpace(format) == "json" {
					jsonData, err := json.Marshal(mm)
					if err != nil {
						log.Fatal(err)
					}
					if i == rows {
						d = "]"
					}
					if i > 1 {
						b = ""
					}
					query = b + string(jsonData) + d
				} else {
					query = fmt.Sprintf("INSERT INTO maladie_medicament (id, maladie_id, medicament_id) "+
						"VALUES (%d,%d,%d);\n",
						i,
						mm.MaladieID,
						mm.MedicamentID,
					)
				}
			}
		default:
			return fmt.Errorf("unsupported table name: %s", table)
		}
		_, err = file.WriteString(query)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	fmt.Printf("Data successfully written to %s\n", filename)
	return nil

}
