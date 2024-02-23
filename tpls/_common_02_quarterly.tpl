\begin{minipage}[t][\remainingHeight]{\myLenTriCol}
{{- range $j, $month := .Body.Quarter.Months -}}
{\noindent\renewcommand{\arraystretch}{0}%
{{- template "monthTabularV2.tpl" dict "Month" $month "TableType" "tabular" -}}
{{- if ne $j 2 -}} \vfill {{- end -}}
{{- end -}}
\end{minipage}%
\hspace{\myLenTriColSep}%
\begin{minipage}[t][\remainingHeight]{\dimexpr2\myLenTriCol+\myLenTriColSep}
  \vskip-\myLenLineHeightButLine
  \vbox to \dimexpr\textheight-\pagetotal+3pt{\myMash[\myQuarterlySpring]{\myNumQuarterlyLines}{\myNumDotWidthTwoThirds}\vspace{0pt plus 1pt}}
\end{minipage}