<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dummy Data Generator - README</title>
   
</head>
<body>
    <header>
        <h1>Dummy Data Generator - README</h1>
    </header>
    <main>
        <h2>Overview</h2>
        <p>
            This dummy data generator allows you to create simple, random data for various database tables. It is designed 
            to be used in BI (Business Intelligence) projects for testing and simulation purposes.
        </p>

        <h2>Usage</h2>
        <p>Run the generator using the following syntax:</p>
        <pre><code>.\generator.exe -t table_name_in_db -r nbr_lines -f output_file.txt</code></pre>
        <p>Here is a breakdown of the command-line arguments:</p>
        <ul>
            <li><code>-t</code>: The name of the table for which data should be generated.</li>
            <li><code>-r</code>: The number of rows of data to generate.</li>
            <li><code>-f</code>: The file where the generated data will be saved.</li>
        </ul>

        <h2>Available Tables</h2>
        <p>The following tables are supported for data generation:</p>
        <ul>
            <li><strong>assuré</strong></li>
            <li><strong>maladie</strong> (29 predefined entries)</li>
            <li><strong>medicament</strong> (81 predefined entries)</li>
            <li><strong>prestation</strong></li>
            <li><strong>assurance_maladie_prestation</strong></li>
            <li><strong>assurance_maladie_medicament</strong></li>
            <li><strong>maladie_medicament</strong></li>
        </ul>

        <h2>Example</h2>
        <p>To generate 100 lines of data for the table <code>assuré</code> and save it to <code>output.txt</code>:</p>
        <pre><code>.\generator.exe -t assuré -r 100 -f output.txt</code></pre>

        <h2>Output</h2>
        <p>The generated data will be saved in the specified file in plain text format.</p>
    </main>
    <footer>
        <p>Dummy Data Generator &copy; 2024</p>
    </footer>
</body>
</html>
