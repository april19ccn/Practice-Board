package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// 类型转换
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

// 可视化列表
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。\
// 定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。

type statusSort struct {
	tracks []*Track
	less   func(i, j *Track) bool
}

func (x statusSort) Len() int           { return len(x.tracks) }
func (x statusSort) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x statusSort) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

func customLess(key ...string) func(i, j *Track) bool {
	return func(i, j *Track) bool {
		for _, k := range key {
			switch k {
			case "Title":
				if i.Title != j.Title {
					return i.Title < j.Title
				}
			case "Artist":
				if i.Artist != j.Artist {
					return i.Artist < j.Artist
				}
			case "Album":
				if i.Album != j.Album {
					return i.Album < j.Album
				}
			case "Year":
				if i.Year != j.Year {
					return i.Year < j.Year
				}
			case "Length":
				if i.Length != j.Length {
					return i.Length < j.Length
				}
			}
		}
		return false
	}
}

// 练习 7.9： 使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

const templ = `
<h1>List -- </h1>
<table>
<tr style='text-align: left'>
    <th><a href="?sort=Title">Title</a></th>
    <th><a href="?sort=Artist">Artist</a></th>
    <th><a href="?sort=Album">Album</a></th>
    <th><a href="?sort=Year">Year</a></th>
    <th><a href="?sort=Length">Length</a></th>
</tr>
{{range .}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
<a href="?sort=Title&sort=Year">按照title主排序和year次排序</a>
<a href="?sort=Title&sort=Album">按照title主排序和album次排序</a>
`

func renderTable(out io.Writer, result statusSort) {
	list := template.Must(template.New("list").Parse(templ))
	if err := list.Execute(out, result.tracks); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 复制原始切片以避免修改全局数据
	tracksCopy := make([]*Track, len(tracks))
	copy(tracksCopy, tracks)

	custom := statusSort{tracksCopy, nil}

	// custom.less = customLess("Title", "Year")
	// sort.Sort(custom)
	// printTracks(custom.tracks)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 使用 r.URL.Query()["sort"] 获取所有 sort 参数
		sortKeys := r.URL.Query()["sort"]
		if len(sortKeys) > 0 {
			fmt.Println(sortKeys)
			custom.less = customLess(sortKeys...)
			sort.Sort(custom)
		}

		renderTable(w, custom)
	})
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
