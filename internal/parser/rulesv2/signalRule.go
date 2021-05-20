package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func parseSignal(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.FATAL && nxt.Token != tokens.SIGNAL && nxt.Token != tokens.WARN {
		return nil, UnexpectedError(nxt, "fatal", "warn", "signal")
	}
	signalNode := &ast.SignalNode{}
	lvl := "signal"
	if nxt.Token == tokens.FATAL {
		lvl = "fatal"
	} else if nxt.Token == tokens.WARN {
		lvl = "warn"
	}
	signalNode.Level = lvl
	expr, err := ExpressionRule(p)
	signalNode.ValueToRaise = expr
	return signalNode, err
}
