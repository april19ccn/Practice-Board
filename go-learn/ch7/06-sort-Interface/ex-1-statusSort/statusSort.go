package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

func customLess(first, second string) func(i, j *Track) bool {
	return func(i, j *Track) bool {
		switch first {
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
		switch second {
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
			}
		}
		return false
	}
}

func main() {
	custom := statusSort{tracks, nil}

	custom.less = customLess("Title", "Year")
	sort.Sort(custom)
	printTracks(custom.tracks)

	// sort.Stable()
}

// 比较这个实现方式和重复使用sort.Stable来排序的方式。
