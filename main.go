package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 2 {
		check(os.Chdir(os.Args[1]))
	}
	b, err := exec.Command("go", "list", "-f", "{{.Dir}}", "./...").CombinedOutput()
	check(err)
	dirs := strings.Split(strings.TrimSpace(string(b)), "\n")
	m := map[string]struct{}{}
	for _, dir := range dirs {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, dir, nil, 0)
		check(err)
		for name, pkg := range pkgs {
			if strings.HasSuffix(name, "_test") {
				continue
			}
			for name, file := range pkg.Files {
				if strings.HasSuffix(name, "_test.go") {
					continue
				}
				ast.Inspect(file, func(n ast.Node) bool {
					switch x := n.(type) {
					case
						*ast.Field,      // skip struct tags
						*ast.ImportSpec: // skip imports
						return false
					case *ast.BasicLit:
						if x.Kind != token.STRING {
							return false
						}
						s, err := strconv.Unquote(x.Value)
						check(err)
						s = strings.NewReplacer(
							"\r", "\n",
							"\r\n", "\n",
							"\n\r", "\n",
						).Replace(s)
						s = strings.TrimSpace(s)
						var lines []string
						for _, line := range strings.Split(s, "\n") {
							line = strings.TrimSpace(line)
							if line != "" {
								lines = append(lines, line)
							}
						}
						s = strings.Join(lines, "\n")
						if s != "" {
							if _, ok := m[s]; !ok {
								m[s] = struct{}{}
								fmt.Println(s)
							}
						}
						return false
					}
					return true
				})
			}
		}
	}
}
