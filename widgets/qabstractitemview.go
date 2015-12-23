package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractItemView struct {
	QAbstractScrollArea
}

type QAbstractItemView_ITF interface {
	QAbstractScrollArea_ITF
	QAbstractItemView_PTR() *QAbstractItemView
}

func PointerFromQAbstractItemView(ptr QAbstractItemView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemView_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemViewFromPointer(ptr unsafe.Pointer) *QAbstractItemView {
	var n = new(QAbstractItemView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractItemView_") {
		n.SetObjectName("QAbstractItemView_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractItemView) QAbstractItemView_PTR() *QAbstractItemView {
	return ptr
}

//QAbstractItemView::CursorAction
type QAbstractItemView__CursorAction int64

const (
	QAbstractItemView__MoveUp       = QAbstractItemView__CursorAction(0)
	QAbstractItemView__MoveDown     = QAbstractItemView__CursorAction(1)
	QAbstractItemView__MoveLeft     = QAbstractItemView__CursorAction(2)
	QAbstractItemView__MoveRight    = QAbstractItemView__CursorAction(3)
	QAbstractItemView__MoveHome     = QAbstractItemView__CursorAction(4)
	QAbstractItemView__MoveEnd      = QAbstractItemView__CursorAction(5)
	QAbstractItemView__MovePageUp   = QAbstractItemView__CursorAction(6)
	QAbstractItemView__MovePageDown = QAbstractItemView__CursorAction(7)
	QAbstractItemView__MoveNext     = QAbstractItemView__CursorAction(8)
	QAbstractItemView__MovePrevious = QAbstractItemView__CursorAction(9)
)

//QAbstractItemView::DragDropMode
type QAbstractItemView__DragDropMode int64

const (
	QAbstractItemView__NoDragDrop   = QAbstractItemView__DragDropMode(0)
	QAbstractItemView__DragOnly     = QAbstractItemView__DragDropMode(1)
	QAbstractItemView__DropOnly     = QAbstractItemView__DragDropMode(2)
	QAbstractItemView__DragDrop     = QAbstractItemView__DragDropMode(3)
	QAbstractItemView__InternalMove = QAbstractItemView__DragDropMode(4)
)

//QAbstractItemView::DropIndicatorPosition
type QAbstractItemView__DropIndicatorPosition int64

const (
	QAbstractItemView__OnItem     = QAbstractItemView__DropIndicatorPosition(0)
	QAbstractItemView__AboveItem  = QAbstractItemView__DropIndicatorPosition(1)
	QAbstractItemView__BelowItem  = QAbstractItemView__DropIndicatorPosition(2)
	QAbstractItemView__OnViewport = QAbstractItemView__DropIndicatorPosition(3)
)

//QAbstractItemView::EditTrigger
type QAbstractItemView__EditTrigger int64

const (
	QAbstractItemView__NoEditTriggers  = QAbstractItemView__EditTrigger(0)
	QAbstractItemView__CurrentChanged  = QAbstractItemView__EditTrigger(1)
	QAbstractItemView__DoubleClicked   = QAbstractItemView__EditTrigger(2)
	QAbstractItemView__SelectedClicked = QAbstractItemView__EditTrigger(4)
	QAbstractItemView__EditKeyPressed  = QAbstractItemView__EditTrigger(8)
	QAbstractItemView__AnyKeyPressed   = QAbstractItemView__EditTrigger(16)
	QAbstractItemView__AllEditTriggers = QAbstractItemView__EditTrigger(31)
)

//QAbstractItemView::ScrollHint
type QAbstractItemView__ScrollHint int64

const (
	QAbstractItemView__EnsureVisible    = QAbstractItemView__ScrollHint(0)
	QAbstractItemView__PositionAtTop    = QAbstractItemView__ScrollHint(1)
	QAbstractItemView__PositionAtBottom = QAbstractItemView__ScrollHint(2)
	QAbstractItemView__PositionAtCenter = QAbstractItemView__ScrollHint(3)
)

//QAbstractItemView::ScrollMode
type QAbstractItemView__ScrollMode int64

const (
	QAbstractItemView__ScrollPerItem  = QAbstractItemView__ScrollMode(0)
	QAbstractItemView__ScrollPerPixel = QAbstractItemView__ScrollMode(1)
)

//QAbstractItemView::SelectionBehavior
type QAbstractItemView__SelectionBehavior int64

const (
	QAbstractItemView__SelectItems   = QAbstractItemView__SelectionBehavior(0)
	QAbstractItemView__SelectRows    = QAbstractItemView__SelectionBehavior(1)
	QAbstractItemView__SelectColumns = QAbstractItemView__SelectionBehavior(2)
)

//QAbstractItemView::SelectionMode
type QAbstractItemView__SelectionMode int64

const (
	QAbstractItemView__NoSelection         = QAbstractItemView__SelectionMode(0)
	QAbstractItemView__SingleSelection     = QAbstractItemView__SelectionMode(1)
	QAbstractItemView__MultiSelection      = QAbstractItemView__SelectionMode(2)
	QAbstractItemView__ExtendedSelection   = QAbstractItemView__SelectionMode(3)
	QAbstractItemView__ContiguousSelection = QAbstractItemView__SelectionMode(4)
)

//QAbstractItemView::State
type QAbstractItemView__State int64

const (
	QAbstractItemView__NoState            = QAbstractItemView__State(0)
	QAbstractItemView__DraggingState      = QAbstractItemView__State(1)
	QAbstractItemView__DragSelectingState = QAbstractItemView__State(2)
	QAbstractItemView__EditingState       = QAbstractItemView__State(3)
	QAbstractItemView__ExpandingState     = QAbstractItemView__State(4)
	QAbstractItemView__CollapsingState    = QAbstractItemView__State(5)
	QAbstractItemView__AnimatingState     = QAbstractItemView__State(6)
)

func (ptr *QAbstractItemView) AlternatingRowColors() bool {
	defer qt.Recovering("QAbstractItemView::alternatingRowColors")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_AlternatingRowColors(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) AutoScrollMargin() int {
	defer qt.Recovering("QAbstractItemView::autoScrollMargin")

	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_AutoScrollMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DefaultDropAction() core.Qt__DropAction {
	defer qt.Recovering("QAbstractItemView::defaultDropAction")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QAbstractItemView_DefaultDropAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropMode() QAbstractItemView__DragDropMode {
	defer qt.Recovering("QAbstractItemView::dragDropMode")

	if ptr.Pointer() != nil {
		return QAbstractItemView__DragDropMode(C.QAbstractItemView_DragDropMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropOverwriteMode() bool {
	defer qt.Recovering("QAbstractItemView::dragDropOverwriteMode")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragDropOverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) DragEnabled() bool {
	defer qt.Recovering("QAbstractItemView::dragEnabled")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractItemView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractItemViewDragLeaveEvent
func callbackQAbstractItemViewDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) EditTriggers() QAbstractItemView__EditTrigger {
	defer qt.Recovering("QAbstractItemView::editTriggers")

	if ptr.Pointer() != nil {
		return QAbstractItemView__EditTrigger(C.QAbstractItemView_EditTriggers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) HasAutoScroll() bool {
	defer qt.Recovering("QAbstractItemView::hasAutoScroll")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_HasAutoScroll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) HorizontalScrollMode() QAbstractItemView__ScrollMode {
	defer qt.Recovering("QAbstractItemView::horizontalScrollMode")

	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_HorizontalScrollMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) IconSize() *core.QSize {
	defer qt.Recovering("QAbstractItemView::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractItemView_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	defer qt.Recovering("QAbstractItemView::selectionBehavior")

	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionBehavior(C.QAbstractItemView_SelectionBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	defer qt.Recovering("QAbstractItemView::selectionMode")

	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionMode(C.QAbstractItemView_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	defer qt.Recovering("QAbstractItemView::setAlternatingRowColors")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAlternatingRowColors(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScroll(enable bool) {
	defer qt.Recovering("QAbstractItemView::setAutoScroll")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScroll(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScrollMargin(margin int) {
	defer qt.Recovering("QAbstractItemView::setAutoScrollMargin")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScrollMargin(ptr.Pointer(), C.int(margin))
	}
}

func (ptr *QAbstractItemView) SetDefaultDropAction(dropAction core.Qt__DropAction) {
	defer qt.Recovering("QAbstractItemView::setDefaultDropAction")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDefaultDropAction(ptr.Pointer(), C.int(dropAction))
	}
}

func (ptr *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView__DragDropMode) {
	defer qt.Recovering("QAbstractItemView::setDragDropMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropMode(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	defer qt.Recovering("QAbstractItemView::setDragDropOverwriteMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QAbstractItemView) SetDragEnabled(enable bool) {
	defer qt.Recovering("QAbstractItemView::setDragEnabled")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	defer qt.Recovering("QAbstractItemView::setDropIndicatorShown")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDropIndicatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView__EditTrigger) {
	defer qt.Recovering("QAbstractItemView::setEditTriggers")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetEditTriggers(ptr.Pointer(), C.int(triggers))
	}
}

func (ptr *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView__ScrollMode) {
	defer qt.Recovering("QAbstractItemView::setHorizontalScrollMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetHorizontalScrollMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetIconSize(size core.QSize_ITF) {
	defer qt.Recovering("QAbstractItemView::setIconSize")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	defer qt.Recovering("QAbstractItemView::setSelectionBehavior")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionBehavior(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	defer qt.Recovering("QAbstractItemView::setSelectionMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	defer qt.Recovering("QAbstractItemView::setTabKeyNavigation")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTabKeyNavigation(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetTextElideMode(mode core.Qt__TextElideMode) {
	defer qt.Recovering("QAbstractItemView::setTextElideMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTextElideMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView__ScrollMode) {
	defer qt.Recovering("QAbstractItemView::setVerticalScrollMode")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetVerticalScrollMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) ShowDropIndicator() bool {
	defer qt.Recovering("QAbstractItemView::showDropIndicator")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_ShowDropIndicator(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TabKeyNavigation() bool {
	defer qt.Recovering("QAbstractItemView::tabKeyNavigation")

	if ptr.Pointer() != nil {
		return C.QAbstractItemView_TabKeyNavigation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TextElideMode() core.Qt__TextElideMode {
	defer qt.Recovering("QAbstractItemView::textElideMode")

	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QAbstractItemView_TextElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) VerticalScrollMode() QAbstractItemView__ScrollMode {
	defer qt.Recovering("QAbstractItemView::verticalScrollMode")

	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_VerticalScrollMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) ConnectActivated(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::activated")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QAbstractItemView) DisconnectActivated() {
	defer qt.Recovering("disconnect QAbstractItemView::activated")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQAbstractItemViewActivated
func callbackQAbstractItemViewActivated(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ClearSelection() {
	defer qt.Recovering("QAbstractItemView::clearSelection")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ConnectClicked(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::clicked")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectClicked() {
	defer qt.Recovering("disconnect QAbstractItemView::clicked")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractItemViewClicked
func callbackQAbstractItemViewClicked(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QAbstractItemView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QAbstractItemView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QAbstractItemView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQAbstractItemViewCloseEditor
func callbackQAbstractItemViewCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ClosePersistentEditor(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemView::closePersistentEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClosePersistentEditor(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QAbstractItemView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QAbstractItemView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QAbstractItemView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQAbstractItemViewCommitData
func callbackQAbstractItemViewCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QAbstractItemView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QAbstractItemView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQAbstractItemViewCurrentChanged
func callbackQAbstractItemViewCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) CurrentIndex() *core.QModelIndex {
	defer qt.Recovering("QAbstractItemView::currentIndex")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectDoubleClicked(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::doubleClicked")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "doubleClicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDoubleClicked() {
	defer qt.Recovering("disconnect QAbstractItemView::doubleClicked")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "doubleClicked")
	}
}

//export callbackQAbstractItemViewDoubleClicked
func callbackQAbstractItemViewDoubleClicked(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::doubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "doubleClicked"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractItemView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractItemViewDragEnterEvent
func callbackQAbstractItemViewDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractItemView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractItemViewDragMoveEvent
func callbackQAbstractItemViewDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractItemView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractItemViewDropEvent
func callbackQAbstractItemViewDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) Edit(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemView::edit")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QAbstractItemView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QAbstractItemView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QAbstractItemView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQAbstractItemViewEditorDestroyed
func callbackQAbstractItemViewEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectEntered(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::entered")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectEntered() {
	defer qt.Recovering("disconnect QAbstractItemView::entered")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractItemViewEntered
func callbackQAbstractItemViewEntered(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::entered")

	if signal := qt.GetSignal(C.GoString(ptrName), "entered"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractItemView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQAbstractItemViewFocusInEvent
func callbackQAbstractItemViewFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractItemView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQAbstractItemViewFocusOutEvent
func callbackQAbstractItemViewFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectIconSizeChanged(f func(size *core.QSize)) {
	defer qt.Recovering("connect QAbstractItemView::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectIconSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconSizeChanged", f)
	}
}

func (ptr *QAbstractItemView) DisconnectIconSizeChanged() {
	defer qt.Recovering("disconnect QAbstractItemView::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectIconSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconSizeChanged")
	}
}

//export callbackQAbstractItemViewIconSizeChanged
func callbackQAbstractItemViewIconSizeChanged(ptrName *C.char, size unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::iconSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

func (ptr *QAbstractItemView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QAbstractItemView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QAbstractItemView) IndexWidget(index core.QModelIndex_ITF) *QWidget {
	defer qt.Recovering("QAbstractItemView::indexWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractItemView_IndexWidget(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractItemView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQAbstractItemViewInputMethodEvent
func callbackQAbstractItemViewInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QAbstractItemView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractItemView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	defer qt.Recovering("QAbstractItemView::itemDelegate")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegate2(index core.QModelIndex_ITF) *QAbstractItemDelegate {
	defer qt.Recovering("QAbstractItemView::itemDelegate")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegate2(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	defer qt.Recovering("QAbstractItemView::itemDelegateForColumn")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegateForColumn(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	defer qt.Recovering("QAbstractItemView::itemDelegateForRow")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegateForRow(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractItemView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractItemViewKeyPressEvent
func callbackQAbstractItemViewKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QAbstractItemView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QAbstractItemView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QAbstractItemView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQAbstractItemViewKeyboardSearch
func callbackQAbstractItemViewKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QAbstractItemView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) Model() *core.QAbstractItemModel {
	defer qt.Recovering("QAbstractItemView::model")

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QAbstractItemView_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractItemView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractItemViewMouseDoubleClickEvent
func callbackQAbstractItemViewMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractItemView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractItemViewMouseMoveEvent
func callbackQAbstractItemViewMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractItemView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractItemViewMousePressEvent
func callbackQAbstractItemViewMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractItemView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractItemViewMouseReleaseEvent
func callbackQAbstractItemViewMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) OpenPersistentEditor(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemView::openPersistentEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_OpenPersistentEditor(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectPressed(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::pressed")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractItemView) DisconnectPressed() {
	defer qt.Recovering("disconnect QAbstractItemView::pressed")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractItemViewPressed
func callbackQAbstractItemViewPressed(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemView::pressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "pressed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectReset(f func()) {
	defer qt.Recovering("connect QAbstractItemView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QAbstractItemView) DisconnectReset() {
	defer qt.Recovering("disconnect QAbstractItemView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQAbstractItemViewReset
func callbackQAbstractItemViewReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractItemView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractItemView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractItemViewResizeEvent
func callbackQAbstractItemViewResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) RootIndex() *core.QModelIndex {
	defer qt.Recovering("QAbstractItemView::rootIndex")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_RootIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QAbstractItemView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QAbstractItemView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QAbstractItemView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQAbstractItemViewRowsAboutToBeRemoved
func callbackQAbstractItemViewRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QAbstractItemView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QAbstractItemView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QAbstractItemView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQAbstractItemViewRowsInserted
func callbackQAbstractItemViewRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QAbstractItemView::scrollTo")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QAbstractItemView) ScrollToBottom() {
	defer qt.Recovering("QAbstractItemView::scrollToBottom")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ScrollToTop() {
	defer qt.Recovering("QAbstractItemView::scrollToTop")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QAbstractItemView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QAbstractItemView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQAbstractItemViewSelectAll
func callbackQAbstractItemViewSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractItemView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractItemView) SelectionModel() *core.QItemSelectionModel {
	defer qt.Recovering("QAbstractItemView::selectionModel")

	if ptr.Pointer() != nil {
		return core.NewQItemSelectionModelFromPointer(C.QAbstractItemView_SelectionModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) SetCurrentIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemView::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) SetIndexWidget(index core.QModelIndex_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QAbstractItemView::setIndexWidget")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIndexWidget(ptr.Pointer(), core.PointerFromQModelIndex(index), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractItemView) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QAbstractItemView::setItemDelegate")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForColumn(column int, delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QAbstractItemView::setItemDelegateForColumn")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForColumn(ptr.Pointer(), C.int(column), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForRow(row int, delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QAbstractItemView::setItemDelegateForRow")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForRow(ptr.Pointer(), C.int(row), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QAbstractItemView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QAbstractItemView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQAbstractItemViewSetModel
func callbackQAbstractItemViewSetModel(ptrName *C.char, model unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QAbstractItemView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQAbstractItemViewSetRootIndex
func callbackQAbstractItemViewSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QAbstractItemView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QAbstractItemView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQAbstractItemViewSetSelectionModel
func callbackQAbstractItemViewSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) SizeHintForColumn(column int) int {
	defer qt.Recovering("QAbstractItemView::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForColumn(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QAbstractItemView) SizeHintForIndex(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QAbstractItemView::sizeHintForIndex")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractItemView_SizeHintForIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) SizeHintForRow(row int) int {
	defer qt.Recovering("QAbstractItemView::sizeHintForRow")

	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForRow(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QAbstractItemView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QAbstractItemView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QAbstractItemView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QAbstractItemView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQAbstractItemViewStartDrag
func callbackQAbstractItemViewStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractItemView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractItemViewTimerEvent
func callbackQAbstractItemViewTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) Update(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemView::update")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QAbstractItemView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QAbstractItemView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QAbstractItemView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQAbstractItemViewUpdateGeometries
func callbackQAbstractItemViewUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractItemView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectViewportEntered(f func()) {
	defer qt.Recovering("connect QAbstractItemView::viewportEntered")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectViewportEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "viewportEntered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectViewportEntered() {
	defer qt.Recovering("disconnect QAbstractItemView::viewportEntered")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectViewportEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "viewportEntered")
	}
}

//export callbackQAbstractItemViewViewportEntered
func callbackQAbstractItemViewViewportEntered(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractItemView::viewportEntered")

	if signal := qt.GetSignal(C.GoString(ptrName), "viewportEntered"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractItemView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QAbstractItemView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QAbstractItemView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) DestroyQAbstractItemView() {
	defer qt.Recovering("QAbstractItemView::~QAbstractItemView")

	if ptr.Pointer() != nil {
		C.QAbstractItemView_DestroyQAbstractItemView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractItemView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractItemViewPaintEvent
func callbackQAbstractItemViewPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractItemView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractItemViewContextMenuEvent
func callbackQAbstractItemViewContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QAbstractItemView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QAbstractItemView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QAbstractItemView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQAbstractItemViewScrollContentsBy
func callbackQAbstractItemViewScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QAbstractItemView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QAbstractItemView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQAbstractItemViewSetupViewport
func callbackQAbstractItemViewSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractItemView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractItemViewWheelEvent
func callbackQAbstractItemViewWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QAbstractItemView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQAbstractItemViewChangeEvent
func callbackQAbstractItemViewChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractItemView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQAbstractItemViewActionEvent
func callbackQAbstractItemViewActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractItemView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQAbstractItemViewEnterEvent
func callbackQAbstractItemViewEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractItemView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQAbstractItemViewHideEvent
func callbackQAbstractItemViewHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractItemView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQAbstractItemViewLeaveEvent
func callbackQAbstractItemViewLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractItemView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQAbstractItemViewMoveEvent
func callbackQAbstractItemViewMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractItemView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QAbstractItemView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractItemView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQAbstractItemViewSetVisible
func callbackQAbstractItemViewSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractItemView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractItemView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQAbstractItemViewShowEvent
func callbackQAbstractItemViewShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractItemView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQAbstractItemViewCloseEvent
func callbackQAbstractItemViewCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QAbstractItemView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QAbstractItemView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QAbstractItemView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQAbstractItemViewInitPainter
func callbackQAbstractItemViewInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractItemView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQAbstractItemViewKeyReleaseEvent
func callbackQAbstractItemViewKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractItemView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQAbstractItemViewTabletEvent
func callbackQAbstractItemViewTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractItemView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractItemViewChildEvent
func callbackQAbstractItemViewChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractItemView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractItemView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractItemView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractItemView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractItemViewCustomEvent
func callbackQAbstractItemViewCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
