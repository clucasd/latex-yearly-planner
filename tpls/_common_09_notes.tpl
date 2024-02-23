{{ if $.Cfg.Dotted -}} \vskip-.5\myLenLineHeightButLine\myMash{\myNumDotHeightFull}{\myNumDotWidthFull} {{- else -}}
\vbox to \dimexpr\textheight-\pagetotal\relax {%
  \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine+.6pt}\hrulefill}}\vfil
}
{{end}}