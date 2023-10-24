package xql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_Int(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []int{1, 2, 3, 4},
			},
			out: false,
		},
		{
			rxql: `a.sum() >= 10`,
			data: map[string]interface{}{
				"a": []int{1, 2, 3, 4},
			},
			out: true,
		},
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []int{2, 2, 3, 4},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestSum_Int64(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []int64{1, 2, 3, 4},
			},
			out: false,
		},
		{
			rxql: `a.sum() >= 10`,
			data: map[string]interface{}{
				"a": []int64{1, 2, 3, 4},
			},
			out: true,
		},
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []int64{2, 2, 3, 4},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestSum_Float64(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.sum() >= 10`,
			data: map[string]interface{}{
				"a": []float64{1.0, 2.0, 3.0, 4.0},
			},
			out: true,
		},
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []float64{1.0, 2.0, 3.0, 4.0},
			},
			out: false,
		},
		{
			rxql: `a.sum() > 10`,
			data: map[string]interface{}{
				"a": []float64{2.0, 2.0, 3.0, 4.0},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestSum(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.sum() > 10 and foo=true`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": true,
			},
			out: true,
		},
		{
			rxql: `a.sum() > 10 and foo=false`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": true,
			},
			out: false,
		},
		{
			rxql: `a.sum() > 10 or foo=true`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": false,
			},
			out: true,
		},
		{
			rxql: `a.sum() > 10 and a.sum() < 12 or foo=true`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": false,
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestCount(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.count()=4`,
			data: map[string]interface{}{
				"a": []int{2, 2, 3, 4},
			},
			out: true,
		},
		{
			rxql: `a.count()=4`,
			data: map[string]interface{}{
				"a": []int64{2, 2, 3, 4},
			},
			out: true,
		},
		{
			rxql: `a.count()=4`,
			data: map[string]interface{}{
				"a": []float64{2, 2, 3, 4},
			},
			out: true,
		},
		{
			rxql: `a.count()=4`,
			data: map[string]interface{}{
				"a": []string{"1", "2", "3", "4"},
			},
			out: true,
		},
		{
			rxql: `a.count()=1`,
			data: map[string]interface{}{
				"a": "hello",
			},
			out: true,
		},
		{
			rxql: `a.count() > 3 and foo=false`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": true,
			},
			out: false,
		},
		{
			rxql: `a.count() > 3 or foo=true`,
			data: map[string]interface{}{
				"a":   []int{2, 2, 3, 4},
				"foo": false,
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestAvg(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.avg()>3`,
			data: map[string]interface{}{
				"a": []float64{4, 4},
			},
			out: true,
		},
		{
			rxql: `a.avg()=4.5`,
			data: map[string]interface{}{
				"a": []float64{4, 5},
			},
			out: true,
		},
		{
			rxql: `a.avg()>3.99999`,
			data: map[string]interface{}{
				"a": []float64{4, 4},
			},
			out: true,
		},
		{
			rxql: `a.avg()>4`,
			data: map[string]interface{}{
				"a": []int64{4, 4},
			},
			out: false,
		},
		{
			rxql: `a.avg()>4`,
			data: map[string]interface{}{
				"a": []int{4, 5},
			},
			out: true,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestMax(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.max()=3`,
			data: map[string]interface{}{
				"a": []float64{1, 4},
			},
			out: false,
		},
		{
			rxql: `a.max()=4.5`,
			data: map[string]interface{}{
				"a": []float64{4.5, 2},
			},
			out: true,
		},
		{
			rxql: `a.max()<10`,
			data: map[string]interface{}{
				"a": []int64{4, 2, 8, 9},
			},
			out: true,
		},
		{
			rxql: `a.max()<10`,
			data: map[string]interface{}{
				"a": []int{4, 2, 10, 8, 9},
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}

func TestMin(t *testing.T) {
	var testData = []struct {
		rxql string
		data map[string]interface{}
		out  bool
	}{
		{
			rxql: `a.min()=4`,
			data: map[string]interface{}{
				"a": []float64{1, 4},
			},
			out: false,
		},
		{
			rxql: `a.min()=4.5`,
			data: map[string]interface{}{
				"a": []float64{4.5, 2},
			},
			out: false,
		},
		{
			rxql: `a.min()=4.5`,
			data: map[string]interface{}{
				"a": []float64{4.5, 4.51},
			},
			out: true,
		},
		{
			rxql: `a.min()<10`,
			data: map[string]interface{}{
				"a": []int64{40, 20, 80, 9, 100},
			},
			out: true,
		},
		{
			rxql: `a.min()>10`,
			data: map[string]interface{}{
				"a": []int{4, 2, 10, 8, 9},
			},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		actual, err := Match(tc.rxql, tc.data)
		ass.NoError(err)
		ass.Equal(tc.out, actual, "rawXQL:%s||data:%+v", tc.rxql, tc.data)
	}
}
