package main

func page1Html(incorrectMsg string) string {
	return `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<Title>Camp Challenge 1: Go Fishing for Phish</Title>
		<link rel="stylesheet" href="/styles.css">
	</head>
	<body>
	<h1>Level 1: Phishing Challenge</h1>
	<p>Welcome to the first level of the cybersecurity camp! In this challenge, you will learn about phishing attacks and how to identify them.</p>
	<p>Phishing is a type of cyber attack where attackers impersonate legitimate organizations or individuals to trick victims into providing sensitive information, such as passwords, credit card numbers, or personal details. These attacks often come in the form of emails, messages, or fake websites.</p>
    <p>At Cyber Camp, phish are the ones you avoid. Read each message, spot the fishy bait, and mark the ones that are trying to trick you.</p>
	
	<form method="post" action="/checklevel1">
         <div class="email-list">
                <article class="email-card">
                    <h3>Urgent: Your Account is Suspended</h3>
                    <p>Dear user, your account has been suspended due to suspicious activity. Click here to verify your identity: http://fakebank.com/verify</p>
                    <label><input type="checkbox" name="phishing_1"> Mark as Phishing</label>
                </article>

                <article class="email-card">
                    <h3>Hey Eloise! Welcome to Our Newsletter</h3>
                    <p>Thank you for subscribing to Costco's newsletter. Please confirm your email by clicking this link: http://costco.com/confirm</p>
                    <label><input type="checkbox" name="phishing_2"> Mark as Phishing</label>
                </article>

                <article class="email-card">
                    <h3>You've Won a Prize!</h3>
                    <p>Congratulations! You've won $1000. Claim it now at http://prizeclaim.com</p>
                    <label><input type="checkbox" name="phishing_3"> Mark as Phishing</label>
                </article>

                <article class="email-card">
                    <h3>Meeting Reminder</h3>
                    <p>Don't forget our meeting tomorrow at 10 AM. Agenda attached.</p>
                    <label><input type="checkbox" name="phishing_4"> Mark as Phishing</label>
                </article>
            </div>

            <div class="action-row">
                <button type="submit">Check Answers</button>` + incorrectMsg + `
                <a class="trail-link" href="index.html">Back to Camp Entrance</a>
            </div>
        </form>
	`

}