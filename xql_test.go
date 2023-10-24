package xql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleSyntaxErr(t *testing.T) {
	var testData = []string{
		"asdasdasd",
		"((a=1",
		"((a='foo'",
		"((a='foo')",
		"a=foo",
		"a in (foo,bar)",
		"a in (foo)",
		"a ∩ (foo)",
		"a !=   foo",
	}
	data := map[string]interface{}{
		"a": "foo",
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc, data)
		ass.False(ok)
		ass.Error(err)
	}
}

func TestMatch_Int(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": 9,
			},
			out: false,
		},
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": 10,
			},
			out: true,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": 10,
			},
			out: false,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": 11,
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": 10,
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": 11,
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": 1,
			},
			out: false,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": 1,
			},
			out: true,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": 10,
			},
			out: false,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": 11,
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_Int64(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": int64(9),
			},
			out: false,
		},
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": int64(10),
			},
			out: true,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": int64(10),
			},
			out: false,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": int64(11),
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": int64(10),
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": int64(11),
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": int64(1),
			},
			out: false,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": int64(1),
			},
			out: true,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": int64(10),
			},
			out: false,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": int64(11),
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_String(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": "10",
			},
			out: true,
		},
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": "010",
			},
			out: false,
		},
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": "",
			},
			out: false,
		},
		{
			rawXql: `a>1`,
			data: map[string]interface{}{
				"a": "1",
			},
			out: false,
		},
		{
			rawXql: `a>1`,
			data: map[string]interface{}{
				"a": "2",
			},
			out: true,
		},
		{
			rawXql: `a>=1`,
			data: map[string]interface{}{
				"a": "1",
			},
			out: true,
		},
		{
			rawXql: `a>=1`,
			data: map[string]interface{}{
				"a": "2",
			},
			out: true,
		},
		{
			rawXql: `a>=1`,
			data: map[string]interface{}{
				"a": "",
			},
			out: false,
		},
		{
			rawXql: `a>20`,
			data: map[string]interface{}{
				"a": "21",
			},
			out: true,
		},
		{
			rawXql: `a>20`,
			data: map[string]interface{}{
				"a": "3",
			},
			out: true,
		},
		{
			rawXql: `a>20`,
			data: map[string]interface{}{
				"a": "1",
			},
			out: false,
		},
		{
			rawXql: `a!=20`,
			data: map[string]interface{}{
				"a": "3",
			},
			out: true,
		},
		{
			rawXql: `a!=20`,
			data: map[string]interface{}{
				"a": "20",
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch__Float(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": float64(10),
			},
			out: true,
		},
		{
			rawXql: `a=10`,
			data: map[string]interface{}{
				"a": float64(10 - epsilon*0.1),
			},
			out: true,
		},
		{
			rawXql: `a!=10`,
			data: map[string]interface{}{
				"a": float64(10),
			},
			out: false,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": float64(10),
			},
			out: false,
		},
		{
			rawXql: `a>10`,
			data: map[string]interface{}{
				"a": float64(10.000000001),
			},
			out: true,
		},
		{
			rawXql: `a<10`,
			data: map[string]interface{}{
				"a": float64(10 - epsilon*0.1),
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": float64(10.0),
			},
			out: true,
		},
		{
			rawXql: `a>=10`,
			data: map[string]interface{}{
				"a": float64(9.0),
			},
			out: false,
		},
		{
			rawXql: `a<=10`,
			data: map[string]interface{}{
				"a": float64(10.0),
			},
			out: true,
		},
		{
			rawXql: `a<=10`,
			data: map[string]interface{}{
				"a": float64(9.0),
			},
			out: true,
		},
		{
			rawXql: `a<=10`,
			data: map[string]interface{}{
				"a": float64(11.0),
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, _ := Match(tc.rawXql, tc.data)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_Boolean(t *testing.T) {
	// comparison between booleans only support =,!=
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=true`,
			data: map[string]interface{}{
				"a": true,
			},
			out: true,
		},
		{
			rawXql: `a=true`,
			data: map[string]interface{}{
				"a": false,
			},
			out: false,
		},
		{
			rawXql: `a=false`,
			data: map[string]interface{}{
				"a": true,
			},
			out: false,
		},
		{
			rawXql: `a=false`,
			data: map[string]interface{}{
				"a": false,
			},
			out: true,
		},
		{
			rawXql: `a!=false`,
			data: map[string]interface{}{
				"a": true,
			},
			out: true,
		},
		{
			rawXql: `a>false`,
			data: map[string]interface{}{
				"a": true,
			},
			out: false,
		},
		{
			rawXql: `a<=false`,
			data: map[string]interface{}{
				"a": true,
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_In(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a in ('swim')`,
			data: map[string]interface{}{
				"a": "swim",
			},
			out: true,
		},
		{
			rawXql: `a in ('soccer')`,
			data: map[string]interface{}{
				"a": "swim",
			},
			out: false,
		},
		{
			rawXql: `a !in ('swim')`,
			data: map[string]interface{}{
				"a": "swim",
			},
			out: false,
		},
		{
			rawXql: `a !in ('soccer')`,
			data: map[string]interface{}{
				"a": "swim",
			},
			out: true,
		},
		{
			rawXql: `a !in (1,2,3)`,
			data: map[string]interface{}{
				"a": 1,
			},
			out: false,
		},
		{
			rawXql: `a in (1,2,3)`,
			data: map[string]interface{}{
				"a": 3,
			},
			out: true,
		},
		{
			rawXql: `a in (1,5,9)`,
			data: map[string]interface{}{
				"a": 5,
			},
			out: true,
		},
		{
			rawXql: `a !in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": int64(9),
			},
			out: true,
		},
		{
			rawXql: `a !in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": 9,
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": 10,
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []int{10, 2},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []string{"1", "5"},
			},
			out: true,
		},
		{
			rawXql: `a !in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []string{"1", "5"},
			},
			out: false,
		},
		{
			rawXql: `a !in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": "1",
			},
			out: false,
		},
		{
			rawXql: `a !in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []string{"1", "5", "3"},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []int64{1, 5},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10.000,   5.000)`,
			data: map[string]interface{}{
				"a": []float64{2.000000000000001, 5.000000000000000002},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2,3, 10.00000000001)`,
			data: map[string]interface{}{
				"a": float64(10.0),
			},
			out: true,
		},
		{
			rawXql: `a !in (1,2,3, 10.00001)`,
			data: map[string]interface{}{
				"a": float64(10.0),
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10)`,
			data: map[string]interface{}{
				"a": []int64{2, 3},
			},
			out: false,
		},
		{
			rawXql: `a in (1,2,3, 10)`,
			data: map[string]interface{}{
				"a": int64(10),
			},
			out: true,
		},
		{
			rawXql: `a in (1,2,3, 10)`,
			data: map[string]interface{}{
				"a": []int64{2, 3},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10)`,
			data: map[string]interface{}{
				"a": []int64{2, 3},
			},
			out: false,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []float64{2, 5},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []float64{1, 10, 2, 5},
			},
			out: true,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []float64{1, 10, 2, 5, 3},
			},
			out: false,
		},
		{
			rawXql: `a in (1,2, 10,   5)`,
			data: map[string]interface{}{
				"a": []string{"1", "5"},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, _ := Match(tc.rawXql, tc.data)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_And(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10 and b>'2' and c<9 and d!=2`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
				"c": int64(-1),
				"d": int64(2),
			},
			out: false,
		},
		{
			rawXql: `a=10 and b>'2' and c<9 and d!=2`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
				"c": int64(-1),
			},
			out: false,
		},
		{
			rawXql: `a=10 and b>'2'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
			},
			out: true,
		},
		{
			rawXql: `a=10 and b>'2'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(2),
			},
			out: false,
		},
		{
			rawXql: `a=10 and b>'2'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
			},
			out: true,
		},
		{
			rawXql: `a=10 and b>'2' and c<9`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
				"c": int64(-1),
			},
			out: true,
		},
		{
			rawXql: `a=10 and b>'2' and c<9 and d!=2`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
				"c": int64(-1),
				"d": int64(0),
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, _ := Match(tc.rawXql, tc.data)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_Or(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=10 or b>'2'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
			},
			out: true,
		},
		{
			rawXql: `a=10 or b>'2'`,
			data: map[string]interface{}{
				"a": int64(9),
				"b": int64(2),
			},
			out: false,
		},
		{
			rawXql: `a=10 or b>'2'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(3),
			},
			out: true,
		},
		{
			rawXql: `a=10 or b>'2' or c<9`,
			data: map[string]interface{}{
				"a": int64(1),
				"b": int64(3),
				"c": int64(100),
			},
			out: true,
		},
		{
			rawXql: `a=10 or b>'2' or c<9 or d!=2`,
			data: map[string]interface{}{
				"a": int64(1),
				"b": int64(2),
				"c": int64(10),
				"d": int64(0),
			},
			out: true,
		},
		{
			rawXql: `a=10 or b>'2' or c<9 or d!=2`,
			data: map[string]interface{}{
				"a": int64(1),
				"b": int64(1),
				"c": int64(10),
				"d": int64(2),
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_Or_And(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `a=9 or c=1 and b!='1'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
				"c": int64(1),
			},
			out: false,
		},
		{
			rawXql: `a=10 and b>'2' or c=1`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
				"c": int64(1),
			},
			out: true,
		},
		{
			rawXql: `a=10 or c=1 and b!='1'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
				"c": int64(1),
			},
			out: true,
		},
		{
			rawXql: `a=10 and (c=1 or b!='1')`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
				"c": int64(1),
			},
			out: true,
		},
		{
			rawXql: `a=10 and (c=1 or b!='1') and d='123'`,
			data: map[string]interface{}{
				"a": int64(10),
				"b": int64(1),
				"c": int64(1),
				"d": "123",
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch_Inter(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `letter ∩ ('a','b','c','d','e')`,
			data: map[string]interface{}{
				"letter": []string{"a", "e"},
			},
			out: true,
		},
		{
			rawXql: `letter !∩ (1,2,3)`,
			data: map[string]interface{}{
				"letter": []float64{0.5, 3.01},
			},
			out: true,
		},
		{
			rawXql: `letter ∩ ('a', 'b','c','d',  'e')`,
			data: map[string]interface{}{
				"letter": []string{"a", "e", "f"},
			},
			out: true,
		},
		{
			rawXql: `letter ∩ ('a','b','c', 'd',  'e')`,
			data: map[string]interface{}{
				"letter": []string{"f"},
			},
			out: false,
		},
		{
			rawXql: `letter ∩ ('a','b','c','d')`,
			data: map[string]interface{}{
				"letter": "c",
			},
			out: true,
		},
		{
			rawXql: `letter ∩ (1,2,3)`,
			data: map[string]interface{}{
				"letter": []float64{2.0, 3.0},
			},
			out: true,
		},
		{
			rawXql: `letter ∩ (1,2,3)`,
			data: map[string]interface{}{
				"letter": []int64{2, 5},
			},
			out: true,
		},
		{
			rawXql: `letter !∩ (1,2,3)`,
			data: map[string]interface{}{
				"letter": []int64{4, 5},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestMatch(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other') and score>=95 and rank !in ('b','c','d')`,
			data: map[string]interface{}{
				"age":   int64(24),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "s",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other')`,
			data: map[string]interface{}{
				"age": int64(24),
				"sex": "other",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other')`,
			data: map[string]interface{}{
				"age": int64(24),
				"sex": "boy",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or some!=5) and words='hello world'`,
			data: map[string]interface{}{
				"age":   int64(211),
				"sex":   "boy",
				"some":  int64(6),
				"words": "hello world",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or some!=5) and words='hello world'`,
			data: map[string]interface{}{
				"age":   int64(21),
				"sex":   "boy",
				"some":  int64(6),
				"words": "hello world",
			},
			out: false,
		},
		{
			rawXql: `tag in (1,3,5) and status!=0`,
			data: map[string]interface{}{
				"tag":    []int64{1, 5},
				"status": int64(3),
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ok, err := Match(tc.rawXql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestRule_Match(t *testing.T) {
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other') and score>=95 and rank !in ('b','c','d')`,
			data: map[string]interface{}{
				"age":   int64(24),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "s",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other')`,
			data: map[string]interface{}{
				"age": int64(24),
				"sex": "other",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or sex='other')`,
			data: map[string]interface{}{
				"age": int64(24),
				"sex": "boy",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or some!=5) and words='hello world'`,
			data: map[string]interface{}{
				"age":   int64(211),
				"sex":   "boy",
				"some":  int64(6),
				"words": "hello world",
			},
			out: true,
		},
		{
			rawXql: `age>23 and (sex in ('boy','girl') or some!=5) and words='hello world'`,
			data: map[string]interface{}{
				"age":   int64(21),
				"sex":   "boy",
				"some":  int64(6),
				"words": "hello world",
			},
			out: false,
		},
		{
			rawXql: `tag in (1,3,5) and status!=0`,
			data: map[string]interface{}{
				"tag":    []int64{1, 5},
				"status": int64(3),
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ruler, err := Rule(tc.rawXql)
		ass.NoError(err)
		ok, err := ruler.Match(tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "rawXql=%s||data=%+v", tc.rawXql, tc.data)
	}
}

func TestRule_Match_Multi(t *testing.T) {
	var testData = []struct {
		data map[string]interface{}
		out  bool
	}{
		{
			data: map[string]interface{}{
				"age":   int64(24),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "s",
			},
			out: true,
		},
		{
			data: map[string]interface{}{
				"age":   int64(23),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "s",
			},
			out: false,
		},
		{
			data: map[string]interface{}{
				"age":   int64(24),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "",
			},
			out: true,
		},
		{
			data: map[string]interface{}{
				"age":   int64(24),
				"sex":   "boy",
				"score": int64(95),
				"rank":  "a",
			},
			out: true,
		},
		{
			data: map[string]interface{}{
				"age":   int64(23),
				"sex":   "boy",
				"score": int64(92),
				"rank":  "s",
			},
			out: false,
		},
		{
			data: map[string]interface{}{
				"age":   int64(23),
				"sex":   "other",
				"score": int64(97),
				"rank":  "s",
			},
			out: false,
		},
	}
	ass := assert.New(t)
	ruler, err := Rule(`age>23 and (sex in ('boy','girl') or sex='other') and score>=95 and rank !in ('b','c','d')`)
	if !ass.NoError(err) {
		t.FailNow()
	}
	for _, tc := range testData {
		ok, err := ruler.Match(tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, ok, "data=%+v", tc.data)
	}
}

func Test_Compare_Slice_And_One_Element(t *testing.T) {
	should := require.New(t)
	var testData = []struct {
		rawXql string
		data   map[string]interface{}
		out    bool
	}{
		{
			rawXql: `letter ∩ ('a')`,
			data: map[string]interface{}{
				"letter": []string{"a", "b", "c"},
			},
			out: true,
		},
		{
			rawXql: `letter ∩ ('a', 'b')`,
			data: map[string]interface{}{
				"letter": []string{"a", "b", "c"},
			},
			out: true,
		},
		{
			rawXql: `letter ∩ ('d')`,
			data: map[string]interface{}{
				"letter": []string{"a", "b", "c"},
			},
			out: false,
		},
		{
			rawXql: `letter in ('a')`,
			data: map[string]interface{}{
				"letter": []string{"a", "b", "c"},
			},
			out: false,
		},
	}
	for _, tc := range testData {
		actual, err := Match(tc.rawXql, tc.data)
		should.NoError(err)
		should.Equal(tc.out, actual)
	}
}
