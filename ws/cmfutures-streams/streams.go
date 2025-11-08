package cmfuturesstreams

import (
	"fmt"
	"strings"
	"github.com/openxapi/binance-go/ws/cmfutures-streams/models"
)

// buildFromPattern replaces {placeholders} in pattern with provided values
func buildFromPattern(pattern string, values map[string]string) (string, error) {
	if values == nil { values = map[string]string{} }
	var b strings.Builder
	for i := 0; i < len(pattern); {
		open := strings.IndexByte(pattern[i:], '{')
		if open < 0 { b.WriteString(pattern[i:]); break }
		open += i
		b.WriteString(pattern[i:open])
		close := strings.IndexByte(pattern[open+1:], '}')
		if close < 0 { 
			// unmatched '{' - write the rest and stop
			b.WriteString(pattern[open:])
			break
		}
		close += open + 1
		key := pattern[open+1:close]
		val, ok := values[key]
		if !ok { return "", fmt.Errorf("missing value for placeholder '%s'", key) }
		b.WriteString(val)
		i = close + 1
	}
	return b.String(), nil
}

func containsString(list []string, v string) bool {
	for _, s := range list { if s == v { return true } }
	return false
}

// AggregateTradeEvent stream metadata (patterns, examples, update speeds)
var AggregateTradeEventStreamPatterns = []string{ "{symbol}@aggTrade" }
var AggregateTradeEventStreamExamples = []string{ "btcusd_perp@aggTrade" }
var AggregateTradeEventUpdateSpeeds = []string{ "100ms" }

// AggregateTradeEventSpeed is a typed alias for supported update speeds
type AggregateTradeEventSpeed string
const AggregateTradeEventSpeed100ms AggregateTradeEventSpeed = "100ms"

var ValidAggregateTradeEventSpeeds = []AggregateTradeEventSpeed{ AggregateTradeEventSpeed100ms }

func (s AggregateTradeEventSpeed) String() string { return string(s) }

// AggregateTradeEventStreamParams defines placeholders for stream patterns of AggregateTradeEvent
type AggregateTradeEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for AggregateTradeEvent patterns
func (p AggregateTradeEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildAggregateTradeEventStream builds a AggregateTradeEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AggregateTradeEventStreamPatterns).
func BuildAggregateTradeEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AggregateTradeEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AggregateTradeEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAggregateTradeEventStreams attempts to build all AggregateTradeEvent stream names satisfiable by provided values
func BuildAggregateTradeEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AggregateTradeEventStreamPatterns))
	for i := range AggregateTradeEventStreamPatterns {
		if s, err := BuildAggregateTradeEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// IndexPriceEvent stream metadata (patterns, examples, update speeds)
var IndexPriceEventStreamPatterns = []string{ "{pair}@indexPrice", "{pair}@indexPrice@{indexPriceInterval}" }
var IndexPriceEventStreamExamples = []string{ "btcusd@indexPrice", "btcusd@indexPrice@1s" }
var IndexPriceEventUpdateSpeeds = []string{ "3000ms", "1000ms" }

// IndexPriceEventSpeed is a typed alias for supported update speeds
type IndexPriceEventSpeed string
const IndexPriceEventSpeed3000ms IndexPriceEventSpeed = "3000ms"
const IndexPriceEventSpeed1000ms IndexPriceEventSpeed = "1000ms"

var ValidIndexPriceEventSpeeds = []IndexPriceEventSpeed{ IndexPriceEventSpeed3000ms, IndexPriceEventSpeed1000ms }

func (s IndexPriceEventSpeed) String() string { return string(s) }

// IndexPriceEventStreamParams defines placeholders for stream patterns of IndexPriceEvent
type IndexPriceEventStreamParams struct {
	Pair models.Pair // pair
	IndexPriceInterval models.IndexPriceInterval // indexPriceInterval
}

// Values returns non-empty placeholder values from params for IndexPriceEvent patterns
func (p IndexPriceEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Pair); s != "" { out["pair"] = s }
	if s := fmt.Sprint(p.IndexPriceInterval); s != "" { out["indexPriceInterval"] = s }
	return out
}

