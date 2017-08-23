package layout

import "strings"

type (
	Token struct {
		text string
		pos  int
	}

	PosByte struct {
		isEnd bool
		val   byte
	}

	LayoutRenderer struct {
		text string
	}
)

func (t *Token) Peek() int {
	if t.pos < len(t.text) {
		return (int)(t.text[t.pos])
	} else {
		return -1
	}
}

func (t *Token) Read() int {
	if t.pos < len(t.text) {
		r := (int)(t.text[t.pos])
		t.pos++
		return r
	} else {
		return -1
	}
}

func (l *LayoutRenderer) Text() string {
	return GetVariable().ConvertVariable(l.text)
}

func parseLayoutUnit(t *Token) string {
	var val string

	for ch := t.Peek(); ch != -1; ch = t.Peek() {
		if ch == '{' {
			val += "{"
			t.Read()
			continue
		}
		if ch == '}' {
			val += "}"
			break
		}
		val += string(ch)
		t.Read()
	}
	return val
}

//ReplaceLogLevelLayout use real level replace Key_LogLevel in layout
func ReplaceLogLevelLayout(layout, level string) string {
	return strings.Replace(layout, Key_LogLevel, strings.ToUpper(level), -1)
}

func CompileLayout(layout string) string {
	if layout == "" {
		return ""
	}
	var renderers []*LayoutRenderer
	var buf, layoutString string
	t := &Token{text: layout}

	for ch := t.Peek(); ch != -1; ch = t.Peek() {
		if ch == '{' {
			if len(buf) > 0 {
				renderers = append(renderers, &LayoutRenderer{text: buf})
				buf = ""
			}
			renderers = append(renderers, &LayoutRenderer{text: parseLayoutUnit(t)})
		} else {
			buf += string(ch)
		}
		t.Read()
	}
	if len(buf) > 0 {
		renderers = append(renderers, &LayoutRenderer{text: buf})
		buf = ""
	}

	for i := 0; i < len(renderers); i++ {
		layoutString += renderers[i].Text()
	}

	return layoutString
}
