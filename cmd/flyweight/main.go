package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/flyweight/glyph"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/flyweight/glyph_context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/flyweight/glyph_factory"
	"github.com/sirupsen/logrus"
)

const charactersNumber = 100

func main() {
	logger := logrus.New()

	glyphFactory := glyph_factory.NewGlyphFactory(glyph.NewCharacter, charactersNumber, logger)
	glyphContext := glyph_context.NewGlyphContext([]string{"Times New Roman 12", "Calibri 10"})

	glyph1 := glyphFactory.GetCharacter(20)
	glyph2 := glyphFactory.GetCharacter(21)
	glyph3 := glyphFactory.GetCharacter(20)

	glyph1.Draw(glyphContext)
	glyph2.Draw(glyphContext)
	glyph3.Draw(glyphContext)
}
