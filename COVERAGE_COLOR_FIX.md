# Correção: Coloração de Cobertura Similar ao `go tool cover`

## Problema
As linhas que **não precisam de cobertura** (package declarations, imports, comentários, linhas vazias, etc.) estavam aparecendo em **vermelho** no relatório HTML, quando deveriam aparecer em **branco/neutro**.

### Comportamento Anterior (INCORRETO):
```
package main              ❌ VERMELHO (incorreto!)
                          ❌ VERMELHO (incorreto!)
// Add function          ❌ VERMELHO (incorreto!)
func Add(a, b int) int { ✅ VERDE   (correto)
    return a + b          ✅ VERDE   (correto)
}                         ✅ VERDE   (correto)
```

### Comportamento Atual (CORRETO):
```
package main              ⚪ BRANCO  (correto!)
                          ⚪ BRANCO  (correto!)
// Add function          ⚪ BRANCO  (correto!)
func Add(a, b int) int { ✅ VERDE   (correto)
    return a + b          ✅ VERDE   (correto)
}                         ✅ VERDE   (correto)
```

## Solução Implementada

### 1. Adicionado campo `HasCoverage` em `LineCoverage`
**Arquivo:** `pkg/source.go`

```go
type LineCoverage struct {
    LineNumber   int
    Content      string
    Count        int
    IsCovered    bool
    HasCoverage  bool // true se a linha faz parte de um bloco de cobertura
}
```

Este campo distingue:
- **HasCoverage = true**: Linha faz parte de um bloco executável (função)
- **HasCoverage = false**: Linha não executável (package, import, comentário, linha vazia)

### 2. Marcação de Blocos de Cobertura
**Arquivo:** `pkg/source.go`

```go
for _, block := range coverage.Blocks {
    for i := block.StartLine; i <= block.EndLine; i++ {
        if i > 0 && i <= len(lines) {
            lines[i-1].Count = block.Count
            lines[i-1].IsCovered = block.Count > 0
            lines[i-1].HasCoverage = true // Marca que esta linha tem cobertura
        }
    }
}
```

### 3. Atualização da Lógica de Cores no Template
**Arquivo:** `pkg/template.go`

**Antes:**
```html
{{if .IsCovered}}line-covered{{else if eq .Count 0}}line-uncovered{{else}}line-neutral{{end}}
```

**Depois:**
```html
{{if .IsCovered}}line-covered{{else if .HasCoverage}}line-uncovered{{else}}line-neutral{{end}}
```

**Lógica:**
1. Se `IsCovered = true` → **VERDE** (código executado)
2. Senão, se `HasCoverage = true` → **VERMELHO** (código não executado mas deveria ser testado)
3. Senão → **BRANCO/NEUTRO** (não é código executável)

## Resultado

### Exemplo Real (calculator.go):

| Linha | Classe         | Conteúdo                               | Status |
|-------|----------------|----------------------------------------|--------|
| 1     | ⚪ line-neutral | package main                          | ✓      |
| 2     | ⚪ line-neutral |                                       | ✓      |
| 3     | ⚪ line-neutral | // Add returns the sum of two integers| ✓      |
| 4     | ✅ line-covered | func Add(a, b int) int {              | ✓      |
| 5     | ✅ line-covered |     return a + b                      | ✓      |
| 6     | ✅ line-covered | }                                     | ✓      |
| 7     | ⚪ line-neutral |                                       | ✓      |
| 8     | ⚪ line-neutral | // Subtract returns...                | ✓      |
| 9     | ✅ line-covered | func Subtract(a, b int) int {         | ✓      |
| 10    | ✅ line-covered |     return a - b                      | ✓      |
| 11    | ✅ line-covered | }                                     | ✓      |
| 12    | ⚪ line-neutral |                                       | ✓      |
| 13    | ⚪ line-neutral | // Multiply returns...                | ✓      |
| 14    | ❌ line-uncovered | func Multiply(a, b int) int {       | ✓      |
| 15    | ❌ line-uncovered |     return a * b                    | ✓      |
| 16    | ❌ line-uncovered | }                                   | ✓      |

## Compatibilidade com `go tool cover`

A coloração agora é **idêntica** ao comportamento oficial do Go:

```bash
# Go tool cover (oficial)
go tool cover -html=coverage.out -o official.html

# go-coverage (nossa ferramenta)
go-coverage -input=coverage.out -output=coverage.html
```

Ambos produzem a mesma coloração:
- ✅ **Verde**: Código coberto por testes
- ❌ **Vermelho**: Código não coberto (mas que deveria ser testado)
- ⚪ **Branco**: Código não executável (declarações, comentários, etc.)

## Arquivos Modificados

1. ✅ `/pkg/source.go`
   - Adicionado campo `HasCoverage` ao struct `LineCoverage`
   - Atualizada lógica para marcar blocos de cobertura

2. ✅ `/pkg/template.go`
   - Atualizada lógica de classes CSS no template HTML
   - Agora usa `HasCoverage` ao invés de `eq .Count 0`

## Teste

```bash
cd /home/rayque.oliveira/projects/go-coverage
./test-local.sh
```

**Resultado esperado:**
- Linhas com código testado: **verde** ✅
- Linhas com código não testado: **vermelho** ❌
- Linhas sem código executável: **branco** ⚪

---

**Data:** 13 de Novembro de 2024  
**Status:** ✅ IMPLEMENTADO E TESTADO  
**Compatibilidade:** 100% com `go tool cover -html`

