// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/ListView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type ListView struct {
	View
}

func NewListView(id string) interface{} {
	obj := ListView{NewView(id).(View)}

	return obj
}


func init() {
	android.ViewTypeConstructors["ListView"] = NewListView
}

func (obj ListView) AreFooterDividersEnabled() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"areFooterDividersEnabled",
	)
}

func (obj ListView) AreHeaderDividersEnabled() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"areHeaderDividersEnabled",
	)
}

func (obj ListView) GetDividerHeight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"getDividerHeight",
	)
}

func (obj ListView) GetFooterViewsCount() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"getFooterViewsCount",
	)
}

func (obj ListView) GetHeaderViewsCount() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"getHeaderViewsCount",
	)
}

func (obj ListView) GetItemsCanFocus() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"getItemsCanFocus",
	)
}

func (obj ListView) GetMaxScrollAmount() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"getMaxScrollAmount",
	)
}

func (obj ListView) IsOpaque() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"isOpaque",
	)
}

func (obj ListView) SetCacheColorHint(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setCacheColorHint",
		color_,
	)
}

func (obj ListView) SetDividerHeight(height_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setDividerHeight",
		height_,
	)
}

func (obj ListView) SetFooterDividersEnabled(footerDividersEnabled_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setFooterDividersEnabled",
		footerDividersEnabled_,
	)
}

func (obj ListView) SetHeaderDividersEnabled(headerDividersEnabled_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setHeaderDividersEnabled",
		headerDividersEnabled_,
	)
}

func (obj ListView) SetItemsCanFocus(itemsCanFocus_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setItemsCanFocus",
		itemsCanFocus_,
	)
}

func (obj ListView) SetSelection(position_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setSelection",
		position_,
	)
}

func (obj ListView) SetSelectionAfterHeaderView() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"setSelectionAfterHeaderView",
	)
}

func (obj ListView) SmoothScrollByOffset(offset_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"smoothScrollByOffset",
		offset_,
	)
}

func (obj ListView) SmoothScrollToPosition(position_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ListView",
		"smoothScrollToPosition",
		position_,
	)
}
