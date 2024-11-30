package model

var (
	BaseNames = []string{
		"Amine", "Samira", "Yacine", "Lyna", "Rawnak", "Ali", "Sarah", "Omar", "Mouad", "Khaled",
		"Rachid", "Farida", "Sofia", "Meryem", "Hassan", "Hicham", "Karim", "Mounir",
		"Fares", "Ibtissam", "Mouna", "Laila", "Samir", "Amina", "Nadia", "Rachid", "Kenza",
	}
	BaseLastNames = []string{
		"Benali", "Khaled", "Mouhoub", "Ait", "Fellah", "Bouras", "Amrani", "Zerhouni", "Ouchene",
		"Merabet", "Bensalem", "Mahmoudi", "Djellal", "Charef", "Boudiaf", "Bennacer", "Sahnoun",
		"Larbi", "Mokrane", "Meziani", "Kaddour", "Ammar", "Sidi", "Ibrahim", "Boussad", "Messaoudi",
	}
	Genres      = []string{"homme", "Femme"}
	Professions = []string{"ingénieur", "Médecin", "Enseignant", "Commerçant", "Artisan"}
	Secteurs    = []string{"technologie", "Santé", "Éducation", "Commerce", "Industrie"}
	EtatCivils  = []string{"célibataire", "Marié(e)", "Divorcé(e)", "Veuf/Veuve"}
	Wilayas     = []Wilaya{
		{"Adrar", "Sud", []string{"Adrar", "Aoulef", "Zaouiat Kounta"}},
		{"Chlef", "Nord", []string{"Chlef", "Tiaret", "Oum El Bouaghi"}},
		{"Laghouat", "Sud", []string{"Laghouat", "Aflou", "Khenchela"}},
		{"Oum El Bouaghi", "Est", []string{"Oum El Bouaghi", "Foum Toub", "Boudouaou"}},
		{"Batna", "Est", []string{"Batna", "Boughezoul", "Mila"}},
		{"Béjaïa", "Nord", []string{"Béjaïa", "Sidi Aïch", "Toudja"}},
		{"Biskra", "Sud", []string{"Biskra", "Oued Djer", "Zeribet el Oued"}},
		{"Béchar", "Sud", []string{"Béchar", "Mecheria", "Taghit"}},
		{"Blida", "Nord", []string{"Blida", "Ouled Yaich", "El Affroun"}},
		{"Bouira", "Nord", []string{"Bouira", "Maatkas", "Lakhdaria"}},
		{"Tamanrasset", "Sud", []string{"Tamanrasset", "In Salah", "Djanet"}},
		{"Tébessa", "Est", []string{"Tébessa", "Ouenza", "Sidi Khaled"}},
		{"Tlemcen", "Ouest", []string{"Tlemcen", "Nedroma", "El Fehoul"}},
		{"Tiaret", "Ouest", []string{"Tiaret", "Frenda", "Mechraa Bel Hassan"}},
		{"Tizi Ouzou", "Nord", []string{"Tizi Ouzou", "Azazga", "Azeffoun"}},
		{"Alger", "Nord", []string{"Alger", "El Harrach", "Bab El Oued"}},
		{"Djelfa", "Sud", []string{"Djelfa", "Messaoud", "Aïn Oussera"}},
		{"Jijel", "Nord", []string{"Jijel", "El Milia", "Taher"}},
		{"Sétif", "Est", []string{"Sétif", "Aïn Azel", "El Eulma"}},
		{"Saïda", "Ouest", []string{"Saïda", "Aïn Sefra", "El Khemis"}},
		{"Skikda", "Nord", []string{"Skikda", "Azzaba", "El Hadaik"}},
		{"Sidi Bel Abbès", "Ouest", []string{"Sidi Bel Abbès", "Aïn Tédelès", "El Ksar"}},
		{"Annaba", "Nord", []string{"Annaba", "El Hadjar", "Seraïdi"}},
		{"Guelma", "Est", []string{"Guelma", "Tamalous", "El Tarf"}},
		{"Constantine", "Est", []string{"Constantine", "Hamma Bouziane", "El Khroub"}},
		{"Médéa", "Nord", []string{"Médéa", "Aïn Dehlil", "Aïn Bouziane"}},
		{"Mostaganem", "Nord", []string{"Mostaganem", "Sidi Lakhdar", "Ouled Boudjemaa"}},
		{"M'Sila", "Sud", []string{"M'Sila", "Boudouaou", "Beniane"}},
		{"Mascara", "Ouest", []string{"Mascara", "El Bordj", "Sidi Khaled"}},
		{"Ouargla", "Sud", []string{"Ouargla", "El Guerrara", "Touggourt"}},
		{"Oran", "Ouest", []string{"Oran", "Aïn El Türck", "Es Senia"}},
		{"El Bayadh", "Sud", []string{"El Bayadh", "Boualem", "Bordj El Kiffan"}},
		{"Illizi", "Sud", []string{"Illizi", "Djanet", "In Salah"}},
		{"Bordj Bou Arréridj", "Est", []string{"Bordj Bou Arréridj", "Aïn Taghrout", "El Madher"}},
		{"Boumerdès", "Nord", []string{"Boumerdès", "Khemis El Khechna", "Amizour"}},
		{"El Tarf", "Nord", []string{"El Tarf", "Berrihane", "Aïn Kechra"}},
		{"Tindouf", "Sud", []string{"Tindouf", "Oum El Araïs", "Boudouaou"}},
		{"Tissemsilt", "Ouest", []string{"Tissemsilt", "Sidi Slimane", "Khemis Miliana"}},
		{"El Oued", "Sud", []string{"El Oued", "Guemar", "Sidi Aïssa"}},
		{"Khenchela", "Est", []string{"Khenchela", "Aïn El Hammam", "Taourirt"}},
		{"Souk Ahras", "Est", []string{"Souk Ahras", "Oued Souf", "Bordj Bou Arréridj"}},
		{"Tipaza", "Nord", []string{"Tipaza", "Merouana", "Bou Ismail"}},
		{"Mila", "Est", []string{"Mila", "Aïn Beïda", "Sidi Merouane"}},
		{"Aïn Defla", "Nord", []string{"Aïn Defla", "El Ksar", "Sidi Abderrahmane"}},
		{"Naâma", "Sud", []string{"Naâma", "El Hassi", "Chlef"}},
		{"Aïn Témouchent", "Ouest", []string{"Aïn Témouchent", "El Amria", "Boudouaou"}},
		{"Ghardaïa", "Sud", []string{"Ghardaïa", "Berriane", "Metlili"}},
		{"Relizane", "Ouest", []string{"Relizane", "Aïn Tédelès", "Oued Rhiou"}},
	}
	MaladiesChroniques = []Maladie{
		{"C01", "Tuberculose", "100"},
		{"C02", "Psychonévrose graves", "100"},
		{"C03", "Maladies cancéreuses", "100"},
		{"C04", "Hémopathies", "100"},
		{"C05", "La sarcoïdose", "100"},
		{"C06", "Hypertension artérielle maligne", "100"},
		{"C07", "Maladies cardiaques et vasculaires (c07-1 à c07-9)", "100"},
		{"C08", "Maladies neurologiques (c08-1 à c08-4)", "100"},
		{"C09", "Maladies musculaires (c09-1 à c09-4)", "100"},
		{"C10", "Encéphalopathies", "100"},
		{"C11", "Néphropathies", "100"},
		{"C12", "Rhumatismes chroniques (c12-1 à c12-3)", "100"},
		{"C13", "Périartérite noueuse", "100"},
		{"C14", "Lupus-érythémateux", "100"},
		{"C15", "Insuffisance respiratoire chronique", "100"},
		{"C16", "Poliomyélite antérieure aiguë", "100"},
		{"C17", "Maladies métaboliques: diabètes, dyslipidémies", "100"},
		{"C18", "Cardiopathies congénitales", "100"},
		{"C19", "Affections endocriniennes complexes", "100"},
		{"C20", "Rhumatisme-articulaire aigu", "100"},
		{"C21", "Ostéomyélite-chronique", "100"},
		{"C22", "Complice. Grave des gastrectomie et de la maladie", "100"},
		{"C23", "Cirrhoses du foie", "100"},
		{"C24", "Recto-colite hémorragique", "100"},
		{"C25", "Psoriasis-pemphigus malin", "100"},
		{"C26", "Hydatides et ses complications", "100"},
		{"T15", "Hypertension artérielle", "80"},
		{"T16", "Maladie de Crohn", "80"},
		{"T17", "Asthme bronchique", "80"},
	}
	Medicaments = []Medicament{
		{"Amlohexal", "Antihypertenseur", "0.85", "PharmAlger"},
		{"Cardensiel", "Antihypertenseur", "0.90", "MedAlgeria"},
		{"Glucophage", "Antidiabétique", "0.80", "GlobalPharma"},
		{"Atorvastatine", "Statines", "0.75", "VitalPharm"},
		{"Xarelto", "Anticoagulant", "0.95", "MediPharm"},
		{"Lantus", "Antidiabétique", "0.88", "MediCare"},
		{"Spiriva", "Bronchodilatateur", "0.90", "PharmaTech"},
		{"Bisoprolol", "Bêta-bloquant", "0.70", "ZenithPharma"},
		{"Diabrel", "Antidiabétique", "0.65", "SantePlus"},
		{"Enalapril", "Antihypertenseur", "0.80", "AlgerMed"},
		{"Metformine", "Antidiabétique", "0.85", "HealthPro"},
		{"Lipitor", "Statines", "0.78", "DrugPharma"},
		{"Lasilix", "Diurétique", "0.75", "MedEx"},
		{"Captopril", "Antihypertenseur", "0.70", "AlgerCare"},
		{"Accupril", "Antihypertenseur", "0.68", "GlobalMed"},
		{"Levofloxacine", "Antibiotique", "0.80", "MediCure"},
		{"Crestor", "Statines", "0.76", "CarePharm"},
		{"Ramipril", "Antihypertenseur", "0.84", "MediCare"},
		{"Icosapent", "Antihypertenseur", "0.77", "PharmNow"},
		{"Pantoprazole", "Antiacide", "0.82", "HealthCo"},
		{"Procoralan", "Antihypertenseur", "0.74", "MedPlus"},
		{"Actos", "Antidiabétique", "0.81", "MediCare"},
		{"Losartan", "Antihypertenseur", "0.83", "AlgerCare"},
		{"Exforge", "Antihypertenseur", "0.79", "VitalCare"},
		{"Combivent", "Bronchodilatateur", "0.88", "MediMax"},
		{"Symbicort", "Bronchodilatateur", "0.85", "MedPlus"},
		{"Furosemide", "Diurétique", "0.77", "PharmaMed"},
		{"Valsartan", "Antihypertenseur", "0.78", "HealthPro"},
		{"Metoprolol", "Bêta-bloquant", "0.70", "CareMed"},
		{"Pravastatine", "Statines", "0.72", "HealthPlus"},
		{"Rivaroxaban", "Anticoagulant", "0.80", "MediX"},
		{"Euthyrox", "Thyroïde", "0.90", "MediPro"},
		{"Rapamune", "Immunosuppresseur", "0.85", "PharmCare"},
		{"Ticagrelor", "Anticoagulant", "0.88", "CarePharma"},
		{"Salbutamol", "Bronchodilatateur", "0.82", "MediCare"},
		{"Diovan", "Antihypertenseur", "0.77", "PharmaX"},
		{"Novorapid", "Insuline", "0.92", "NovoPharma"},
		{"Lisinopril", "Antihypertenseur", "0.70", "AlgerMed"},
		{"Pradaxa", "Anticoagulant", "0.79", "CardioPharm"},
		{"Amoxicilline", "Antibiotique", "0.78", "GlobalCare"},
		{"Acide folique", "Supplément", "0.88", "MedBio"},
		{"Flixotide", "Bronchodilatateur", "0.85", "HealthPharm"},
		{"Seroquel", "Antipsychotique", "0.72", "PharmX"},
		{"Imdur", "Vasodilatateur", "0.80", "MedX"},
		{"Valium", "Anxiolytique", "0.75", "HealthCure"},
		{"Epilim", "Anticonvulsivant", "0.80", "MediCo"},
		{"Zestril", "Antihypertenseur", "0.76", "MediHealth"},
		{"Xenical", "Anti-obésité", "0.70", "MedEx"},
		{"Propranolol", "Bêta-bloquant", "0.78", "MediPro"},
		{"Brilique", "Antithrombotique", "0.75", "CardioPlus"},
		{"Furosemide", "Diurétique", "0.80", "MediCare"},
		{"Cordarone", "Antiarythmique", "0.85", "MediCure"},
		{"Sinemet", "Anti-Parkinsonien", "0.90", "MedAlger"},
		{"Ventolin", "Bronchodilatateur", "0.88", "AlgerPharma"},
		{"Karbinal", "Antihistaminique", "0.72", "HealthPlus"},
		{"Inderal", "Bêta-bloquant", "0.79", "MediCo"},
		{"Glibenclamide", "Antidiabétique", "0.75", "HealthX"},
		{"Aldactone", "Diurétique", "0.70", "MedX"},
		{"Actonel", "Antirhumatismal", "0.80", "CareMeds"},
		{"Levothyroxine", "Thyroïde", "0.90", "MediPharm"},
		{"Symbiva", "Antihypertenseur", "0.77", "MediCure"},
		{"Colchicine", "Anti-inflammatoire", "0.80", "HealthPro"},
		{"Trelegy", "Bronchodilatateur", "0.85", "MedCure"},
		{"Symbicort", "Bronchodilatateur", "0.83", "PharmaCare"},
		{"Lasix", "Diurétique", "0.78", "GlobalPharm"},
		{"Norvasc", "Antihypertenseur", "0.80", "MedPlus"},
		{"Ivabradine", "Antiarythmique", "0.85", "CareAlger"},
		{"Celestamine", "Antihistaminique", "0.75", "PharmNow"},
		{"Amlodipine", "Antihypertenseur", "0.80", "VitalPharm"},
		{"Remicade", "Immunosuppresseur", "0.90", "AlgerMed"},
		{"Imuran", "Immunosuppresseur", "0.88", "MedHealth"},
		{"Moduretic", "Diurétique", "0.80", "MediCure"},
		{"Valtrex", "Antiviral", "0.83", "GlobalCare"},
		{"Duphaston", "Progestatif", "0.75", "MediCure"},
		{"Ketoconazole", "Antifongique", "0.78", "PharmaPro"},
		{"Mirtazapine", "Antidépresseur", "0.80", "MediX"},
		{"Vytorin", "Statines", "0.85", "MedX"},
		{"Zocor", "Statines", "0.77", "MediCare"},
		{"Trulicity", "Antidiabétique", "0.90", "NovoCare"},
		{"Jardiance", "Antidiabétique", "0.85", "MediCare"},
		{"Fludrocortisone", "Hormonal", "0.80", "HealthPlus"},
		{"Azor", "Antihypertenseur", "0.75", "GlobalMed"},
	}

	Libelles = []string{
		"Consultation médicale", "Analyse de laboratoire", "Radiographie", "Chirurgie mineure",
		"Échographie", "Examen clinique", "Opération chirurgicale", "Consultation en cardiologie",
		"Test sanguin", "Consultation ophtalmologique", "Imagerie par résonance magnétique", 
		"Test de grossesse", "Analyse génétique", "Consultation en neurologie", "Rééducation physique",
	}
	Types = []string{
		"Consultation", "Analyse", "Examen", "Intervention", "Test", "Opération", "Suivi", "Diagnostic",
	}
)