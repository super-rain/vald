//
// Copyright (C) 2019-2020 Vdaas.org Vald team ( kpango, rinx, kmrmt )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package rest provides rest api logic
package rest

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/vdaas/vald/apis/grpc/gateway/vald"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	type want struct {
		want Handler
	}
	type test struct {
		name       string
		args       args
		want       want
		checkFunc  func(want, Handler) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, got Handler) error {
		if !reflect.DeepEqual(got, w.want) {
			return errors.Errorf("got = %v, want %v", got, w.want)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           opts: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           opts: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}

			got := New(test.args.opts...)
			if err := test.checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Index(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		want int
		err  error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, got int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(got, w.want) {
			return errors.Errorf("got = %v, want %v", got, w.want)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			got, err := h.Index(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, got, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Search(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.Search(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_SearchByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.SearchByID(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Insert(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.Insert(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_MultiInsert(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.MultiInsert(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Update(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.Update(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_MultiUpdate(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.MultiUpdate(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Remove(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.Remove(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_MultiRemove(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.MultiRemove(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_GetObject(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.GetObject(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}

func Test_handler_Exists(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		vald vald.ValdServer
	}
	type want struct {
		wantCode int
		err      error
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int, error) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotCode int, err error) error {
		if !errors.Is(err, w.err) {
			return errors.Errorf("got error = %v, want %v", err, w.err)
		}
		if !reflect.DeepEqual(gotCode, w.wantCode) {
			return errors.Errorf("got = %v, want %v", gotCode, w.wantCode)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           w: nil,
		           r: nil,
		       },
		       fields: fields {
		           vald: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           args: args {
		           w: nil,
		           r: nil,
		           },
		           fields: fields {
		           vald: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			h := &handler{
				vald: test.fields.vald,
			}

			gotCode, err := h.Exists(test.args.w, test.args.r)
			if err := test.checkFunc(test.want, gotCode, err); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}
