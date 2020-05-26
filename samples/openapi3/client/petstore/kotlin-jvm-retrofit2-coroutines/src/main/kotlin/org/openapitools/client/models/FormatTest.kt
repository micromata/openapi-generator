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


import com.google.gson.annotations.SerializedName
import java.io.Serializable
/**
 * 
 * @param number 
 * @param byte 
 * @param date 
 * @param password 
 * @param integer 
 * @param int32 
 * @param int64 
 * @param float 
 * @param double 
 * @param string 
 * @param binary 
 * @param dateTime 
 * @param uuid 
 * @param patternWithDigits A string that is a 10 digit number. Can have leading zeros.
 * @param patternWithDigitsAndDelimiter A string starting with 'image_' (case insensitive) and one to three digits following i.e. Image_01.
 */

data class FormatTest (
    @SerializedName("number")
    val number: java.math.BigDecimal,
    @SerializedName("byte")
    val byte: kotlin.ByteArray,
    @SerializedName("date")
    val date: java.time.LocalDate,
    @SerializedName("password")
    val password: kotlin.String,
    @SerializedName("integer")
    val integer: kotlin.Int? = null,
    @SerializedName("int32")
    val int32: kotlin.Int? = null,
    @SerializedName("int64")
    val int64: kotlin.Long? = null,
    @SerializedName("float")
    val float: kotlin.Float? = null,
    @SerializedName("double")
    val double: kotlin.Double? = null,
    @SerializedName("string")
    val string: kotlin.String? = null,
    @SerializedName("binary")
    val binary: java.io.File? = null,
    @SerializedName("dateTime")
    val dateTime: java.time.OffsetDateTime? = null,
    @SerializedName("uuid")
    val uuid: java.util.UUID? = null,
    /* A string that is a 10 digit number. Can have leading zeros. */
    @SerializedName("pattern_with_digits")
    val patternWithDigits: kotlin.String? = null,
    /* A string starting with 'image_' (case insensitive) and one to three digits following i.e. Image_01. */
    @SerializedName("pattern_with_digits_and_delimiter")
    val patternWithDigitsAndDelimiter: kotlin.String? = null
) : Serializable {
	companion object {
		private const val serialVersionUID: Long = 123
	}

}

