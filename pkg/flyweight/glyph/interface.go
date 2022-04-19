package glyph

type glyphContext = interface {
	GetFont() string
}

type Glyph = interface {
	Draw(glyphContext glyphContext)
}
