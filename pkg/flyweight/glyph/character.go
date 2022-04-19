package glyph

import "github.com/sirupsen/logrus"

type character struct {
	code   int
	logger *logrus.Logger
}

func (g *character) Draw(glyphContext glyphContext) {
	font := glyphContext.GetFont()
	g.logger.Infof("Glyph with code %d draw with font %s", g.code, font)
}

func NewCharacter(code int, logger *logrus.Logger) Glyph {
	return &character{
		code:   code,
		logger: logger,
	}
}
