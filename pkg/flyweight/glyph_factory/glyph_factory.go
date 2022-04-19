package glyph_factory

import "github.com/sirupsen/logrus"

type glyphContext = interface {
	GetFont() string
}

type glyph = interface {
	Draw(glyphContext glyphContext)
}

type charactersCreator func(code int, logger *logrus.Logger) glyph

type GlyphFactory interface {
	GetCharacter(code int) glyph
}

type glyphFactory struct {
	characters        []glyph
	charactersCreator charactersCreator
	logger            *logrus.Logger
}

func (g *glyphFactory) GetCharacter(code int) glyph {
	if g.characters[code] == nil {
		g.logger.Infof("Create new character with code %d", code)
		g.characters[code] = g.charactersCreator(code, g.logger)
	}
	g.logger.Infof("Return character with code %d", code)
	return g.characters[code]
}

func NewGlyphFactory(charactersCreator charactersCreator, charactersNumber int, logger *logrus.Logger) GlyphFactory {
	return &glyphFactory{
		characters:        make([]glyph, charactersNumber),
		charactersCreator: charactersCreator,
		logger:            logger,
	}
}
