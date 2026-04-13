package main

func welcomeHtml() string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Welcome to Cybersecurity Camp!</title>
		<link rel="stylesheet" href="/styles.css">
	</head>
	<body class="welcome">
	<main>
		<header>
		<h1>Welcome to Cybersecurity Camp!</h1>
		<p>We are excited to have you here. This camp is designed to introduce you to the world of cybersecurity 
		and help you develop the skills needed to protect yourself and others online.</p>
		</header>

		<section class="what-to-expect">
		<h2>What to Expect</h2>
		<p>Throughout the camp, you will participate in hands-on activities, workshops, and challenges that will 
		teach you about various cybersecurity topics such as network security, cryptography, ethical hacking, 
		and more. You will also have the opportunity to collaborate with your peers and learn from experienced 
		instructors.</p>
		</section>

		<section class="hero-card">
        <div class="hero-scene lake-scene">
            <div class="sun"></div>
            <div class="hill hill-left"></div>
            <div class="hill hill-right"></div>
            <div class="lake">
                <div class="wave wave1"></div>
                <div class="wave wave2"></div>
                <div class="fish"></div>
            </div>
        </div>
        <a class="button-link" href="/level1">Start Fishing</a>
        <p class="small-note">The camp trail begins at the lake. Catch the phish before they reel you in.</p>
    	</section>
		`
}