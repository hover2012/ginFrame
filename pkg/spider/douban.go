package spider

import (
	"fmt"
	"gin/models"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type DoubanMove struct {
	Title    string
	Subtitle string
	Other    string
	Description     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}

func GetPages(url string) (pages []Page) {
	doc := PareseDoc(url)
	return ParsePages(doc)
}

func GetMovies(url string)(movies []DoubanMove)  {
	doc := PareseDoc(url)
	return ParseMovies(doc)
}

func PareseDoc( url string ) (doc *goquery.Document)  {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("艾玛错了")
	}
	for key, value := range header {
		reqest.Header.Add(key, value)
	}
	pageNode, err := client.Do(reqest)
	if err != nil {
		panic("艾玛又错了")
	}
	defer pageNode.Body.Close()
	//body, err := ioutil.ReadAll(pageNode.Body)

	if err != nil {
		fmt.Println("err")
	}
	//	fmt.Println(string(body))
	doc, err = goquery.NewDocumentFromReader(pageNode.Body)
	if err != nil {
		panic(err)
	}
	return doc
}


func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find(".paginator>a").Each(func(i int, selection *goquery.Selection) {
		page, _ := strconv.Atoi(selection.Text())
		url, _ := selection.Attr("href")
		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})
	return pages
	//doc.Find()
}

func ParseMovies(doc *goquery.Document) (movies []DoubanMove) {
	doc.Find(".grid_view >li").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find(".hd a span").Eq(0).Text()
		subTitle := selection.Find(".hd a span").Eq(1).Text()
		subTitle = strings.TrimLeft(subTitle, " /")

		other := selection.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, " /")

		des := selection.Find(".bd p").Eq(0).Text()
		des = strings.TrimSpace(des)
		dess := strings.Split(des, "\n")
		des = dess[0]

		movieDesc := strings.Split(dess[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := selection.Find(".bd .star .rating_num").Text()

		comment := selection.Find(".bd .star span").Eq(3).Text()
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := strings.TrimSpace(selection.Find(".bd .quote .inq").Text())

		movie := DoubanMove{
			Title: title,
			Subtitle: subTitle,
			Other: other,
			Description: des,
			Year: year,
			Area: area,
			Tag: tag,
			Star: star,
			Comment: comment,
			Quote: quote,
		}
		moviemodel := &models.MovieModel{
			Title: title,
			Subtitle: subTitle,
			Other: other,
			Description: des,
			Year: year,
			Area: area,
			Tag: tag,
			Star: star,
			Comment: comment,
			Quote: quote,
		}
		//log.Printf("i:%d movie:v%",i,movie)
		 models.AddMovie(moviemodel)
		movies = append(movies, movie)

	})
	return movies
}