// BuildIndexPriceEventStream builds a IndexPriceEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see IndexPriceEventStreamPatterns).
func BuildIndexPriceEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(IndexPriceEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := IndexPriceEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildIndexPriceEventStreams attempts to build all IndexPriceEvent stream names satisfiable by provided values
func BuildIndexPriceEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(IndexPriceEventStreamPatterns))
	for i := range IndexPriceEventStreamPatterns {
		if s, err := BuildIndexPriceEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// MarkPriceEvent stream metadata (patterns, examples, update speeds)
var MarkPriceEventStreamPatterns = []string{ "{symbol}@markPrice", "{symbol}@markPrice@{markPriceInterval}" }
var MarkPriceEventStreamExamples = []string{ "btcusd_perp@markPrice", "btcusd_200626@markPrice@1s" }
var MarkPriceEventUpdateSpeeds = []string{ "3000ms", "1000ms" }

// MarkPriceEventSpeed is a typed alias for supported update speeds
type MarkPriceEventSpeed string
const MarkPriceEventSpeed3000ms MarkPriceEventSpeed = "3000ms"
const MarkPriceEventSpeed1000ms MarkPriceEventSpeed = "1000ms"

var ValidMarkPriceEventSpeeds = []MarkPriceEventSpeed{ MarkPriceEventSpeed3000ms, MarkPriceEventSpeed1000ms }

func (s MarkPriceEventSpeed) String() string { return string(s) }

// MarkPriceEventStreamParams defines placeholders for stream patterns of MarkPriceEvent
type MarkPriceEventStreamParams struct {
	Symbol models.Symbol // symbol
	MarkPriceInterval models.MarkPriceInterval // markPriceInterval
}

// Values returns non-empty placeholder values from params for MarkPriceEvent patterns
func (p MarkPriceEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.MarkPriceInterval); s != "" { out["markPriceInterval"] = s }
	return out
}

// BuildMarkPriceEventStream builds a MarkPriceEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see MarkPriceEventStreamPatterns).
func BuildMarkPriceEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(MarkPriceEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := MarkPriceEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildMarkPriceEventStreams attempts to build all MarkPriceEvent stream names satisfiable by provided values
func BuildMarkPriceEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(MarkPriceEventStreamPatterns))
	for i := range MarkPriceEventStreamPatterns {
		if s, err := BuildMarkPriceEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllMarkPricesEvent stream metadata (patterns, examples, update speeds)
var AllMarkPricesEventStreamPatterns = []string{ "!markPrice@arr", "!markPrice@arr@{markPriceInterval}" }
var AllMarkPricesEventStreamExamples = []string{ "!markPrice@arr", "!markPrice@arr@3s" }
var AllMarkPricesEventUpdateSpeeds = []string{ "3000ms", "1000ms" }

// AllMarkPricesEventSpeed is a typed alias for supported update speeds
type AllMarkPricesEventSpeed string
const AllMarkPricesEventSpeed3000ms AllMarkPricesEventSpeed = "3000ms"
const AllMarkPricesEventSpeed1000ms AllMarkPricesEventSpeed = "1000ms"

var ValidAllMarkPricesEventSpeeds = []AllMarkPricesEventSpeed{ AllMarkPricesEventSpeed3000ms, AllMarkPricesEventSpeed1000ms }

func (s AllMarkPricesEventSpeed) String() string { return string(s) }

// AllMarkPricesEventStreamParams defines placeholders for stream patterns of AllMarkPricesEvent
type AllMarkPricesEventStreamParams struct {
	MarkPriceInterval models.MarkPriceInterval // markPriceInterval
}

// Values returns non-empty placeholder values from params for AllMarkPricesEvent patterns
func (p AllMarkPricesEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.MarkPriceInterval); s != "" { out["markPriceInterval"] = s }
	return out
}

