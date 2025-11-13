package coverage

const htmlTemplateContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Coverage Report</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; background: #f6f8fa; color: #24292e; line-height: 1.5; }
        .header { background: #24292e; color: white; padding: 20px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .header h1 { font-size: 28px; font-weight: 600; margin-bottom: 10px; }
        .overall-stats { display: flex; gap: 30px; margin-top: 15px; font-size: 14px; }
        .stat { display: flex; align-items: center; gap: 8px; }
        .stat-label { opacity: 0.8; }
        .stat-value { font-weight: 600; font-size: 16px; }
        .coverage-badge { display: inline-block; padding: 4px 12px; border-radius: 12px; font-weight: 600; font-size: 14px; color: white; }
        .container { display: flex; max-width: 100%; margin: 0 auto; min-height: calc(100vh - 120px); }
        .sidebar { width: 300px; background: white; border-right: 1px solid #e1e4e8; overflow-y: auto; position: sticky; top: 0; height: 100vh; }
        .sidebar-header { padding: 15px 20px; border-bottom: 1px solid #e1e4e8; font-weight: 600; background: #f6f8fa; }
        .file-tree { padding: 10px 0; }
        .tree-node { padding: 6px 20px; cursor: pointer; display: flex; align-items: center; gap: 8px; transition: background 0.2s; }
        .tree-node:hover { background: #f6f8fa; }
        .tree-node.active { background: #e1e4e8; font-weight: 600; }
        .tree-icon { width: 16px; font-size: 12px; }
        .tree-coverage { margin-left: auto; font-size: 12px; padding: 2px 6px; border-radius: 6px; font-weight: 600; color: white; }
        .content { flex: 1; padding: 20px; overflow-x: auto; }
        .file-section { background: white; border-radius: 6px; margin-bottom: 20px; border: 1px solid #e1e4e8; overflow: hidden; }
        .file-header { padding: 15px 20px; background: #f6f8fa; border-bottom: 1px solid #e1e4e8; display: flex; justify-content: space-between; align-items: center; }
        .file-name { font-weight: 600; font-size: 16px; font-family: monospace; }
        .file-stats { display: flex; gap: 15px; font-size: 13px; align-items: center; }
        .code-container { overflow-x: auto; }
        .code-table { width: 100%; border-collapse: collapse; font-family: monospace; font-size: 13px; }
        .code-table td { padding: 0; vertical-align: top; }
        .line-number { width: 50px; text-align: right; padding: 2px 10px; color: #6a737d; user-select: none; background: #f6f8fa; border-right: 1px solid #e1e4e8; }
        .line-content { padding: 2px 10px; white-space: pre; overflow-x: auto; }
        .line-covered { background: #e6ffed; }
        .line-uncovered { background: #ffeef0; }
        .line-neutral { background: white; }
        .no-source { padding: 40px; text-align: center; color: #6a737d; }
        .summary-table { width: 100%; border-collapse: collapse; background: white; border-radius: 6px; overflow: hidden; }
        .summary-table th { background: #f6f8fa; padding: 12px 15px; text-align: left; font-weight: 600; border-bottom: 1px solid #e1e4e8; }
        .summary-table td { padding: 10px 15px; border-bottom: 1px solid #e1e4e8; }
        .summary-table tr:last-child td { border-bottom: none; }
        .summary-table tr:hover { background: #f6f8fa; }
        .path-cell { font-family: monospace; font-size: 13px; }
        .coverage-cell { text-align: center; width: 100px; }
        .statements-cell { text-align: center; width: 120px; font-size: 13px; color: #6a737d; }
        .section-title { font-size: 20px; font-weight: 600; margin-bottom: 15px; padding-bottom: 10px; border-bottom: 2px solid #e1e4e8; }
    </style>
</head>
<body>
    <div class="header">
        <h1>📊 Go Coverage Report</h1>
        <div class="overall-stats">
            <div class="stat">
                <span class="stat-label">Overall Coverage:</span>
                <span class="coverage-badge" style="background: {{.OverallColor}}">{{formatPct .OverallPct}}</span>
            </div>
            <div class="stat">
                <span class="stat-label">Statements:</span>
                <span class="stat-value">{{.CoveredStmts}} / {{.TotalStmts}}</span>
            </div>
            <div class="stat">
                <span class="stat-label">Mode:</span>
                <span class="stat-value">{{.Mode}}</span>
            </div>
        </div>
    </div>
    <div class="container">
        <div class="sidebar">
            <div class="sidebar-header">📁 Files</div>
            <div class="file-tree">
                {{range .Files}}
                <div class="tree-node" onclick="scrollToFile('{{.Path}}')" data-file="{{.Path}}">
                    <span class="tree-icon">📄</span>
                    <span>{{.Name}}</span>
                    <span class="tree-coverage" style="background: {{.Color}}">{{formatPct .Coverage}}</span>
                </div>
                {{end}}
            </div>
        </div>
        <div class="content">
            <div class="section-title">Coverage Summary</div>
            <div class="file-section">
                <table class="summary-table">
                    <thead>
                        <tr>
                            <th>File</th>
                            <th class="coverage-cell">Coverage</th>
                            <th class="statements-cell">Statements</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Files}}
                        <tr onclick="scrollToFile('{{.Path}}')" style="cursor: pointer;">
                            <td class="path-cell">{{.Path}}</td>
                            <td class="coverage-cell">
                                <span class="coverage-badge" style="background: {{.Color}}">{{formatPct .Coverage}}</span>
                            </td>
                            <td class="statements-cell">{{.Covered}} / {{.Total}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
            <div class="section-title" style="margin-top: 40px;">File Details</div>
            {{range .Files}}
            <div class="file-section" id="file-{{.Path}}">
                <div class="file-header">
                    <div class="file-name">{{.Path}}</div>
                    <div class="file-stats">
                        <span>{{.Covered}} / {{.Total}} statements</span>
                        <span class="coverage-badge" style="background: {{.Color}}">{{formatPct .Coverage}}</span>
                    </div>
                </div>
                {{if .HasSource}}
                <div class="code-container">
                    <table class="code-table">
                        {{range .Lines}}
                        <tr class="{{if .IsCovered}}line-covered{{else if .HasCoverage}}line-uncovered{{else}}line-neutral{{end}}">
                            <td class="line-number">{{.LineNumber}}</td>
                            <td class="line-content">{{.Content}}</td>
                        </tr>
                        {{end}}
                    </table>
                </div>
                {{else}}
                <div class="no-source">⚠️ Source file not found<br><small style="font-size: 12px; opacity: 0.8;">Run go-coverage from your project root directory</small></div>
                {{end}}
            </div>
            {{end}}
        </div>
    </div>
    <script>
        function scrollToFile(filePath) {
            const element = document.getElementById('file-' + filePath);
            if (element) {
                element.scrollIntoView({ behavior: 'smooth', block: 'start' });
                document.querySelectorAll('.tree-node').forEach(node => node.classList.remove('active'));
                document.querySelector('.tree-node[data-file="' + filePath + '"]').classList.add('active');
            }
        }
    </script>
</body>
</html>`
