<h1>Stori Technical Challenge</h1>

This project consists on a system that reads a csv file from the mounted directory /app/directory, saves it to the database and then builds a formatted email which is sent to the customer to let him know a brief summary of his transactions in a certain period of time.
As a feature, it's able to save the csv file in AWS S3.
<hr>
<br>
<strong>GitHub repository:</strong><br> 
https://github.com/perlavianey/stori<br>
Release v1.0.0: https://github.com/perlavianey/stori/releases/tag/storiChallengeDocker
<br>

<strong>Built With:</strong><br>
<li>Golang as development language</li>
<li>Postgres as database</li>
<li>Docker as container</li>
<li>AWS S3 as storage</li>
<li>HTML for the email template</li>
<br>

<strong>Algorithm Description:</strong><br><br>
The algorithm reads any file (or files) placed in the directory /app/directory in docker, saves it to S3 using [ULID](https://github.com/ulid/spec) as filename identifiers:<br><br>

![Alt text](screenshot_s3_aws.png?raw=true "AWS S3 screenshot")

Then it saves the records data into the database table <strong>transaction</strong>. <br>The algorithm assumes that we have the account information stored in a database table (named <strong>account</strong>) which contains the accountId, the name and the email address of the customers. It's linked to the transaction with the table field <strong>account_id</strong> as foreign key:<br><br>

![Alt text](screenshot_database_model.png?raw=true "Postgres Database Model")

<br>
Then it makes the corresponding calculations to get the total balance, transactions per month and average debit and credit amounts and then builds a formatted email with the summary, which is sent to the customer in a pretty format. The csv file is also sent as an attachment to the email:<br><br>

![Alt text](screenshot_email.png?raw=true "Email screenshot")<br>

If any of the records is not stored successfully in the database, an error is logged so the system administrator can make the corresponding troubleshooting.<br><br>
<strong>Prerrequisites:</strong><br>
To run this project you need:<br>
<ol>
<li>To have Docker and Go installed.</li>
<li>To have an AWS account with S3 PUTObject permission.</li>
<li>To have the AWS access key and secret access key to that AWS account.</li>
<li>To have a S3 bucket named  <strong>transactions-stori-pv</strong> in that AWS account.</li>
</ol>
<br>
<strong>Usage:</strong><br>
To run this project:<br>
<ol>
<li>Open the file docker-compose.yml and edit the environment variables AWS_ACCESS_KEY_ID and   AWS_SECRET_ACCESS_KEY to yours.</li>
<li>Open the file stori/docker_postgres_init.sql  and edit it adding your name and email address (line 18).</li>
<li>Run the command <strong>docker compose up --build</strong> from stori location.</li>
<li>Open a new console terminal and run the command <strong>docker run stori</strong></li>
</ol>

You'll get an email to the address specified in step 2, formatted with the Stori's logo, which has the Summary of the transactions indicated in the file directory/input.csv and this file as an attachment. This summary contains the total balance, the number of transactions per month and the average debit and credit amounts.