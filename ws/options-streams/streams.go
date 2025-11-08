package optionsstreams

import (
	"fmt"
	"strings"
	"github.com/openxapi/binance-go/ws/options-streams/models"
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

// NewSymbolInfoEvent stream metadata (patterns, examples, update speeds)
var NewSymbolInfoEventStreamPatterns = []string{ "option_pair" }
var NewSymbolInfoEventStreamExamples = []string{  }
var NewSymbolInfoEventUpdateSpeeds = []string{ "50ms" }

// NewSymbolInfoEventSpeed is a typed alias for supported update speeds
type NewSymbolInfoEventSpeed string
const NewSymbolInfoEventSpeed50ms NewSymbolInfoEventSpeed = "50ms"

var ValidNewSymbolInfoEventSpeeds = []NewSymbolInfoEventSpeed{ NewSymbolInfoEventSpeed50ms }

func (s NewSymbolInfoEventSpeed) String() string { return string(s) }

// BuildNewSymbolInfoEventStream builds a NewSymbolInfoEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see NewSymbolInfoEventStreamPatterns).
func BuildNewSymbolInfoEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(NewSymbolInfoEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := NewSymbolInfoEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildNewSymbolInfoEventStreams attempts to build all NewSymbolInfoEvent stream names satisfiable by provided values
func BuildNewSymbolInfoEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(NewSymbolInfoEventStreamPatterns))
	for i := range NewSymbolInfoEventStreamPatterns {
		if s, err := BuildNewSymbolInfoEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// OpenInterestEvent stream metadata (patterns, examples, update speeds)
var OpenInterestEventStreamPatterns = []string{ "{underlyingAsset}@openInterest@{expirationDate}" }
var OpenInterestEventStreamExamples = []string{ "ETH@openInterest@221125" }
var OpenInterestEventUpdateSpeeds = []string{ "60s" }

// OpenInterestEventSpeed is a typed alias for supported update speeds
type OpenInterestEventSpeed string
const OpenInterestEventSpeed60s OpenInterestEventSpeed = "60s"

var ValidOpenInterestEventSpeeds = []OpenInterestEventSpeed{ OpenInterestEventSpeed60s }

func (s OpenInterestEventSpeed) String() string { return string(s) }

// OpenInterestEventStreamParams defines placeholders for stream patterns of OpenInterestEvent
type OpenInterestEventStreamParams struct {
	UnderlyingAsset models.UnderlyingAsset // underlyingAsset
	ExpirationDate models.ExpirationDate // expirationDate
}

// Values returns non-empty placeholder values from params for OpenInterestEvent patterns
func (p OpenInterestEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.UnderlyingAsset); s != "" { out["underlyingAsset"] = s }
	if s := fmt.Sprint(p.ExpirationDate); s != "" { out["expirationDate"] = s }
	return out
}

// BuildOpenInterestEventStream builds a OpenInterestEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see OpenInterestEventStreamPatterns).
func BuildOpenInterestEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(OpenInterestEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := OpenInterestEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildOpenInterestEventStreams attempts to build all OpenInterestEvent stream names satisfiable by provided values
func BuildOpenInterestEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(OpenInterestEventStreamPatterns))
	for i := range OpenInterestEventStreamPatterns {
		if s, err := BuildOpenInterestEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// MarkPriceEvent stream metadata (patterns, examples, update speeds)
var MarkPriceEventStreamPatterns = []string{ "{underlyingAsset}@markPrice" }
var MarkPriceEventStreamExamples = []string{ "ETH@markPrice" }
var MarkPriceEventUpdateSpeeds = []string{ "1000ms" }

// MarkPriceEventSpeed is a typed alias for supported update speeds
type MarkPriceEventSpeed string
const MarkPriceEventSpeed1000ms MarkPriceEventSpeed = "1000ms"

var ValidMarkPriceEventSpeeds = []MarkPriceEventSpeed{ MarkPriceEventSpeed1000ms }

func (s MarkPriceEventSpeed) String() string { return string(s) }

// MarkPriceEventStreamParams defines placeholders for stream patterns of MarkPriceEvent
type MarkPriceEventStreamParams struct {
	UnderlyingAsset models.UnderlyingAsset // underlyingAsset
}

// Values returns non-empty placeholder values from params for MarkPriceEvent patterns
func (p MarkPriceEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.UnderlyingAsset); s != "" { out["underlyingAsset"] = s }
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

// KlineEvent stream metadata (patterns, examples, update speeds)
var KlineEventStreamPatterns = []string{ "{symbol}@kline_{interval}" }
var KlineEventStreamExamples = []string{ "BTC-200630-9000-P@kline_1m" }
var KlineEventUpdateSpeeds = []string{ "1000ms" }

// KlineEventSpeed is a typed alias for supported update speeds
type KlineEventSpeed string
const KlineEventSpeed1000ms KlineEventSpeed = "1000ms"

var ValidKlineEventSpeeds = []KlineEventSpeed{ KlineEventSpeed1000ms }

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

// TickerByUnderlyingEvent stream metadata (patterns, examples, update speeds)
var TickerByUnderlyingEventStreamPatterns = []string{ "{underlyingAsset}@ticker@{expirationDate}" }
var TickerByUnderlyingEventStreamExamples = []string{ "ETH@ticker@220930" }
var TickerByUnderlyingEventUpdateSpeeds = []string{ "1000ms" }

// TickerByUnderlyingEventSpeed is a typed alias for supported update speeds
type TickerByUnderlyingEventSpeed string
const TickerByUnderlyingEventSpeed1000ms TickerByUnderlyingEventSpeed = "1000ms"

var ValidTickerByUnderlyingEventSpeeds = []TickerByUnderlyingEventSpeed{ TickerByUnderlyingEventSpeed1000ms }

func (s TickerByUnderlyingEventSpeed) String() string { return string(s) }

// TickerByUnderlyingEventStreamParams defines placeholders for stream patterns of TickerByUnderlyingEvent
type TickerByUnderlyingEventStreamParams struct {
	UnderlyingAsset models.UnderlyingAsset // underlyingAsset
	ExpirationDate models.ExpirationDate // expirationDate
}

// Values returns non-empty placeholder values from params for TickerByUnderlyingEvent patterns
func (p TickerByUnderlyingEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.UnderlyingAsset); s != "" { out["underlyingAsset"] = s }
	if s := fmt.Sprint(p.ExpirationDate); s != "" { out["expirationDate"] = s }
	return out
}

// BuildTickerByUnderlyingEventStream builds a TickerByUnderlyingEvent stream name using a pattern index and placeholder values
// Required placeholders depend on the selected pattern (see TickerByUnderlyingEventStreamPatterns).
func BuildTickerByUnderlyingEventStream(patternIndex int, values map[string]string) (string, error) {
	if patternIndex < 0 || patternIndex >= len(TickerByUnderlyingEventStreamPatterns) { return "", fmt.Errorf("invalid pattern index") }
	pat := TickerByUnderlyingEventStreamPatterns[patternIndex]
	return buildFromPattern(pat, values)
}

// BuildTickerByUnderlyingEventStreams attempts to build all TickerByUnderlyingEvent stream names satisfiable by provided values
func BuildTickerByUnderlyingEventStreams(values map[string]string) ([]string, error) {
	out := make([]string, 0, len(TickerByUnderlyingEventStreamPatterns))
	for i := range TickerByUnderlyingEventStreamPatterns {
		if s, err := BuildTickerByUnderlyingEventStream(i, values); err == nil { out = append(out, s) }
	}
	if len(out) == 0 { return nil, fmt.Errorf("no patterns satisfied by provided values") }
	return out, nil
}

// IndexPriceEvent stream metadata (patterns, examples, update speeds)
var IndexPriceEventStreamPatterns = []string{ "{symbol}@index" }
var IndexPriceEventStreamExamples = []string{ "ETHUSDT@index" }
var IndexPriceEventUpdateSpeeds = []string{ "1000ms" }

// IndexPriceEventSpeed is a typed alias for supported update speeds
type IndexPriceEventSpeed string
const IndexPriceEventSpeed1000ms IndexPriceEventSpeed = "1000ms"

var ValidIndexPriceEventSpeeds = []IndexPriceEventSpeed{ IndexPriceEventSpeed1000ms }

func (s IndexPriceEventSpeed) String() string { return string(s) }

// IndexPriceEventStreamParams defines placeholders for stream patterns of IndexPriceEvent
type IndexPriceEventStreamParams struct {
	Symbol models.Symbol // symbol
}

// Values returns non-empty placeholder values from params for IndexPriceEvent patterns
func (p IndexPriceEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
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

// TickerEvent stream metadata (patterns, examples, update speeds)
var TickerEventStreamPatterns = []string{ "{symbol}@ticker" }
var TickerEventStreamExamples = []string{ "BTC-210630-9000-P@ticker" }
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

// TradeEvent stream metadata (patterns, examples, update speeds)
var TradeEventStreamPatterns = []string{ "{symbol}@trade", "{underlyingAsset}@trade" }
var TradeEventStreamExamples = []string{ "BTC-210630-9000-P@trade", "ETH@trade" }
var TradeEventUpdateSpeeds = []string{ "50ms" }

// TradeEventSpeed is a typed alias for supported update speeds
type TradeEventSpeed string
const TradeEventSpeed50ms TradeEventSpeed = "50ms"

var ValidTradeEventSpeeds = []TradeEventSpeed{ TradeEventSpeed50ms }

func (s TradeEventSpeed) String() string { return string(s) }

// TradeEventStreamParams defines placeholders for stream patterns of TradeEvent
type TradeEventStreamParams struct {
	Symbol models.Symbol // symbol
	UnderlyingAsset models.UnderlyingAsset // underlyingAsset
}

// Values returns non-empty placeholder values from params for TradeEvent patterns
func (p TradeEventStreamParams) Values() map[string]string {
	out := make(map[string]string)
	if s := fmt.Sprint(p.Symbol); s != "" { out["symbol"] = s }
	if s := fmt.Sprint(p.UnderlyingAsset); s != "" { out["underlyingAsset"] = s }
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

// PartialDepthEvent stream metadata (patterns, examples, update speeds)
var PartialDepthEventStreamPatterns = []string{ "{symbol}@depth{levels}", "{symbol}@depth{levels}@{speed}" }
var PartialDepthEventStreamExamples = []string{ "BTC-210630-9000-P@depth10", "BTC-210630-9000-P@depth10@100ms" }
var PartialDepthEventUpdateSpeeds = []string{ "100ms", "1000ms", "500ms" }

// PartialDepthEventSpeed is a typed alias for supported update speeds
type PartialDepthEventSpeed string
const PartialDepthEventSpeed100ms PartialDepthEventSpeed = "100ms"
const PartialDepthEventSpeed1000ms PartialDepthEventSpeed = "1000ms"
const PartialDepthEventSpeed500ms PartialDepthEventSpeed = "500ms"

var ValidPartialDepthEventSpeeds = []PartialDepthEventSpeed{ PartialDepthEventSpeed100ms, PartialDepthEventSpeed1000ms, PartialDepthEventSpeed500ms }

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


