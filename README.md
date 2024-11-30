# Dummy Data Generator

This project is a simple tool for generating dummy data for Business Intelligence (BI) projects. Use the generator tool to create data for predefined database tables and export it to a file.

## Usage

To use the generator, clone the project and run the following command:

.\generator.exe -t <table_name_in_db> -r <nbr_lines> -f <output_file.txt>


Example:

.\generator.exe -t assuré -r 100 -f output_assuré.txt


## Available Tables

The following tables are supported by the generator:

- **assuré**
- **maladie** (29 predefined entries)
- **medicament** (81 predefined entries)
- **prestation**
- **assurance_maladie_prestation**
- **assurance_maladie_medicament**
- **maladie_medicament**

## Notes

- Ensure the database schema matches the table names for seamless data insertion.
- The generated file is in a simple text format, which can be parsed for database import.

---

## Database schema

CREATE TABLE assuré (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque assuré
	nom_prenom VARCHAR(100) NOT NULL, -- Nom et prénom de l'assuré
	dateN DATE NOT NULL, -- Date de naissance
	lieuN VARCHAR(100) NOT NULL, -- Lieu de naissance
	genre NOT NULL, -- Genre de l'assuré
	lieu_res VARCHAR(150), -- Lieu de résidence
	commune VARCHAR(100), -- Commune de résidence
	wilaya VARCHAR(100), -- Wilaya de résidence
	region VARCHAR(100), -- Région de résidence
	num_tel VARCHAR(15), -- Numéro de téléphone
	profession VARCHAR(100), -- Profession de l'assuré
	secteur VARCHAR(100), -- Secteur d'activité
	revenu DECIMAL(10, 2), -- Revenu mensuel
	etat_civil NOT NULL -- État civil
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE maladie (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque maladie
	code_maladie VARCHAR(20) NOT NULL UNIQUE, -- Code unique de la maladie
	libelle VARCHAR(100) NOT NULL, -- Libellé ou nom de la maladie
	taux DECIMAL(5, 2) NOT NULL, -- Taux associé à la maladie (exemple : taux de prévalence)
	type_maladie NOT NULL -- Type de maladie
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE medicament (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque médicament
	nom VARCHAR(100) NOT NULL, -- Nom du médicament
	famille VARCHAR(50), -- Famille du médicament (exemple : Antibiotique, Analgésique, etc.)
	TR DECIMAL(5, 2) NOT NULL, -- Taux de remboursement (TR) du médicament
	marque VARCHAR(50) -- Marque du médicament
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE prestation (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque prestation
	libelle VARCHAR(100) NOT NULL, -- Libellé ou description de la prestation
	type_pres NOT NULL -- Type de prestation
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE assurance_maladie_prestation (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque association
	assure_id INT NOT NULL, -- Clé étrangère vers la table assuré
	maladie_id INT NOT NULL, -- Clé étrangère vers la table maladie
	prestation_id INT NOT NULL, -- Clé étrangère vers la table prestation
	prix_prest DECIMAL(10, 2) NOT NULL, -- Prix du médicament associé à la prestation
	FOREIGN KEY (assure_id) REFERENCES assuré(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Référence à assuré
	FOREIGN KEY (maladie_id) REFERENCES maladie(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Référence à maladie
	FOREIGN KEY (prestation_id) REFERENCES prestation(id) ON DELETE CASCADE ON UPDATE CASCADE -- Référence à prestation
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE assurance_maladie_medicament (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque association
	assure_id INT NOT NULL, -- Clé étrangère vers la table assuré
	maladie_id INT NOT NULL, -- Clé étrangère vers la table maladie
	medicament_id INT NOT NULL, -- Clé étrangère vers la table prestation
	prix_medic DECIMAL(10, 2) NOT NULL, -- Prix du médicament associé à la prestation
	FOREIGN KEY (assure_id) REFERENCES assuré(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Référence à assuré
	FOREIGN KEY (maladie_id) REFERENCES maladie(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Référence à maladie
	FOREIGN KEY (medicament_id) REFERENCES medicament(id) ON DELETE CASCADE ON UPDATE CASCADE -- Référence à prestation
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE maladie_medicament (
	id INT AUTO_INCREMENT PRIMARY KEY, -- Identifiant unique pour chaque association
	maladie_id INT NOT NULL, -- Clé étrangère vers la table maladie
	medicament_id INT NOT NULL, -- Clé étrangère vers la table médicament
	FOREIGN KEY (maladie_id) REFERENCES maladie(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Référence à maladie
	FOREIGN KEY (medicament_id) REFERENCES medicament(id) ON DELETE CASCADE ON UPDATE CASCADE -- Référence à médicament
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;







Created by [ZEBDA Ahmed Yassine] | Dummy Data Generator Project
