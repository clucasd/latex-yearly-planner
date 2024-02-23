package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/task"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func TasksIndexed(cfg config.Config, tpls []string) (page.Modules, error) {
	index := task.NewIndex(cfg.Year, cfg.Layout.Numbers.TasksOnPage, cfg.Layout.Numbers.TasksIndexPages)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	for idx, indexPage := range index.Pages {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Tasks":        indexPage,
				"Breadcrumb":   indexPage.Breadcrumb(cfg.Year, idx),
				"HeadingMOS":   indexPage.HeadingMOS(idx+1, len(index.Pages)),
				"SideQuarters": year.SideQuarters(0),
				"SideMonths":   year.SideMonths(0),
				"Extra":        index.PrevNext(idx).WithTopRightCorner(cfg.ClearTopRightCorner),
				"Extra2":       extra2(cfg.ClearTopRightCorner, false, true, false, nil, 0),
			},
		})
	}

	for idxPage, tasks := range index.Pages {
		for _, tk := range tasks {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[1],
				Body: map[string]interface{}{
					"Task":         tk,
					"Breadcrumb":   tk.Breadcrumb(),
					"HeadingMOS":   tk.HeadingMOS(idxPage),
					"SideQuarters": year.SideQuarters(0),
					"SideMonths":   year.SideMonths(0),
					"Extra": tk.
						PrevNext(cfg.Layout.Numbers.TasksOnPage * cfg.Layout.Numbers.TasksIndexPages).
						WithTopRightCorner(cfg.ClearTopRightCorner),
					"Extra2": extra2(cfg.ClearTopRightCorner, false, false, false, nil, idxPage+1),
				},
			})
		}
	}

	return modules, nil
}
