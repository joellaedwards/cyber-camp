package main

func page2Html() string {
	return `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<Title>Camp Challenge 2: Hike the http/https Forest</Title>
		<link rel="stylesheet" href="/styles.css">
	</head>

	<body class="forest" data-page="level2">
		<main>
		<header>
			<div class="progress-badge">Camp Challenge 2 of 3</div>
			<h1>Camp Challenge 2: Hike the http/https Forest</h1>
			<p>Sometimes it's important to stay on the train, but sometimes you can find cool hidden
			stuff by venturing off the beaten path. In this challenge, try going off the path a bit to try to
			find coolOffTrailStuff.</p>	
		</header>

		<section class="card">
			<p class="small-note">Venture off the trail to find coolOffTrailStuff.</p>
		</section>

		<footer>
		    <p>Hint: you can always change your (URL) path to /coolOffTrailStuff</p>
   	 	</footer>
	</main>
	</html>
	`
}