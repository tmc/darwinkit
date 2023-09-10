// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource?language=objc
type PCNNGroupNormalizationDataSource interface {
	// optional
	UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer metal.CommandBufferObject, groupNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState
	HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch() bool

	// optional
	Beta() *float32
	HasBeta() bool

	// optional
	Epsilon() float64
	HasEpsilon() bool

	// optional
	Gamma() *float32
	HasGamma() bool

	// optional
	UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch *foundation.Array) bool
	HasUpdateGammaAndBetaWithGroupNormalizationStateBatch() bool

	// optional
	InitWithCoder(aDecoder foundation.Coder) objc.Object
	HasInitWithCoder() bool

	// optional
	Epsilon() float32
	HasEpsilon() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	SetNumberOfGroups(value uint)
	HasSetNumberOfGroups() bool

	// optional
	NumberOfGroups() uint
	HasNumberOfGroups() bool

	// optional
	NumberOfFeatureChannels() uint
	HasNumberOfFeatureChannels() bool
}

// ensure impl type implements protocol interface
var _ PCNNGroupNormalizationDataSource = (*CNNGroupNormalizationDataSourceObject)(nil)

// A concrete type for the [PCNNGroupNormalizationDataSource] protocol.
type CNNGroupNormalizationDataSourceObject struct {
	objc.Object
}

func (c_ CNNGroupNormalizationDataSourceObject) HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithCommandBuffer:groupNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152553-updategammaandbetawithcommandbuf?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer metal.CommandBufferObject, groupNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("updateGammaAndBetaWithCommandBuffer:groupNormalizationStateBatch:"), po0, groupNormalizationStateBatch)
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasBeta() bool {
	return c_.RespondsToSelector(objc.Sel("beta"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152543-beta?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) Beta() *float32 {
	rv := objc.Call[*float32](c_, objc.Sel("beta"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasEpsilon() bool {
	return c_.RespondsToSelector(objc.Sel("epsilon"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152546-epsilon?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152547-gamma?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) Gamma() *float32 {
	rv := objc.Call[*float32](c_, objc.Sel("gamma"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasUpdateGammaAndBetaWithGroupNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithGroupNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152554-updategammaandbetawithgroupnorma?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch *foundation.Array) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateGammaAndBetaWithGroupNormalizationStateBatch:"), groupNormalizationStateBatch)
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152548-initwithcoder?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) InitWithCoder(aDecoder foundation.Coder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasGamma() bool {
	return c_.RespondsToSelector(objc.Sel("gamma"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152546-epsilon?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) Epsilon() float32 {
	rv := objc.Call[float32](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152549-label?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152545-encodewithcoder?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) EncodeWithCoder(aCoder foundation.Coder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), aCoder)
}

func (c_ CNNGroupNormalizationDataSourceObject) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152550-numberoffeaturechannels?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasSetNumberOfGroups() bool {
	return c_.RespondsToSelector(objc.Sel("setNumberOfGroups:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152551-numberofgroups?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) SetNumberOfGroups(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfGroups:"), value)
}

func (c_ CNNGroupNormalizationDataSourceObject) HasNumberOfGroups() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfGroups"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152551-numberofgroups?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) NumberOfGroups() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfGroups"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceObject) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152550-numberoffeaturechannels?language=objc
func (c_ CNNGroupNormalizationDataSourceObject) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
