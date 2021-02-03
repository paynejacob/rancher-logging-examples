package index

import "time"

const listPageTemplate = `
<html>
	<table>
		<thead>
			<td>Name</td>
			<td>First Log</td>
			<td>Last Log</td>
			<td>Log Count</td>
		</thead>
		<tbody>
			{{range .}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.FirstLog | DateFormat }}</td>
				<td>{{.LastLog | DateFormat}}</td>
				<td>{{.LogCount}}</td>
				<td><a href="/{{.Name}}/">view logs</a></td>
			</tr>
			{{ end }}
		</tbody>
	</table>
</html>
`

func DateFormat(t time.Time) string {
	return t.Format(time.RFC3339)
}