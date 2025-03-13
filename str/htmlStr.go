package str

import (
	"bytes"
	"io"

	"strings"

	"golang.org/x/net/html"
)

func RemoveImgNode(str string) string {
	doc, _ := html.Parse(strings.NewReader(str))
	// doc.RemoveChild("img")
	doc = RemoveImg(doc)

	// for _, img := range imgs {
	// 	fmt.Println(renderNode(img))
	// 	// doc.RemoveChild(img)
	// }
	newStr := renderNode(doc)

	return newStr
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func RemoveImg(doc *html.Node) *html.Node {

	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "img" {
			// imgs = append(imgs, node)
			node.Parent.RemoveChild(node)
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)

	return doc
}
