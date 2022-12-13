package main

func main() {
	connector := connector{"admin", "test", "tcp(192.168.1.17:8000)", "lbr"}
	authors := connector.getAuthors()
	authorHtml := HTML{"author.html", "author", ""}
	buildFromTemplate(authors, authorHtml)
	for _, element := range authors {
		buildFromTemplate(connector.getAuthorPage(element), HTML{element.Name + ".html", "authorPage", "authors/"})
	}
}
