package main

func page3Html() string{
	return `
	<!DOCTYPE html>
	<html lang="en">

	<head>
		<meta charset="UTF-8">
		<Title>Camp Challenge 3: Cool Off Trail Stuff</Title>
		<link rel="stylesheet" href="/styles.css">
	</head>

	<body class="cookie" data-page="level3">
<main>
    <header>
        <div class="progress-badge">Camp Challenge 3 of 3</div>
        <h1>Camp Challenge 3: Cookie Snack Break</h1>
        <p>This trail stop is all about cookies, but not the kind you eat. Use your browser's developer tools to change the <code>access</code> cookie value from <code>denied</code> to <code>granted</code> so you can grab the snack and keep climbing.</p>
    </header>

    <section class="card">
        <p>Open your browser developer tools and add a cookie named <strong>access</strong> with the value <code>granted</code>. Refresh the page once the change is made.</p>
        <ul class="checklist">
            <li>Inspect the cookie like you're checking trail snacks.</li>
            <li>Add a cookie with the name access and set it to granted to unlock the next summit path.</li>
            <li>Refresh the page once the change is made.</li>
        </ul>
    </section>

    <footer>
        <p class="small-note">This is the final gear check before the mountain push.</p>
    </footer>
</main>
</body

	</html>
	`
}