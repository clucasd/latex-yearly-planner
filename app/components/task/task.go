package task

import (
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

type Tasks []*Task
type Task struct {
	Year   int
	Number int
	Page   int
}

func NewTask(year, page, number int) *Task {
	return &Task{Year: year, Page: page, Number: number}
}

func (p Tasks) Breadcrumb(year, idx int) string {
	postfix := ""
	if idx > 0 {
		postfix = " " + strconv.Itoa(idx+1)
	}

	return header.Items{
		header.NewIntItem(year),
		header.NewTextItem("Tasks Index" + postfix).Ref(true),
	}.Table(true)
}

func (p Tasks) HeadingMOS(page, pages int) string {
	var out string

	if page > 1 && 1==0 {
		out += tex.Hyperlink(p.ref(page-1), tex.ResizeBoxW(`\myLenHeaderResizeBox`, `$\langle$`)) + " "
	}

	out += tex.Hypertarget(p.ref(page), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Tasks Index`)

	if page < pages && 1==0 {
		out += " " + tex.Hyperlink(p.ref(page+1), tex.ResizeBoxW(`\myLenHeaderResizeBox`, `$\rangle$`))
	}

	return out
}

func (p Tasks) ref(page int) string {
	var suffix string

	if page > 1 {
		suffix = " " + strconv.Itoa(page)
	}

	return "Tasks Index" + suffix
}

func (n Task) HyperLink() string {
	return hyper.Link(n.ref(), fmt.Sprintf("%02d", n.Number))
}

func (n Task) Breadcrumb() string {
	page := ""

	if n.Page > 1 {
		page = " " + strconv.Itoa(n.Page)
	}

	return header.Items{
		header.NewIntItem(n.Year),
		header.NewTextItem("Tasks Index" + page),
		header.NewTextItem(n.ref()).Ref(true),
	}.Table(true)
}

func (n Task) PrevNext(Tasks int) header.Items {
	items := header.Items{}

	if n.Number > 1 {
		items = append(items, header.NewTextItem("Task "+strconv.Itoa(n.Number-1)))
	}

	if n.Number < Tasks {
		items = append(items, header.NewTextItem("Task "+strconv.Itoa(n.Number+1)))
	}

	return items
}

func (n Task) HeadingMOS(page int) string {
	num := strconv.Itoa(n.Number)

	return tex.Hypertarget(n.ref(), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Task `+num+`\myDummyQ`)
}

func (n Task) ref() string {
	return "Task " + strconv.Itoa(n.Number)
}