// BuildAllMarkPricesEventStream builds a AllMarkPricesEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllMarkPricesEventStreamPatterns).
func BuildAllMarkPricesEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllMarkPricesEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllMarkPricesEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllMarkPricesEventStreams attempts to build all AllMarkPricesEvent stream names satisfiable by provided values
func BuildAllMarkPricesEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllMarkPricesEventStreamPatterns))
	for i := range AllMarkPricesEventStreamPatterns {
		if s, err := BuildAllMarkPricesEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// KlineEvent stream metadata (patterns, examples, update speeds)
var KlineEventStreamPatterns = []string{ "{symbol}@kline_{interval}" }
var KlineEventStreamExamples = []string{ "btcusd_200626@kline_1m" }
var KlineEventUpdateSpeeds = []string{ "250ms" }

// KlineEventSpeed is a typed alias for supported update speeds
type KlineEventSpeed string
const KlineEventSpeed250ms KlineEventSpeed = "250ms"

var ValidKlineEventSpeeds = []KlineEventSpeed{ KlineEventSpeed250ms }

func (s KlineEventSpeed) String() string { return string(s) }

// KlineEventStreamParams defines placeholders for stream patterns of KlineEvent
type KlineEventStreamParams struct {
	Symbol models.Symbol // symbol
	Interval models.Interval // interval
}

// Values returns non-empty placeholder values from params for KlineEvent patterns
func (p KlineEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.Interval); s != "" { out["interval"] = s }
	return out
}

// BuildKlineEventStream builds a KlineEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see KlineEventStreamPatterns).
func BuildKlineEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(KlineEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := KlineEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildKlineEventStreams attempts to build all KlineEvent stream names satisfiable by provided values
func BuildKlineEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(KlineEventStreamPatterns))
	for i := range KlineEventStreamPatterns {
		if s, err := BuildKlineEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// ContinuousKlineEvent stream metadata (patterns, examples, update speeds)
var ContinuousKlineEventStreamPatterns = []string{ "{pair}_{contractType}@continuousKline_{interval}" }
var ContinuousKlineEventStreamExamples = []string{ "btcusd_next_quarter@continuousKline_1m" }
var ContinuousKlineEventUpdateSpeeds = []string{ "250ms" }

// ContinuousKlineEventSpeed is a typed alias for supported update speeds
type ContinuousKlineEventSpeed string
const ContinuousKlineEventSpeed250ms ContinuousKlineEventSpeed = "250ms"

var ValidContinuousKlineEventSpeeds = []ContinuousKlineEventSpeed{ ContinuousKlineEventSpeed250ms }

func (s ContinuousKlineEventSpeed) String() string { return string(s) }

// ContinuousKlineEventStreamParams defines placeholders for stream patterns of ContinuousKlineEvent
type ContinuousKlineEventStreamParams struct {
	Pair models.Pair // pair
	ContractType models.ContractType // contractType
	Interval models.Interval // interval
}

// Values returns non-empty placeholder values from params for ContinuousKlineEvent patterns
func (p ContinuousKlineEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Pair); s != "" { out["pair"] = s }
	if s := fmt.Sprint(p.ContractType); s != "" { out["contractType"] = s }
	if s := fmt.Sprint(p.Interval); s != "" { out["interval"] = s }
	return out
}

