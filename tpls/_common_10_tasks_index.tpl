\begin{tabularx}{\linewidth}{l|X}
  \arrayrulecolor{\myColorGray}
{{ range $task := .Body.Tasks }}
  {{ $task.HyperLink }} & \myLineHeightButLine{} \\ \hline
{{ end }}
\end{tabularx}
