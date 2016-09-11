/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-11 18:55:47.106260962 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// Combines input and output parameters with an interpolation algorithm to define a keyframe graph (but not its target).
type AnimationSampler struct {
	GlTFProperty
	Input         GlTFid `json:"input,omitempty" validator:"required"`
	Interpolation string `json:"interpolation"`
	Output        GlTFid `json:"output,omitempty" validator:"required"`
}
