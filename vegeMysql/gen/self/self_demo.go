package self

import (
	"github.com/Godyu97/vege9/vegeTools"
)

type Subscriber struct {
	Region   *string //[]int64
	Keywords *string //[]string
}

func (r *Subscriber) RegionValue() ([]int64, error) {
	s := make([]int64, 0)
	if r.Region == nil {
		return s, nil
	}
	err := vegeTools.JsonUnmarshalFromString(*r.Region, &s)
	return s, err
}

func (r *Subscriber) RegionScan(in []int64) error {
	if in == nil {
		in = make([]int64, 0, 0)
	}
	js, err := vegeTools.JsonMarshalToString(in)
	if err != nil {
		return err
	}
	r.Region = &js
	return nil
}

func (r *Subscriber) KeywordsValue() ([]string, error) {
	s := make([]string, 0)
	if r.Keywords == nil {
		return s, nil
	}
	err := vegeTools.JsonUnmarshalFromString(*r.Keywords, &s)
	return s, err
}

func (r *Subscriber) KeywordsScan(in []string) error {
	if in == nil {
		in = make([]string, 0, 0)
	}
	js, err := vegeTools.JsonMarshalToString(in)
	if err != nil {
		return err
	}
	r.Keywords = &js
	return nil
}
