# Dummy Data Generator

This project is a simple tool for generating dummy data for Business Intelligence (BI) projects. Use the generator tool to create data for predefined database tables and export it to a file.
there is tow type of generation josn or sql format

## Usage

To use the generator, clone the project and run the following command:

.\generator.exe -t <table_name_in_db> -r <nbr_lines> -f <output_file.txt>


Example:

.\generator.exe -t assure -r 100 -f output_assure.txt


## Available Tables

The following tables are supported by the generator:

- **assure**
- **maladie** (28 predefined entries)
- **medicament** (80 predefined entries)
- **prestation**
- **assurance_maladie_prestation**
- **assurance_maladie_medicament**
- **maladie_medicament**

## Notes

- Ensure the database schema matches the table names for seamless data insertion.
- The generated file is in a simple text format, which can be parsed for database import.

---

## Database schema
in mysql run the schema found in db.sql


