// Generated from Xql.g4 by ANTLR 4.7.

package grammar // Xql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseXqlListener is a complete listener for a parse tree produced by XqlParser.
type BaseXqlListener struct{}

var _ XqlListener = &BaseXqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseXqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseXqlListener) ExitQuery(ctx *QueryContext) {}

// EnterEmbbedExpr is called when production embbedExpr is entered.
func (s *BaseXqlListener) EnterEmbbedExpr(ctx *EmbbedExprContext) {}

// ExitEmbbedExpr is called when production embbedExpr is exited.
func (s *BaseXqlListener) ExitEmbbedExpr(ctx *EmbbedExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseXqlListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseXqlListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseXqlListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseXqlListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseXqlListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseXqlListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterBooleanExpr is called when production booleanExpr is entered.
func (s *BaseXqlListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production booleanExpr is exited.
func (s *BaseXqlListener) ExitBooleanExpr(ctx *BooleanExprContext) {}

// EnterLeftexpr is called when production leftexpr is entered.
func (s *BaseXqlListener) EnterLeftexpr(ctx *LeftexprContext) {}

// ExitLeftexpr is called when production leftexpr is exited.
func (s *BaseXqlListener) ExitLeftexpr(ctx *LeftexprContext) {}

// EnterValue is called when production value is entered.
func (s *BaseXqlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseXqlListener) ExitValue(ctx *ValueContext) {}
