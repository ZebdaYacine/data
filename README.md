<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dummy Data Generator - README</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 0;
            background-color: #f9f9f9;
        }

        header {
            background: #333;
            color: #fff;
            padding: 1rem 0;
            text-align: center;
        }

        main {
            padding: 20px;
            max-width: 800px;
            margin: 20px auto;
            background: #fff;
            border-radius: 5px;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        }

        h1, h2 {
            color: #333;
        }

        ul {
            list-style: square;
            margin: 0;
            padding: 0 0 0 20px;
        }

        li {
            margin-bottom: 10px;
        }

        pre {
            background: #f4f4f4;
            padding: 10px;
            border-radius: 5px;
            overflow-x: auto;
        }

        code {
            font-family: Consolas, Monaco, "Andale Mono", "Ubuntu Mono", monospace;
            color: #d6336c;
        }

        footer {
            text-align: center;
            margin-top: 20px;
            font-size: 0.9em;
            color: #666;
        }
    </style>
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
