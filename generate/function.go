package generate

import (
	"fmt"
	"log"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

func (db *Generator) ToFunction(fw string, sym Symbol) *codegen.Function {
	// these functions have known declparse failures
	knownIssues := map[string]bool{
		"AUListenerAddParameter":                                     true,
		"AUParameterListenerNotify":                                  true,
		"AudioCodecGetProperty":                                      true,
		"AudioDeviceAddIOProc":                                       true,
		"AudioDeviceAddPropertyListener":                             true,
		"AudioDeviceCreateIOProcID":                                  true,
		"AudioDeviceGetProperty":                                     true,
		"AudioDeviceSetProperty":                                     true,
		"AudioDriverPlugInDeviceGetProperty":                         true,
		"AudioDriverPlugInDeviceSetProperty":                         true,
		"AudioDriverPlugInStreamGetProperty":                         true,
		"AudioDriverPlugInStreamSetProperty":                         true,
		"AudioFileComponentGetGlobalInfoSize":                        true,
		"AudioFileComponentOpenWithCallbacks":                        true,
		"AudioFileComponentReadPacketData":                           true,
		"AudioFileComponentReadPackets":                              true,
		"AudioFileComponentSetProperty":                              true,
		"AudioFileGetProperty":                                       true,
		"AudioHardwareAddPropertyListener":                           true,
		"AudioHardwareGetProperty":                                   true,
		"AudioHardwareSetProperty":                                   true,
		"AudioObjectAddPropertyListener":                             true,
		"AudioObjectGetPropertyData":                                 true,
		"AudioObjectGetPropertyDataSize":                             true,
		"AudioObjectRemovePropertyListener":                          true,
		"AudioObjectSetPropertyData":                                 true,
		"AudioQueueNewOutput":                                        true,
		"AudioQueueProcessingTapNew":                                 true,
		"AudioStreamAddPropertyListener":                             true,
		"AudioStreamGetProperty":                                     true,
		"AudioStreamSetProperty":                                     true,
		"AudioUnitAddPropertyListener":                               true,
		"AudioUnitAddRenderNotify":                                   true,
		"AudioUnitGetProperty":                                       true,
		"AudioUnitRemoveRenderNotify":                                true,
		"CAClockAddListener":                                         true,
		"CAClockRemoveListener":                                      true,
		"CAShow":                                                     true,
		"CAShowFile":                                                 true,
		"CFAllocatorDeallocate":                                      true,
		"CFAllocatorReallocate":                                      true,
		"CFArrayAppendValue":                                         true,
		"CFArrayApplyFunction":                                       true,
		"CFArrayBSearchValues":                                       true,
		"CFArrayContainsValue":                                       true,
		"CFArrayCreate":                                              true,
		"CFArrayGetCountOfValue":                                     true,
		"CFArrayGetFirstIndexOfValue":                                true,
		"CFArrayGetLastIndexOfValue":                                 true,
		"CFArrayGetValues":                                           true,
		"CFArrayInsertValueAtIndex":                                  true,
		"CFArrayReplaceValues":                                       true,
		"CFArraySetValueAtIndex":                                     true,
		"CFArraySortValues":                                          true,
		"CFBagAddValue":                                              true,
		"CFBagApplyFunction":                                         true,
		"CFBagContainsValue":                                         true,
		"CFBagCreate":                                                true,
		"CFBagGetCountOfValue":                                       true,
		"CFBagGetValue":                                              true,
		"CFBagGetValueIfPresent":                                     true,
		"CFBagGetValues":                                             true,
		"CFBagRemoveValue":                                           true,
		"CFBagReplaceValue":                                          true,
		"CFBagSetValue":                                              true,
		"CFBinaryHeapAddValue":                                       true,
		"CFBinaryHeapApplyFunction":                                  true,
		"CFBinaryHeapContainsValue":                                  true,
		"CFBinaryHeapGetCountOfValue":                                true,
		"CFBinaryHeapGetMinimumIfPresent":                            true,
		"CFBinaryHeapGetValues":                                      true,
		"CFBundleGetDataPointersForNames":                            true,
		"CFBundleGetFunctionPointersForNames":                        true,
		"CFDateCompare":                                              true,
		"CFDictionaryAddValue":                                       true,
		"CFDictionaryApplyFunction":                                  true,
		"CFDictionaryContainsKey":                                    true,
		"CFDictionaryContainsValue":                                  true,
		"CFDictionaryCreate":                                         true,
		"CFDictionaryGetCountOfKey":                                  true,
		"CFDictionaryGetCountOfValue":                                true,
		"CFDictionaryGetKeysAndValues":                               true,
		"CFDictionaryGetValue":                                       true,
		"CFDictionaryGetValueIfPresent":                              true,
		"CFDictionaryRemoveValue":                                    true,
		"CFDictionaryReplaceValue":                                   true,
		"CFDictionarySetValue":                                       true,
		"CFErrorCreateWithUserInfoKeysAndValues":                     true,
		"CFNotificationCenterAddObserver":                            true,
		"CFNotificationCenterPostNotification":                       true,
		"CFNotificationCenterRemoveEveryObserver":                    true,
		"CFNotificationCenterRemoveObserver":                         true,
		"CFNumberCompare":                                            true,
		"CFNumberCreate":                                             true,
		"CFNumberFormatterGetValueFromString":                        true,
		"CFNumberGetValue":                                           true,
		"CFPlugInInstanceGetInterfaceFunctionTable":                  true,
		"CFSetAddValue":                                              true,
		"CFSetApplyFunction":                                         true,
		"CFSetContainsValue":                                         true,
		"CFSetCreate":                                                true,
		"CFSetGetCountOfValue":                                       true,
		"CFSetGetValue":                                              true,
		"CFSetGetValueIfPresent":                                     true,
		"CFSetGetValues":                                             true,
		"CFSetRemoveValue":                                           true,
		"CFSetReplaceValue":                                          true,
		"CFSetSetValue":                                              true,
		"CFSocketCopyRegisteredValue":                                true,
		"CFTreeApplyFunctionToChildren":                              true,
		"CFTreeSortChildren":                                         true,
		"CFURLCopyResourcePropertyForKey":                            true,
		"CFXMLNodeCreate":                                            true,
		"CGBitmapContextCreate":                                      true,
		"CGPDFScannerPopName":                                        true,
		"CGBitmapContextCreateWithData":                              true,
		"CGColorSpaceCreateIndexed":                                  true,
		"CGColorSpaceCreateWithPlatformColorSpace":                   true,
		"CGConvertColorDataWithFormat":                               true,
		"CGDataConsumerCreate":                                       true,
		"CGDataProviderCreateDirect":                                 true,
		"CGDataProviderCreateSequential":                             true,
		"CGDataProviderCreateWithData":                               true,
		"CGDisplayRegisterReconfigurationCallback":                   true,
		"CGDisplayRemoveReconfigurationCallback":                     true,
		"CGEventPostToPSN":                                           true,
		"CGEventTapCreate":                                           true,
		"CGEventTapCreateForPSN":                                     true,
		"CGEventTapCreateForPid":                                     true,
		"CGEventTapPostEvent":                                        true,
		"CGFontCreateWithPlatformFont":                               true,
		"CGFunctionCreate":                                           true,
		"CGPDFArrayApplyBlock":                                       true,
		"CGPDFArrayGetName":                                          true,
		"CGPDFDictionaryApplyBlock":                                  true,
		"CGPDFDictionaryApplyFunction":                               true,
		"CGPDFDictionaryGetName":                                     true,
		"CGPDFObjectGetValue":                                        true,
		"CGPDFScannerCreate":                                         true,
		"CGPSConverterCreate":                                        true,
		"CGPathApply":                                                true,
		"CGPatternCreate":                                            true,
		"CGRegisterScreenRefreshCallback":                            true,
		"CGScreenRegisterMoveCallback":                               true,
		"CGScreenUnregisterMoveCallback":                             true,
		"CGUnregisterScreenRefreshCallback":                          true,
		"CMAudioFormatDescriptionCreate":                             true,
		"CMBlockBufferAccessDataBytes":                               true,
		"CMBlockBufferAppendMemoryBlock":                             true,
		"CMBlockBufferCopyDataBytes":                                 true,
		"CMBlockBufferCreateWithMemoryBlock":                         true,
		"CMBlockBufferReplaceDataBytes":                              true,
		"CMBufferQueueCallForEachBuffer":                             true,
		"CMBufferQueueInstallTrigger":                                true,
		"CMBufferQueueInstallTriggerWithIntegerThreshold":            true,
		"CMBufferQueueResetWithCallback":                             true,
		"CMBufferQueueSetValidationCallback":                         true,
		"CMIOObjectAddPropertyListener":                              true,
		"CMIOObjectGetPropertyData":                                  true,
		"CMIOObjectGetPropertyDataSize":                              true,
		"CMIOObjectPropertiesChanged":                                true,
		"CMIOObjectRemovePropertyListener":                           true,
		"CMIOObjectSetPropertyData":                                  true,
		"CMIOObjectsPublishedAndDied":                                true,
		"CMIOStreamClockCreate":                                      true,
		"CMIOStreamCopyBufferQueue":                                  true,
		"CMSampleBufferCallForEachSample":                            true,
		"CMSampleBufferCreate":                                       true,
		"CMSimpleQueueEnqueue":                                       true,
		"CVDisplayLinkSetOutputCallback":                             true,
		"CVPixelBufferCreateWithBytes":                               true,
		"CVPixelBufferCreateWithPlanarBytes":                         true,
		"IOBluetoothL2CAPChannelRegisterForChannelCloseNotification": true,
		"IOBluetoothOBEXSessionOpenTransportConnection":              true,
		"IOBluetoothPackData":                                        true,
		"IOBluetoothPackDataList":                                    true,
		"IOBluetoothUnpackData":                                      true,
		"IOBluetoothUnpackDataList":                                  true,
		"MIDIClientCreate":                                           true,
		"MIDIDestinationCreate":                                      true,
		"MIDIEndpointGetRefCons":                                     true,
		"MIDIEndpointSetRefCons":                                     true,
		"MIDIEventListForEachEvent":                                  true,
		"MIDIInputPortCreate":                                        true,
		"MIDIPortConnectSource":                                      true,
		"NSBeginAlertSheet":                                          true,
		"NSBeginCriticalAlertSheet":                                  true,
		"NSBeginInformationalAlertSheet":                             true,
		"NSCopyMemoryPages":                                          true,
		"NSDeallocateMemoryPages":                                    true,
		"NSHashGet":                                                  true,
		"NSHashInsert":                                               true,
		"NSHashInsertIfAbsent":                                       true,
		"NSHashInsertKnownAbsent":                                    true,
		"NSHashRemove":                                               true,
		"NSMapGet":                                                   true,
		"NSMapInsert":                                                true,
		"NSMapInsertIfAbsent":                                        true,
		"NSMapInsertKnownAbsent":                                     true,
		"NSMapMember":                                                true,
		"NSMapRemove":                                                true,
		"NSNextMapEnumeratorPair":                                    true,
		"NSReallocateCollectable":                                    true,
		"NSShowAnimationEffect":                                      true,
		"NSZoneFree":                                                 true,
		"NSZoneFromPointer":                                          true,
		"NSZoneRealloc":                                              true,
		"OBEXAddApplicationParameterHeader":                          true,
		"OBEXAddAuthorizationChallengeHeader":                        true,
		"OBEXAddAuthorizationResponseHeader":                         true,
		"OBEXAddBodyHeader":                                          true,
		"OBEXAddByteSequenceHeader":                                  true,
		"OBEXAddConnectionIDHeader":                                  true,
		"OBEXAddHTTPHeader":                                          true,
		"OBEXAddObjectClassHeader":                                   true,
		"OBEXAddTargetHeader":                                        true,
		"OBEXAddTimeISOHeader":                                       true,
		"OBEXAddUserDefinedHeader":                                   true,
		"OBEXAddWhoHeader":                                           true,
		"OBEXCreateVCard":                                            true,
		"OBEXGetHeaders":                                             true,
		"OBEXSessionAbort":                                           true,
		"OBEXSessionAbortResponse":                                   true,
		"OBEXSessionConnect":                                         true,
		"OBEXSessionConnectResponse":                                 true,
		"OBEXSessionDisconnect":                                      true,
		"OBEXSessionDisconnectResponse":                              true,
		"OBEXSessionGet":                                             true,
		"OBEXSessionGetResponse":                                     true,
		"OBEXSessionPut":                                             true,
		"OBEXSessionPutResponse":                                     true,
		"OBEXSessionSetPath":                                         true,
		"OBEXSessionSetPathResponse":                                 true,
		"OBEXSessionSetServerCallback":                               true,
	}
	if knownIssues[sym.Name] {
		_, err := sym.Parse()
		log.Printf("skipping function %s %s because of known issue: decl='%s' err='%v'\n", fw, sym.Name, sym.Declaration, err)
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	fntyp, ok := typ.(*typing.FunctionType)
	if !ok {
		return nil
	}
	fn := &codegen.Function{
		Name:        sym.Name,
		Deprecated:  sym.Deprecated,
		GoName:      modules.TrimPrefix(sym.Name),
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        fntyp,
	}
	// temporary skip for things deprecated in 14.0
	// check if macOS platform is DeprecatedAt 14.0
	for _, p := range sym.Platforms {
		if p.Name == "macOS" && p.Deprecated {
			fn.Deprecated = true
		}
	}

	// populate params:
	log.Printf("decl: %v %s\n", sym.Name, sym.Declaration)
	for i, p := range fntyp.Parameters {
		if p.Type != nil {
			log.Printf(" param %#v: %v %+v\n", i, p.Name, p.Type.ObjcName())
		}
	}
	for _, p := range fntyp.Parameters {
		if p.Type == nil {
			fmt.Printf("skipping %s: %s because of nil type\n", sym.Name, p.Name)
			return nil
		}
		// skip pointers to ref types (for now)
		if pt, ok := p.Type.(*typing.PointerType); ok {
			if _, ok := pt.Type.(*typing.RefType); ok {
				fmt.Printf("skipping %s: %s because of pointer to ref type\n", sym.Name, p.Name)
				return nil
			}
		}
		fn.Parameters = append(fn.Parameters, &codegen.Param{
			Name: p.Name,
			Type: p.Type,
		})
	}
	// populate return type
	if fntyp.ReturnType != nil {
		fn.ReturnType = fntyp.ReturnType
	}

	return fn

}

/*
 */
