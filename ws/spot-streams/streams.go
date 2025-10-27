package spotstreams

import (
	"fmt"
	"strings"
	"github.com/openxapi/binance-go/ws/spot-streams/models"
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

// AggregateTradeEvent stream metadata (patterns, examples)
var AggregateTradeEventStreamPatterns = []string{ "{symbol}@aggTrade" }
var AggregateTradeEventStreamExamples = []string{ "btcusdt@aggTrade" }

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

// TradeEvent stream metadata (patterns, examples)
var TradeEventStreamPatterns = []string{ "{symbol}@trade" }
var TradeEventStreamExamples = []string{ "btcusdt@trade" }

// TradeEventStreamParams defines placeholders for stream patterns of TradeEvent
type TradeEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for TradeEvent patterns
func (p TradeEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildTradeEventStream builds a TradeEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see TradeEventStreamPatterns).
func BuildTradeEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(TradeEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := TradeEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildTradeEventStreams attempts to build all TradeEvent stream names satisfiable by provided values
func BuildTradeEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(TradeEventStreamPatterns))
	for i := range TradeEventStreamPatterns {
		if s, err := BuildTradeEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// KlineEvent stream metadata (patterns, examples, update speeds)
var KlineEventStreamPatterns = []string{ "{symbol}@kline_{interval}", "{symbol}@kline_{interval}@+08:00" }
var KlineEventStreamExamples = []string{ "btcusdt@kline_1m" }
var KlineEventUpdateSpeeds = []string{ "1000ms", "2000ms" }

// KlineEventSpeed is a typed alias for supported update speeds
type KlineEventSpeed string
const KlineEventSpeed1000ms KlineEventSpeed = "1000ms"
const KlineEventSpeed2000ms KlineEventSpeed = "2000ms"

var ValidKlineEventSpeeds = []KlineEventSpeed{ KlineEventSpeed1000ms, KlineEventSpeed2000ms }

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

// MiniTickerEvent stream metadata (patterns, examples, update speeds)
var MiniTickerEventStreamPatterns = []string{ "{symbol}@miniTicker" }
var MiniTickerEventStreamExamples = []string{ "btcusdt@miniTicker" }
var MiniTickerEventUpdateSpeeds = []string{ "1000ms" }

// MiniTickerEventSpeed is a typed alias for supported update speeds
type MiniTickerEventSpeed string
const MiniTickerEventSpeed1000ms MiniTickerEventSpeed = "1000ms"

var ValidMiniTickerEventSpeeds = []MiniTickerEventSpeed{ MiniTickerEventSpeed1000ms }

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

// AllMiniTickersEvent stream metadata (patterns, examples, update speeds)
var AllMiniTickersEventStreamPatterns = []string{ "!miniTicker@arr" }
var AllMiniTickersEventStreamExamples = []string{ "!miniTicker@arr" }
var AllMiniTickersEventUpdateSpeeds = []string{ "1000ms" }

// AllMiniTickersEventSpeed is a typed alias for supported update speeds
type AllMiniTickersEventSpeed string
const AllMiniTickersEventSpeed1000ms AllMiniTickersEventSpeed = "1000ms"

var ValidAllMiniTickersEventSpeeds = []AllMiniTickersEventSpeed{ AllMiniTickersEventSpeed1000ms }

func (s AllMiniTickersEventSpeed) String() string { return string(s) }

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
var TickerEventStreamExamples = []string{ "btcusdt@ticker" }
var TickerEventUpdateSpeeds = []string{ "1000ms" }

// TickerEventSpeed is a typed alias for supported update speeds
type TickerEventSpeed string
const TickerEventSpeed1000ms TickerEventSpeed = "1000ms"

var ValidTickerEventSpeeds = []TickerEventSpeed{ TickerEventSpeed1000ms }

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

// AllTickersEvent stream metadata (patterns, examples, update speeds)
var AllTickersEventStreamPatterns = []string{ "!ticker@arr" }
var AllTickersEventStreamExamples = []string{ "!ticker@arr" }
var AllTickersEventUpdateSpeeds = []string{ "1000ms" }

// AllTickersEventSpeed is a typed alias for supported update speeds
type AllTickersEventSpeed string
const AllTickersEventSpeed1000ms AllTickersEventSpeed = "1000ms"

var ValidAllTickersEventSpeeds = []AllTickersEventSpeed{ AllTickersEventSpeed1000ms }

func (s AllTickersEventSpeed) String() string { return string(s) }

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

// RollingWindowTickerEvent stream metadata (patterns, examples, update speeds)
var RollingWindowTickerEventStreamPatterns = []string{ "{symbol}@ticker_{windowSize}" }
var RollingWindowTickerEventStreamExamples = []string{ "btcusdt@ticker_1h" }
var RollingWindowTickerEventUpdateSpeeds = []string{ "1000ms" }

// RollingWindowTickerEventSpeed is a typed alias for supported update speeds
type RollingWindowTickerEventSpeed string
const RollingWindowTickerEventSpeed1000ms RollingWindowTickerEventSpeed = "1000ms"

var ValidRollingWindowTickerEventSpeeds = []RollingWindowTickerEventSpeed{ RollingWindowTickerEventSpeed1000ms }

func (s RollingWindowTickerEventSpeed) String() string { return string(s) }

// RollingWindowTickerEventStreamParams defines placeholders for stream patterns of RollingWindowTickerEvent
type RollingWindowTickerEventStreamParams struct {
	Symbol models.Symbol // symbol
	WindowSize models.WindowSize // windowSize
}

// Values returns non-empty placeholder values from params for RollingWindowTickerEvent patterns
func (p RollingWindowTickerEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.WindowSize); s != "" { out["windowSize"] = s }
	return out
}

// BuildRollingWindowTickerEventStream builds a RollingWindowTickerEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see RollingWindowTickerEventStreamPatterns).
func BuildRollingWindowTickerEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(RollingWindowTickerEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := RollingWindowTickerEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildRollingWindowTickerEventStreams attempts to build all RollingWindowTickerEvent stream names satisfiable by provided values
func BuildRollingWindowTickerEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(RollingWindowTickerEventStreamPatterns))
	for i := range RollingWindowTickerEventStreamPatterns {
		if s, err := BuildRollingWindowTickerEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// AllRollingWindowTickersEvent stream metadata (patterns, examples, update speeds)
var AllRollingWindowTickersEventStreamPatterns = []string{ "!ticker_{windowSize}@arr" }
var AllRollingWindowTickersEventStreamExamples = []string{ "!ticker_1h@arr" }
var AllRollingWindowTickersEventUpdateSpeeds = []string{ "1000ms" }

// AllRollingWindowTickersEventSpeed is a typed alias for supported update speeds
type AllRollingWindowTickersEventSpeed string
const AllRollingWindowTickersEventSpeed1000ms AllRollingWindowTickersEventSpeed = "1000ms"

var ValidAllRollingWindowTickersEventSpeeds = []AllRollingWindowTickersEventSpeed{ AllRollingWindowTickersEventSpeed1000ms }

func (s AllRollingWindowTickersEventSpeed) String() string { return string(s) }

// AllRollingWindowTickersEventStreamParams defines placeholders for stream patterns of AllRollingWindowTickersEvent
type AllRollingWindowTickersEventStreamParams struct {
	WindowSize models.WindowSize // windowSize
}

// Values returns non-empty placeholder values from params for AllRollingWindowTickersEvent patterns
func (p AllRollingWindowTickersEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.WindowSize); s != "" { out["windowSize"] = s }
	return out
}

// BuildAllRollingWindowTickersEventStream builds a AllRollingWindowTickersEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AllRollingWindowTickersEventStreamPatterns).
func BuildAllRollingWindowTickersEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AllRollingWindowTickersEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AllRollingWindowTickersEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAllRollingWindowTickersEventStreams attempts to build all AllRollingWindowTickersEvent stream names satisfiable by provided values
func BuildAllRollingWindowTickersEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AllRollingWindowTickersEventStreamPatterns))
	for i := range AllRollingWindowTickersEventStreamPatterns {
		if s, err := BuildAllRollingWindowTickersEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// BookTickerEvent stream metadata (patterns, examples)
var BookTickerEventStreamPatterns = []string{ "{symbol}@bookTicker" }
var BookTickerEventStreamExamples = []string{ "btcusdt@bookTicker" }

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

// AveragePriceEvent stream metadata (patterns, examples, update speeds)
var AveragePriceEventStreamPatterns = []string{ "{symbol}@avgPrice" }
var AveragePriceEventStreamExamples = []string{ "btcusdt@avgPrice" }
var AveragePriceEventUpdateSpeeds = []string{ "1000ms" }

// AveragePriceEventSpeed is a typed alias for supported update speeds
type AveragePriceEventSpeed string
const AveragePriceEventSpeed1000ms AveragePriceEventSpeed = "1000ms"

var ValidAveragePriceEventSpeeds = []AveragePriceEventSpeed{ AveragePriceEventSpeed1000ms }

func (s AveragePriceEventSpeed) String() string { return string(s) }

// AveragePriceEventStreamParams defines placeholders for stream patterns of AveragePriceEvent
type AveragePriceEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for AveragePriceEvent patterns
func (p AveragePriceEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	return out
}

// BuildAveragePriceEventStream builds a AveragePriceEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see AveragePriceEventStreamPatterns).
func BuildAveragePriceEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(AveragePriceEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := AveragePriceEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildAveragePriceEventStreams attempts to build all AveragePriceEvent stream names satisfiable by provided values
func BuildAveragePriceEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(AveragePriceEventStreamPatterns))
	for i := range AveragePriceEventStreamPatterns {
		if s, err := BuildAveragePriceEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// PartialDepthEvent stream metadata (patterns, examples, update speeds)
var PartialDepthEventStreamPatterns = []string{ "{symbol}@depth{levels}", "{symbol}@depth{levels}@{speed}" }
var PartialDepthEventStreamExamples = []string{ "btcusdt@depth5@100ms" }
var PartialDepthEventUpdateSpeeds = []string{ "100ms", "1000ms" }

// PartialDepthEventSpeed is a typed alias for supported update speeds
type PartialDepthEventSpeed string
const PartialDepthEventSpeed100ms PartialDepthEventSpeed = "100ms"
const PartialDepthEventSpeed1000ms PartialDepthEventSpeed = "1000ms"

var ValidPartialDepthEventSpeeds = []PartialDepthEventSpeed{ PartialDepthEventSpeed100ms, PartialDepthEventSpeed1000ms }

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
	if p.Levels != 0 { out["levels"] = fmt.Sprint(p.Levels) }
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
var DiffDepthEventStreamExamples = []string{ "btcusdt@depth@100ms" }
var DiffDepthEventUpdateSpeeds = []string{ "100ms", "1000ms" }

// DiffDepthEventSpeed is a typed alias for supported update speeds
type DiffDepthEventSpeed string
const DiffDepthEventSpeed100ms DiffDepthEventSpeed = "100ms"
const DiffDepthEventSpeed1000ms DiffDepthEventSpeed = "1000ms"

var ValidDiffDepthEventSpeeds = []DiffDepthEventSpeed{ DiffDepthEventSpeed100ms, DiffDepthEventSpeed1000ms }

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


