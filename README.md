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

Created by [ZEBDA Ahmed Yassine] | Dummy Data Generator Project