// BuildContinuousKlineEventStream builds a ContinuousKlineEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see ContinuousKlineEventStreamPatterns).
func BuildContinuousKlineEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(ContinuousKlineEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := ContinuousKlineEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildContinuousKlineEventStreams attempts to build all ContinuousKlineEvent stream names satisfiable by provided values
func BuildContinuousKlineEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(ContinuousKlineEventStreamPatterns))
	for i := range ContinuousKlineEventStreamPatterns {
		if s, err := BuildContinuousKlineEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// IndexKlineEvent stream metadata (patterns, examples, update speeds)
var IndexKlineEventStreamPatterns = []string{ "{pair}@indexPriceKline_{interval}" }
var IndexKlineEventStreamExamples = []string{ "btcusd@indexPriceKline_1h" }
var IndexKlineEventUpdateSpeeds = []string{ "250ms" }

// IndexKlineEventSpeed is a typed alias for supported update speeds
type IndexKlineEventSpeed string
const IndexKlineEventSpeed250ms IndexKlineEventSpeed = "250ms"

var ValidIndexKlineEventSpeeds = []IndexKlineEventSpeed{ IndexKlineEventSpeed250ms }

func (s IndexKlineEventSpeed) String() string { return string(s) }

// IndexKlineEventStreamParams defines placeholders for stream patterns of IndexKlineEvent
type IndexKlineEventStreamParams struct {
	Pair models.Pair // pair
	Interval models.Interval // interval
}

// Values returns non-empty placeholder values from params for IndexKlineEvent patterns
func (p IndexKlineEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Pair); s != "" { out["pair"] = s }
	if s := fmt.Sprint(p.Interval); s != "" { out["interval"] = s }
	return out
}

// BuildIndexKlineEventStream builds a IndexKlineEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see IndexKlineEventStreamPatterns).
func BuildIndexKlineEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(IndexKlineEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := IndexKlineEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildIndexKlineEventStreams attempts to build all IndexKlineEvent stream names satisfiable by provided values
func BuildIndexKlineEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(IndexKlineEventStreamPatterns))
	for i := range IndexKlineEventStreamPatterns {
		if s, err := BuildIndexKlineEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// MarkPriceKlineEvent stream metadata (patterns, examples, update speeds)
var MarkPriceKlineEventStreamPatterns = []string{ "{symbol}@markPriceKline_{interval}" }
var MarkPriceKlineEventStreamExamples = []string{ "btcusd_perp@markPriceKline_4h" }
var MarkPriceKlineEventUpdateSpeeds = []string{ "250ms" }

// MarkPriceKlineEventSpeed is a typed alias for supported update speeds
type MarkPriceKlineEventSpeed string
const MarkPriceKlineEventSpeed250ms MarkPriceKlineEventSpeed = "250ms"

var ValidMarkPriceKlineEventSpeeds = []MarkPriceKlineEventSpeed{ MarkPriceKlineEventSpeed250ms }

func (s MarkPriceKlineEventSpeed) String() string { return string(s) }

// MarkPriceKlineEventStreamParams defines placeholders for stream patterns of MarkPriceKlineEvent
type MarkPriceKlineEventStreamParams struct {
	Symbol models.Symbol // symbol
	Interval models.Interval // interval
}

// Values returns non-empty placeholder values from params for MarkPriceKlineEvent patterns
func (p MarkPriceKlineEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.Interval); s != "" { out["interval"] = s }
	return out
}

