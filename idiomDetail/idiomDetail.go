package idiomDetail

import "github.com/antchfx/htmlquery"

func GetIdiomDetail(url string) string {
  doc, err := htmlquery.LoadURL(url)
  nodes, err := htmlquery.QueryAll(doc, `//*[@class="lemma-summary"]`)
  detail:= htmlquery.InnerText(nodes[0])
  if err != nil {
    panic(err)
  }
  return detail
}
