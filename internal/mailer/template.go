package mailer

const newsHtmlTemplate = `
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
