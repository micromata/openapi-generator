/**
* OpenAPI Petstore
* This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
*
* The version of the OpenAPI document: 1.0.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/
package org.openapitools.client.models


import kotlinx.serialization.*
import kotlinx.serialization.internal.CommonEnumSerializer
/**
 * 
 * @param justSymbol 
 * @param arrayEnum 
 */
@Serializable
data class EnumArrays (
    @SerialName(value = "just_symbol") val justSymbol: EnumArrays.JustSymbol? = null,
    @SerialName(value = "array_enum") val arrayEnum: kotlin.Array<EnumArrays.ArrayEnum>? = null
) {

    /**
    * 
    * Values: greaterThanEqual,dollar
    */
    @Serializable(with = JustSymbol.Serializer::class)
    enum class JustSymbol(val value: kotlin.String){
        greaterThanEqual(">="),
        dollar("$");

        object Serializer : CommonEnumSerializer<JustSymbol>("JustSymbol", values(), values().map { it.value.toString() }.toTypedArray())
    }
    /**
    * 
    * Values: fish,crab
    */
    @Serializable(with = ArrayEnum.Serializer::class)
    enum class ArrayEnum(val value: kotlin.String){
        fish("fish"),
        crab("crab");

        object Serializer : CommonEnumSerializer<ArrayEnum>("ArrayEnum", values(), values().map { it.value.toString() }.toTypedArray())
    }
}

