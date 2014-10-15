<!DOCTYPE html>
<html>
<head>
<title>Go Pastebin</title>
</head>
<body>
<h3>Create a Pastebin</h3>
<div>
<form action="/new" method="post">
<p><label for="title">Title</label>
<input type="text" name="title" required /></p>
<p><label for="content">Content</label>
<textarea name="content" required></textarea></p>
<select name="language" required>
	{{range .Languages}}
	<option value={{.Id}}>{{.Name}}</option>
	{{end}}
</select>
<p><input type="submit" value="Create Pastebin" /></p>
</form>
</div>
<ul>
{{range .Pastes}}
<li><a href="/paste/{{.Id}}">{{.Title}}</a></li>
{{end}}
</ul>
</body>
</html>