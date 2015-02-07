// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/StackView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type StackView struct {
	View
}

func NewStackView(id string) interface{} {
	obj := StackView{NewView(id).(View)}

	return obj
}

func (obj StackView) GetClassName() string {
	return "android.widget.StackView"
}

func init() {
	android.ViewTypeConstructors["android.widget.StackView"] = NewStackView
}

func (obj StackView) Advance() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.StackView",
		"advance",
	))
}

func (obj StackView) ShowNext() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.StackView",
		"showNext",
	))
}

func (obj StackView) ShowPrevious() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.StackView",
		"showPrevious",
	))
}
