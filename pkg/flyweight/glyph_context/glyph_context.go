package glyph_context

const startCurrentPosition = 0

type GlyphContext = interface {
	GetFont() string
}

type glyphContext struct {
	fonts           []string
	currentPosition int
}

func (g *glyphContext) GetFont() string {
	var font string
	font = g.fonts[g.currentPosition%len(g.fonts)]
	g.currentPosition++
	return font
}

func NewGlyphContext(fonts []string) GlyphContext {
	return &glyphContext{
		fonts:           fonts,
		currentPosition: startCurrentPosition,
	}
}
