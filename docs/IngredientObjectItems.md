# IngredientObjectItems

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Item name as provided by brand owner or as shown on packaging | [optional] [default to null]
**Categories** | **[]string** |  | [optional] [default to null]
**Nutrients** | [**[]IngredientObjectNutrients**](IngredientObject_nutrients.md) | An array containing nutrient informatio objects for this food item | [optional] [default to null]
**CalorieConversionFactor** | [***IngredientObjectCalorieConversionFactor**](IngredientObject_calorie_conversion_factor.md) |  | [optional] [default to null]
**ProteinConversionFactor** | **float64** | The multiplication factor used to calculate protein from nitrogen | [optional] [default to null]
**Components** | [**[]IngredientObjectComponents**](IngredientObject_components.md) | An array of objects containing the constituent parts of a food (e.g. bone is a component of meat) | [optional] [default to null]
**Portions** | [**[]IngredientObjectPortions**](IngredientObject_portions.md) | An array of objects containing information on discrete amounts of a food found in this item | [optional] [default to null]
**CommonName** | **string** | Common name associated with this item. These generally clarify what the item is (e.g. when the brand name is \&quot;BRAND&#x27;s Spicy Enchilada\&quot; the common name may be \&quot;Chicken enchilada\&quot;) | [optional] [default to null]
**Footnote** | **string** | Comments on any unusual aspects of this item. Examples might include unusual aspects of the food overall | [optional] [default to null]
**SearchTerm** | **string** | The original search term that found this food item | [optional] [default to null]
**Score** | **string** | A value that represents how similar the name of this food item is to the original search term. The lower the value the closer this item&#x27;s name is to the original search term. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

