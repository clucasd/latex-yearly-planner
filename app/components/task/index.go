package task

import (
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

type Pages []Tasks

type Index struct {
	Pages Pages
}

func NewIndex(year, TasksOnPage, pages int) *Index {
	pgs := make(Pages, 0, pages)

	for pageNum := 1; pageNum <= pages; pageNum++ {
		pg := make(Tasks, 0, TasksOnPage)

		for TaskNum := 1; TaskNum <= TasksOnPage; TaskNum++ {
			pg = append(pg, NewTask(year, pageNum, (pageNum-1)*TasksOnPage+TaskNum))
		}

		pgs = append(pgs, pg)
	}

	return &Index{Pages: pgs}
}

func (i Index) PrevNext(currIdx int) header.Items {
	if len(i.Pages) <= 1 {
		return header.Items{}
	}

	list := header.Items{}

	if currIdx > 0 {
		postfix := " " + strconv.Itoa(currIdx)
		if currIdx == 1 {
			postfix = ""
		}

		list = append(list, header.NewTextItem("Tasks Index"+postfix))
	}

	if currIdx+1 < len(i.Pages) {
		list = append(list, header.NewTextItem("Tasks Index "+strconv.Itoa(currIdx+2)))
	}

	return list
}
