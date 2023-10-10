package mailer

const newsHtmlTemplate = `
<body style="color: black; width: 1000px;">
	<img src="cid:logo.png" alt="Company Logo" style="margin: 16px auto; width: 200px; display: block;"/>
	<div style="width: 480px; margin: auto;">
		<p>Dear {{.S.Name}},</p>
		<p>The Ai News Letter on {{.T.Format "2006/01/02 03:04:05 pm"}} is in now! Lets take a look at the lastest news.</p>
		{{range $i, $n := .Ns}}
		<div style="margin-top: 32px; margin-bottom: 32px; color: black;">
			<img src="{{$n.ImgUrl}}" alt="Article Image" width="480" height="270"/>
			<h3 style="font-size: 1.8rem; margin: 8px 0;"><a href="{{$n.Url}}">{{$n.Title}}</a></h3>
			<p style="color: black;">{{$n.AISummary}}</p>
			<small style="color: black;">Published at: {{$n.PublishedAt.Format "2006/01/02 03:04:05 pm"}}</small>
		</div>
		{{end}}
		<i style="color: black; font-size: 0.8rem;">Your Most Reliable AI Mailer</i>
		<div style="background-color: #F1EFEF; width: 480px; padding: 36px; margin-top:16px;">
			<a style="display: block;" href="http://locahost/unsubscribe">Unsubscribe</a>
			<small style="color: black;">TMC &copy;2023 All Rights Reserved</small>
		</div>
	</div>
</body>
`

const newsTextTemplate = `
	Hi {{.S.Name}}, 
		The AI newsletter on {{.T.Format "2006/01/02 03:04:05 pm"}} is here now!
		{{range $i, $n := .Ns}}
		Article {{$i}}:
			Title: {{$n.Title}}
			Summary: {{$n.AISummary}}
			Published Time: {{$n.PublishedAt.Format "2006/01/02 03:04:05 pm"}}
		{{end}}
	Your Reliable TMC Mailer
`