// BuildMarkPriceKlineEventStream builds a MarkPriceKlineEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see MarkPriceKlineEventStreamPatterns).
func BuildMarkPriceKlineEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(MarkPriceKlineEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := MarkPriceKlineEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildMarkPriceKlineEventStreams attempts to build all MarkPriceKlineEvent stream names satisfiable by provided values
func BuildMarkPriceKlineEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(MarkPriceKlineEventStreamPatterns))
	for i := range MarkPriceKlineEventStreamPatterns {
		if s, err := BuildMarkPriceKlineEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// MiniTickerEvent stream metadata (patterns, examples, update speeds)
var MiniTickerEventStreamPatterns = []string{ "{symbol}@miniTicker" }
var MiniTickerEventStreamExamples = []string{ "btcusd_perp@miniTicker" }
var MiniTickerEventUpdateSpeeds = []string{ "500ms" }

// MiniTickerEventSpeed is a typed alias for supported update speeds
type MiniTickerEventSpeed string
const MiniTickerEventSpeed500ms MiniTickerEventSpeed = "500ms"

var ValidMiniTickerEventSpeeds = []MiniTickerEventSpeed{ MiniTickerEventSpeed500ms }

func (s MiniTickerEventSpeed) String() string { return string(s) }

// MiniTickerEventStreamParams defines placeholders for stream patterns of MiniTickerEvent
type MiniTickerEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for MiniTickerEvent patterns
func (p MiniTickerEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildMiniTickerEventStream builds a MiniTickerEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see MiniTickerEventStreamPatterns).
func BuildMiniTickerEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(MiniTickerEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := MiniTickerEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildMiniTickerEventStreams attempts to build all MiniTickerEvent stream names satisfiable by provided values
func BuildMiniTickerEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(MiniTickerEventStreamPatterns))
	for i := range MiniTickerEventStreamPatterns {
		if s, err := BuildMiniTickerEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllMiniTickersEvent stream metadata (patterns, examples)
var AllMiniTickersEventStreamPatterns = []string{ "!miniTicker@arr" }
var AllMiniTickersEventStreamExamples = []string{ "!miniTicker@arr" }

// BuildAllMiniTickersEventStream builds a AllMiniTickersEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllMiniTickersEventStreamPatterns).
func BuildAllMiniTickersEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllMiniTickersEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllMiniTickersEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllMiniTickersEventStreams attempts to build all AllMiniTickersEvent stream names satisfiable by provided values
func BuildAllMiniTickersEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllMiniTickersEventStreamPatterns))
	for i := range AllMiniTickersEventStreamPatterns {
		if s, err := BuildAllMiniTickersEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// TickerEvent stream metadata (patterns, examples, update speeds)
var TickerEventStreamPatterns = []string{ "{symbol}@ticker" }
var TickerEventStreamExamples = []string{ "btcusd_perp@ticker" }
var TickerEventUpdateSpeeds = []string{ "500ms" }

// TickerEventSpeed is a typed alias for supported update speeds
type TickerEventSpeed string
const TickerEventSpeed500ms TickerEventSpeed = "500ms"

var ValidTickerEventSpeeds = []TickerEventSpeed{ TickerEventSpeed500ms }

func (s TickerEventSpeed) String() string { return string(s) }

// TickerEventStreamParams defines placeholders for stream patterns of TickerEvent
type TickerEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for TickerEvent patterns
func (p TickerEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildTickerEventStream builds a TickerEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see TickerEventStreamPatterns).
func BuildTickerEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(TickerEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := TickerEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildTickerEventStreams attempts to build all TickerEvent stream names satisfiable by provided values
func BuildTickerEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(TickerEventStreamPatterns))
	for i := range TickerEventStreamPatterns {
		if s, err := BuildTickerEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllTickersEvent stream metadata (patterns, examples)
var AllTickersEventStreamPatterns = []string{ "!ticker@arr" }
var AllTickersEventStreamExamples = []string{ "!ticker@arr" }

// BuildAllTickersEventStream builds a AllTickersEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllTickersEventStreamPatterns).
func BuildAllTickersEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllTickersEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllTickersEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllTickersEventStreams attempts to build all AllTickersEvent stream names satisfiable by provided values
func BuildAllTickersEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllTickersEventStreamPatterns))
	for i := range AllTickersEventStreamPatterns {
		if s, err := BuildAllTickersEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// BookTickerEvent stream metadata (patterns, examples)
var BookTickerEventStreamPatterns = []string{ "{symbol}@bookTicker" }
var BookTickerEventStreamExamples = []string{ "btcusd_perp@bookTicker" }

// BookTickerEventStreamParams defines placeholders for stream patterns of BookTickerEvent
type BookTickerEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for BookTickerEvent patterns
func (p BookTickerEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildBookTickerEventStream builds a BookTickerEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see BookTickerEventStreamPatterns).
func BuildBookTickerEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(BookTickerEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := BookTickerEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildBookTickerEventStreams attempts to build all BookTickerEvent stream names satisfiable by provided values
func BuildBookTickerEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(BookTickerEventStreamPatterns))
	for i := range BookTickerEventStreamPatterns {
		if s, err := BuildBookTickerEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllBookTickersEvent stream metadata (patterns, examples)
var AllBookTickersEventStreamPatterns = []string{ "!bookTicker" }
var AllBookTickersEventStreamExamples = []string{ "!bookTicker" }

// BuildAllBookTickersEventStream builds a AllBookTickersEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllBookTickersEventStreamPatterns).
func BuildAllBookTickersEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllBookTickersEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllBookTickersEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllBookTickersEventStreams attempts to build all AllBookTickersEvent stream names satisfiable by provided values
func BuildAllBookTickersEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllBookTickersEventStreamPatterns))
	for i := range AllBookTickersEventStreamPatterns {
		if s, err := BuildAllBookTickersEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// LiquidationEvent stream metadata (patterns, examples, update speeds)
var LiquidationEventStreamPatterns = []string{ "{symbol}@forceOrder" }
var LiquidationEventStreamExamples = []string{ "btcusd_perp@forceOrder" }
var LiquidationEventUpdateSpeeds = []string{ "1000ms" }

// LiquidationEventSpeed is a typed alias for supported update speeds
type LiquidationEventSpeed string
const LiquidationEventSpeed1000ms LiquidationEventSpeed = "1000ms"

var ValidLiquidationEventSpeeds = []LiquidationEventSpeed{ LiquidationEventSpeed1000ms }

func (s LiquidationEventSpeed) String() string { return string(s) }

// LiquidationEventStreamParams defines placeholders for stream patterns of LiquidationEvent
type LiquidationEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for LiquidationEvent patterns
func (p LiquidationEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildLiquidationEventStream builds a LiquidationEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see LiquidationEventStreamPatterns).
func BuildLiquidationEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(LiquidationEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := LiquidationEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildLiquidationEventStreams attempts to build all LiquidationEvent stream names satisfiable by provided values
func BuildLiquidationEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(LiquidationEventStreamPatterns))
	for i := range LiquidationEventStreamPatterns {
		if s, err := BuildLiquidationEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllLiquidationsEvent stream metadata (patterns, examples, update speeds)
var AllLiquidationsEventStreamPatterns = []string{ "!forceOrder@arr" }
var AllLiquidationsEventStreamExamples = []string{ "!forceOrder@arr" }
var AllLiquidationsEventUpdateSpeeds = []string{ "1000ms" }

// AllLiquidationsEventSpeed is a typed alias for supported update speeds
type AllLiquidationsEventSpeed string
const AllLiquidationsEventSpeed1000ms AllLiquidationsEventSpeed = "1000ms"

var ValidAllLiquidationsEventSpeeds = []AllLiquidationsEventSpeed{ AllLiquidationsEventSpeed1000ms }

func (s AllLiquidationsEventSpeed) String() string { return string(s) }

// BuildAllLiquidationsEventStream builds a AllLiquidationsEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllLiquidationsEventStreamPatterns).
func BuildAllLiquidationsEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllLiquidationsEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllLiquidationsEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllLiquidationsEventStreams attempts to build all AllLiquidationsEvent stream names satisfiable by provided values
func BuildAllLiquidationsEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllLiquidationsEventStreamPatterns))
	for i := range AllLiquidationsEventStreamPatterns {
		if s, err := BuildAllLiquidationsEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// ContractInfoEvent stream metadata (patterns, examples)
var ContractInfoEventStreamPatterns = []string{ "!contractInfo" }
var ContractInfoEventStreamExamples = []string{ "!contractInfo" }

// BuildContractInfoEventStream builds a ContractInfoEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see ContractInfoEventStreamPatterns).
func BuildContractInfoEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(ContractInfoEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := ContractInfoEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildContractInfoEventStreams attempts to build all ContractInfoEvent stream names satisfiable by provided values
func BuildContractInfoEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(ContractInfoEventStreamPatterns))
	for i := range ContractInfoEventStreamPatterns {
		if s, err := BuildContractInfoEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// PartialDepthEvent stream metadata (patterns, examples, update speeds)
var PartialDepthEventStreamPatterns = []string{ "{symbol}@depth{levels}", "{symbol}@depth{levels}@{speed}" }
var PartialDepthEventStreamExamples = []string{ "btcusd_perp@depth5@100ms" }
var PartialDepthEventUpdateSpeeds = []string{ "100ms", "250ms", "500ms" }

// PartialDepthEventSpeed is a typed alias for supported update speeds
type PartialDepthEventSpeed string
const PartialDepthEventSpeed100ms PartialDepthEventSpeed = "100ms"
const PartialDepthEventSpeed250ms PartialDepthEventSpeed = "250ms"
const PartialDepthEventSpeed500ms PartialDepthEventSpeed = "500ms"

var ValidPartialDepthEventSpeeds = []PartialDepthEventSpeed{ PartialDepthEventSpeed100ms, PartialDepthEventSpeed250ms, PartialDepthEventSpeed500ms }

func (s PartialDepthEventSpeed) String() string { return string(s) }

// PartialDepthEventStreamParams defines placeholders for stream patterns of PartialDepthEvent
type PartialDepthEventStreamParams struct {
	Symbol models.Symbol // symbol
	Levels models.DepthLevels // levels
	Speed models.DepthSpeed // speed
}

// Values returns non-empty placeholder values from params for PartialDepthEvent patterns
func (p PartialDepthEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.Levels); s != "" { out["levels"] = s }
	if s := fmt.Sprint(p.Speed); s != "" { out["speed"] = s }
	return out
}

// BuildPartialDepthEventStream builds a PartialDepthEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see PartialDepthEventStreamPatterns).
func BuildPartialDepthEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(PartialDepthEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := PartialDepthEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildPartialDepthEventStreams attempts to build all PartialDepthEvent stream names satisfiable by provided values
func BuildPartialDepthEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(PartialDepthEventStreamPatterns))
	for i := range PartialDepthEventStreamPatterns {
		if s, err := BuildPartialDepthEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// DiffDepthEvent stream metadata (patterns, examples, update speeds)
var DiffDepthEventStreamPatterns = []string{ "{symbol}@depth", "{symbol}@depth@{speed}" }
var DiffDepthEventStreamExamples = []string{ "btcusd_perp@depth@100ms" }
var DiffDepthEventUpdateSpeeds = []string{ "100ms", "250ms", "500ms" }

// DiffDepthEventSpeed is a typed alias for supported update speeds
type DiffDepthEventSpeed string
const DiffDepthEventSpeed100ms DiffDepthEventSpeed = "100ms"
const DiffDepthEventSpeed250ms DiffDepthEventSpeed = "250ms"
const DiffDepthEventSpeed500ms DiffDepthEventSpeed = "500ms"

var ValidDiffDepthEventSpeeds = []DiffDepthEventSpeed{ DiffDepthEventSpeed100ms, DiffDepthEventSpeed250ms, DiffDepthEventSpeed500ms }

func (s DiffDepthEventSpeed) String() string { return string(s) }

// DiffDepthEventStreamParams defines placeholders for stream patterns of DiffDepthEvent
type DiffDepthEventStreamParams struct {
	Symbol models.Symbol // symbol
	Speed models.DepthSpeed // speed
}

// Values returns non-empty placeholder values from params for DiffDepthEvent patterns
func (p DiffDepthEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.Speed); s != "" { out["speed"] = s }
	return out
}

// BuildDiffDepthEventStream builds a DiffDepthEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see DiffDepthEventStreamPatterns).
func BuildDiffDepthEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(DiffDepthEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := DiffDepthEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildDiffDepthEventStreams attempts to build all DiffDepthEvent stream names satisfiable by provided values
func BuildDiffDepthEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(DiffDepthEventStreamPatterns))
	for i := range DiffDepthEventStreamPatterns {
		if s, err := BuildDiffDepthEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}


